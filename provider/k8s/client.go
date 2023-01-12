package k8s

import (
	"os"
	"path/filepath"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/infraboard/mpaas/provider/k8s/batch"
	"github.com/infraboard/mpaas/provider/k8s/config"
	"github.com/infraboard/mpaas/provider/k8s/event"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

var (
	DEFAULT_NAMESPACE = "default"
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
		kubeconf:   kubeConf,
		restconf:   restConf,
		client:     client,
		log:        zap.L().Named("provider.k8s"),
		appVersion: "apps/v1",
	}, nil
}

type Client struct {
	kubeconf   *clientcmdapi.Config
	restconf   *rest.Config
	client     *kubernetes.Clientset
	log        logger.Logger
	appVersion string
}

func (c *Client) ServerVersion() (string, error) {
	si, err := c.client.ServerVersion()
	if err != nil {
		return "", err
	}

	return si.String(), nil
}

func (c *Client) ServerResources() ([]*metav1.APIResourceList, error) {
	return c.client.ServerResources()
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

func (c *Client) WorkLoad() *workload.Workload {
	return workload.NewWorkload(c.client)
}

func (c *Client) Config() *config.Config {
	return config.NewConfig(c.client)
}

func (c *Client) Batch() *batch.Batch {
	return batch.NewBatch(c.client)
}

func (c *Client) Event() *event.Event {
	return event.NewEvent(c.client)
}
