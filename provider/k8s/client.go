package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewClient(kubeConfigYaml string) (*Client, error) {
	// 加载配置
	kubeConf, err := clientcmd.Load([]byte(kubeConfigYaml))
	if err != nil {
		return nil, err
	}

	// 构造Restclient Config
	restConf, err := clientcmd.BuildConfigFromKubeconfigGetter("",
		func() (*clientcmdapi.Config, error) {
			return kubeConf, nil
		},
	)
	if err != nil {
		return nil, err
	}

	// 初始化客户端
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}

	return &Client{
		kubeconf: kubeConf,
		restconf: restConf,
		client:   client,
	}, nil
}

type Client struct {
	kubeconf *clientcmdapi.Config
	restconf *rest.Config
	client   *kubernetes.Clientset
}

func (c *Client) ServerVersion() (string, error) {
	si, err := c.client.ServerVersion()
	if err != nil {
		return "", err
	}

	return si.String(), nil
}

func (c *Client) GetContexts() map[string]*clientcmdapi.Context {
	return c.kubeconf.Contexts
}

func (c *Client) CurrentContext() *clientcmdapi.Context {
	return c.kubeconf.Contexts[c.kubeconf.CurrentContext]
}

func (c *Client) CurrentCluster() *clientcmdapi.Cluster {
	ctx := c.CurrentContext()
	if ctx == nil {
		return nil
	}

	return c.kubeconf.Clusters[ctx.Cluster]
}

type GetRequest struct {
	Namespace string
	Name      string
	Opts      metav1.GetOptions
}

type DeleteRequest struct {
	Namespace string
	Name      string
	Opts      metav1.DeleteOptions
}
