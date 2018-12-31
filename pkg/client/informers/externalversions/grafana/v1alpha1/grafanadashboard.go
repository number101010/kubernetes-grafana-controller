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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	grafanav1alpha1 "kubernetes-grafana-controller/pkg/apis/grafana/v1alpha1"
	versioned "kubernetes-grafana-controller/pkg/client/clientset/versioned"
	internalinterfaces "kubernetes-grafana-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubernetes-grafana-controller/pkg/client/listers/grafana/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GrafanaDashboardInformer provides access to a shared informer and lister for
// GrafanaDashboards.
type GrafanaDashboardInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GrafanaDashboardLister
}

type grafanaDashboardInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGrafanaDashboardInformer constructs a new informer for GrafanaDashboard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGrafanaDashboardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGrafanaDashboardInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGrafanaDashboardInformer constructs a new informer for GrafanaDashboard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGrafanaDashboardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GrafanaV1alpha1().GrafanaDashboards(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GrafanaV1alpha1().GrafanaDashboards(namespace).Watch(options)
			},
		},
		&grafanav1alpha1.GrafanaDashboard{},
		resyncPeriod,
		indexers,
	)
}

func (f *grafanaDashboardInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGrafanaDashboardInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *grafanaDashboardInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&grafanav1alpha1.GrafanaDashboard{}, f.defaultInformer)
}

func (f *grafanaDashboardInformer) Lister() v1alpha1.GrafanaDashboardLister {
	return v1alpha1.NewGrafanaDashboardLister(f.Informer().GetIndexer())
}