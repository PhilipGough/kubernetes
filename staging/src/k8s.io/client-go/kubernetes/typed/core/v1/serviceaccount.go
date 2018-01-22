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

package v1

import (
	context "context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// ServiceAccountsGetter has a method to return a ServiceAccountInterface.
// A group's client should implement this interface.
type ServiceAccountsGetter interface {
	ServiceAccounts(namespace string) ServiceAccountInterface
}

// ServiceAccountInterface has methods to work with ServiceAccount resources.
type ServiceAccountInterface interface {
	Create(ctx context.Context, obj *v1.ServiceAccount) (*v1.ServiceAccount, error)
	Update(ctx context.Context, obj *v1.ServiceAccount) (*v1.ServiceAccount, error)
	Delete(ctx context.Context, name string, options *metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.ServiceAccount, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ServiceAccountList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ServiceAccount, err error)
	ServiceAccountExpansion
}

// serviceAccounts implements ServiceAccountInterface
type serviceAccounts struct {
	client rest.Interface
	ns     string
}

// newServiceAccounts returns a ServiceAccounts
func newServiceAccounts(c *CoreV1Client, namespace string) *serviceAccounts {
	return &serviceAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceAccount, and returns the corresponding serviceAccount object, and an error if there is any.
func (c *serviceAccounts) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ServiceAccount, err error) {
	result = &v1.ServiceAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceAccounts that match those selectors.
func (c *serviceAccounts) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ServiceAccountList, err error) {
	result = &v1.ServiceAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceAccounts.
func (c *serviceAccounts) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Watch()
}

// Create takes the representation of a serviceAccount and creates it.  Returns the server's representation of the serviceAccount, and an error, if there is any.
func (c *serviceAccounts) Create(ctx context.Context, serviceAccount *v1.ServiceAccount) (result *v1.ServiceAccount, err error) {
	result = &v1.ServiceAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceaccounts").
		Body(serviceAccount).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceAccount and updates it. Returns the server's representation of the serviceAccount, and an error, if there is any.
func (c *serviceAccounts) Update(ctx context.Context, serviceAccount *v1.ServiceAccount) (result *v1.ServiceAccount, err error) {
	result = &v1.ServiceAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceaccounts").
		Name(serviceAccount.Name).
		Body(serviceAccount).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceAccount and deletes it. Returns an error if one occurs.
func (c *serviceAccounts) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceaccounts").
		Name(name).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceAccounts) DeleteCollection(ctx context.Context, options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceAccount.
func (c *serviceAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ServiceAccount, err error) {
	result = &v1.ServiceAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Context(ctx).
		Do().
		Into(result)
	return
}
