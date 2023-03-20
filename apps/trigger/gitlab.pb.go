// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.6
// source: apps/trigger/pb/gitlab.proto

package trigger

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

// https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html
type EVENT_TYPE int32

const (
	// 推送事件
	EVENT_TYPE_PUSH EVENT_TYPE = 0
	// 打标事件
	EVENT_TYPE_TAG EVENT_TYPE = 1
	// bug事件
	EVENT_TYPE_ISSUE EVENT_TYPE = 2
	// 评论事件
	EVENT_TYPE_COMMENT EVENT_TYPE = 3
	// MR事件
	EVENT_TYPE_MERGE_REQUEST EVENT_TYPE = 4
)

// Enum value maps for EVENT_TYPE.
var (
	EVENT_TYPE_name = map[int32]string{
		0: "PUSH",
		1: "TAG",
		2: "ISSUE",
		3: "COMMENT",
		4: "MERGE_REQUEST",
	}
	EVENT_TYPE_value = map[string]int32{
		"PUSH":          0,
		"TAG":           1,
		"ISSUE":         2,
		"COMMENT":       3,
		"MERGE_REQUEST": 4,
	}
)

func (x EVENT_TYPE) Enum() *EVENT_TYPE {
	p := new(EVENT_TYPE)
	*p = x
	return p
}

func (x EVENT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EVENT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_trigger_pb_gitlab_proto_enumTypes[0].Descriptor()
}

func (EVENT_TYPE) Type() protoreflect.EnumType {
	return &file_apps_trigger_pb_gitlab_proto_enumTypes[0]
}

func (x EVENT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EVENT_TYPE.Descriptor instead.
func (EVENT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_trigger_pb_gitlab_proto_rawDescGZIP(), []int{0}
}

type GitlabWebHookEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件类型
	// @gotags: bson:"event_type" json:"event_type"
	EventType EVENT_TYPE `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=infraboard.mpaas.trigger.EVENT_TYPE" json:"event_type" bson:"event_type"`
	// 事件类型
	// @gotags: bson:"object_kind" json:"object_kind" validate:"required"
	ObjectKind string `protobuf:"bytes,2,opt,name=object_kind,json=objectKind,proto3" json:"object_kind" bson:"object_kind" validate:"required"`
	// 事件名称
	// @gotags: bson:"event_name" json:"event_name" validate:"required"
	EventName string `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name" bson:"event_name" validate:"required"`
	// 关联分支
	// @gotags: bson:"ref" json:"ref" validate:"required"
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref" bson:"ref" validate:"required"`
	// 触发者用户ID
	// @gotags: bson:"user_id" json:"user_id"
	UserId int64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id" bson:"user_id"`
	// 触发者用户名称
	// @gotags: bson:"user_name" json:"user_name"
	UserName string `protobuf:"bytes,6,opt,name=user_name,json=userName,proto3" json:"user_name" bson:"user_name"`
	// 用户头像
	// @gotags: bson:"user_avatar" json:"user_avatar"
	UserAvatar string `protobuf:"bytes,7,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar" bson:"user_avatar"`
	// 用户头像
	// @gotags: bson:"project" json:"project"
	Project *Project `protobuf:"bytes,8,opt,name=project,proto3" json:"project" bson:"project"`
	// Commit信息
	// @gotags: bson:"commits" json:"commits"
	Commits []*Commit `protobuf:"bytes,9,rep,name=commits,proto3" json:"commits" bson:"commits"`
}

func (x *GitlabWebHookEvent) Reset() {
	*x = GitlabWebHookEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitlabWebHookEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabWebHookEvent) ProtoMessage() {}

func (x *GitlabWebHookEvent) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabWebHookEvent.ProtoReflect.Descriptor instead.
func (*GitlabWebHookEvent) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_gitlab_proto_rawDescGZIP(), []int{0}
}

func (x *GitlabWebHookEvent) GetEventType() EVENT_TYPE {
	if x != nil {
		return x.EventType
	}
	return EVENT_TYPE_PUSH
}

func (x *GitlabWebHookEvent) GetObjectKind() string {
	if x != nil {
		return x.ObjectKind
	}
	return ""
}

func (x *GitlabWebHookEvent) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *GitlabWebHookEvent) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *GitlabWebHookEvent) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GitlabWebHookEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GitlabWebHookEvent) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *GitlabWebHookEvent) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *GitlabWebHookEvent) GetCommits() []*Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项目id
	// @gotags: bson:"id" json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" bson:"id"`
	// 描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description" bson:"description"`
	// 名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" bson:"name"`
	// 项目的Web访问地址
	// @gotags: bson:"web_url" json:"web_url"
	WebUrl string `protobuf:"bytes,7,opt,name=web_url,json=webUrl,proto3" json:"web_url" bson:"web_url"`
	// 项目的Logo地址
	// @gotags: bson:"avatar_url" json:"avatar_url"
	AvatarUrl string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url" bson:"avatar_url"`
	// ssh 地址
	// @gotags: bson:"git_ssh_url" json:"git_ssh_url"
	GitSshUrl string `protobuf:"bytes,4,opt,name=git_ssh_url,json=gitSshUrl,proto3" json:"git_ssh_url" bson:"git_ssh_url"`
	// http 地址
	// @gotags: bson:"git_http_url" json:"git_http_url"
	GitHttpUrl string `protobuf:"bytes,5,opt,name=git_http_url,json=gitHttpUrl,proto3" json:"git_http_url" bson:"git_http_url"`
	// namespace
	// @gotags: bson:"path_with_namespace" json:"path_with_namespace"
	NamespacePath string `protobuf:"bytes,6,opt,name=namespace_path,json=namespacePath,proto3" json:"path_with_namespace" bson:"path_with_namespace"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_gitlab_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetWebUrl() string {
	if x != nil {
		return x.WebUrl
	}
	return ""
}

func (x *Project) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Project) GetGitSshUrl() string {
	if x != nil {
		return x.GitSshUrl
	}
	return ""
}

func (x *Project) GetGitHttpUrl() string {
	if x != nil {
		return x.GitHttpUrl
	}
	return ""
}

func (x *Project) GetNamespacePath() string {
	if x != nil {
		return x.NamespacePath
	}
	return ""
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	// @gotags: bson:"id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	// commit message
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" bson:"message"`
	// title
	// @gotags: bson:"title" json:"title"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title" bson:"title"`
	// 文本格式时间
	// @gotags: bson:"timestamp" json:"timestamp"
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp" bson:"timestamp"`
	// commit对应的url
	// @gotags: bson:"url" json:"url"
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url" bson:"url"`
	// 作者
	// @gotags: bson:"author" json:"author"
	Author *Author `protobuf:"bytes,6,opt,name=author,proto3" json:"author" bson:"author"`
	// 新加的文件
	// @gotags: bson:"added" json:"added"
	Added []string `protobuf:"bytes,7,rep,name=added,proto3" json:"added" bson:"added"`
	// 修改的文件
	// @gotags: bson:"modified" json:"modified"
	Modified []string `protobuf:"bytes,8,rep,name=modified,proto3" json:"modified" bson:"modified"`
	// 删除的文件
	// @gotags: bson:"removed" json:"removed"
	Removed []string `protobuf:"bytes,9,rep,name=removed,proto3" json:"removed" bson:"removed"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_gitlab_proto_rawDescGZIP(), []int{2}
}

func (x *Commit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Commit) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Commit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Commit) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Commit) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Commit) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Commit) GetAdded() []string {
	if x != nil {
		return x.Added
	}
	return nil
}

func (x *Commit) GetModified() []string {
	if x != nil {
		return x.Modified
	}
	return nil
}

func (x *Commit) GetRemoved() []string {
	if x != nil {
		return x.Removed
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 作者名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 作者邮箱
	// @gotags: bson:"email" json:"email"
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email" bson:"email"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_gitlab_proto_rawDescGZIP(), []int{3}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Merge Request时相关数据
type ObjectAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项目id
	// @gotags: bson:"id" json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" bson:"id"`
	// 项目id
	// @gotags: bson:"iid" json:"iid"
	Iid int64 `protobuf:"varint,2,opt,name=iid,proto3" json:"iid" bson:"iid"`
	// 目标项目
	// @gotags: bson:"target_project_id" json:"target_project_id"
	TargetProjectId int64 `protobuf:"varint,3,opt,name=target_project_id,json=targetProjectId,proto3" json:"target_project_id" bson:"target_project_id"`
	// 目标分支
	// @gotags: bson:"target_branch" json:"target_branch"
	TargetBranch string `protobuf:"bytes,4,opt,name=target_branch,json=targetBranch,proto3" json:"target_branch" bson:"target_branch"`
	// 源项目
	// @gotags: bson:"source_project_id" json:"source_project_id"
	SourceProjectId int64 `protobuf:"varint,5,opt,name=source_project_id,json=sourceProjectId,proto3" json:"source_project_id" bson:"source_project_id"`
	// 源分支
	// @gotags: bson:"source_branch" json:"source_branch"
	SourceBranch string `protobuf:"bytes,6,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch" bson:"source_branch"`
	// 作者
	// @gotags: bson:"author_id" json:"author_id"
	AuthorId int64 `protobuf:"varint,7,opt,name=author_id,json=authorId,proto3" json:"author_id" bson:"author_id"`
	// 接收方
	// @gotags: bson:"assignee_ids" json:"assignee_ids"
	AssigneeIds []int64 `protobuf:"varint,8,rep,packed,name=assignee_ids,json=assigneeIds,proto3" json:"assignee_ids" bson:"assignee_ids"`
	// 作者
	// @gotags: bson:"assignee_id" json:"assignee_id"
	AssigneeId int64 `protobuf:"varint,9,opt,name=assignee_id,json=assigneeId,proto3" json:"assignee_id" bson:"assignee_id"`
	// 接收方
	// @gotags: bson:"reviewer_ids" json:"reviewer_ids"
	ReviewerIds []int64 `protobuf:"varint,10,rep,packed,name=reviewer_ids,json=reviewerIds,proto3" json:"reviewer_ids" bson:"reviewer_ids"`
	// 标题
	// @gotags: bson:"title" json:"title"
	Title string `protobuf:"bytes,11,opt,name=title,proto3" json:"title" bson:"title"`
	// 具体描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description" bson:"description"`
	// 创建时间
	// @gotags: bson:"created_at" json:"created_at"
	CreatedAt string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at" bson:"created_at"`
	// 更新时间
	// @gotags: bson:"updated_at" json:"updated_at"
	UpdatedAt string `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at" bson:"updated_at"`
	// 状态
	// @gotags: bson:"state" json:"state"
	State string `protobuf:"bytes,15,opt,name=state,proto3" json:"state" bson:"state"`
	// 阻塞到解决
	// @gotags: bson:"blocking_discussions_resolved" json:"blocking_discussions_resolved"
	BlockingDiscussionsResolved bool `protobuf:"varint,16,opt,name=blocking_discussions_resolved,json=blockingDiscussionsResolved,proto3" json:"blocking_discussions_resolved" bson:"blocking_discussions_resolved"`
	// 解决中
	// @gotags: bson:"work_in_progress" json:"work_in_progress"
	WorkInProgress bool `protobuf:"varint,17,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress" bson:"work_in_progress"`
	// 第一次贡献
	// @gotags: bson:"first_contribution" json:"first_contribution"
	FirstContribution bool `protobuf:"varint,18,opt,name=first_contribution,json=firstContribution,proto3" json:"first_contribution" bson:"first_contribution"`
	// merge状态
	// @gotags: bson:"merge_status" json:"merge_status"
	MergeStatus string `protobuf:"bytes,19,opt,name=merge_status,json=mergeStatus,proto3" json:"merge_status" bson:"merge_status"`
	// MR URL
	// @gotags: bson:"url" json:"url"
	Url string `protobuf:"bytes,20,opt,name=url,proto3" json:"url" bson:"url"`
	// 最近一次commit
	// @gotags: bson:"last_commit" json:"last_commit"
	LastCommit *Commit `protobuf:"bytes,21,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit" bson:"last_commit"`
	// 具体操作
	// @gotags: bson:"action" json:"action"
	Action string `protobuf:"bytes,22,opt,name=action,proto3" json:"action" bson:"action"`
	// 状态的详细描述
	// @gotags: bson:"detailed_merge_status" json:"detailed_merge_status"
	DetailedMergeStatus string `protobuf:"bytes,23,opt,name=detailed_merge_status,json=detailedMergeStatus,proto3" json:"detailed_merge_status" bson:"detailed_merge_status"`
}

func (x *ObjectAttributes) Reset() {
	*x = ObjectAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectAttributes) ProtoMessage() {}

func (x *ObjectAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_gitlab_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectAttributes.ProtoReflect.Descriptor instead.
func (*ObjectAttributes) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_gitlab_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectAttributes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjectAttributes) GetIid() int64 {
	if x != nil {
		return x.Iid
	}
	return 0
}

func (x *ObjectAttributes) GetTargetProjectId() int64 {
	if x != nil {
		return x.TargetProjectId
	}
	return 0
}

func (x *ObjectAttributes) GetTargetBranch() string {
	if x != nil {
		return x.TargetBranch
	}
	return ""
}

func (x *ObjectAttributes) GetSourceProjectId() int64 {
	if x != nil {
		return x.SourceProjectId
	}
	return 0
}

func (x *ObjectAttributes) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

func (x *ObjectAttributes) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *ObjectAttributes) GetAssigneeIds() []int64 {
	if x != nil {
		return x.AssigneeIds
	}
	return nil
}

func (x *ObjectAttributes) GetAssigneeId() int64 {
	if x != nil {
		return x.AssigneeId
	}
	return 0
}

func (x *ObjectAttributes) GetReviewerIds() []int64 {
	if x != nil {
		return x.ReviewerIds
	}
	return nil
}

func (x *ObjectAttributes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ObjectAttributes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ObjectAttributes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ObjectAttributes) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ObjectAttributes) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ObjectAttributes) GetBlockingDiscussionsResolved() bool {
	if x != nil {
		return x.BlockingDiscussionsResolved
	}
	return false
}

func (x *ObjectAttributes) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

func (x *ObjectAttributes) GetFirstContribution() bool {
	if x != nil {
		return x.FirstContribution
	}
	return false
}

func (x *ObjectAttributes) GetMergeStatus() string {
	if x != nil {
		return x.MergeStatus
	}
	return ""
}

func (x *ObjectAttributes) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ObjectAttributes) GetLastCommit() *Commit {
	if x != nil {
		return x.LastCommit
	}
	return nil
}

func (x *ObjectAttributes) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ObjectAttributes) GetDetailedMergeStatus() string {
	if x != nil {
		return x.DetailedMergeStatus
	}
	return ""
}

var File_apps_trigger_pb_gitlab_proto protoreflect.FileDescriptor

var file_apps_trigger_pb_gitlab_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x22, 0xfb, 0x02, 0x0a, 0x12, 0x47, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x43, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x3b, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c,
	0x12, 0x1e, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x73, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x53, 0x73, 0x68, 0x55, 0x72, 0x6c,
	0x12, 0x20, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x48, 0x74, 0x74, 0x70, 0x55,
	0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0xfe, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x32, 0x0a, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xc7,
	0x06, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x1d, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x41, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x4a, 0x0a, 0x0a, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x54, 0x41, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x53, 0x53,
	0x55, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x04, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_trigger_pb_gitlab_proto_rawDescOnce sync.Once
	file_apps_trigger_pb_gitlab_proto_rawDescData = file_apps_trigger_pb_gitlab_proto_rawDesc
)

func file_apps_trigger_pb_gitlab_proto_rawDescGZIP() []byte {
	file_apps_trigger_pb_gitlab_proto_rawDescOnce.Do(func() {
		file_apps_trigger_pb_gitlab_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_trigger_pb_gitlab_proto_rawDescData)
	})
	return file_apps_trigger_pb_gitlab_proto_rawDescData
}

var file_apps_trigger_pb_gitlab_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_trigger_pb_gitlab_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_trigger_pb_gitlab_proto_goTypes = []interface{}{
	(EVENT_TYPE)(0),            // 0: infraboard.mpaas.trigger.EVENT_TYPE
	(*GitlabWebHookEvent)(nil), // 1: infraboard.mpaas.trigger.GitlabWebHookEvent
	(*Project)(nil),            // 2: infraboard.mpaas.trigger.Project
	(*Commit)(nil),             // 3: infraboard.mpaas.trigger.Commit
	(*Author)(nil),             // 4: infraboard.mpaas.trigger.Author
	(*ObjectAttributes)(nil),   // 5: infraboard.mpaas.trigger.ObjectAttributes
}
var file_apps_trigger_pb_gitlab_proto_depIdxs = []int32{
	0, // 0: infraboard.mpaas.trigger.GitlabWebHookEvent.event_type:type_name -> infraboard.mpaas.trigger.EVENT_TYPE
	2, // 1: infraboard.mpaas.trigger.GitlabWebHookEvent.project:type_name -> infraboard.mpaas.trigger.Project
	3, // 2: infraboard.mpaas.trigger.GitlabWebHookEvent.commits:type_name -> infraboard.mpaas.trigger.Commit
	4, // 3: infraboard.mpaas.trigger.Commit.author:type_name -> infraboard.mpaas.trigger.Author
	3, // 4: infraboard.mpaas.trigger.ObjectAttributes.last_commit:type_name -> infraboard.mpaas.trigger.Commit
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_trigger_pb_gitlab_proto_init() }
func file_apps_trigger_pb_gitlab_proto_init() {
	if File_apps_trigger_pb_gitlab_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_trigger_pb_gitlab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitlabWebHookEvent); i {
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
		file_apps_trigger_pb_gitlab_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_apps_trigger_pb_gitlab_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_apps_trigger_pb_gitlab_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_apps_trigger_pb_gitlab_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectAttributes); i {
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
			RawDescriptor: file_apps_trigger_pb_gitlab_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_trigger_pb_gitlab_proto_goTypes,
		DependencyIndexes: file_apps_trigger_pb_gitlab_proto_depIdxs,
		EnumInfos:         file_apps_trigger_pb_gitlab_proto_enumTypes,
		MessageInfos:      file_apps_trigger_pb_gitlab_proto_msgTypes,
	}.Build()
	File_apps_trigger_pb_gitlab_proto = out.File
	file_apps_trigger_pb_gitlab_proto_rawDesc = nil
	file_apps_trigger_pb_gitlab_proto_goTypes = nil
	file_apps_trigger_pb_gitlab_proto_depIdxs = nil
}
