module sigs.k8s.io/cluster-api/test/infrastructure/docker

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	sigs.k8s.io/cluster-api v0.3.0-rc.2.0.20200302175844-3011d8c2580c
	sigs.k8s.io/cluster-api/test/framework v0.0.0-20200304170348-97097699f713
	sigs.k8s.io/controller-runtime v0.5.0
	sigs.k8s.io/kind v0.7.1-0.20200303021537-981bd80d3802
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191121015412-41065c7a8c2a
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	sigs.k8s.io/cluster-api => ../../..
	sigs.k8s.io/cluster-api/test/framework => ../../framework
)
