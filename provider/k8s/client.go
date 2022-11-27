package k8s

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	DEFAULT_NAMESPACE = "default"
)

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

	// 基于Interface封装的客户端
	// client.AppsV1().Deployments("default").Create(nil, nil, metav1.CreateOptions{})
	// RESTclient https://github.com/jindezgm/k8s-src-analysis/blob/master/client-go/rest/Client.md
	// client.RESTClient().Post().Namespace("ns").Body("body").Do(nil).Into(resp).Error()

	return &Client{
		kubeconf: kubeConf,
		restconf: restConf,
		client:   client,
		log:      zap.L().Named("provider.k8s"),
	}, nil
}

type Client struct {
	kubeconf *clientcmdapi.Config
	restconf *rest.Config
	client   *kubernetes.Clientset
	log      logger.Logger
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

func NewGetRequest(name string) *GetRequest {
	return &GetRequest{
		Namespace: DEFAULT_NAMESPACE,
		Name:      name,
	}
}

type GetRequest struct {
	Namespace string
	Name      string
	Opts      metav1.GetOptions
}

func NewDeleteRequest(name string) *DeleteRequest {
	return &DeleteRequest{
		Namespace: DEFAULT_NAMESPACE,
		Name:      name,
	}
}

type DeleteRequest struct {
	Namespace string
	Name      string
	Opts      metav1.DeleteOptions
}

func NewListRequest() *ListRequest {
	return &ListRequest{}
}

type ListRequest struct {
	Namespace         string
	SkipManagedFields bool
	Opts              metav1.ListOptions
}

func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}

type CreateRequest struct {
	Namespace string
	Opts      metav1.CreateOptions
}
