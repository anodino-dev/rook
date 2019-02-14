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

package fake

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeISGWs implements ISGWInterface
type FakeISGWs struct {
	Fake *FakeEdgefsV1alpha1
	ns   string
}

var isgwsResource = schema.GroupVersionResource{Group: "edgefs.rook.io", Version: "v1alpha1", Resource: "isgws"}

var isgwsKind = schema.GroupVersionKind{Group: "edgefs.rook.io", Version: "v1alpha1", Kind: "ISGW"}

// Get takes name of the iSGW, and returns the corresponding iSGW object, and an error if there is any.
func (c *FakeISGWs) Get(name string, options v1.GetOptions) (result *v1alpha1.ISGW, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(isgwsResource, c.ns, name), &v1alpha1.ISGW{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ISGW), err
}

// List takes label and field selectors, and returns the list of ISGWs that match those selectors.
func (c *FakeISGWs) List(opts v1.ListOptions) (result *v1alpha1.ISGWList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(isgwsResource, isgwsKind, c.ns, opts), &v1alpha1.ISGWList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ISGWList{ListMeta: obj.(*v1alpha1.ISGWList).ListMeta}
	for _, item := range obj.(*v1alpha1.ISGWList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iSGWs.
func (c *FakeISGWs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(isgwsResource, c.ns, opts))

}

// Create takes the representation of a iSGW and creates it.  Returns the server's representation of the iSGW, and an error, if there is any.
func (c *FakeISGWs) Create(iSGW *v1alpha1.ISGW) (result *v1alpha1.ISGW, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(isgwsResource, c.ns, iSGW), &v1alpha1.ISGW{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ISGW), err
}

// Update takes the representation of a iSGW and updates it. Returns the server's representation of the iSGW, and an error, if there is any.
func (c *FakeISGWs) Update(iSGW *v1alpha1.ISGW) (result *v1alpha1.ISGW, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(isgwsResource, c.ns, iSGW), &v1alpha1.ISGW{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ISGW), err
}

// Delete takes name of the iSGW and deletes it. Returns an error if one occurs.
func (c *FakeISGWs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(isgwsResource, c.ns, name), &v1alpha1.ISGW{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeISGWs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(isgwsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ISGWList{})
	return err
}

// Patch applies the patch and returns the patched iSGW.
func (c *FakeISGWs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ISGW, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(isgwsResource, c.ns, name, data, subresources...), &v1alpha1.ISGW{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ISGW), err
}