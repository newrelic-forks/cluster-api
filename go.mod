module sigs.k8s.io/cluster-api

go 1.13

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/coredns/corefile-migration v1.0.6
	github.com/coreos/etcd v3.3.15+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.7.1+incompatible
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/go-logr/logr v0.1.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/gofuzz v1.1.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.0
	github.com/prometheus/client_model v0.2.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/grpc v1.23.1
	k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/apiserver v0.17.2
	k8s.io/client-go v0.17.2
	k8s.io/cluster-bootstrap v0.0.0-20190516232516-d7d78ab2cfe7
	k8s.io/component-base v0.17.2
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.1.6 // indirect
	sigs.k8s.io/cluster-api/test/framework v0.0.0-20200304170348-97097699f713 // indirect
	sigs.k8s.io/cluster-api/test/infrastructure/docker v0.0.0-20200311154737-54260061ddce // indirect
	sigs.k8s.io/controller-runtime v0.5.0
	sigs.k8s.io/yaml v1.2.0
)
