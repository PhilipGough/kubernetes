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
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

// DeploymentsGetter has a method to return a DeploymentInterface.
// A group's client should implement this interface.
type DeploymentsGetter interface {
	Deployments(namespace string) DeploymentInterface
}

// DeploymentInterface has methods to work with Deployment resources.
type DeploymentInterface interface {
	Create(ctx context.Context, obj *extensions.Deployment) (*extensions.Deployment, error)
	Update(ctx context.Context, obj *extensions.Deployment) (*extensions.Deployment, error)
	UpdateStatus(ctx context.Context, obj *extensions.Deployment) (*extensions.Deployment, error)
	Delete(ctx context.Context, name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context, name string, options v1.GetOptions) (*extensions.Deployment, error)
	List(ctx context.Context, opts v1.ListOptions) (*extensions.DeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.Deployment, err error)
	GetScale(ctx context.Context, deploymentName string, options v1.GetOptions) (*autoscaling.Scale, error)
	UpdateScale(ctx context.Context, deploymentName string, scale *autoscaling.Scale) (*autoscaling.Scale, error)

	DeploymentExpansion
}

// deployments implements DeploymentInterface
type deployments struct {
	client rest.Interface
	ns     string
}

// newDeployments returns a Deployments
func newDeployments(c *ExtensionsClient, namespace string) *deployments {
	return &deployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *deployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *extensions.Deployment, err error) {
	result = &extensions.Deployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deployments) List(ctx context.Context, opts v1.ListOptions) (result *extensions.DeploymentList, err error) {
	result = &extensions.DeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *deployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Watch()
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Create(ctx context.Context, deployment *extensions.Deployment) (result *extensions.Deployment, err error) {
	result = &extensions.Deployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deployments").
		Body(deployment).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Update(ctx context.Context, deployment *extensions.Deployment) (result *extensions.Deployment, err error) {
	result = &extensions.Deployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		Body(deployment).
		Context(ctx).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deployments) UpdateStatus(ctx context.Context, deployment *extensions.Deployment) (result *extensions.Deployment, err error) {
	result = &extensions.Deployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		SubResource("status").
		Body(deployment).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *deployments) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployments) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deployment.
func (c *deployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.Deployment, err error) {
	result = &extensions.Deployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Context(ctx).
		Do().
		Into(result)
	return
}

// GetScale takes name of the deployment, and returns the corresponding autoscaling.Scale object, and an error if there is any.
func (c *deployments) GetScale(ctx context.Context, deploymentName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *deployments) UpdateScale(ctx context.Context, deploymentName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		Body(scale).
		Context(ctx).
		Do().
		Into(result)
	return
}
