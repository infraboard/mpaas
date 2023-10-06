// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/audit/pb/rpc.proto

package audit

import (
	request "github.com/infraboard/mcube/http/request"
	resource "github.com/infraboard/mcube/pb/resource"
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

type QueryRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// 资源标签过滤
	// @gotags: json:"filters"
	Filters []*resource.LabelRequirement `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters"`
	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page"`
	// 集群 Id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,4,rep,name=ids,proto3" json:"ids"`
	// 集群 标签
	// @gotags: json:"label"
	Label map[string]string `protobuf:"bytes,15,rep,name=label,proto3" json:"label" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QueryRecordRequest) Reset() {
	*x = QueryRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_audit_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRecordRequest) ProtoMessage() {}

func (x *QueryRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_audit_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryRecordRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_audit_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRecordRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *QueryRecordRequest) GetFilters() []*resource.LabelRequirement {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *QueryRecordRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRecordRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *QueryRecordRequest) GetLabel() map[string]string {
	if x != nil {
		return x.Label
	}
	return nil
}

var File_mpaas_apps_audit_pb_rpc_proto protoreflect.FileDescriptor

var file_mpaas_apps_audit_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x1a, 0x1f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe4, 0x02, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x45,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x4b, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x38, 0x0a, 0x0a,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xbc, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x57,
	0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x5c, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x53, 0x65, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpaas_apps_audit_pb_rpc_proto_rawDescOnce sync.Once
	file_mpaas_apps_audit_pb_rpc_proto_rawDescData = file_mpaas_apps_audit_pb_rpc_proto_rawDesc
)

func file_mpaas_apps_audit_pb_rpc_proto_rawDescGZIP() []byte {
	file_mpaas_apps_audit_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_audit_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_audit_pb_rpc_proto_rawDescData)
	})
	return file_mpaas_apps_audit_pb_rpc_proto_rawDescData
}

var file_mpaas_apps_audit_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mpaas_apps_audit_pb_rpc_proto_goTypes = []interface{}{
	(*QueryRecordRequest)(nil),        // 0: infraboard.mpaas.audit.QueryRecordRequest
	nil,                               // 1: infraboard.mpaas.audit.QueryRecordRequest.LabelEntry
	(*resource.Scope)(nil),            // 2: infraboard.mcube.resource.Scope
	(*resource.LabelRequirement)(nil), // 3: infraboard.mcube.resource.LabelRequirement
	(*request.PageRequest)(nil),       // 4: infraboard.mcube.page.PageRequest
	(*SaveRecordRequest)(nil),         // 5: infraboard.mpaas.audit.SaveRecordRequest
	(*Record)(nil),                    // 6: infraboard.mpaas.audit.Record
	(*RecordSet)(nil),                 // 7: infraboard.mpaas.audit.RecordSet
}
var file_mpaas_apps_audit_pb_rpc_proto_depIdxs = []int32{
	2, // 0: infraboard.mpaas.audit.QueryRecordRequest.scope:type_name -> infraboard.mcube.resource.Scope
	3, // 1: infraboard.mpaas.audit.QueryRecordRequest.filters:type_name -> infraboard.mcube.resource.LabelRequirement
	4, // 2: infraboard.mpaas.audit.QueryRecordRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // 3: infraboard.mpaas.audit.QueryRecordRequest.label:type_name -> infraboard.mpaas.audit.QueryRecordRequest.LabelEntry
	5, // 4: infraboard.mpaas.audit.RPC.SaveRecord:input_type -> infraboard.mpaas.audit.SaveRecordRequest
	0, // 5: infraboard.mpaas.audit.RPC.QueryRecord:input_type -> infraboard.mpaas.audit.QueryRecordRequest
	6, // 6: infraboard.mpaas.audit.RPC.SaveRecord:output_type -> infraboard.mpaas.audit.Record
	7, // 7: infraboard.mpaas.audit.RPC.QueryRecord:output_type -> infraboard.mpaas.audit.RecordSet
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mpaas_apps_audit_pb_rpc_proto_init() }
func file_mpaas_apps_audit_pb_rpc_proto_init() {
	if File_mpaas_apps_audit_pb_rpc_proto != nil {
		return
	}
	file_mpaas_apps_audit_pb_audit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_audit_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRecordRequest); i {
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
			RawDescriptor: file_mpaas_apps_audit_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpaas_apps_audit_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_audit_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mpaas_apps_audit_pb_rpc_proto_msgTypes,
	}.Build()
	File_mpaas_apps_audit_pb_rpc_proto = out.File
	file_mpaas_apps_audit_pb_rpc_proto_rawDesc = nil
	file_mpaas_apps_audit_pb_rpc_proto_goTypes = nil
	file_mpaas_apps_audit_pb_rpc_proto_depIdxs = nil
}
