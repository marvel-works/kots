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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/replicatedhq/kots/kotskinds/apis/kots/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAirgaps implements AirgapInterface
type FakeAirgaps struct {
	Fake *FakeKotsV1beta1
	ns   string
}

var airgapsResource = schema.GroupVersionResource{Group: "kots.io", Version: "v1beta1", Resource: "airgaps"}

var airgapsKind = schema.GroupVersionKind{Group: "kots.io", Version: "v1beta1", Kind: "Airgap"}

// Get takes name of the airgap, and returns the corresponding airgap object, and an error if there is any.
func (c *FakeAirgaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Airgap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(airgapsResource, c.ns, name), &v1beta1.Airgap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Airgap), err
}

// List takes label and field selectors, and returns the list of Airgaps that match those selectors.
func (c *FakeAirgaps) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AirgapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(airgapsResource, airgapsKind, c.ns, opts), &v1beta1.AirgapList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AirgapList{ListMeta: obj.(*v1beta1.AirgapList).ListMeta}
	for _, item := range obj.(*v1beta1.AirgapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested airgaps.
func (c *FakeAirgaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(airgapsResource, c.ns, opts))

}

// Create takes the representation of a airgap and creates it.  Returns the server's representation of the airgap, and an error, if there is any.
func (c *FakeAirgaps) Create(ctx context.Context, airgap *v1beta1.Airgap, opts v1.CreateOptions) (result *v1beta1.Airgap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(airgapsResource, c.ns, airgap), &v1beta1.Airgap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Airgap), err
}

// Update takes the representation of a airgap and updates it. Returns the server's representation of the airgap, and an error, if there is any.
func (c *FakeAirgaps) Update(ctx context.Context, airgap *v1beta1.Airgap, opts v1.UpdateOptions) (result *v1beta1.Airgap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(airgapsResource, c.ns, airgap), &v1beta1.Airgap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Airgap), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAirgaps) UpdateStatus(ctx context.Context, airgap *v1beta1.Airgap, opts v1.UpdateOptions) (*v1beta1.Airgap, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(airgapsResource, "status", c.ns, airgap), &v1beta1.Airgap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Airgap), err
}

// Delete takes name of the airgap and deletes it. Returns an error if one occurs.
func (c *FakeAirgaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(airgapsResource, c.ns, name), &v1beta1.Airgap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAirgaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(airgapsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AirgapList{})
	return err
}

// Patch applies the patch and returns the patched airgap.
func (c *FakeAirgaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Airgap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(airgapsResource, c.ns, name, pt, data, subresources...), &v1beta1.Airgap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Airgap), err
}
