// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/task/pb/job_task.proto

package task

import (
	job "github.com/infraboard/mpaas/apps/job"
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
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

type STAGE int32

const (
	// 等待执行
	STAGE_PENDDING STAGE = 0
	// 运行中
	STAGE_ACTIVE STAGE = 1
	// 回滚中, 回滚是一种特殊的运行中状态, 回滚时会运行回滚参数
	STAGE_ROLLBACKING STAGE = 2
	// 取消中
	STAGE_CANCELING STAGE = 3
	// 任务被取消
	STAGE_CANCELED STAGE = 10
	// 运行失败
	STAGE_FAILED STAGE = 11
	// 运行成功
	STAGE_SUCCEEDED STAGE = 12
)

// Enum value maps for STAGE.
var (
	STAGE_name = map[int32]string{
		0:  "PENDDING",
		1:  "ACTIVE",
		2:  "ROLLBACKING",
		3:  "CANCELING",
		10: "CANCELED",
		11: "FAILED",
		12: "SUCCEEDED",
	}
	STAGE_value = map[string]int32{
		"PENDDING":    0,
		"ACTIVE":      1,
		"ROLLBACKING": 2,
		"CANCELING":   3,
		"CANCELED":    10,
		"FAILED":      11,
		"SUCCEEDED":   12,
	}
)

func (x STAGE) Enum() *STAGE {
	p := new(STAGE)
	*p = x
	return p
}

func (x STAGE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STAGE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_task_pb_job_task_proto_enumTypes[0].Descriptor()
}

func (STAGE) Type() protoreflect.EnumType {
	return &file_apps_task_pb_job_task_proto_enumTypes[0]
}

func (x STAGE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STAGE.Descriptor instead.
func (STAGE) EnumDescriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{0}
}

type EVENT_LEVEL int32

const (
	// debug信息, 无影响
	EVENT_LEVEL_DEBUG EVENT_LEVEL = 0
	// info, 普通通知
	EVENT_LEVEL_INFO EVENT_LEVEL = 1
	// warn, 可能导致任务异常
	EVENT_LEVEL_WARN EVENT_LEVEL = 2
	// error, 任务的一些错误信息
	EVENT_LEVEL_ERROR EVENT_LEVEL = 3
)

// Enum value maps for EVENT_LEVEL.
var (
	EVENT_LEVEL_name = map[int32]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARN",
		3: "ERROR",
	}
	EVENT_LEVEL_value = map[string]int32{
		"DEBUG": 0,
		"INFO":  1,
		"WARN":  2,
		"ERROR": 3,
	}
)

func (x EVENT_LEVEL) Enum() *EVENT_LEVEL {
	p := new(EVENT_LEVEL)
	*p = x
	return p
}

func (x EVENT_LEVEL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EVENT_LEVEL) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_task_pb_job_task_proto_enumTypes[1].Descriptor()
}

func (EVENT_LEVEL) Type() protoreflect.EnumType {
	return &file_apps_task_pb_job_task_proto_enumTypes[1]
}

func (x EVENT_LEVEL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EVENT_LEVEL.Descriptor instead.
func (EVENT_LEVEL) EnumDescriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{1}
}

type JobTaskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 清单
	// @gotags: bson:"items" json:"items"
	Items []*JobTask `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *JobTaskSet) Reset() {
	*x = JobTaskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTaskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskSet) ProtoMessage() {}

func (x *JobTaskSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskSet.ProtoReflect.Descriptor instead.
func (*JobTaskSet) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{0}
}

func (x *JobTaskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *JobTaskSet) GetItems() []*JobTask {
	if x != nil {
		return x.Items
	}
	return nil
}

type JobTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// task定义, job运行时参数
	// @gotags: bson:",inline" json:"spec"
	Spec *pipeline.RunJobRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 任务当前状态
	// @gotags: bson:"status" json:"status"
	Status *JobTaskStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status" bson:"status"`
	// 关联Job
	// @gotags: bson:"job" json:"job"
	Job *job.Job `protobuf:"bytes,4,opt,name=job,proto3" json:"job" bson:"job"`
}

func (x *JobTask) Reset() {
	*x = JobTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTask) ProtoMessage() {}

func (x *JobTask) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTask.ProtoReflect.Descriptor instead.
func (*JobTask) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{1}
}

func (x *JobTask) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *JobTask) GetSpec() *pipeline.RunJobRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *JobTask) GetStatus() *JobTaskStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *JobTask) GetJob() *job.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,1,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,2,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 更新人
	// @gotags: bson:"update_by" json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Meta) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Meta) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

type JobTaskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务当前状态
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,1,opt,name=stage,proto3,enum=infraboard.mpaas.task.STAGE" json:"stage" bson:"stage"`
	// 任务开始时间
	// @gotags: bson:"start_at" json:"start_at"
	StartAt int64 `protobuf:"varint,2,opt,name=start_at,json=startAt,proto3" json:"start_at" bson:"start_at"`
	// 任务结束时间
	// @gotags: bson:"end_at" json:"end_at"
	EndAt int64 `protobuf:"varint,3,opt,name=end_at,json=endAt,proto3" json:"end_at" bson:"end_at"`
	// 状态描述
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message" bson:"message"`
	// 任务状态详细描述, 用于Debug
	// @gotags: bson:"detail" json:"detail"
	Detail string `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail" bson:"detail"`
	// Job Task运行时环境变量, 大写开头的变量会更新到pipline中, 注入到后续执行的任务中
	// 小写开头的变量, 作为Task运行后的输出, task自己保留
	// @gotags: bson:"runtime_envs" json:"runtime_envs"
	RuntimeEnvs []*RuntimeEnv `protobuf:"bytes,6,rep,name=runtime_envs,json=runtimeEnvs,proto3" json:"runtime_envs" bson:"runtime_envs"`
	// 任务运行后的产物信息, 用于界面展示
	// @gotags: bson:"markdown_output" json:"markdown_output"
	MarkdownOutput string `protobuf:"bytes,7,opt,name=markdown_output,json=markdownOutput,proto3" json:"markdown_output" bson:"markdown_output"`
	// 任务所需的临时资源
	// @gotags: bson:"temporary_resources" json:"temporary_resources"
	TemporaryResources []*TemporaryResource `protobuf:"bytes,8,rep,name=temporary_resources,json=temporaryResources,proto3" json:"temporary_resources" bson:"temporary_resources"`
	// 任务事件
	// @gotags: bson:"events" json:"events"
	Events []*Event `protobuf:"bytes,9,rep,name=events,proto3" json:"events" bson:"events"`
}

func (x *JobTaskStatus) Reset() {
	*x = JobTaskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTaskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskStatus) ProtoMessage() {}

func (x *JobTaskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskStatus.ProtoReflect.Descriptor instead.
func (*JobTaskStatus) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{3}
}

func (x *JobTaskStatus) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_PENDDING
}

func (x *JobTaskStatus) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *JobTaskStatus) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *JobTaskStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JobTaskStatus) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *JobTaskStatus) GetRuntimeEnvs() []*RuntimeEnv {
	if x != nil {
		return x.RuntimeEnvs
	}
	return nil
}

func (x *JobTaskStatus) GetMarkdownOutput() string {
	if x != nil {
		return x.MarkdownOutput
	}
	return ""
}

func (x *JobTaskStatus) GetTemporaryResources() []*TemporaryResource {
	if x != nil {
		return x.TemporaryResources
	}
	return nil
}

func (x *JobTaskStatus) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件发生的时间
	// @gotags: bson:"time" json:"time"
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time" bson:"time"`
	// 事件等级
	// @gotags: bson:"level" json:"level"
	Level EVENT_LEVEL `protobuf:"varint,2,opt,name=level,proto3,enum=infraboard.mpaas.task.EVENT_LEVEL" json:"level" bson:"level"`
	// 事件描述
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message" bson:"message"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Event) GetLevel() EVENT_LEVEL {
	if x != nil {
		return x.Level
	}
	return EVENT_LEVEL_DEBUG
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// task 运行过程中 产生的临时环境变量, 即用记录Job的输出(将输出保存为环境变量)
// 在Task结束时同步到Pipeline Task中, 后续Task执行时 读取并注入
// 注意只有大写的变量才会被导出
type RuntimeEnv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 环境变量名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 环境变量值
	// @gotags: bson:"value" json:"value"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value" bson:"value"`
	// 是否只读 不允许修改
	// @gotags: bson:"read_only" json:"read_only"
	ReadOnly bool `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"read_only" bson:"read_only"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 被谁更新的, 当值被修改后，追踪下谁修改的
	// @gotags: bson:"update_task_id" json:"update_task_id"
	UpdateTaskId string `protobuf:"bytes,5,opt,name=update_task_id,json=updateTaskId,proto3" json:"update_task_id" bson:"update_task_id"`
}

func (x *RuntimeEnv) Reset() {
	*x = RuntimeEnv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeEnv) ProtoMessage() {}

func (x *RuntimeEnv) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeEnv.ProtoReflect.Descriptor instead.
func (*RuntimeEnv) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{5}
}

func (x *RuntimeEnv) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RuntimeEnv) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RuntimeEnv) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *RuntimeEnv) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *RuntimeEnv) GetUpdateTaskId() string {
	if x != nil {
		return x.UpdateTaskId
	}
	return ""
}

// 临时资源, 在Pipline允许结束时,调用释放
type TemporaryResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源的类型, 比如 configmap
	// @gotags: bson:"kind" json:"kind"
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind" bson:"kind"`
	// 资源的名字, 资源的集群和Namespace 由job维护
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name"`
	// 资源的详情数据
	// @gotags: bson:"detail" json:"detail"
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail" bson:"detail"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 创建时间
	// @gotags: bson:"release_at" json:"release_at"
	ReleaseAt int64 `protobuf:"varint,5,opt,name=release_at,json=releaseAt,proto3" json:"release_at" bson:"release_at"`
}

func (x *TemporaryResource) Reset() {
	*x = TemporaryResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporaryResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporaryResource) ProtoMessage() {}

func (x *TemporaryResource) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporaryResource.ProtoReflect.Descriptor instead.
func (*TemporaryResource) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{6}
}

func (x *TemporaryResource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *TemporaryResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemporaryResource) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *TemporaryResource) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *TemporaryResource) GetReleaseAt() int64 {
	if x != nil {
		return x.ReleaseAt
	}
	return 0
}

type RunTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// job名称定义
	// @gotags: bson:"job_spec" json:"job_spec"
	JobSpec string `protobuf:"bytes,2,opt,name=job_spec,json=jobSpec,proto3" json:"job_spec" bson:"job_spec"`
	// 试运行, 并不会真正执行Job
	// @gotags: bson:"dry_run" json:"dry_run"
	DryRun bool `protobuf:"varint,3,opt,name=dry_run,json=dryRun,proto3" json:"dry_run" bson:"dry_run"`
	// 手动更新Job的状态, 默认由job runner的operator更新
	// @gotags: bson:"manual_update_status" json:"manual_update_status"
	ManualUpdateStatus bool `protobuf:"varint,4,opt,name=manual_update_status,json=manualUpdateStatus,proto3" json:"manual_update_status" bson:"manual_update_status"`
	// job运行时参数
	// @gotags: bson:"params" json:"params"
	Params *job.VersionedRunParam `protobuf:"bytes,5,opt,name=params,proto3" json:"params" bson:"params"`
	// 标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *RunTaskRequest) Reset() {
	*x = RunTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_job_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskRequest) ProtoMessage() {}

func (x *RunTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_job_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskRequest.ProtoReflect.Descriptor instead.
func (*RunTaskRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_job_task_proto_rawDescGZIP(), []int{7}
}

func (x *RunTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunTaskRequest) GetJobSpec() string {
	if x != nil {
		return x.JobSpec
	}
	return ""
}

func (x *RunTaskRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *RunTaskRequest) GetManualUpdateStatus() bool {
	if x != nil {
		return x.ManualUpdateStatus
	}
	return false
}

func (x *RunTaskRequest) GetParams() *job.VersionedRunParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *RunTaskRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_apps_task_pb_job_task_proto protoreflect.FileDescriptor

var file_apps_task_pb_job_task_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6a,
	0x6f, 0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70,
	0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0a,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xe3, 0x01, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x75,
	0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2b, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x5d, 0x0a, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x22, 0xa7, 0x03, 0x0a, 0x0d,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x44, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x52, 0x0b,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d,
	0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x59, 0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72,
	0x79, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x12, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6f, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x45, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22,
	0x8f, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41,
	0x74, 0x22, 0xd1, 0x02, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x30, 0x0a, 0x14,
	0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6d, 0x61, 0x6e, 0x75,
	0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52,
	0x75, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x6a, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x47, 0x45, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x45, 0x4e, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x4c,
	0x42, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x0c, 0x2a, 0x37, 0x0a, 0x0b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_task_pb_job_task_proto_rawDescOnce sync.Once
	file_apps_task_pb_job_task_proto_rawDescData = file_apps_task_pb_job_task_proto_rawDesc
)

func file_apps_task_pb_job_task_proto_rawDescGZIP() []byte {
	file_apps_task_pb_job_task_proto_rawDescOnce.Do(func() {
		file_apps_task_pb_job_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_task_pb_job_task_proto_rawDescData)
	})
	return file_apps_task_pb_job_task_proto_rawDescData
}

var file_apps_task_pb_job_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_task_pb_job_task_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_task_pb_job_task_proto_goTypes = []interface{}{
	(STAGE)(0),                     // 0: infraboard.mpaas.task.STAGE
	(EVENT_LEVEL)(0),               // 1: infraboard.mpaas.task.EVENT_LEVEL
	(*JobTaskSet)(nil),             // 2: infraboard.mpaas.task.JobTaskSet
	(*JobTask)(nil),                // 3: infraboard.mpaas.task.JobTask
	(*Meta)(nil),                   // 4: infraboard.mpaas.task.Meta
	(*JobTaskStatus)(nil),          // 5: infraboard.mpaas.task.JobTaskStatus
	(*Event)(nil),                  // 6: infraboard.mpaas.task.Event
	(*RuntimeEnv)(nil),             // 7: infraboard.mpaas.task.RuntimeEnv
	(*TemporaryResource)(nil),      // 8: infraboard.mpaas.task.TemporaryResource
	(*RunTaskRequest)(nil),         // 9: infraboard.mpaas.task.RunTaskRequest
	nil,                            // 10: infraboard.mpaas.task.RunTaskRequest.LabelsEntry
	(*pipeline.RunJobRequest)(nil), // 11: infraboard.mpaas.pipeline.RunJobRequest
	(*job.Job)(nil),                // 12: infraboard.mpaas.job.Job
	(*job.VersionedRunParam)(nil),  // 13: infraboard.mpaas.job.VersionedRunParam
}
var file_apps_task_pb_job_task_proto_depIdxs = []int32{
	3,  // 0: infraboard.mpaas.task.JobTaskSet.items:type_name -> infraboard.mpaas.task.JobTask
	4,  // 1: infraboard.mpaas.task.JobTask.meta:type_name -> infraboard.mpaas.task.Meta
	11, // 2: infraboard.mpaas.task.JobTask.spec:type_name -> infraboard.mpaas.pipeline.RunJobRequest
	5,  // 3: infraboard.mpaas.task.JobTask.status:type_name -> infraboard.mpaas.task.JobTaskStatus
	12, // 4: infraboard.mpaas.task.JobTask.job:type_name -> infraboard.mpaas.job.Job
	0,  // 5: infraboard.mpaas.task.JobTaskStatus.stage:type_name -> infraboard.mpaas.task.STAGE
	7,  // 6: infraboard.mpaas.task.JobTaskStatus.runtime_envs:type_name -> infraboard.mpaas.task.RuntimeEnv
	8,  // 7: infraboard.mpaas.task.JobTaskStatus.temporary_resources:type_name -> infraboard.mpaas.task.TemporaryResource
	6,  // 8: infraboard.mpaas.task.JobTaskStatus.events:type_name -> infraboard.mpaas.task.Event
	1,  // 9: infraboard.mpaas.task.Event.level:type_name -> infraboard.mpaas.task.EVENT_LEVEL
	13, // 10: infraboard.mpaas.task.RunTaskRequest.params:type_name -> infraboard.mpaas.job.VersionedRunParam
	10, // 11: infraboard.mpaas.task.RunTaskRequest.labels:type_name -> infraboard.mpaas.task.RunTaskRequest.LabelsEntry
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_apps_task_pb_job_task_proto_init() }
func file_apps_task_pb_job_task_proto_init() {
	if File_apps_task_pb_job_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_task_pb_job_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTaskSet); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTask); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTaskStatus); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeEnv); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporaryResource); i {
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
		file_apps_task_pb_job_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTaskRequest); i {
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
			RawDescriptor: file_apps_task_pb_job_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_task_pb_job_task_proto_goTypes,
		DependencyIndexes: file_apps_task_pb_job_task_proto_depIdxs,
		EnumInfos:         file_apps_task_pb_job_task_proto_enumTypes,
		MessageInfos:      file_apps_task_pb_job_task_proto_msgTypes,
	}.Build()
	File_apps_task_pb_job_task_proto = out.File
	file_apps_task_pb_job_task_proto_rawDesc = nil
	file_apps_task_pb_job_task_proto_goTypes = nil
	file_apps_task_pb_job_task_proto_depIdxs = nil
}
