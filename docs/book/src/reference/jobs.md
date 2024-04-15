# Jobs

This document intents to provide an overview over our jobs running via Prow, GitHub actions and Google Cloud Build.
It also documents the cluster-api specific configuration in test-infra.

## Builds and Tests running on the main branch

> NOTE: To see which test jobs execute which tests or e2e tests, you can click on the links which lead to the respective test overviews in testgrid.

### Presubmits

Prow Presubmits:
<<<<<<< HEAD
* [pull-cluster-api-build-main] `./scripts/ci-build.sh`
* ✳️️ [pull-cluster-api-apidiff-main] `./scripts/ci-apidiff.sh`
* [pull-cluster-api-verify-main] `./scripts/ci-verify.sh`
* [pull-cluster-api-test-main] `./scripts/ci-test.sh`
* [pull-cluster-api-test-mink8s-main] `./scripts/ci-test.sh`
* [pull-cluster-api-e2e-main] `./scripts/ci-e2e.sh`
  * GINKGO_FOCUS: `[PR-Blocking]`
* ✳️️ [pull-cluster-api-e2e-informing-main] `./scripts/ci-e2e.sh`
  * GINKGO_FOCUS: `[PR-Informing]`, GINKGO_SKIP: `[IPv6]`
* ✳️️ [pull-cluster-api-e2e-informing-ipv6-main] `./scripts/ci-e2e.sh`
  * GINKGO_FOCUS: `[IPv6] [PR-Informing]`, IP_FAMILY: `IPv6`
* ✳️️ ✴️ [pull-cluster-api-e2e-full-main] `./scripts/ci-e2e.sh`
  * GINKGO_SKIP: `[PR-Blocking] [Conformance] [K8s-Upgrade] [IPv6]` (i.e. "no tags")
* ✳️️ ✴️ [pull-cluster-api-e2e-workload-upgrade-1-25-latest-main] `./scripts/ci-e2e.sh` FROM: `stable-1.25` TO: `ci/latest-1.26`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
=======
* mandatory for merge, always run:
  * [pull-cluster-api-build-main] `./scripts/ci-build.sh`
  * [pull-cluster-api-verify-main] `./scripts/ci-verify.sh`
* mandatory for merge, run if go code changes:
  * [pull-cluster-api-test-main] `./scripts/ci-test.sh`
  * [pull-cluster-api-e2e-main] `./scripts/ci-e2e.sh`
    * GINKGO_FOCUS: `[PR-Blocking]`
* optional for merge, run if go code changes:
  * [pull-cluster-api-apidiff-main] `./scripts/ci-apidiff.sh`
  * [pull-cluster-api-e2e-informing-main] `./scripts/ci-e2e.sh`
    * GINKGO_FOCUS: `[PR-Informing]`
* mandatory for merge, run if manually triggered:
  * [pull-cluster-api-test-mink8s-main] `./scripts/ci-test.sh`
    * KUBEBUILDER_ENVTEST_KUBERNETES_VERSION: `1.24.2`
  * [pull-cluster-api-e2e-mink8s-main] `./scripts/ci-e2e.sh`
    * GINKGO_SKIP: `[Conformance] [K8s-Upgrade]|[IPv6]`
    * KUBERNETES_VERSION_MANAGEMENT: `stable-1.24`
  * [pull-cluster-api-e2e-full-dualstack-and-ipv6-main] `./scripts/ci-e2e.sh`
    * DOCKER_IN_DOCKER_IPV6_ENABLED: `true`
    * GINKGO_SKIP: `[PR-Blocking] [Conformance] [K8s-Upgrade]`
  * [pull-cluster-api-e2e-full-main] `./scripts/ci-e2e.sh`
    * GINKGO_SKIP: `[PR-Blocking] [Conformance] [K8s-Upgrade]|[IPv6]`
  * [pull-cluster-api-e2e-workload-upgrade-1-28-latest-main] `./scripts/ci-e2e.sh` FROM: `stable-1.28` TO: `ci/latest-1.29`
    * GINKGO_FOCUS: `[K8s-Upgrade]`
* optional for merge, run if manually triggered:
  * [pull-cluster-api-e2e-scale-main-experimental] `./scripts/ci-e2e-scale.sh`
>>>>>>> v1.5.7

GitHub Presubmit Workflows:
* golangci-lint: golangci/golangci-lint-action
  * Runs golangci-lint. Can be run locally via `make lint`.
* verify: kubernetes-sigs/kubebuilder-release-tools verifier
  * Verifies the PR titles have a valid format, i.e. contains one of the valid icons.
  * Verifies the PR description is valid, i.e. is long enough.
* Check PR Markdown links (run when markdown files changed)
<<<<<<< HEAD
  * Checks markdown for broken links. 
=======
  * Checks markdown for broken links.
>>>>>>> v1.5.7
* dependabot (run on dependabot PRs)
  * Regenerates Go modules and code.
* release (run on tags)
  * Creates a GitHub release with release notes for the tag.
<<<<<<< HEAD
=======

GitHub Weekly Workflows:
* golangci-lint: golangci/golangci-lint-action
  * Weekly check all Markdown links
* scan-images:
  * Scan all images for vulnerabilities. Can be run locally via `make verify-container-images`
>>>>>>> v1.5.7

### Postsubmits

Prow Postsubmits:
* [post-cluster-api-push-images] Google Cloud Build: `make release-staging`

### Periodics

Prow Periodics:
* [periodic-cluster-api-test-main] `./scripts/ci-test.sh`
* [periodic-cluster-api-test-mink8s-main] `./scripts/ci-test.sh`
<<<<<<< HEAD
  * KUBEBUILDER_ENVTEST_KUBERNETES_VERSION: `1.20.2` 
* [periodic-cluster-api-e2e-main] `./scripts/ci-e2e.sh`
  * GINKGO_SKIP: `[Conformance] [K8s-Upgrade]|[IPv6]`
* [periodic-cluster-api-e2e-mink8s-main] `./scripts/ci-e2e.sh`
  * GINKGO_SKIP: `[Conformance] [K8s-Upgrade]|[IPv6]|[ClusterClass]`
  * KUBERNETES_VERSION_MANAGEMENT: `stable-1.20`, CLUSTER_TOPOLOGY: `false`
* [periodic-cluster-api-e2e-workload-upgrade-1-18-1-19-main] `./scripts/ci-e2e.sh` FROM: `stable-1.18` TO: `stable-1.19`
=======
  * KUBEBUILDER_ENVTEST_KUBERNETES_VERSION: `1.24.2`
* [periodic-cluster-api-e2e-main] `./scripts/ci-e2e.sh`
  * GINKGO_SKIP: `[Conformance] [K8s-Upgrade]|[IPv6]`
* [periodic-cluster-api-e2e-mink8s-main] `./scripts/ci-e2e.sh`
  * GINKGO_SKIP: `[Conformance] [K8s-Upgrade]|[IPv6]`
  * KUBERNETES_VERSION_MANAGEMENT: `stable-1.24`
* [periodic-cluster-api-e2e-dualstack-and-ipv6-main] `./scripts/ci-e2e.sh`
  * DOCKER_IN_DOCKER_IPV6_ENABLED: `true`
  * GINKGO_SKIP: `[Conformance] [K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-23-1-24-main] `./scripts/ci-e2e.sh` FROM: `stable-1.23` TO: `stable-1.24`
>>>>>>> v1.5.7
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-24-1-25-main] `./scripts/ci-e2e.sh` FROM: `stable-1.24` TO: `stable-1.25`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-25-1-26-main] `./scripts/ci-e2e.sh` FROM: `stable-1.25` TO: `stable-1.26`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-26-1-27-main] `./scripts/ci-e2e.sh` FROM: `stable-1.26` TO: `stable-1.27`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-27-1-28-main] `./scripts/ci-e2e.sh` FROM: `stable-1.27` TO: `stable-1.28`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
<<<<<<< HEAD
* [periodic-cluster-api-e2e-workload-upgrade-1-23-1-24-main] `./scripts/ci-e2e.sh` FROM: `stable-1.23` TO: `stable-1.24`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-24-1-25-main] `./scripts/ci-e2e.sh` FROM: `stable-1.24` TO: `stable-1.25`
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [periodic-cluster-api-e2e-workload-upgrade-1-25-latest-main] `./scripts/ci-e2e.sh` FROM: `stable-1.25` TO: `ci/latest-1.26`
=======
* [periodic-cluster-api-e2e-workload-upgrade-1-28-latest-main] `./scripts/ci-e2e.sh` FROM: `stable-1.28` TO: `ci/latest-1.29`
>>>>>>> v1.5.7
  * GINKGO_FOCUS: `[K8s-Upgrade]`
* [cluster-api-push-images-nightly] Google Cloud Build: `make release-staging-nightly`

## Test-infra configuration

* config/jobs/image-pushing/k8s-staging-cluster-api.yaml
<<<<<<< HEAD
  * Configures nightly and postsubmit jobs to push images and manifests. 
* config/jobs/kubernetes-sigs/cluster-api/
  * Configures Cluster API  presubmit and periodic jobs.
* config/testgrids/kubernetes/sig-cluster-lifecycle/config.yaml
  * Configures Cluster API testgrid dashboards.  
=======
  * Configures nightly and postsubmit jobs to push images and manifests.
* config/jobs/kubernetes-sigs/cluster-api/
  * Configures Cluster API  presubmit and periodic jobs.
* config/testgrids/kubernetes/sig-cluster-lifecycle/config.yaml
  * Configures Cluster API testgrid dashboards.
>>>>>>> v1.5.7
* config/prow/config.yaml
  * `branch-protection` and `tide` are configured to make the golangci-lint GitHub action mandatory for merge
* config/prow/plugins.yaml
  * `triggers`: configures `/ok-to-test`
  * `approve`: disable auto-approval of PR authors, ignore GitHub reviews (/approve is explicitly required)
  * `milestone_applier`: configures that merged PRs are automatically added to the correct milestone after merge
  * `repo_milestone`: configures `cluster-api-maintainers` as maintainers
  * `require_matching_label`: configures `needs-triage`
  * `plugins`: enables `milestone`, `override` and `require-matching-label` plugins
  * `external_plugins`: enables `cherrypicker`
* label_sync/labels.yaml
<<<<<<< HEAD
  * Configures labels for the `cluster-api` repository. 
=======
  * Configures labels for the `cluster-api` repository.
>>>>>>> v1.5.7


<!-- links -->
[pull-cluster-api-build-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-build-main
[pull-cluster-api-apidiff-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-apidiff-main
[pull-cluster-api-verify-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-verify-main
[pull-cluster-api-test-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-test-main
[pull-cluster-api-test-mink8s-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-test-mink8s-main
[pull-cluster-api-e2e-mink8s-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-mink8s-main
[pull-cluster-api-e2e-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-main
[pull-cluster-api-e2e-informing-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-informing-main
[pull-cluster-api-e2e-full-dualstack-and-ipv6-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-full-dualstack-and-ipv6-main
[pull-cluster-api-e2e-full-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-full-main
<<<<<<< HEAD
[pull-cluster-api-e2e-workload-upgrade-1-25-latest-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-main-1-25-latest
=======
[pull-cluster-api-e2e-workload-upgrade-1-28-latest-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-main-1-28-latest
[pull-cluster-api-e2e-scale-main-experimental]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-pr-e2e-scale-main-experimental
>>>>>>> v1.5.7
[periodic-cluster-api-test-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-test-main
[periodic-cluster-api-test-mink8s-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-test-mink8s-main
[periodic-cluster-api-e2e-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main
[periodic-cluster-api-e2e-mink8s-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-mink8s-main
<<<<<<< HEAD
[periodic-cluster-api-e2e-workload-upgrade-1-18-1-19-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-18-1-19
[periodic-cluster-api-e2e-workload-upgrade-1-19-1-20-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-19-1-20
[periodic-cluster-api-e2e-workload-upgrade-1-20-1-21-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-20-1-21
[periodic-cluster-api-e2e-workload-upgrade-1-21-1-22-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-21-1-22
[periodic-cluster-api-e2e-workload-upgrade-1-22-1-23-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-22-1-23
[periodic-cluster-api-e2e-workload-upgrade-1-23-1-24-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-23-1-24
[periodic-cluster-api-e2e-workload-upgrade-1-24-1-25-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-24-1-25
[periodic-cluster-api-e2e-workload-upgrade-1-25-latest-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-25-latest
=======
[periodic-cluster-api-e2e-dualstack-and-ipv6-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-dualstack-and-ipv6-main
[periodic-cluster-api-e2e-workload-upgrade-1-23-1-24-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-23-1-24
[periodic-cluster-api-e2e-workload-upgrade-1-24-1-25-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-24-1-25
[periodic-cluster-api-e2e-workload-upgrade-1-25-1-26-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-25-1-26
[periodic-cluster-api-e2e-workload-upgrade-1-26-1-27-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-26-1-27
[periodic-cluster-api-e2e-workload-upgrade-1-27-1-28-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-27-1-28
[periodic-cluster-api-e2e-workload-upgrade-1-28-latest-main]: https://testgrid.k8s.io/sig-cluster-lifecycle-cluster-api#capi-e2e-main-1-28-latest
>>>>>>> v1.5.7
[cluster-api-push-images-nightly]: https://testgrid.k8s.io/sig-cluster-lifecycle-image-pushes#cluster-api-push-images-nightly
[post-cluster-api-push-images]: https://testgrid.k8s.io/sig-cluster-lifecycle-image-pushes#post-cluster-api-push-images
