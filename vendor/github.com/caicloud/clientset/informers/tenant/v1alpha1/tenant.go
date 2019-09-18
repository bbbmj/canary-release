/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/tenant/v1alpha1"
	tenantv1alpha1 "github.com/caicloud/clientset/pkg/apis/tenant/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// TenantInformer provides access to a shared informer and lister for
// Tenants.
type TenantInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TenantLister
}

type tenantInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTenantInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTenantInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTenantInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenantV1alpha1().Tenants().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenantV1alpha1().Tenants().Watch(options)
			},
		},
		&tenantv1alpha1.Tenant{},
		resyncPeriod,
		indexers,
	)
}

func (f *tenantInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTenantInformer(client.(kubernetes.Interface), resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tenantInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tenantv1alpha1.Tenant{}, f.defaultInformer)
}

func (f *tenantInformer) Lister() v1alpha1.TenantLister {
	return v1alpha1.NewTenantLister(f.Informer().GetIndexer())
}
