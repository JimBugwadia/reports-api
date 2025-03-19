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

package v1beta2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta2 "sigs.k8s.io/wg-policy-prototypes/policy-report/apis/reports.x-k8s.io/v1beta2"
	scheme "sigs.k8s.io/wg-policy-prototypes/policy-report/pkg/client/clientset/versioned/scheme"
)

// ClusterReportsGetter has a method to return a ClusterReportInterface.
// A group's client should implement this interface.
type ClusterReportsGetter interface {
	ClusterReports() ClusterReportInterface
}

// ClusterReportInterface has methods to work with ClusterReport resources.
type ClusterReportInterface interface {
	Create(ctx context.Context, clusterReport *v1beta2.ClusterReport, opts v1.CreateOptions) (*v1beta2.ClusterReport, error)
	Update(ctx context.Context, clusterReport *v1beta2.ClusterReport, opts v1.UpdateOptions) (*v1beta2.ClusterReport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.ClusterReport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.ClusterReportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ClusterReport, err error)
	ClusterReportExpansion
}

// clusterReports implements ClusterReportInterface
type clusterReports struct {
	client rest.Interface
}

// newClusterReports returns a ClusterReports
func newClusterReports(c *ReportsV1beta2Client) *clusterReports {
	return &clusterReports{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterReport, and returns the corresponding clusterReport object, and an error if there is any.
func (c *clusterReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.ClusterReport, err error) {
	result = &v1beta2.ClusterReport{}
	err = c.client.Get().
		Resource("clusterreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterReports that match those selectors.
func (c *clusterReports) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ClusterReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.ClusterReportList{}
	err = c.client.Get().
		Resource("clusterreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterReports.
func (c *clusterReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterReport and creates it.  Returns the server's representation of the clusterReport, and an error, if there is any.
func (c *clusterReports) Create(ctx context.Context, clusterReport *v1beta2.ClusterReport, opts v1.CreateOptions) (result *v1beta2.ClusterReport, err error) {
	result = &v1beta2.ClusterReport{}
	err = c.client.Post().
		Resource("clusterreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterReport and updates it. Returns the server's representation of the clusterReport, and an error, if there is any.
func (c *clusterReports) Update(ctx context.Context, clusterReport *v1beta2.ClusterReport, opts v1.UpdateOptions) (result *v1beta2.ClusterReport, err error) {
	result = &v1beta2.ClusterReport{}
	err = c.client.Put().
		Resource("clusterreports").
		Name(clusterReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterReport and deletes it. Returns an error if one occurs.
func (c *clusterReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterreports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterreports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterReport.
func (c *clusterReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ClusterReport, err error) {
	result = &v1beta2.ClusterReport{}
	err = c.client.Patch(pt).
		Resource("clusterreports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
