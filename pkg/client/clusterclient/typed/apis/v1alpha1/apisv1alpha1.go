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

package v1alpha1

import (
	"context"
	"fmt"
	apisapiv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type WrappedApisV1alpha1 struct {
	Cluster  logicalcluster.LogicalCluster
	Delegate apisv1alpha1.ApisV1alpha1Interface
}

func (w *WrappedApisV1alpha1) RESTClient() rest.Interface {
	return w.Delegate.RESTClient()
}

func (w *WrappedApisV1alpha1) APIBindings() apisv1alpha1.APIBindingInterface {
	return &wrappedAPIBinding{
		cluster:  w.Cluster,
		delegate: w.Delegate.APIBindings(),
	}
}

type wrappedAPIBinding struct {
	cluster  logicalcluster.LogicalCluster
	delegate apisv1alpha1.APIBindingInterface
}

func (w *wrappedAPIBinding) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

func (w *wrappedAPIBinding) Create(ctx context.Context, aPIBinding *apisapiv1alpha1.APIBinding, opts metav1.CreateOptions) (*apisapiv1alpha1.APIBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, aPIBinding, opts)
}

func (w *wrappedAPIBinding) Update(ctx context.Context, aPIBinding *apisapiv1alpha1.APIBinding, opts metav1.UpdateOptions) (*apisapiv1alpha1.APIBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, aPIBinding, opts)
}

func (w *wrappedAPIBinding) UpdateStatus(ctx context.Context, aPIBinding *apisapiv1alpha1.APIBinding, opts metav1.UpdateOptions) (*apisapiv1alpha1.APIBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.UpdateStatus(ctx, aPIBinding, opts)
}

func (w *wrappedAPIBinding) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrappedAPIBinding) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

func (w *wrappedAPIBinding) Get(ctx context.Context, name string, opts metav1.GetOptions) (*apisapiv1alpha1.APIBinding, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w *wrappedAPIBinding) List(ctx context.Context, opts metav1.ListOptions) (*apisapiv1alpha1.APIBindingList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedAPIBinding) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrappedAPIBinding) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *apisapiv1alpha1.APIBinding, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}

func (w *WrappedApisV1alpha1) APIExports() apisv1alpha1.APIExportInterface {
	return &wrappedAPIExport{
		cluster:  w.Cluster,
		delegate: w.Delegate.APIExports(),
	}
}

type wrappedAPIExport struct {
	cluster  logicalcluster.LogicalCluster
	delegate apisv1alpha1.APIExportInterface
}

func (w *wrappedAPIExport) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

func (w *wrappedAPIExport) Create(ctx context.Context, aPIExport *apisapiv1alpha1.APIExport, opts metav1.CreateOptions) (*apisapiv1alpha1.APIExport, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, aPIExport, opts)
}

func (w *wrappedAPIExport) Update(ctx context.Context, aPIExport *apisapiv1alpha1.APIExport, opts metav1.UpdateOptions) (*apisapiv1alpha1.APIExport, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, aPIExport, opts)
}

func (w *wrappedAPIExport) UpdateStatus(ctx context.Context, aPIExport *apisapiv1alpha1.APIExport, opts metav1.UpdateOptions) (*apisapiv1alpha1.APIExport, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.UpdateStatus(ctx, aPIExport, opts)
}

func (w *wrappedAPIExport) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrappedAPIExport) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

func (w *wrappedAPIExport) Get(ctx context.Context, name string, opts metav1.GetOptions) (*apisapiv1alpha1.APIExport, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w *wrappedAPIExport) List(ctx context.Context, opts metav1.ListOptions) (*apisapiv1alpha1.APIExportList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedAPIExport) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrappedAPIExport) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *apisapiv1alpha1.APIExport, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}

func (w *WrappedApisV1alpha1) APIResourceSchemas() apisv1alpha1.APIResourceSchemaInterface {
	return &wrappedAPIResourceSchema{
		cluster:  w.Cluster,
		delegate: w.Delegate.APIResourceSchemas(),
	}
}

type wrappedAPIResourceSchema struct {
	cluster  logicalcluster.LogicalCluster
	delegate apisv1alpha1.APIResourceSchemaInterface
}

func (w *wrappedAPIResourceSchema) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

func (w *wrappedAPIResourceSchema) Create(ctx context.Context, aPIResourceSchema *apisapiv1alpha1.APIResourceSchema, opts metav1.CreateOptions) (*apisapiv1alpha1.APIResourceSchema, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, aPIResourceSchema, opts)
}

func (w *wrappedAPIResourceSchema) Update(ctx context.Context, aPIResourceSchema *apisapiv1alpha1.APIResourceSchema, opts metav1.UpdateOptions) (*apisapiv1alpha1.APIResourceSchema, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, aPIResourceSchema, opts)
}

func (w *wrappedAPIResourceSchema) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrappedAPIResourceSchema) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

func (w *wrappedAPIResourceSchema) Get(ctx context.Context, name string, opts metav1.GetOptions) (*apisapiv1alpha1.APIResourceSchema, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w *wrappedAPIResourceSchema) List(ctx context.Context, opts metav1.ListOptions) (*apisapiv1alpha1.APIResourceSchemaList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedAPIResourceSchema) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrappedAPIResourceSchema) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *apisapiv1alpha1.APIResourceSchema, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}