// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/traffic/pb/rpc.proto

package traffic

import (
	request "github.com/infraboard/mcube/http/request"
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

type QueryRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryRuleRequest) Reset() {
	*x = QueryRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_traffic_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRuleRequest) ProtoMessage() {}

func (x *QueryRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_traffic_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRuleRequest.ProtoReflect.Descriptor instead.
func (*QueryRuleRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_traffic_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRuleRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type DescribeRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeRuleRequest) Reset() {
	*x = DescribeRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_traffic_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRuleRequest) ProtoMessage() {}

func (x *DescribeRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_traffic_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRuleRequest.ProtoReflect.Descriptor instead.
func (*DescribeRuleRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_traffic_pb_rpc_proto_rawDescGZIP(), []int{1}
}

var File_mpaas_apps_traffic_pb_rpc_proto protoreflect.FileDescriptor

var file_mpaas_apps_traffic_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x1a, 0x20, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc0, 0x01, 0x0a, 0x03, 0x52,
	0x50, 0x43, 0x12, 0x5a, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x5d,
	0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2d,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mpaas_apps_traffic_pb_rpc_proto_rawDescOnce sync.Once
	file_mpaas_apps_traffic_pb_rpc_proto_rawDescData = file_mpaas_apps_traffic_pb_rpc_proto_rawDesc
)

func file_mpaas_apps_traffic_pb_rpc_proto_rawDescGZIP() []byte {
	file_mpaas_apps_traffic_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_traffic_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_traffic_pb_rpc_proto_rawDescData)
	})
	return file_mpaas_apps_traffic_pb_rpc_proto_rawDescData
}

var file_mpaas_apps_traffic_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mpaas_apps_traffic_pb_rpc_proto_goTypes = []interface{}{
	(*QueryRuleRequest)(nil),    // 0: infraboard.mpaas.traffic.QueryRuleRequest
	(*DescribeRuleRequest)(nil), // 1: infraboard.mpaas.traffic.DescribeRuleRequest
	(*request.PageRequest)(nil), // 2: infraboard.mcube.page.PageRequest
	(*RuleSet)(nil),             // 3: infraboard.mpaas.traffic.RuleSet
	(*Rule)(nil),                // 4: infraboard.mpaas.traffic.Rule
}
var file_mpaas_apps_traffic_pb_rpc_proto_depIdxs = []int32{
	2, // 0: infraboard.mpaas.traffic.QueryRuleRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 1: infraboard.mpaas.traffic.RPC.QueryRule:input_type -> infraboard.mpaas.traffic.QueryRuleRequest
	1, // 2: infraboard.mpaas.traffic.RPC.DescribeRule:input_type -> infraboard.mpaas.traffic.DescribeRuleRequest
	3, // 3: infraboard.mpaas.traffic.RPC.QueryRule:output_type -> infraboard.mpaas.traffic.RuleSet
	4, // 4: infraboard.mpaas.traffic.RPC.DescribeRule:output_type -> infraboard.mpaas.traffic.Rule
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mpaas_apps_traffic_pb_rpc_proto_init() }
func file_mpaas_apps_traffic_pb_rpc_proto_init() {
	if File_mpaas_apps_traffic_pb_rpc_proto != nil {
		return
	}
	file_mpaas_apps_traffic_pb_rule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_traffic_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRuleRequest); i {
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
		file_mpaas_apps_traffic_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRuleRequest); i {
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
			RawDescriptor: file_mpaas_apps_traffic_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpaas_apps_traffic_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_traffic_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mpaas_apps_traffic_pb_rpc_proto_msgTypes,
	}.Build()
	File_mpaas_apps_traffic_pb_rpc_proto = out.File
	file_mpaas_apps_traffic_pb_rpc_proto_rawDesc = nil
	file_mpaas_apps_traffic_pb_rpc_proto_goTypes = nil
	file_mpaas_apps_traffic_pb_rpc_proto_depIdxs = nil
}