/*
Copyright 2022 The Kubernetes Authors.

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

package cluster

import (
<<<<<<< HEAD
=======
	"context"
>>>>>>> v1.5.7
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
<<<<<<< HEAD
	"golang.org/x/net/context"
=======
>>>>>>> v1.5.7
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/cluster-api/cmd/clusterctl/internal/test"
)

func Test_CRDMigrator(t *testing.T) {
	tests := []struct {
		name               string
		CRs                []unstructured.Unstructured
		currentCRD         *apiextensionsv1.CustomResourceDefinition
		newCRD             *apiextensionsv1.CustomResourceDefinition
		wantIsMigrated     bool
		wantStoredVersions []string
		wantErr            bool
	}{
		{
			name:           "No-op if current CRD does not exists",
			currentCRD:     &apiextensionsv1.CustomResourceDefinition{ObjectMeta: metav1.ObjectMeta{Name: "something else"}}, // There is currently no "foo" CRD
			newCRD:         &apiextensionsv1.CustomResourceDefinition{ObjectMeta: metav1.ObjectMeta{Name: "foo"}},
			wantIsMigrated: false,
		},
		{
			name: "Error if current CRD does not have a storage version",
			currentCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1alpha1"}, // No storage version as storage is not set.
=======
						{Name: "v1alpha1", Served: true}, // No storage version as storage is not set.
>>>>>>> v1.5.7
					},
				},
				Status: apiextensionsv1.CustomResourceDefinitionStatus{StoredVersions: []string{"v1alpha1"}},
			},
			newCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1alpha1"},
=======
						{Name: "v1alpha1", Served: true},
>>>>>>> v1.5.7
					},
				},
			},
			wantErr:        true,
			wantIsMigrated: false,
		},
		{
			name: "No-op if new CRD supports same versions",
			currentCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1alpha1", Storage: true},
=======
						{Name: "v1alpha1", Storage: true, Served: true},
>>>>>>> v1.5.7
					},
				},
				Status: apiextensionsv1.CustomResourceDefinitionStatus{StoredVersions: []string{"v1alpha1"}},
			},
			newCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1alpha1", Storage: true},
=======
						{Name: "v1alpha1", Storage: true, Served: true},
>>>>>>> v1.5.7
					},
				},
			},
			wantIsMigrated: false,
		},
		{
			name: "No-op if new CRD adds a new versions",
			currentCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1alpha1", Storage: true},
=======
						{Name: "v1alpha1", Storage: true, Served: true},
>>>>>>> v1.5.7
					},
				},
				Status: apiextensionsv1.CustomResourceDefinitionStatus{StoredVersions: []string{"v1alpha1"}},
			},
			newCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1beta1", Storage: true}, // v1beta1 is being added
						{Name: "v1alpha1"},               // v1alpha1 still exists
=======
						{Name: "v1beta1", Storage: true, Served: true}, // v1beta1 is being added
						{Name: "v1alpha1", Served: true},               // v1alpha1 still exists
>>>>>>> v1.5.7
					},
				},
			},
			wantIsMigrated: false,
		},
		{
			name: "Fails if new CRD drops current storage version",
			currentCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1alpha1", Storage: true},
=======
						{Name: "v1alpha1", Storage: true, Served: true},
>>>>>>> v1.5.7
					},
				},
				Status: apiextensionsv1.CustomResourceDefinitionStatus{StoredVersions: []string{"v1alpha1"}},
			},
			newCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1", Storage: true}, // CRD is jumping to v1, but dropping current storage version without allowing migration.
=======
						{Name: "v1", Storage: true, Served: true}, // CRD is jumping to v1, but dropping current storage version without allowing migration.
>>>>>>> v1.5.7
					},
				},
			},
			wantErr: true,
		},
		{
<<<<<<< HEAD
			name: "Migrate",
=======
			name: "Migrate CRs if their storage version is removed from the CRD",
>>>>>>> v1.5.7
			CRs: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "foo/v1beta1",
						"kind":       "Foo",
						"metadata": map[string]interface{}{
							"name":      "cr1",
							"namespace": metav1.NamespaceDefault,
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "foo/v1beta1",
						"kind":       "Foo",
						"metadata": map[string]interface{}{
							"name":      "cr2",
							"namespace": metav1.NamespaceDefault,
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "foo/v1beta1",
						"kind":       "Foo",
						"metadata": map[string]interface{}{
							"name":      "cr3",
							"namespace": metav1.NamespaceDefault,
						},
					},
				},
			},
			currentCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Group: "foo",
					Names: apiextensionsv1.CustomResourceDefinitionNames{Kind: "Foo", ListKind: "FooList"},
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1beta1", Storage: true},
						{Name: "v1alpha1"},
=======
						{Name: "v1beta1", Storage: true, Served: true},
						{Name: "v1alpha1", Served: true},
>>>>>>> v1.5.7
					},
				},
				Status: apiextensionsv1.CustomResourceDefinitionStatus{StoredVersions: []string{"v1beta1", "v1alpha1"}},
			},
			newCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Group: "foo",
					Names: apiextensionsv1.CustomResourceDefinitionNames{Kind: "Foo", ListKind: "FooList"},
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
<<<<<<< HEAD
						{Name: "v1", Storage: true}, // v1 is being added
						{Name: "v1beta1"},           // v1beta1 still there (required for migration)
=======
						{Name: "v1", Storage: true, Served: true}, // v1 is being added
						{Name: "v1beta1", Served: true},           // v1beta1 still there (required for migration)
>>>>>>> v1.5.7
						// v1alpha1 is being dropped
					},
				},
			},
			wantStoredVersions: []string{"v1beta1"}, // v1alpha1 should be dropped from current CRD's stored versions
			wantIsMigrated:     true,
		},
<<<<<<< HEAD
=======
		{
			name: "Migrate the CR if their storage version is no longer served by the CRD",
			CRs: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "foo/v1beta1",
						"kind":       "Foo",
						"metadata": map[string]interface{}{
							"name":      "cr1",
							"namespace": metav1.NamespaceDefault,
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "foo/v1beta1",
						"kind":       "Foo",
						"metadata": map[string]interface{}{
							"name":      "cr2",
							"namespace": metav1.NamespaceDefault,
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "foo/v1beta1",
						"kind":       "Foo",
						"metadata": map[string]interface{}{
							"name":      "cr3",
							"namespace": metav1.NamespaceDefault,
						},
					},
				},
			},
			currentCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Group: "foo",
					Names: apiextensionsv1.CustomResourceDefinitionNames{Kind: "Foo", ListKind: "FooList"},
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
						{Name: "v1beta1", Storage: true, Served: true},
						{Name: "v1alpha1", Served: true},
					},
				},
				Status: apiextensionsv1.CustomResourceDefinitionStatus{StoredVersions: []string{"v1beta1", "v1alpha1"}},
			},
			newCRD: &apiextensionsv1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "foo"},
				Spec: apiextensionsv1.CustomResourceDefinitionSpec{
					Group: "foo",
					Names: apiextensionsv1.CustomResourceDefinitionNames{Kind: "Foo", ListKind: "FooList"},
					Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
						{Name: "v1", Storage: true, Served: true}, // v1 is being added
						{Name: "v1beta1", Served: true},           // v1beta1 still there (required for migration)
						{Name: "v1alpha1", Served: false},         // v1alpha1 is no longer being served (required for migration)
					},
				},
			},
			wantStoredVersions: []string{"v1beta1"}, // v1alpha1 should be dropped from current CRD's stored versions
			wantIsMigrated:     true,
		},
>>>>>>> v1.5.7
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			objs := []client.Object{tt.currentCRD}
			for i := range tt.CRs {
				objs = append(objs, &tt.CRs[i])
			}

			c, err := test.NewFakeProxy().WithObjs(objs...).NewClient()
			g.Expect(err).ToNot(HaveOccurred())
			countingClient := newUpgradeCountingClient(c)

			m := crdMigrator{
				Client: countingClient,
			}

			isMigrated, err := m.run(ctx, tt.newCRD)
			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
			g.Expect(isMigrated).To(Equal(tt.wantIsMigrated))

			if isMigrated {
				storageVersion, err := storageVersionForCRD(tt.currentCRD)
				g.Expect(err).ToNot(HaveOccurred())

				// Check all the objects has been migrated.
				g.Expect(countingClient.count).To(HaveKeyWithValue(fmt.Sprintf("%s/%s, Kind=%s", tt.currentCRD.Spec.Group, storageVersion, tt.currentCRD.Spec.Names.Kind), len(tt.CRs)))

				// Check storage versions has been cleaned up.
				currentCRD := &apiextensionsv1.CustomResourceDefinition{}
				err = c.Get(ctx, client.ObjectKeyFromObject(tt.newCRD), currentCRD)
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(currentCRD.Status.StoredVersions).To(Equal(tt.wantStoredVersions))
			}
		})
	}
}

type UpgradeCountingClient struct {
	count map[string]int
	client.Client
}

func newUpgradeCountingClient(inner client.Client) UpgradeCountingClient {
	return UpgradeCountingClient{
		count:  map[string]int{},
		Client: inner,
	}
}

func (u UpgradeCountingClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	u.count[obj.GetObjectKind().GroupVersionKind().String()]++
	return u.Client.Update(ctx, obj, opts...)
}
