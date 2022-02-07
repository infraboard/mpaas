package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewClient(kubeConfigYaml string) error {
	kc, err := clientcmd.BuildConfigFromKubeconfigGetter("",
		func() (*clientcmdapi.Config, error) {
			return clientcmd.Load([]byte(kubeConfigYaml))
		},
	)
	if err != nil {
		return err
	}

	// 初始化客户端
	kubernetes.NewForConfig(kc)

	return nil
}
