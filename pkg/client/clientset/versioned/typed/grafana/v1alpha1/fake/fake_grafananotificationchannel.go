/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "kubernetes-grafana-controller/pkg/apis/grafana/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGrafanaNotificationChannels implements GrafanaNotificationChannelInterface
type FakeGrafanaNotificationChannels struct {
	Fake *FakeGrafanaV1alpha1
	ns   string
}

var grafananotificationchannelsResource = schema.GroupVersionResource{Group: "grafana.k8s.io", Version: "v1alpha1", Resource: "grafananotificationchannels"}

var grafananotificationchannelsKind = schema.GroupVersionKind{Group: "grafana.k8s.io", Version: "v1alpha1", Kind: "GrafanaNotificationChannel"}

// Get takes name of the grafanaNotificationChannel, and returns the corresponding grafanaNotificationChannel object, and an error if there is any.
func (c *FakeGrafanaNotificationChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.GrafanaNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grafananotificationchannelsResource, c.ns, name), &v1alpha1.GrafanaNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaNotificationChannel), err
}

// List takes label and field selectors, and returns the list of GrafanaNotificationChannels that match those selectors.
func (c *FakeGrafanaNotificationChannels) List(opts v1.ListOptions) (result *v1alpha1.GrafanaNotificationChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grafananotificationchannelsResource, grafananotificationchannelsKind, c.ns, opts), &v1alpha1.GrafanaNotificationChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GrafanaNotificationChannelList{ListMeta: obj.(*v1alpha1.GrafanaNotificationChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.GrafanaNotificationChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested grafanaNotificationChannels.
func (c *FakeGrafanaNotificationChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grafananotificationchannelsResource, c.ns, opts))

}

// Create takes the representation of a grafanaNotificationChannel and creates it.  Returns the server's representation of the grafanaNotificationChannel, and an error, if there is any.
func (c *FakeGrafanaNotificationChannels) Create(grafanaNotificationChannel *v1alpha1.GrafanaNotificationChannel) (result *v1alpha1.GrafanaNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grafananotificationchannelsResource, c.ns, grafanaNotificationChannel), &v1alpha1.GrafanaNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaNotificationChannel), err
}

// Update takes the representation of a grafanaNotificationChannel and updates it. Returns the server's representation of the grafanaNotificationChannel, and an error, if there is any.
func (c *FakeGrafanaNotificationChannels) Update(grafanaNotificationChannel *v1alpha1.GrafanaNotificationChannel) (result *v1alpha1.GrafanaNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grafananotificationchannelsResource, c.ns, grafanaNotificationChannel), &v1alpha1.GrafanaNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaNotificationChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGrafanaNotificationChannels) UpdateStatus(grafanaNotificationChannel *v1alpha1.GrafanaNotificationChannel) (*v1alpha1.GrafanaNotificationChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(grafananotificationchannelsResource, "status", c.ns, grafanaNotificationChannel), &v1alpha1.GrafanaNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaNotificationChannel), err
}

// Delete takes name of the grafanaNotificationChannel and deletes it. Returns an error if one occurs.
func (c *FakeGrafanaNotificationChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(grafananotificationchannelsResource, c.ns, name), &v1alpha1.GrafanaNotificationChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGrafanaNotificationChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grafananotificationchannelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GrafanaNotificationChannelList{})
	return err
}

// Patch applies the patch and returns the patched grafanaNotificationChannel.
func (c *FakeGrafanaNotificationChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GrafanaNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grafananotificationchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GrafanaNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GrafanaNotificationChannel), err
}
