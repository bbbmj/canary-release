/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/tenant/v1alpha1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type TenantV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClusterQuotasGetter
	PartitionsGetter
	TenantsGetter
}

// TenantV1alpha1Client is used to interact with features provided by the tenant.caicloud.io group.
type TenantV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TenantV1alpha1Client) ClusterQuotas() ClusterQuotaInterface {
	return newClusterQuotas(c)
}

func (c *TenantV1alpha1Client) Partitions() PartitionInterface {
	return newPartitions(c)
}

func (c *TenantV1alpha1Client) Tenants() TenantInterface {
	return newTenants(c)
}

// NewForConfig creates a new TenantV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*TenantV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TenantV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TenantV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TenantV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TenantV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TenantV1alpha1Client {
	return &TenantV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TenantV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
