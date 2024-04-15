//go:build e2e
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
	. "github.com/onsi/ginkgo/v2"
	"k8s.io/utils/pointer"
)

var _ = Describe("When testing Cluster API working on self-hosted clusters", func() {
	SelfHostedSpec(ctx, func() SelfHostedSpecInput {
		return SelfHostedSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			InfrastructureProvider: pointer.String("docker"),
		}
	})
})

var _ = Describe("When testing Cluster API working on self-hosted clusters using ClusterClass [ClusterClass]", func() {
	SelfHostedSpec(ctx, func() SelfHostedSpecInput {
		return SelfHostedSpecInput{
			E2EConfig:                e2eConfig,
			ClusterctlConfigPath:     clusterctlConfigPath,
			BootstrapClusterProxy:    bootstrapClusterProxy,
			ArtifactFolder:           artifactFolder,
			SkipCleanup:              skipCleanup,
			Flavor:                   "topology",
			InfrastructureProvider:   pointer.String("docker"),
			ControlPlaneMachineCount: pointer.Int64(1),
			WorkerMachineCount:       pointer.Int64(1),
		}
	})
})

var _ = Describe("When testing Cluster API working on self-hosted clusters using ClusterClass with a HA control plane [ClusterClass]", func() {
	SelfHostedSpec(ctx, func() SelfHostedSpecInput {
		return SelfHostedSpecInput{
			E2EConfig:                e2eConfig,
			ClusterctlConfigPath:     clusterctlConfigPath,
			BootstrapClusterProxy:    bootstrapClusterProxy,
			ArtifactFolder:           artifactFolder,
			SkipCleanup:              skipCleanup,
			Flavor:                   "topology",
			InfrastructureProvider:   pointer.String("docker"),
			ControlPlaneMachineCount: pointer.Int64(3),
			WorkerMachineCount:       pointer.Int64(1),
		}
	})
})

var _ = Describe("When testing Cluster API working on single-node self-hosted clusters using ClusterClass [ClusterClass]", func() {
	SelfHostedSpec(ctx, func() SelfHostedSpecInput {
		return SelfHostedSpecInput{
			E2EConfig:                e2eConfig,
			ClusterctlConfigPath:     clusterctlConfigPath,
			BootstrapClusterProxy:    bootstrapClusterProxy,
			ArtifactFolder:           artifactFolder,
			SkipCleanup:              skipCleanup,
			Flavor:                   "topology-single-node-cluster",
			InfrastructureProvider:   pointer.String("docker"),
			ControlPlaneMachineCount: pointer.Int64(1),
			WorkerMachineCount:       pointer.Int64(0),
		}
	})
})
