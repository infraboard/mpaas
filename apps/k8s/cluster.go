package k8s

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
	v1 "k8s.io/api/core/v1"

	"github.com/infraboard/mpaas/conf"
	"github.com/infraboard/mpaas/provider/k8s"
	"github.com/infraboard/mpaas/provider/k8s/workload"
)

var (
	validate = validator.New()
)

func NewCreateClusterRequest() *CreateClusterRequest {
	return &CreateClusterRequest{
		Domain:    "default",
		Namespace: "default",
	}
}

func NewCluster(req *CreateClusterRequest) (*Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Cluster{
		Meta:   NewMeta(),
		Spec:   req,
		Status: &Status{},
	}, nil
}

func NewMeta() *Meta {
	return &Meta{
		Id:         xid.New().String(),
		CreateAt:   time.Now().UnixMicro(),
		ServerInfo: &ServerInfo{},
	}
}

func (req *CreateClusterRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateClusterRequest) UpdateOwner() {
	req.CreateBy = "default"
	req.Domain = "default"
	req.Namespace = "default"
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

func (s *ClusterSet) DecryptKubeConf(key string) error {
	errs := []string{}
	for i := range s.Items {
		err := s.Items[i].DecryptKubeConf(key)
		if err != nil {
			errs = append(errs, fmt.Sprintf(
				"decrypt %s kubeconf error, %s",
				s.Items[i].Spec.Name,
				err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, ","))
	}

	return nil
}

func NewDefaultCluster() *Cluster {
	return &Cluster{
		Spec: &CreateClusterRequest{},
	}
}

func (i *Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*Meta
		*CreateClusterRequest
		*Status
	}{i.Meta, i.Spec, i.Status})
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
	m := i.Meta
	m.UpdateAt = time.Now().Unix()
	m.UpdateBy = req.UpdateBy
	i.Spec = req.Spec
}

func (i *Cluster) Patch(req *UpdateClusterRequest) error {
	m := i.Meta
	m.UpdateAt = time.Now().Unix()
	m.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Spec, req.Spec)
}

func (i *Cluster) EncryptKubeConf(key string) error {
	// 判断文本是否已经加密
	if strings.HasPrefix(i.Spec.KubeConfig, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text has ciphered")
	}

	cipherText, err := cbc.Encrypt([]byte(i.Spec.KubeConfig), []byte(key))
	if err != nil {
		return err
	}

	base64Str := base64.StdEncoding.EncodeToString(cipherText)
	i.Spec.KubeConfig = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, base64Str)
	return nil
}

func (i *Cluster) DecryptKubeConf(key string) error {
	// 判断文本是否已经是明文
	if !strings.HasPrefix(i.Spec.KubeConfig, conf.CIPHER_TEXT_PREFIX) {
		return nil
	}

	base64CipherText := strings.TrimPrefix(i.Spec.KubeConfig, conf.CIPHER_TEXT_PREFIX)

	cipherText, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return err
	}

	planText, err := cbc.Decrypt([]byte(cipherText), []byte(key))
	if err != nil {
		return err
	}

	i.Spec.KubeConfig = string(planText)
	return nil
}

func (i *Cluster) Desense() {
	if i.Spec.KubeConfig != "" {
		i.Spec.KubeConfig = "****"
	}
}

func (i *Cluster) Client() (*k8s.Client, error) {
	return k8s.NewClient(i.Spec.KubeConfig)
}

func (i *Cluster) KubeConfSecret(mountPath string) *v1.Secret {
	secret := new(v1.Secret)
	secret.Name = fmt.Sprintf("cluster-%s", i.Meta.Id)
	secret.StringData = map[string]string{
		"config": i.Spec.KubeConfig,
	}
	secret.Annotations = map[string]string{
		workload.ANNOTATION_SECRET_MOUNT: mountPath,
	}
	return secret
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
		Vendor:   qs.Get("vendor"),
		Region:   qs.Get("region"),
	}
}

func (req *QueryClusterRequest) UpdateNamespace() {
	req.Domain = "default"
	req.Namespace = "default"
}

func NewPutClusterRequest(id string) *UpdateClusterRequest {
	return &UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Spec:       NewCreateClusterRequest(),
	}
}

func NewPatchClusterRequest(id string) *UpdateClusterRequest {
	return &UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Spec:       NewCreateClusterRequest(),
	}
}

func NewDeleteClusterRequestWithID(id string) *DeleteClusterRequest {
	return &DeleteClusterRequest{
		Id: id,
	}
}
