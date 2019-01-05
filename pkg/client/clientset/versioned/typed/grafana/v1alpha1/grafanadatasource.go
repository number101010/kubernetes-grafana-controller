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

package v1alpha1

import (
	v1alpha1 "kubernetes-grafana-controller/pkg/apis/grafana/v1alpha1"
	scheme "kubernetes-grafana-controller/pkg/client/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GrafanaDataSourcesGetter has a method to return a GrafanaDataSourceInterface.
// A group's client should implement this interface.
type GrafanaDataSourcesGetter interface {
	GrafanaDataSources(namespace string) GrafanaDataSourceInterface
}

// GrafanaDataSourceInterface has methods to work with GrafanaDataSource resources.
type GrafanaDataSourceInterface interface {
	Create(*v1alpha1.GrafanaDataSource) (*v1alpha1.GrafanaDataSource, error)
	Update(*v1alpha1.GrafanaDataSource) (*v1alpha1.GrafanaDataSource, error)
	UpdateStatus(*v1alpha1.GrafanaDataSource) (*v1alpha1.GrafanaDataSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GrafanaDataSource, error)
	List(opts v1.ListOptions) (*v1alpha1.GrafanaDataSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GrafanaDataSource, err error)
	GrafanaDataSourceExpansion
}

// grafanaDataSources implements GrafanaDataSourceInterface
type grafanaDataSources struct {
	client rest.Interface
	ns     string
}

// newGrafanaDataSources returns a GrafanaDataSources
func newGrafanaDataSources(c *GrafanaV1alpha1Client, namespace string) *grafanaDataSources {
	return &grafanaDataSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the grafanaDataSource, and returns the corresponding grafanaDataSource object, and an error if there is any.
func (c *grafanaDataSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GrafanaDataSource, err error) {
	result = &v1alpha1.GrafanaDataSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("grafanadatasources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GrafanaDataSources that match those selectors.
func (c *grafanaDataSources) List(opts v1.ListOptions) (result *v1alpha1.GrafanaDataSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GrafanaDataSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("grafanadatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested grafanaDataSources.
func (c *grafanaDataSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("grafanadatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a grafanaDataSource and creates it.  Returns the server's representation of the grafanaDataSource, and an error, if there is any.
func (c *grafanaDataSources) Create(grafanaDataSource *v1alpha1.GrafanaDataSource) (result *v1alpha1.GrafanaDataSource, err error) {
	result = &v1alpha1.GrafanaDataSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("grafanadatasources").
		Body(grafanaDataSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a grafanaDataSource and updates it. Returns the server's representation of the grafanaDataSource, and an error, if there is any.
func (c *grafanaDataSources) Update(grafanaDataSource *v1alpha1.GrafanaDataSource) (result *v1alpha1.GrafanaDataSource, err error) {
	result = &v1alpha1.GrafanaDataSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("grafanadatasources").
		Name(grafanaDataSource.Name).
		Body(grafanaDataSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *grafanaDataSources) UpdateStatus(grafanaDataSource *v1alpha1.GrafanaDataSource) (result *v1alpha1.GrafanaDataSource, err error) {
	result = &v1alpha1.GrafanaDataSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("grafanadatasources").
		Name(grafanaDataSource.Name).
		SubResource("status").
		Body(grafanaDataSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the grafanaDataSource and deletes it. Returns an error if one occurs.
func (c *grafanaDataSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("grafanadatasources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *grafanaDataSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("grafanadatasources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched grafanaDataSource.
func (c *grafanaDataSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GrafanaDataSource, err error) {
	result = &v1alpha1.GrafanaDataSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("grafanadatasources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}