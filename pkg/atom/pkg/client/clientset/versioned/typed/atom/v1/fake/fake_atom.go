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
	atom_v1 "github.com/convox/convox/pkg/atom/pkg/apis/atom/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAtoms implements AtomInterface
type FakeAtoms struct {
	Fake *FakeAtomV1
	ns   string
}

var atomsResource = schema.GroupVersionResource{Group: "atom.convox.com", Version: "v1", Resource: "atoms"}

var atomsKind = schema.GroupVersionKind{Group: "atom.convox.com", Version: "v1", Kind: "Atom"}

// Get takes name of the atom, and returns the corresponding atom object, and an error if there is any.
func (c *FakeAtoms) Get(name string, options v1.GetOptions) (result *atom_v1.Atom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(atomsResource, c.ns, name), &atom_v1.Atom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*atom_v1.Atom), err
}

// List takes label and field selectors, and returns the list of Atoms that match those selectors.
func (c *FakeAtoms) List(opts v1.ListOptions) (result *atom_v1.AtomList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(atomsResource, atomsKind, c.ns, opts), &atom_v1.AtomList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &atom_v1.AtomList{ListMeta: obj.(*atom_v1.AtomList).ListMeta}
	for _, item := range obj.(*atom_v1.AtomList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested atoms.
func (c *FakeAtoms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(atomsResource, c.ns, opts))

}

// Create takes the representation of a atom and creates it.  Returns the server's representation of the atom, and an error, if there is any.
func (c *FakeAtoms) Create(atom *atom_v1.Atom) (result *atom_v1.Atom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(atomsResource, c.ns, atom), &atom_v1.Atom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*atom_v1.Atom), err
}

// Update takes the representation of a atom and updates it. Returns the server's representation of the atom, and an error, if there is any.
func (c *FakeAtoms) Update(atom *atom_v1.Atom) (result *atom_v1.Atom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(atomsResource, c.ns, atom), &atom_v1.Atom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*atom_v1.Atom), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAtoms) UpdateStatus(atom *atom_v1.Atom) (*atom_v1.Atom, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(atomsResource, "status", c.ns, atom), &atom_v1.Atom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*atom_v1.Atom), err
}

// Delete takes name of the atom and deletes it. Returns an error if one occurs.
func (c *FakeAtoms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(atomsResource, c.ns, name), &atom_v1.Atom{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAtoms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(atomsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &atom_v1.AtomList{})
	return err
}

// Patch applies the patch and returns the patched atom.
func (c *FakeAtoms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *atom_v1.Atom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(atomsResource, c.ns, name, data, subresources...), &atom_v1.Atom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*atom_v1.Atom), err
}