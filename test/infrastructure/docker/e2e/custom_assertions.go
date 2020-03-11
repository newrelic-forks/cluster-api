// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

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

package e2e

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	types "github.com/onsi/gomega/types"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/test/framework"
	"sigs.k8s.io/cluster-api/test/framework/options"
	infrav1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ensureDockerArtifactsDeletedInput is an example of a provider specific assertion.
type ensureDockerArtifactsDeletedInput struct {
	Lister  framework.Lister
	Cluster *clusterv1.Cluster
}

// ensureDockerArtifactsDeleted ensure we have cleaned up provider specific objects.
func ensureDockerArtifactsDeleted(input ensureDockerArtifactsDeletedInput) {
	if options.SkipResourceCleanup {
		return
	}
	By("Ensuring docker artifacts have been deleted")
	ctx := context.Background()

	lbl, err := labels.Parse(fmt.Sprintf("%s=%s", clusterv1.ClusterLabelName, input.Cluster.GetClusterName()))
	Expect(err).ToNot(HaveOccurred())
	opt := &client.ListOptions{LabelSelector: lbl}

	dcl := &infrav1.DockerClusterList{}
	Expect(input.Lister.List(ctx, dcl, opt)).To(Succeed())
	Expect(dcl.Items).To(HaveLen(0))

	dml := &infrav1.DockerMachineList{}
	Expect(input.Lister.List(ctx, dml, opt)).To(Succeed())
	Expect(dml.Items).To(HaveLen(0))

	dmtl := &infrav1.DockerMachineTemplateList{}
	Expect(input.Lister.List(ctx, dmtl, opt)).To(Succeed())
	Expect(dmtl.Items).To(HaveLen(0))
	By("Succeeding in deleting all docker artifacts")
}

type controllerMatch struct {
	kind  string
	owner metav1.Object
}

func (m *controllerMatch) Match(actual interface{}) (success bool, err error) {
	actualMeta, err := meta.Accessor(actual)
	if err != nil {
		return false, fmt.Errorf("unable to read meta for %T: %v", actual, err)
	}

	owner := metav1.GetControllerOf(actualMeta)
	if owner == nil {
		return false, fmt.Errorf("no controller found (owner ref with controller = true) for object %#v", actual)
	}

	match := (owner.Kind == m.kind &&
		owner.Name == m.owner.GetName() && owner.UID == m.owner.GetUID())

	return match, nil
}

func (m *controllerMatch) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected\n\t%#vto have a controller reference pointing to %s/%s (%v)", actual, m.kind, m.owner.GetName(), m.owner.GetUID())
}

func (m *controllerMatch) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected\n\t%#vto not have a controller reference pointing to %s/%s (%v)", actual, m.kind, m.owner.GetName(), m.owner.GetUID())
}

func HaveControllerRef(kind string, owner metav1.Object) types.GomegaMatcher {
	return &controllerMatch{kind, owner}
}
