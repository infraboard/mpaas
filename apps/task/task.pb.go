// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/task/pb/task.proto

package task

import (
	job "github.com/infraboard/mpaas/apps/job"
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

type TaskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 清单
	// @gotags: bson:"items" json:"items"
	Items []*Task `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *TaskSet) Reset() {
	*x = TaskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSet) ProtoMessage() {}

func (x *TaskSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSet.ProtoReflect.Descriptor instead.
func (*TaskSet) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskSet) GetItems() []*Task {
	if x != nil {
		return x.Items
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task定义
	// @gotags: bson:"spec" json:"spec"
	Spec *RunJobRequest `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec" bson:"spec"`
	// 关联Job
	// @gotags: bson:"job" json:"job"
	Job *job.Job `protobuf:"bytes,2,opt,name=job,proto3" json:"job" bson:"job"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetSpec() *RunJobRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Task) GetJob() *job.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task执行的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// task执行的空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// job名称: name
	// @gotags: bson:"job" json:"job"
	Job string `protobuf:"bytes,3,opt,name=job,proto3" json:"job" bson:"job"`
	// job版本
	// @gotags: bson:"version" json:"version"
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version" bson:"version"`
	// job版本对应的参数
	// @gotags: bson:"with" json:"with"
	With map[string]string `protobuf:"bytes,5,rep,name=with,proto3" json:"with" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"with"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{2}
}

func (x *RunJobRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RunJobRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RunJobRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *RunJobRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RunJobRequest) GetWith() map[string]string {
	if x != nil {
		return x.With
	}
	return nil
}

type K8SJobRunParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用于执行k8s job的集群
	// @gotags: bson:"cluster_id" json:"cluster_id"
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id" bson:"cluster_id"`
}

func (x *K8SJobRunParams) Reset() {
	*x = K8SJobRunParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SJobRunParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SJobRunParams) ProtoMessage() {}

func (x *K8SJobRunParams) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SJobRunParams.ProtoReflect.Descriptor instead.
func (*K8SJobRunParams) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{3}
}

func (x *K8SJobRunParams) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

var File_apps_task_pb_task_proto protoreflect.FileDescriptor

var file_apps_task_pb_task_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x1a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x6d, 0x0a, 0x04, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x38, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x2b, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0xee, 0x01, 0x0a, 0x0d, 0x52,
	0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6a, 0x6f, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42,
	0x0a, 0x04, 0x77, 0x69, 0x74, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x77, 0x69,
	0x74, 0x68, 0x1a, 0x37, 0x0a, 0x09, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x0f, 0x4b,
	0x38, 0x73, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_task_pb_task_proto_rawDescOnce sync.Once
	file_apps_task_pb_task_proto_rawDescData = file_apps_task_pb_task_proto_rawDesc
)

func file_apps_task_pb_task_proto_rawDescGZIP() []byte {
	file_apps_task_pb_task_proto_rawDescOnce.Do(func() {
		file_apps_task_pb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_task_pb_task_proto_rawDescData)
	})
	return file_apps_task_pb_task_proto_rawDescData
}

var file_apps_task_pb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_task_pb_task_proto_goTypes = []interface{}{
	(*TaskSet)(nil),         // 0: infraboard.mpaas.task.TaskSet
	(*Task)(nil),            // 1: infraboard.mpaas.task.Task
	(*RunJobRequest)(nil),   // 2: infraboard.mpaas.task.RunJobRequest
	(*K8SJobRunParams)(nil), // 3: infraboard.mpaas.task.K8sJobRunParams
	nil,                     // 4: infraboard.mpaas.task.RunJobRequest.WithEntry
	(*job.Job)(nil),         // 5: infraboard.mpaas.job.Job
}
var file_apps_task_pb_task_proto_depIdxs = []int32{
	1, // 0: infraboard.mpaas.task.TaskSet.items:type_name -> infraboard.mpaas.task.Task
	2, // 1: infraboard.mpaas.task.Task.spec:type_name -> infraboard.mpaas.task.RunJobRequest
	5, // 2: infraboard.mpaas.task.Task.job:type_name -> infraboard.mpaas.job.Job
	4, // 3: infraboard.mpaas.task.RunJobRequest.with:type_name -> infraboard.mpaas.task.RunJobRequest.WithEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_task_pb_task_proto_init() }
func file_apps_task_pb_task_proto_init() {
	if File_apps_task_pb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_task_pb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSet); i {
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
		file_apps_task_pb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_apps_task_pb_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobRequest); i {
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
		file_apps_task_pb_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SJobRunParams); i {
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
			RawDescriptor: file_apps_task_pb_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_task_pb_task_proto_goTypes,
		DependencyIndexes: file_apps_task_pb_task_proto_depIdxs,
		MessageInfos:      file_apps_task_pb_task_proto_msgTypes,
	}.Build()
	File_apps_task_pb_task_proto = out.File
	file_apps_task_pb_task_proto_rawDesc = nil
	file_apps_task_pb_task_proto_goTypes = nil
	file_apps_task_pb_task_proto_depIdxs = nil
}
