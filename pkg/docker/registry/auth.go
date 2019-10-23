package registry

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/docker/distribution/registry/client/auth/challenge"
	"github.com/pkg/errors"
)

func TestPushAccess(endpoint, username, password, org string) error {

	endpoint = sanitizeEndpoint(endpoint)

	// We need to check if we can push images to a repo.
	// We cannot get push permission to an org alone.
	scope := org + "/testrepo"
	basicAuthToken := makeBasicAuthToken(username, password)

	if IsECREndpoint(endpoint) {
		token, err := GetECRBasicAuthToken(endpoint, username, password)
		if err != nil {
			return errors.Wrap(err, "failed to ping registry")
		}
		basicAuthToken = token
		scope = org // ECR has no concept of organization and it should be an empty string
	}

	// TODO: Support http
	pingURL := fmt.Sprintf("https://%s/v2/", endpoint)

	resp, err := http.Get(pingURL)
	if err != nil {
		return errors.Wrap(err, "failed to ping registry")
	}

	if resp.StatusCode == http.StatusOK {
		// Anonymous registry that does not require authentication
		return nil
	}

	if resp.StatusCode != http.StatusUnauthorized {
		return errors.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	challenges := challenge.ResponseChallenges(resp)
	if len(challenges) == 0 {
		return errors.Wrap(err, "no auth challenges found for endpoint")
	}

	if challenges[0].Scheme == "basic" {
		// ecr uses basic auth. not much more we can do here without actually pushing an image
		return nil
	}

	host := challenges[0].Parameters["realm"]
	v := url.Values{}
	v.Set("service", challenges[0].Parameters["service"])
	v.Set("scope", fmt.Sprintf("repository:%s:push", scope))

	authURL := host + "?" + v.Encode()

	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create auth request")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuthToken))

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to execute auth request")
	}
	defer resp.Body.Close()

	authBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to load auth response")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(errorResponseToString(authBody))
	}

	if strings.Contains(endpoint, "gcr.io") {
		// gcr tokens are in a different format. it's enough to test that a token has been issued.
		return nil
	}

	bearerToken, err := newBearerTokenFromJSONBlob(authBody)
	if err != nil {
		return errors.Wrap(err, "failed to parse bearer token")
	}

	jwtToken, err := bearerToken.getJwtToken()
	if err != nil {
		return errors.Wrap(err, "failed to parse JWT token")
	}

	claims, err := getJwtTokenClaims(jwtToken)
	if err != nil {
		return errors.Wrap(err, "failed to get claims")
	}

	for _, access := range claims.Access {
		if access.Type != "repository" {
			continue
		}
		if access.Name != scope {
			continue
		}
		for _, action := range access.Actions {
			if action == "push" {
				return nil
			}
		}
	}

	return errors.Errorf("%q has no push permission in %q", username, org)
}

func makeBasicAuthToken(username, password string) string {
	token := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(token))
}

func sanitizeEndpoint(endpoint string) string {
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "https://")
	endpoint = strings.TrimSuffix(endpoint, "/v2/")
	endpoint = strings.TrimSuffix(endpoint, "/v2")
	endpoint = strings.TrimSuffix(endpoint, "/v1/")
	endpoint = strings.TrimSuffix(endpoint, "/v1")
	return endpoint
}