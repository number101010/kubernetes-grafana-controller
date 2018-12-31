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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubernetes-grafana-controller/pkg/apis/grafana/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GrafanaNotificationChannelLister helps list GrafanaNotificationChannels.
type GrafanaNotificationChannelLister interface {
	// List lists all GrafanaNotificationChannels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GrafanaNotificationChannel, err error)
	// GrafanaNotificationChannels returns an object that can list and get GrafanaNotificationChannels.
	GrafanaNotificationChannels(namespace string) GrafanaNotificationChannelNamespaceLister
	GrafanaNotificationChannelListerExpansion
}

// grafanaNotificationChannelLister implements the GrafanaNotificationChannelLister interface.
type grafanaNotificationChannelLister struct {
	indexer cache.Indexer
}

// NewGrafanaNotificationChannelLister returns a new GrafanaNotificationChannelLister.
func NewGrafanaNotificationChannelLister(indexer cache.Indexer) GrafanaNotificationChannelLister {
	return &grafanaNotificationChannelLister{indexer: indexer}
}

// List lists all GrafanaNotificationChannels in the indexer.
func (s *grafanaNotificationChannelLister) List(selector labels.Selector) (ret []*v1alpha1.GrafanaNotificationChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GrafanaNotificationChannel))
	})
	return ret, err
}

// GrafanaNotificationChannels returns an object that can list and get GrafanaNotificationChannels.
func (s *grafanaNotificationChannelLister) GrafanaNotificationChannels(namespace string) GrafanaNotificationChannelNamespaceLister {
	return grafanaNotificationChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GrafanaNotificationChannelNamespaceLister helps list and get GrafanaNotificationChannels.
type GrafanaNotificationChannelNamespaceLister interface {
	// List lists all GrafanaNotificationChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GrafanaNotificationChannel, err error)
	// Get retrieves the GrafanaNotificationChannel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GrafanaNotificationChannel, error)
	GrafanaNotificationChannelNamespaceListerExpansion
}

// grafanaNotificationChannelNamespaceLister implements the GrafanaNotificationChannelNamespaceLister
// interface.
type grafanaNotificationChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GrafanaNotificationChannels in the indexer for a given namespace.
func (s grafanaNotificationChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GrafanaNotificationChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GrafanaNotificationChannel))
	})
	return ret, err
}

// Get retrieves the GrafanaNotificationChannel from the indexer for a given namespace and name.
func (s grafanaNotificationChannelNamespaceLister) Get(name string) (*v1alpha1.GrafanaNotificationChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("grafananotificationchannel"), name)
	}
	return obj.(*v1alpha1.GrafanaNotificationChannel), nil
}