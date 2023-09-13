// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/cluster/pb/cluster.proto

package cluster

import (
	resource "github.com/infraboard/mcube/pb/resource"
	deploy "github.com/infraboard/mpaas/apps/deploy"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 清单
	Items []*Cluster `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ClusterSet) Reset() {
	*x = ClusterSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterSet) ProtoMessage() {}

func (x *ClusterSet) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterSet.ProtoReflect.Descriptor instead.
func (*ClusterSet) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ClusterSet) GetItems() []*Cluster {
	if x != nil {
		return x.Items
	}
	return nil
}

// 部署集群
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *resource.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 元信息
	// @gotags: bson:",inline" json:"scope"
	Scope *resource.Scope `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope" bson:",inline"`
	// 创建信息
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateClusterRequest `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 关联的部署
	// @gotags: bson:"-" json:"deployments"
	Deployments *deploy.DeploymentSet `protobuf:"bytes,4,opt,name=deployments,proto3" json:"deployments" bson:"-"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Cluster) GetMeta() *resource.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Cluster) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Cluster) GetSpec() *CreateClusterRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Cluster) GetDeployments() *deploy.DeploymentSet {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务Id
	// @gotags: bson:"service_id" json:"service_id" validate:"required"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id" validate:"required"`
	// 集群名称
	// @gotags: bson:"name" json:"name" validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 集群描述
	// @gotags: bson:"describe" json:"describe"
	Describe string `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe" bson:"describe"`
	// 部署标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClusterRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CreateClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClusterRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *CreateClusterRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_mpaas_apps_cluster_pb_cluster_proto protoreflect.FileDescriptor

var file_mpaas_apps_cluster_pb_cluster_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a,
	0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x0a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x84, 0x02,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x36,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x52, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpaas_apps_cluster_pb_cluster_proto_rawDescOnce sync.Once
	file_mpaas_apps_cluster_pb_cluster_proto_rawDescData = file_mpaas_apps_cluster_pb_cluster_proto_rawDesc
)

func file_mpaas_apps_cluster_pb_cluster_proto_rawDescGZIP() []byte {
	file_mpaas_apps_cluster_pb_cluster_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_cluster_pb_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_cluster_pb_cluster_proto_rawDescData)
	})
	return file_mpaas_apps_cluster_pb_cluster_proto_rawDescData
}

var file_mpaas_apps_cluster_pb_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mpaas_apps_cluster_pb_cluster_proto_goTypes = []interface{}{
	(*ClusterSet)(nil),           // 0: infraboard.mpaas.cluster.ClusterSet
	(*Cluster)(nil),              // 1: infraboard.mpaas.cluster.Cluster
	(*CreateClusterRequest)(nil), // 2: infraboard.mpaas.cluster.CreateClusterRequest
	nil,                          // 3: infraboard.mpaas.cluster.CreateClusterRequest.LabelsEntry
	(*resource.Meta)(nil),        // 4: infraboard.mcube.resource.Meta
	(*resource.Scope)(nil),       // 5: infraboard.mcube.resource.Scope
	(*deploy.DeploymentSet)(nil), // 6: infraboard.mpaas.deploy.DeploymentSet
}
var file_mpaas_apps_cluster_pb_cluster_proto_depIdxs = []int32{
	1, // 0: infraboard.mpaas.cluster.ClusterSet.items:type_name -> infraboard.mpaas.cluster.Cluster
	4, // 1: infraboard.mpaas.cluster.Cluster.meta:type_name -> infraboard.mcube.resource.Meta
	5, // 2: infraboard.mpaas.cluster.Cluster.scope:type_name -> infraboard.mcube.resource.Scope
	2, // 3: infraboard.mpaas.cluster.Cluster.spec:type_name -> infraboard.mpaas.cluster.CreateClusterRequest
	6, // 4: infraboard.mpaas.cluster.Cluster.deployments:type_name -> infraboard.mpaas.deploy.DeploymentSet
	3, // 5: infraboard.mpaas.cluster.CreateClusterRequest.labels:type_name -> infraboard.mpaas.cluster.CreateClusterRequest.LabelsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mpaas_apps_cluster_pb_cluster_proto_init() }
func file_mpaas_apps_cluster_pb_cluster_proto_init() {
	if File_mpaas_apps_cluster_pb_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mpaas_apps_cluster_pb_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mpaas_apps_cluster_pb_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mpaas_apps_cluster_pb_cluster_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_cluster_pb_cluster_proto_depIdxs,
		MessageInfos:      file_mpaas_apps_cluster_pb_cluster_proto_msgTypes,
	}.Build()
	File_mpaas_apps_cluster_pb_cluster_proto = out.File
	file_mpaas_apps_cluster_pb_cluster_proto_rawDesc = nil
	file_mpaas_apps_cluster_pb_cluster_proto_goTypes = nil
	file_mpaas_apps_cluster_pb_cluster_proto_depIdxs = nil
}
