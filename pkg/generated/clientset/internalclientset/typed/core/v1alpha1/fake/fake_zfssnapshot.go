/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZFSSnapshots implements ZFSSnapshotInterface
type FakeZFSSnapshots struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var zfssnapshotsResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "zfssnapshots"}

var zfssnapshotsKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "ZFSSnapshot"}

// Get takes name of the zFSSnapshot, and returns the corresponding zFSSnapshot object, and an error if there is any.
func (c *FakeZFSSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.ZFSSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zfssnapshotsResource, c.ns, name), &v1alpha1.ZFSSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSSnapshot), err
}

// List takes label and field selectors, and returns the list of ZFSSnapshots that match those selectors.
func (c *FakeZFSSnapshots) List(opts v1.ListOptions) (result *v1alpha1.ZFSSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zfssnapshotsResource, zfssnapshotsKind, c.ns, opts), &v1alpha1.ZFSSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ZFSSnapshotList{ListMeta: obj.(*v1alpha1.ZFSSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.ZFSSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zFSSnapshots.
func (c *FakeZFSSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zfssnapshotsResource, c.ns, opts))

}

// Create takes the representation of a zFSSnapshot and creates it.  Returns the server's representation of the zFSSnapshot, and an error, if there is any.
func (c *FakeZFSSnapshots) Create(zFSSnapshot *v1alpha1.ZFSSnapshot) (result *v1alpha1.ZFSSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zfssnapshotsResource, c.ns, zFSSnapshot), &v1alpha1.ZFSSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSSnapshot), err
}

// Update takes the representation of a zFSSnapshot and updates it. Returns the server's representation of the zFSSnapshot, and an error, if there is any.
func (c *FakeZFSSnapshots) Update(zFSSnapshot *v1alpha1.ZFSSnapshot) (result *v1alpha1.ZFSSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zfssnapshotsResource, c.ns, zFSSnapshot), &v1alpha1.ZFSSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeZFSSnapshots) UpdateStatus(zFSSnapshot *v1alpha1.ZFSSnapshot) (*v1alpha1.ZFSSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zfssnapshotsResource, "status", c.ns, zFSSnapshot), &v1alpha1.ZFSSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSSnapshot), err
}

// Delete takes name of the zFSSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeZFSSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zfssnapshotsResource, c.ns, name), &v1alpha1.ZFSSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZFSSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zfssnapshotsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ZFSSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched zFSSnapshot.
func (c *FakeZFSSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ZFSSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zfssnapshotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ZFSSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZFSSnapshot), err
}
