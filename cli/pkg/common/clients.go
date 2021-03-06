package common

import (
	"fmt"
	"time"

	glooV1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	istiosecret "github.com/solo-io/supergloo/pkg/api/external/istio/encryption/v1"
	superglooV1 "github.com/solo-io/supergloo/pkg/api/v1"
	factory2 "github.com/solo-io/supergloo/pkg/factory"
	k8sApiExt "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetUpstreamClient() (*glooV1.UpstreamClient, error) {
	config, err := GetKubernetesConfig(0)
	if err != nil {
		return nil, err
	}

	upstreamClient, err := glooV1.NewUpstreamClient(&factory.KubeResourceClientFactory{
		Crd:         glooV1.UpstreamCrd,
		Cfg:         config,
		SharedCache: kube.NewKubeCache(),
	})
	if err != nil {
		return nil, err
	}
	if err = upstreamClient.Register(); err != nil {
		return nil, err
	}
	return &upstreamClient, nil
}

func GetIstioSecretClient() (*istiosecret.IstioCacertsSecretClient, error) {
	clientset, err := GetKubernetesClient()
	secretClient, err := factory2.GetIstioCacertsSecretClient(clientset)
	if err != nil {
		return nil, err
	}
	if err = secretClient.Register(); err != nil {
		return nil, err
	}
	return &secretClient, nil
}

func GetGlooSecretClient() (*glooV1.SecretClient, error) {
	clientset, err := GetKubernetesClient()
	secretClient, err := glooV1.NewSecretClient(&factory.KubeSecretClientFactory{
		Clientset: clientset,
	})
	if err != nil {
		return nil, err
	}
	if err = secretClient.Register(); err != nil {
		return nil, err
	}
	return &secretClient, nil
}

func GetMeshClient() (*superglooV1.MeshClient, error) {
	config, err := GetKubernetesConfig(0)
	if err != nil {
		return nil, err
	}
	meshClient, err := superglooV1.NewMeshClient(&factory.KubeResourceClientFactory{
		Crd:         superglooV1.MeshCrd,
		Cfg:         config,
		SharedCache: kube.NewKubeCache(),
	})
	if err != nil {
		return nil, err
	}
	if err = meshClient.Register(); err != nil {
		return nil, err
	}
	return &meshClient, nil
}

func GetRoutingRuleClient() (*superglooV1.RoutingRuleClient, error) {
	config, err := GetKubernetesConfig(0)
	if err != nil {
		return nil, err
	}
	rrClient, err := superglooV1.NewRoutingRuleClient(&factory.KubeResourceClientFactory{
		Crd:         superglooV1.RoutingRuleCrd,
		Cfg:         config,
		SharedCache: kube.NewKubeCache(),
	})
	if err != nil {
		return nil, err
	}
	if err = rrClient.Register(); err != nil {
		return nil, err
	}
	return &rrClient, nil
}

func GetInstallClient() (*superglooV1.InstallClient, error) {
	cfg, err := kubeutils.GetConfig("", "")
	cache := kube.NewKubeCache()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	installClient, err := superglooV1.NewInstallClient(&factory.KubeResourceClientFactory{
		Crd:         superglooV1.InstallCrd,
		Cfg:         cfg,
		SharedCache: cache,
	})
	if err != nil {
		return nil, err
	}
	if err = installClient.Register(); err != nil {
		return nil, err
	}
	return &installClient, nil
}

func GetKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := GetKubernetesConfig(0)
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return kubeClient, nil
}

func GetKubernetesConfig(timeout time.Duration) (*rest.Config, error) {
	config, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, fmt.Errorf(KubeConfigError, err)
	}
	config.Timeout = timeout
	return config, nil
}

// Return a client to query kubernetes CRDs
func GetKubeCrdClient() (*k8sApiExt.CustomResourceDefinitionInterface, error) {
	config, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, fmt.Errorf(KubeConfigError, err)
	}

	apiExtClient, err := k8sApiExt.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Error building Kubernetes client: %v \n", err)
	}
	crdClient := apiExtClient.CustomResourceDefinitions()
	return &crdClient, nil
}
