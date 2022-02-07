package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewClient(kubeConfigYaml string) (*Client, error) {
	kc, err := clientcmd.BuildConfigFromKubeconfigGetter("",
		func() (*clientcmdapi.Config, error) {
			return clientcmd.Load([]byte(kubeConfigYaml))
		},
	)
	if err != nil {
		return nil, err
	}

	// 初始化客户端
	client, err := kubernetes.NewForConfig(kc)
	if err != nil {
		return nil, err
	}

	return &Client{c: client}, nil
}

type Client struct {
	c *kubernetes.Clientset
}

func (c *Client) ServerVersion() (string, error) {
	si, err := c.c.ServerVersion()
	if err != nil {
		return "", err
	}
	return si.String(), nil
}
