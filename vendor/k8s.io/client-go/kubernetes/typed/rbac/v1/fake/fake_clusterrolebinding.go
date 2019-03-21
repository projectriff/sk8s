/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	rbac_v1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRoleBindings implements ClusterRoleBindingInterface
type FakeClusterRoleBindings struct {
	Fake *FakeRbacV1
}

var clusterrolebindingsResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "clusterrolebindings"}

var clusterrolebindingsKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "ClusterRoleBinding"}

// Get takes name of the clusterRoleBinding, and returns the corresponding clusterRoleBinding object, and an error if there is any.
func (c *FakeClusterRoleBindings) Get(name string, options v1.GetOptions) (result *rbac_v1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterrolebindingsResource, name), &rbac_v1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbac_v1.ClusterRoleBinding), err
}

// List takes label and field selectors, and returns the list of ClusterRoleBindings that match those selectors.
func (c *FakeClusterRoleBindings) List(opts v1.ListOptions) (result *rbac_v1.ClusterRoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterrolebindingsResource, clusterrolebindingsKind, opts), &rbac_v1.ClusterRoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbac_v1.ClusterRoleBindingList{}
	for _, item := range obj.(*rbac_v1.ClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoleBindings.
func (c *FakeClusterRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterrolebindingsResource, opts))
}

// Create takes the representation of a clusterRoleBinding and creates it.  Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *FakeClusterRoleBindings) Create(clusterRoleBinding *rbac_v1.ClusterRoleBinding) (result *rbac_v1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterrolebindingsResource, clusterRoleBinding), &rbac_v1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbac_v1.ClusterRoleBinding), err
}

// Update takes the representation of a clusterRoleBinding and updates it. Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *FakeClusterRoleBindings) Update(clusterRoleBinding *rbac_v1.ClusterRoleBinding) (result *rbac_v1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterrolebindingsResource, clusterRoleBinding), &rbac_v1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbac_v1.ClusterRoleBinding), err
}

// Delete takes name of the clusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterrolebindingsResource, name), &rbac_v1.ClusterRoleBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterrolebindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &rbac_v1.ClusterRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterRoleBinding.
func (c *FakeClusterRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *rbac_v1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrolebindingsResource, name, data, subresources...), &rbac_v1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbac_v1.ClusterRoleBinding), err
}
