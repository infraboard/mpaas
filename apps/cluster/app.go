package cluster

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"

	"github.com/infraboard/mpaas/conf"
)

const (
	AppName = "cluster"
)

var (
	validate = validator.New()
)

func NewCreateClusterRequest() *CreateClusterRequest {
	return &CreateClusterRequest{}
}

func NewCluster(req *CreateClusterRequest) (*Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Cluster{
		Id:         xid.New().String(),
		CreateAt:   time.Now().UnixMicro(),
		Data:       req,
		ServerInfo: &ServerInfo{},
		Status:     &Status{},
	}, nil
}

func (req *CreateClusterRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateClusterRequest) UpdateOwner(tk *token.Token) {
	req.CreateBy = tk.Account
	req.Domain = tk.Domain
	req.Namespace = tk.NamespaceName
}

func NewClusterSet() *ClusterSet {
	return &ClusterSet{
		Items: []*Cluster{},
	}
}

func (s *ClusterSet) Add(item *Cluster) {
	s.Items = append(s.Items, item)
}

func (s *ClusterSet) Desense() {
	for i := range s.Items {
		s.Items[i].Desense()
	}
}

func NewDefaultCluster() *Cluster {
	return &Cluster{
		Data: &CreateClusterRequest{},
	}
}

func (i *Cluster) IsAlive() error {
	if i.Status == nil {
		return fmt.Errorf("status is nil")
	}

	if !i.Status.IsAlive {
		return fmt.Errorf(i.Status.Message)
	}

	return nil
}

func (i *Cluster) Update(req *UpdateClusterRequest) {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	i.Data = req.Data
}

func (i *Cluster) Patch(req *UpdateClusterRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Data, req.Data)
}

func (i *Cluster) EncryptKubeConf(key string) error {
	// 判断文本是否已经加密
	if strings.HasPrefix(i.Data.KubeConfig, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text has ciphered")
	}

	cipherText, err := cbc.Encrypt([]byte(i.Data.KubeConfig), []byte(key))
	if err != nil {
		return err
	}

	base64Str := base64.StdEncoding.EncodeToString(cipherText)
	i.Data.KubeConfig = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, base64Str)
	return nil
}

func (i *Cluster) DecryptKubeConf(key string) error {
	// 判断文本是否已经是明文
	if !strings.HasPrefix(i.Data.KubeConfig, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text is plan text")
	}

	base64CipherText := strings.TrimPrefix(i.Data.KubeConfig, conf.CIPHER_TEXT_PREFIX)

	cipherText, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return err
	}

	planText, err := cbc.Decrypt([]byte(cipherText), []byte(key))
	if err != nil {
		return err
	}

	i.Data.KubeConfig = string(planText)
	return nil
}

func (i *Cluster) Desense() {
	if i.Data.KubeConfig != "" {
		i.Data.KubeConfig = "****"
	}
}

func NewDescribeClusterRequest(id string) *DescribeClusterRequest {
	return &DescribeClusterRequest{
		Id: id,
	}
}

func NewQueryClusterRequest() *QueryClusterRequest {
	return &QueryClusterRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewQueryClusterRequestFromHTTP(r *http.Request) *QueryClusterRequest {
	qs := r.URL.Query()

	return &QueryClusterRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewPutClusterRequest(id string) *UpdateClusterRequest {
	return &UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateClusterRequest(),
	}
}

func NewPatchClusterRequest(id string) *UpdateClusterRequest {
	return &UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateClusterRequest(),
	}
}

func NewDeleteClusterRequestWithID(id string) *DeleteClusterRequest {
	return &DeleteClusterRequest{
		Id: id,
	}
}
