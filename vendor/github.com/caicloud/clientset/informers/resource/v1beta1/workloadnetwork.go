/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1beta1 "github.com/caicloud/clientset/listers/resource/v1beta1"
	resourcev1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// WorkloadNetworkInformer provides access to a shared informer and lister for
// WorkloadNetworks.
type WorkloadNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.WorkloadNetworkLister
}

type workloadNetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkloadNetworkInformer constructs a new informer for WorkloadNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkloadNetworkInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkloadNetworkInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkloadNetworkInformer constructs a new informer for WorkloadNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkloadNetworkInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1beta1().WorkloadNetworks(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1beta1().WorkloadNetworks(namespace).Watch(options)
			},
		},
		&resourcev1beta1.WorkloadNetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *workloadNetworkInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkloadNetworkInformer(client.(kubernetes.Interface), f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workloadNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1beta1.WorkloadNetwork{}, f.defaultInformer)
}

func (f *workloadNetworkInformer) Lister() v1beta1.WorkloadNetworkLister {
	return v1beta1.NewWorkloadNetworkLister(f.Informer().GetIndexer())
}
