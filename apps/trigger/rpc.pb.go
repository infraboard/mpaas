// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/trigger/pb/rpc.proto

package trigger

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apps_trigger_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_trigger_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x76, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x6f, 0x0a, 0x11, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x2c, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x57,
	0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apps_trigger_pb_rpc_proto_goTypes = []interface{}{
	(*GitlabWebHookEvent)(nil), // 0: infraboard.mpaas.trigger.GitlabWebHookEvent
}
var file_apps_trigger_pb_rpc_proto_depIdxs = []int32{
	0, // 0: infraboard.mpaas.trigger.RPC.HandleGitlabEvent:input_type -> infraboard.mpaas.trigger.GitlabWebHookEvent
	0, // 1: infraboard.mpaas.trigger.RPC.HandleGitlabEvent:output_type -> infraboard.mpaas.trigger.GitlabWebHookEvent
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_trigger_pb_rpc_proto_init() }
func file_apps_trigger_pb_rpc_proto_init() {
	if File_apps_trigger_pb_rpc_proto != nil {
		return
	}
	file_apps_trigger_pb_gitlab_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_trigger_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_trigger_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_trigger_pb_rpc_proto_depIdxs,
	}.Build()
	File_apps_trigger_pb_rpc_proto = out.File
	file_apps_trigger_pb_rpc_proto_rawDesc = nil
	file_apps_trigger_pb_rpc_proto_goTypes = nil
	file_apps_trigger_pb_rpc_proto_depIdxs = nil
}
