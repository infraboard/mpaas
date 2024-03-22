// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: mpaas/apps/event/pb/problem.proto

package event

import (
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

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 故障信息
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateProblemRequest `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 故障状态
	// @gotags: bson:"status" json:"status"
	Status *ProblemStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_problem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_problem_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetSpec() *CreateProblemRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Problem) GetStatus() *ProblemStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// 创建故障
type CreateProblemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 对象所在域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 对象所在空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 触发告警的规则名称, 一个规则 会生成一个故障单
	// @gotags: bson:"rule_name" json:"rule_name"
	RuleName string `protobuf:"bytes,3,opt,name=rule_name,json=ruleName,proto3" json:"rule_name" bson:"rule_name"`
	// 规则URL地址, 用于跳转查看
	// @gotags: bson:"rule_url" json:"rule_url"
	RuleUrl string `protobuf:"bytes,4,opt,name=rule_url,json=ruleUrl,proto3" json:"rule_url" bson:"rule_url"`
	// 故障Id, 如果不传会自动生成
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id" bson:"_id"`
	// 故障开始时间
	// @gotags: bson:"start" json:"start"
	Start int64 `protobuf:"varint,6,opt,name=start,proto3" json:"start" bson:"start"`
	// 故障结束时间
	// @gotags: bson:"end" json:"end"
	End int64 `protobuf:"varint,7,opt,name=end,proto3" json:"end" bson:"end"`
	// 故障级别
	// @gotags: bson:"level" json:"level"
	Level LEVEL `protobuf:"varint,8,opt,name=level,proto3,enum=infraboard.mpaas.event.LEVEL" json:"level" bson:"level"`
	// 故障标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateProblemRequest) Reset() {
	*x = CreateProblemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_problem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProblemRequest) ProtoMessage() {}

func (x *CreateProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_problem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProblemRequest.ProtoReflect.Descriptor instead.
func (*CreateProblemRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_problem_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProblemRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateProblemRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateProblemRequest) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *CreateProblemRequest) GetRuleUrl() string {
	if x != nil {
		return x.RuleUrl
	}
	return ""
}

func (x *CreateProblemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProblemRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *CreateProblemRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *CreateProblemRequest) GetLevel() LEVEL {
	if x != nil {
		return x.Level
	}
	return LEVEL_TRACE
}

func (x *CreateProblemRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type ProblemStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProblemStatus) Reset() {
	*x = ProblemStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_problem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemStatus) ProtoMessage() {}

func (x *ProblemStatus) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_problem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemStatus.ProtoReflect.Descriptor instead.
func (*ProblemStatus) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_problem_proto_rawDescGZIP(), []int{2}
}

var File_mpaas_apps_event_pb_problem_proto protoreflect.FileDescriptor

var file_mpaas_apps_event_pb_problem_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xfe, 0x02, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x75, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x50, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpaas_apps_event_pb_problem_proto_rawDescOnce sync.Once
	file_mpaas_apps_event_pb_problem_proto_rawDescData = file_mpaas_apps_event_pb_problem_proto_rawDesc
)

func file_mpaas_apps_event_pb_problem_proto_rawDescGZIP() []byte {
	file_mpaas_apps_event_pb_problem_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_event_pb_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_event_pb_problem_proto_rawDescData)
	})
	return file_mpaas_apps_event_pb_problem_proto_rawDescData
}

var file_mpaas_apps_event_pb_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mpaas_apps_event_pb_problem_proto_goTypes = []interface{}{
	(*Problem)(nil),              // 0: infraboard.mpaas.event.Problem
	(*CreateProblemRequest)(nil), // 1: infraboard.mpaas.event.CreateProblemRequest
	(*ProblemStatus)(nil),        // 2: infraboard.mpaas.event.ProblemStatus
	nil,                          // 3: infraboard.mpaas.event.CreateProblemRequest.LabelsEntry
	(LEVEL)(0),                   // 4: infraboard.mpaas.event.LEVEL
}
var file_mpaas_apps_event_pb_problem_proto_depIdxs = []int32{
	1, // 0: infraboard.mpaas.event.Problem.spec:type_name -> infraboard.mpaas.event.CreateProblemRequest
	2, // 1: infraboard.mpaas.event.Problem.status:type_name -> infraboard.mpaas.event.ProblemStatus
	4, // 2: infraboard.mpaas.event.CreateProblemRequest.level:type_name -> infraboard.mpaas.event.LEVEL
	3, // 3: infraboard.mpaas.event.CreateProblemRequest.labels:type_name -> infraboard.mpaas.event.CreateProblemRequest.LabelsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mpaas_apps_event_pb_problem_proto_init() }
func file_mpaas_apps_event_pb_problem_proto_init() {
	if File_mpaas_apps_event_pb_problem_proto != nil {
		return
	}
	file_mpaas_apps_event_pb_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_event_pb_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
		file_mpaas_apps_event_pb_problem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProblemRequest); i {
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
		file_mpaas_apps_event_pb_problem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemStatus); i {
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
			RawDescriptor: file_mpaas_apps_event_pb_problem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mpaas_apps_event_pb_problem_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_event_pb_problem_proto_depIdxs,
		MessageInfos:      file_mpaas_apps_event_pb_problem_proto_msgTypes,
	}.Build()
	File_mpaas_apps_event_pb_problem_proto = out.File
	file_mpaas_apps_event_pb_problem_proto_rawDesc = nil
	file_mpaas_apps_event_pb_problem_proto_goTypes = nil
	file_mpaas_apps_event_pb_problem_proto_depIdxs = nil
}
