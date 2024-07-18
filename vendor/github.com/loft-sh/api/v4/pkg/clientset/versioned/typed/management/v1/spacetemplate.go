// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SpaceTemplatesGetter has a method to return a SpaceTemplateInterface.
// A group's client should implement this interface.
type SpaceTemplatesGetter interface {
	SpaceTemplates() SpaceTemplateInterface
}

// SpaceTemplateInterface has methods to work with SpaceTemplate resources.
type SpaceTemplateInterface interface {
	Create(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.CreateOptions) (*v1.SpaceTemplate, error)
	Update(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.UpdateOptions) (*v1.SpaceTemplate, error)
	UpdateStatus(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.UpdateOptions) (*v1.SpaceTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SpaceTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SpaceTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpaceTemplate, err error)
	SpaceTemplateExpansion
}

// spaceTemplates implements SpaceTemplateInterface
type spaceTemplates struct {
	client rest.Interface
}

// newSpaceTemplates returns a SpaceTemplates
func newSpaceTemplates(c *ManagementV1Client) *spaceTemplates {
	return &spaceTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the spaceTemplate, and returns the corresponding spaceTemplate object, and an error if there is any.
func (c *spaceTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SpaceTemplate, err error) {
	result = &v1.SpaceTemplate{}
	err = c.client.Get().
		Resource("spacetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpaceTemplates that match those selectors.
func (c *spaceTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SpaceTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SpaceTemplateList{}
	err = c.client.Get().
		Resource("spacetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spaceTemplates.
func (c *spaceTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("spacetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a spaceTemplate and creates it.  Returns the server's representation of the spaceTemplate, and an error, if there is any.
func (c *spaceTemplates) Create(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.CreateOptions) (result *v1.SpaceTemplate, err error) {
	result = &v1.SpaceTemplate{}
	err = c.client.Post().
		Resource("spacetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spaceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a spaceTemplate and updates it. Returns the server's representation of the spaceTemplate, and an error, if there is any.
func (c *spaceTemplates) Update(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.UpdateOptions) (result *v1.SpaceTemplate, err error) {
	result = &v1.SpaceTemplate{}
	err = c.client.Put().
		Resource("spacetemplates").
		Name(spaceTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spaceTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *spaceTemplates) UpdateStatus(ctx context.Context, spaceTemplate *v1.SpaceTemplate, opts metav1.UpdateOptions) (result *v1.SpaceTemplate, err error) {
	result = &v1.SpaceTemplate{}
	err = c.client.Put().
		Resource("spacetemplates").
		Name(spaceTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spaceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the spaceTemplate and deletes it. Returns an error if one occurs.
func (c *spaceTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("spacetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spaceTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("spacetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched spaceTemplate.
func (c *spaceTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SpaceTemplate, err error) {
	result = &v1.SpaceTemplate{}
	err = c.client.Patch(pt).
		Resource("spacetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}