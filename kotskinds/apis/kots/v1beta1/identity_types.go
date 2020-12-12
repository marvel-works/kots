/*
Copyright 2019 Replicated, Inc..

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	dextypes "github.com/replicatedhq/kots/pkg/identity/types/dex"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IdentitySpec struct {
	OIDCRedirectURIs            []string            `json:"oidcRedirectUris" yaml:"oidcRedirectUris"`
	OAUTH2AlwaysShowLoginScreen bool                `json:"oauth2AlwaysShowLoginScreen,omitempty" yaml:"oauth2AlwaysShowLoginScreen,omitempty"`
	SigningKeysExpiration       string              `json:"signingKeysExpiration,omitempty" yaml:"signingKeysExpiration,omitempty"`
	IDTokensExpiration          string              `json:"idTokensExpiration,omitempty" yaml:"idTokensExpiration,omitempty"`
	SupportedProviders          []string            `json:"supportedProviders,omitempty" yaml:"supportedProviders,omitempty"`
	EnableRestrictedGroups      bool                `json:"enableRestrictedGroups,omitempty" yaml:"enableRestrictedGroups,omitempty"`
	Roles                       []string            `json:"roles,omitempty" yaml:"roles,omitempty"`
	Storage                     *StorageConfigValue `json:"storageConfig,omitempty" yaml:"storageConfig,omitempty"`
}

type StorageConfigValue struct {
	Value     *StorageConfig       `json:"value,omitempty" yaml:"value,omitempty"`
	ValueFrom *StorageConfigSource `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}

type StorageConfig dextypes.Storage

type StorageConfigSource struct {
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
}

// IdentityStatus defines the observed state of Identity
type IdentityStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Identity is the Schema for the identity document
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Identity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentitySpec   `json:"spec,omitempty"`
	Status IdentityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityList contains a list of Identities
type IdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Identity `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Identity{}, &IdentityList{})
}