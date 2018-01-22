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

package internalversion

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	settings "k8s.io/kubernetes/pkg/apis/settings"
	scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

// PodPresetsGetter has a method to return a PodPresetInterface.
// A group's client should implement this interface.
type PodPresetsGetter interface {
	PodPresets(namespace string) PodPresetInterface
}

// PodPresetInterface has methods to work with PodPreset resources.
type PodPresetInterface interface {
	Create(ctx context.Context, obj *settings.PodPreset) (*settings.PodPreset, error)
	Update(ctx context.Context, obj *settings.PodPreset) (*settings.PodPreset, error)
	Delete(ctx context.Context, name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context, name string, options v1.GetOptions) (*settings.PodPreset, error)
	List(ctx context.Context, opts v1.ListOptions) (*settings.PodPresetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *settings.PodPreset, err error)
	PodPresetExpansion
}

// podPresets implements PodPresetInterface
type podPresets struct {
	client rest.Interface
	ns     string
}

// newPodPresets returns a PodPresets
func newPodPresets(c *SettingsClient, namespace string) *podPresets {
	return &podPresets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podPreset, and returns the corresponding podPreset object, and an error if there is any.
func (c *podPresets) Get(ctx context.Context, name string, options v1.GetOptions) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podpresets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodPresets that match those selectors.
func (c *podPresets) List(ctx context.Context, opts v1.ListOptions) (result *settings.PodPresetList, err error) {
	result = &settings.PodPresetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podPresets.
func (c *podPresets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Watch()
}

// Create takes the representation of a podPreset and creates it.  Returns the server's representation of the podPreset, and an error, if there is any.
func (c *podPresets) Create(ctx context.Context, podPreset *settings.PodPreset) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podpresets").
		Body(podPreset).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Update takes the representation of a podPreset and updates it. Returns the server's representation of the podPreset, and an error, if there is any.
func (c *podPresets) Update(ctx context.Context, podPreset *settings.PodPreset) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podpresets").
		Name(podPreset.Name).
		Body(podPreset).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Delete takes name of the podPreset and deletes it. Returns an error if one occurs.
func (c *podPresets) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podpresets").
		Name(name).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podPresets) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// Patch applies the patch and returns the patched podPreset.
func (c *podPresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podpresets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Context(ctx).
		Do().
		Into(result)
	return
}
