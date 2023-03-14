// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.6
// source: apps/deploy/pb/rpc.proto

package deploy

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
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

type DESCRIBE_BY int32

const (
	// 用户的ID
	DESCRIBE_BY_ID DESCRIBE_BY = 0
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "ID",
	}
	DESCRIBE_BY_value = map[string]int32{
		"ID": 0,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_deploy_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_apps_deploy_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_deploy_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type UpdateDeploymentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部署Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 更新Token, 校验合法性
	// @gotags: json:"update_token"
	UpdateToken string `protobuf:"bytes,2,opt,name=update_token,json=updateToken,proto3" json:"update_token"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// k8s相关配置更新, 当部署时k8s部署是有效
	// @gotags: json:"updated_k8s_config"
	UpdatedK8SConfig *K8STypeConfig `protobuf:"bytes,4,opt,name=updated_k8s_config,json=updatedK8sConfig,proto3" json:"updated_k8s_config"`
}

func (x *UpdateDeploymentStatusRequest) Reset() {
	*x = UpdateDeploymentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeploymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeploymentStatusRequest) ProtoMessage() {}

func (x *UpdateDeploymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeploymentStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDeploymentStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDeploymentStatusRequest) GetUpdateToken() string {
	if x != nil {
		return x.UpdateToken
	}
	return ""
}

func (x *UpdateDeploymentStatusRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateDeploymentStatusRequest) GetUpdatedK8SConfig() *K8STypeConfig {
	if x != nil {
		return x.UpdatedK8SConfig
	}
	return nil
}

type QueryDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 部署Id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids"`
}

func (x *QueryDeploymentRequest) Reset() {
	*x = QueryDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDeploymentRequest) ProtoMessage() {}

func (x *QueryDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDeploymentRequest.ProtoReflect.Descriptor instead.
func (*QueryDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *QueryDeploymentRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryDeploymentRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DescribeDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=infraboard.mpaas.deploy.DESCRIBE_BY" json:"describe_by"`
	// 查询值
	// @gotags: json:"describe_value"  validate:"required"
	DescribeValue string `protobuf:"bytes,2,opt,name=describe_value,json=describeValue,proto3" json:"describe_value" validate:"required"`
}

func (x *DescribeDeploymentRequest) Reset() {
	*x = DescribeDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDeploymentRequest) ProtoMessage() {}

func (x *DescribeDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDeploymentRequest.ProtoReflect.Descriptor instead.
func (*DescribeDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeDeploymentRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_ID
}

func (x *DescribeDeploymentRequest) GetDescribeValue() string {
	if x != nil {
		return x.DescribeValue
	}
	return ""
}

type UpdateDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 部署Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// 创建信息
	// @gotags: json:"spec"
	Spec *CreateDeploymentRequest `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec"`
}

func (x *UpdateDeploymentRequest) Reset() {
	*x = UpdateDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeploymentRequest) ProtoMessage() {}

func (x *UpdateDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeploymentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDeploymentRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateDeploymentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDeploymentRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateDeploymentRequest) GetSpec() *CreateDeploymentRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部署Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteDeploymentRequest) Reset() {
	*x = DeleteDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeploymentRequest) ProtoMessage() {}

func (x *DeleteDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeploymentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDeploymentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_deploy_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_deploy_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x54, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x4b,
	0x38, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x4b, 0x38, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x62,
	0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x0a, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd3,
	0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x44,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a,
	0x15, 0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x06,
	0x0a, 0x02, 0x49, 0x44, 0x10, 0x00, 0x32, 0xd7, 0x02, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x6a,
	0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x12, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x75, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_deploy_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_deploy_pb_rpc_proto_rawDescData = file_apps_deploy_pb_rpc_proto_rawDesc
)

func file_apps_deploy_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_deploy_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_deploy_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_deploy_pb_rpc_proto_rawDescData)
	})
	return file_apps_deploy_pb_rpc_proto_rawDescData
}

var file_apps_deploy_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_deploy_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_deploy_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),                      // 0: infraboard.mpaas.deploy.DESCRIBE_BY
	(*UpdateDeploymentStatusRequest)(nil), // 1: infraboard.mpaas.deploy.UpdateDeploymentStatusRequest
	(*QueryDeploymentRequest)(nil),        // 2: infraboard.mpaas.deploy.QueryDeploymentRequest
	(*DescribeDeploymentRequest)(nil),     // 3: infraboard.mpaas.deploy.DescribeDeploymentRequest
	(*UpdateDeploymentRequest)(nil),       // 4: infraboard.mpaas.deploy.UpdateDeploymentRequest
	(*DeleteDeploymentRequest)(nil),       // 5: infraboard.mpaas.deploy.DeleteDeploymentRequest
	(*K8STypeConfig)(nil),                 // 6: infraboard.mpaas.deploy.K8sTypeConfig
	(*request.PageRequest)(nil),           // 7: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),              // 8: infraboard.mcube.request.UpdateMode
	(*CreateDeploymentRequest)(nil),       // 9: infraboard.mpaas.deploy.CreateDeploymentRequest
	(*DeploymentSet)(nil),                 // 10: infraboard.mpaas.deploy.DeploymentSet
	(*Deployment)(nil),                    // 11: infraboard.mpaas.deploy.Deployment
}
var file_apps_deploy_pb_rpc_proto_depIdxs = []int32{
	6,  // 0: infraboard.mpaas.deploy.UpdateDeploymentStatusRequest.updated_k8s_config:type_name -> infraboard.mpaas.deploy.K8sTypeConfig
	7,  // 1: infraboard.mpaas.deploy.QueryDeploymentRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0,  // 2: infraboard.mpaas.deploy.DescribeDeploymentRequest.describe_by:type_name -> infraboard.mpaas.deploy.DESCRIBE_BY
	8,  // 3: infraboard.mpaas.deploy.UpdateDeploymentRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	9,  // 4: infraboard.mpaas.deploy.UpdateDeploymentRequest.spec:type_name -> infraboard.mpaas.deploy.CreateDeploymentRequest
	2,  // 5: infraboard.mpaas.deploy.RPC.QueryDeployment:input_type -> infraboard.mpaas.deploy.QueryDeploymentRequest
	3,  // 6: infraboard.mpaas.deploy.RPC.DescribeDeployment:input_type -> infraboard.mpaas.deploy.DescribeDeploymentRequest
	1,  // 7: infraboard.mpaas.deploy.RPC.UpdateDeploymentStatus:input_type -> infraboard.mpaas.deploy.UpdateDeploymentStatusRequest
	10, // 8: infraboard.mpaas.deploy.RPC.QueryDeployment:output_type -> infraboard.mpaas.deploy.DeploymentSet
	11, // 9: infraboard.mpaas.deploy.RPC.DescribeDeployment:output_type -> infraboard.mpaas.deploy.Deployment
	11, // 10: infraboard.mpaas.deploy.RPC.UpdateDeploymentStatus:output_type -> infraboard.mpaas.deploy.Deployment
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apps_deploy_pb_rpc_proto_init() }
func file_apps_deploy_pb_rpc_proto_init() {
	if File_apps_deploy_pb_rpc_proto != nil {
		return
	}
	file_apps_deploy_pb_deploy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_deploy_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeploymentStatusRequest); i {
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
		file_apps_deploy_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDeploymentRequest); i {
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
		file_apps_deploy_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDeploymentRequest); i {
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
		file_apps_deploy_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeploymentRequest); i {
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
		file_apps_deploy_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeploymentRequest); i {
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
			RawDescriptor: file_apps_deploy_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_deploy_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_deploy_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_deploy_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_deploy_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_deploy_pb_rpc_proto = out.File
	file_apps_deploy_pb_rpc_proto_rawDesc = nil
	file_apps_deploy_pb_rpc_proto_goTypes = nil
	file_apps_deploy_pb_rpc_proto_depIdxs = nil
}
