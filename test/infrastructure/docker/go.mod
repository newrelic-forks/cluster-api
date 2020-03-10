module sigs.k8s.io/cluster-api/test/infrastructure/docker

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/pkg/errors v0.9.0
	gopkg.in/yaml.v3 v3.0.0-20191120175047-4206685974f2
	k8s.io/api v0.0.0-20191121015604-11707872ac1c
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	sigs.k8s.io/cluster-api v0.2.6-0.20191223162332-fd807a3d843b
	sigs.k8s.io/cluster-api/test/framework v0.0.0-20191216153338-db9baf9ade0c
	sigs.k8s.io/controller-runtime v0.4.0
	sigs.k8s.io/kind v0.7.0
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191121015412-41065c7a8c2a
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	sigs.k8s.io/cluster-api => ../../..
	sigs.k8s.io/cluster-api/test/framework => ../../framework
)
