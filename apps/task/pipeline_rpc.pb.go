// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/task/pb/pipeline_rpc.proto

package task

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

type DeletePipelineTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeletePipelineTaskRequest) Reset() {
	*x = DeletePipelineTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePipelineTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePipelineTaskRequest) ProtoMessage() {}

func (x *DeletePipelineTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePipelineTaskRequest.ProtoReflect.Descriptor instead.
func (*DeletePipelineTaskRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DeletePipelineTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RunPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *RunPipelineRequest) Reset() {
	*x = RunPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunPipelineRequest) ProtoMessage() {}

func (x *RunPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunPipelineRequest.ProtoReflect.Descriptor instead.
func (*RunPipelineRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RunPipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QueryPipelineTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 任务Id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids"`
}

func (x *QueryPipelineTaskRequest) Reset() {
	*x = QueryPipelineTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPipelineTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPipelineTaskRequest) ProtoMessage() {}

func (x *QueryPipelineTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPipelineTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryPipelineTaskRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPipelineTaskRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPipelineTaskRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DescribePipelineTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePipelineTaskRequest) Reset() {
	*x = DescribePipelineTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePipelineTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePipelineTaskRequest) ProtoMessage() {}

func (x *DescribePipelineTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePipelineTaskRequest.ProtoReflect.Descriptor instead.
func (*DescribePipelineTaskRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePipelineTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_task_pb_pipeline_rpc_proto protoreflect.FileDescriptor

var file_apps_task_pb_pipeline_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x20, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a,
	0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xb8, 0x03, 0x0a, 0x0b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x50, 0x43, 0x12, 0x5d, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x6c, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12,
	0x6f, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x6b, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_task_pb_pipeline_rpc_proto_rawDescOnce sync.Once
	file_apps_task_pb_pipeline_rpc_proto_rawDescData = file_apps_task_pb_pipeline_rpc_proto_rawDesc
)

func file_apps_task_pb_pipeline_rpc_proto_rawDescGZIP() []byte {
	file_apps_task_pb_pipeline_rpc_proto_rawDescOnce.Do(func() {
		file_apps_task_pb_pipeline_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_task_pb_pipeline_rpc_proto_rawDescData)
	})
	return file_apps_task_pb_pipeline_rpc_proto_rawDescData
}

var file_apps_task_pb_pipeline_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_task_pb_pipeline_rpc_proto_goTypes = []interface{}{
	(*DeletePipelineTaskRequest)(nil),   // 0: infraboard.mpaas.task.DeletePipelineTaskRequest
	(*RunPipelineRequest)(nil),          // 1: infraboard.mpaas.task.RunPipelineRequest
	(*QueryPipelineTaskRequest)(nil),    // 2: infraboard.mpaas.task.QueryPipelineTaskRequest
	(*DescribePipelineTaskRequest)(nil), // 3: infraboard.mpaas.task.DescribePipelineTaskRequest
	(*request.PageRequest)(nil),         // 4: infraboard.mcube.page.PageRequest
	(*PipelineTask)(nil),                // 5: infraboard.mpaas.task.PipelineTask
	(*PipelineTaskSet)(nil),             // 6: infraboard.mpaas.task.PipelineTaskSet
}
var file_apps_task_pb_pipeline_rpc_proto_depIdxs = []int32{
	4, // 0: infraboard.mpaas.task.QueryPipelineTaskRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // 1: infraboard.mpaas.task.PipelineRPC.RunPipeline:input_type -> infraboard.mpaas.task.RunPipelineRequest
	2, // 2: infraboard.mpaas.task.PipelineRPC.QueryPipelineTask:input_type -> infraboard.mpaas.task.QueryPipelineTaskRequest
	3, // 3: infraboard.mpaas.task.PipelineRPC.DescribePipelineTask:input_type -> infraboard.mpaas.task.DescribePipelineTaskRequest
	0, // 4: infraboard.mpaas.task.PipelineRPC.DeletePipelineTask:input_type -> infraboard.mpaas.task.DeletePipelineTaskRequest
	5, // 5: infraboard.mpaas.task.PipelineRPC.RunPipeline:output_type -> infraboard.mpaas.task.PipelineTask
	6, // 6: infraboard.mpaas.task.PipelineRPC.QueryPipelineTask:output_type -> infraboard.mpaas.task.PipelineTaskSet
	5, // 7: infraboard.mpaas.task.PipelineRPC.DescribePipelineTask:output_type -> infraboard.mpaas.task.PipelineTask
	5, // 8: infraboard.mpaas.task.PipelineRPC.DeletePipelineTask:output_type -> infraboard.mpaas.task.PipelineTask
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_task_pb_pipeline_rpc_proto_init() }
func file_apps_task_pb_pipeline_rpc_proto_init() {
	if File_apps_task_pb_pipeline_rpc_proto != nil {
		return
	}
	file_apps_task_pb_pipeline_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_task_pb_pipeline_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePipelineTaskRequest); i {
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
		file_apps_task_pb_pipeline_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPipelineRequest); i {
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
		file_apps_task_pb_pipeline_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPipelineTaskRequest); i {
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
		file_apps_task_pb_pipeline_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePipelineTaskRequest); i {
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
			RawDescriptor: file_apps_task_pb_pipeline_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_task_pb_pipeline_rpc_proto_goTypes,
		DependencyIndexes: file_apps_task_pb_pipeline_rpc_proto_depIdxs,
		MessageInfos:      file_apps_task_pb_pipeline_rpc_proto_msgTypes,
	}.Build()
	File_apps_task_pb_pipeline_rpc_proto = out.File
	file_apps_task_pb_pipeline_rpc_proto_rawDesc = nil
	file_apps_task_pb_pipeline_rpc_proto_goTypes = nil
	file_apps_task_pb_pipeline_rpc_proto_depIdxs = nil
}
