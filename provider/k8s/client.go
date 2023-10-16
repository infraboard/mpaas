package k8s

import (
	"os"
	"path/filepath"

	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/infraboard/mpaas/provider/k8s/admin"
	"github.com/infraboard/mpaas/provider/k8s/config"
	"github.com/infraboard/mpaas/provider/k8s/event"
	"github.com/infraboard/mpaas/provider/k8s/network"
	"github.com/infraboard/mpaas/provider/k8s/storage"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

func NewClientFromFile(kubeConfPath string) (*Client, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	kc, err := os.ReadFile(filepath.Join(wd, kubeConfPath))
	if err != nil {
		return nil, err
	}

	return NewClient(string(kc))
}

func NewClient(kubeConfigYaml string) (*Client, error) {
	// 加载kubeconfig配置
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
		log:      logger.Sub("provider.k8s"),
	}, nil
}

type Client struct {
	kubeconf *clientcmdapi.Config
	restconf *rest.Config
	client   *kubernetes.Clientset
	log      *zerolog.Logger
}

func (c *Client) ServerVersion() (string, error) {
	si, err := c.client.ServerVersion()
	if err != nil {
		return "", err
	}

	return si.String(), nil
}

func (c *Client) ServerResources() ([]*metav1.APIResourceList, error) {
	return c.client.ServerPreferredResources()
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

// 应用负载
func (c *Client) WorkLoad() *workload.Client {
	return workload.NewWorkload(c.client, c.restconf)
}

// 应用配置
func (c *Client) Config() *config.Client {
	return config.NewConfig(c.client)
}

// 应用存储
func (c *Client) Storage() *storage.Client {
	return storage.NewStorage(c.client)
}

// 应用网络
func (c *Client) Network() *network.Client {
	return network.NewNetwork(c.client)
}

// 应用事件
func (c *Client) Event() *event.Client {
	return event.NewEvent(c.client)
}

// 集群管理
func (c *Client) Admin() *admin.Client {
	return admin.NewAdmin(c.client)
}
