//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code auto-generated. DO NOT EDIT.

package clusterclient

import (
	"fmt"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"

	workloadv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/workload/v1alpha1"
	workloadv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clusterclient/typed/workload/v1alpha1"

	apiresourcev1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apiresource/v1alpha1"
	apiresourcev1alpha1client "github.com/kcp-dev/kcp/pkg/client/clusterclient/typed/apiresource/v1alpha1"

	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1alpha1"
	tenancyv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clusterclient/typed/tenancy/v1alpha1"

	tenancyv1beta1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1beta1"
	tenancyv1beta1client "github.com/kcp-dev/kcp/pkg/client/clusterclient/typed/tenancy/v1beta1"

	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1"
	apisv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clusterclient/typed/apis/v1alpha1"

	schedulingv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	schedulingv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clusterclient/typed/scheduling/v1alpha1"
)

func NewForConfig(config *rest.Config) (*ClusterClient, error) {
	client, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %w", err)
	}

	clusterRoundTripper := kcp.NewClusterRoundTripper(client.Transport)
	client.Transport = clusterRoundTripper

	delegate, err := versioned.NewForConfigAndClient(config, client)
	if err != nil {
		return nil, fmt.Errorf("error creating delegate clientset: %w", err)
	}

	return &ClusterClient{
		delegate: delegate,
	}, nil
}

type ClusterClient struct {
	delegate versioned.Interface
}

func (c *ClusterClient) Cluster(cluster logicalcluster.LogicalCluster) versioned.Interface {
	return &wrappedInterface{
		cluster:  cluster,
		delegate: c.delegate,
	}
}

type wrappedInterface struct {
	cluster  logicalcluster.LogicalCluster
	delegate versioned.Interface
}

func (w *wrappedInterface) Discovery() discovery.DiscoveryInterface {
	return w.delegate.Discovery()
}

func (w *wrappedInterface) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface {
	return &workloadv1alpha1client.WrappedWorkloadV1alpha1{
		Cluster:  w.cluster,
		Delegate: w.delegate.WorkloadV1alpha1(),
	}
}

func (w *wrappedInterface) ApiresourceV1alpha1() apiresourcev1alpha1.ApiresourceV1alpha1Interface {
	return &apiresourcev1alpha1client.WrappedApiresourceV1alpha1{
		Cluster:  w.cluster,
		Delegate: w.delegate.ApiresourceV1alpha1(),
	}
}

func (w *wrappedInterface) TenancyV1alpha1() tenancyv1alpha1.TenancyV1alpha1Interface {
	return &tenancyv1alpha1client.WrappedTenancyV1alpha1{
		Cluster:  w.cluster,
		Delegate: w.delegate.TenancyV1alpha1(),
	}
}

func (w *wrappedInterface) TenancyV1beta1() tenancyv1beta1.TenancyV1beta1Interface {
	return &tenancyv1beta1client.WrappedTenancyV1beta1{
		Cluster:  w.cluster,
		Delegate: w.delegate.TenancyV1beta1(),
	}
}

func (w *wrappedInterface) ApisV1alpha1() apisv1alpha1.ApisV1alpha1Interface {
	return &apisv1alpha1client.WrappedApisV1alpha1{
		Cluster:  w.cluster,
		Delegate: w.delegate.ApisV1alpha1(),
	}
}

func (w *wrappedInterface) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return &schedulingv1alpha1client.WrappedSchedulingV1alpha1{
		Cluster:  w.cluster,
		Delegate: w.delegate.SchedulingV1alpha1(),
	}
}
