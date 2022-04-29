// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: api/application/v1/application.proto

package v1

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

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Replicas  uint32 `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	PlanId    uint32 `protobuf:"varint,4,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	State     string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	StateId   uint32 `protobuf:"varint,6,opt,name=state_id,json=stateId,proto3" json:"state_id,omitempty"`
	CompanyId uint32 `protobuf:"varint,7,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatorId uint32 `protobuf:"varint,8,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	UpdaterId uint32 `protobuf:"varint,9,opt,name=updater_id,json=updaterId,proto3" json:"updater_id,omitempty"`
	CreatedAt string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetReplicas() uint32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Application) GetPlanId() uint32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *Application) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Application) GetStateId() uint32 {
	if x != nil {
		return x.StateId
	}
	return 0
}

func (x *Application) GetCompanyId() uint32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Application) GetCreatorId() uint32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Application) GetUpdaterId() uint32 {
	if x != nil {
		return x.UpdaterId
	}
	return 0
}

func (x *Application) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Application) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AuthServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AuthServer) Reset() {
	*x = AuthServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthServer) ProtoMessage() {}

func (x *AuthServer) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthServer.ProtoReflect.Descriptor instead.
func (*AuthServer) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{1}
}

func (x *AuthServer) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AuthServer) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StateId uint32 `protobuf:"varint,2,opt,name=state_id,json=stateId,proto3" json:"state_id,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{2}
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pod) GetStateId() uint32 {
	if x != nil {
		return x.StateId
	}
	return 0
}

type CreateAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App        *Application `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	AuthServer *AuthServer  `protobuf:"bytes,2,opt,name=auth_server,json=authServer,proto3" json:"auth_server,omitempty"`
}

func (x *CreateAppRequest) Reset() {
	*x = CreateAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRequest) ProtoMessage() {}

func (x *CreateAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRequest.ProtoReflect.Descriptor instead.
func (*CreateAppRequest) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAppRequest) GetApp() *Application {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *CreateAppRequest) GetAuthServer() *AuthServer {
	if x != nil {
		return x.AuthServer
	}
	return nil
}

type CreateAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAppResponse) Reset() {
	*x = CreateAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppResponse) ProtoMessage() {}

func (x *CreateAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppResponse.ProtoReflect.Descriptor instead.
func (*CreateAppResponse) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{4}
}

type DeleteAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAppRequest) Reset() {
	*x = DeleteAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppRequest) ProtoMessage() {}

func (x *DeleteAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppRequest.ProtoReflect.Descriptor instead.
func (*DeleteAppRequest) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAppRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State uint32 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *DeleteAppResponse) Reset() {
	*x = DeleteAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppResponse) ProtoMessage() {}

func (x *DeleteAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppResponse.ProtoReflect.Descriptor instead.
func (*DeleteAppResponse) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAppResponse) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type UpdateAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App        *Application `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	AuthServer *AuthServer  `protobuf:"bytes,2,opt,name=auth_server,json=authServer,proto3" json:"auth_server,omitempty"`
}

func (x *UpdateAppRequest) Reset() {
	*x = UpdateAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppRequest) ProtoMessage() {}

func (x *UpdateAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppRequest.ProtoReflect.Descriptor instead.
func (*UpdateAppRequest) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAppRequest) GetApp() *Application {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *UpdateAppRequest) GetAuthServer() *AuthServer {
	if x != nil {
		return x.AuthServer
	}
	return nil
}

type UpdateAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAppResponse) Reset() {
	*x = UpdateAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppResponse) ProtoMessage() {}

func (x *UpdateAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppResponse.ProtoReflect.Descriptor instead.
func (*UpdateAppResponse) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{8}
}

type FindAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size      uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Companies []uint32 `protobuf:"varint,4,rep,packed,name=companies,proto3" json:"companies,omitempty"`
	States    []uint32 `protobuf:"varint,5,rep,packed,name=states,proto3" json:"states,omitempty"`
}

func (x *FindAppRequest) Reset() {
	*x = FindAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAppRequest) ProtoMessage() {}

func (x *FindAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAppRequest.ProtoReflect.Descriptor instead.
func (*FindAppRequest) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{9}
}

func (x *FindAppRequest) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *FindAppRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FindAppRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindAppRequest) GetCompanies() []uint32 {
	if x != nil {
		return x.Companies
	}
	return nil
}

func (x *FindAppRequest) GetStates() []uint32 {
	if x != nil {
		return x.States
	}
	return nil
}

type FindAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size  uint32         `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Page  uint32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Total uint32         `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	Apps  []*Application `protobuf:"bytes,5,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *FindAppResponse) Reset() {
	*x = FindAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAppResponse) ProtoMessage() {}

func (x *FindAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAppResponse.ProtoReflect.Descriptor instead.
func (*FindAppResponse) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{10}
}

func (x *FindAppResponse) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *FindAppResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FindAppResponse) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindAppResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FindAppResponse) GetApps() []*Application {
	if x != nil {
		return x.Apps
	}
	return nil
}

type GetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAppRequest) Reset() {
	*x = GetAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRequest) ProtoMessage() {}

func (x *GetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRequest.ProtoReflect.Descriptor instead.
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{11}
}

func (x *GetAppRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App        *Application `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	AuthServer *AuthServer  `protobuf:"bytes,2,opt,name=auth_server,json=authServer,proto3" json:"auth_server,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_application_v1_application_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_application_v1_application_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_api_application_v1_application_proto_rawDescGZIP(), []int{12}
}

func (x *GetAppResponse) GetApp() *Application {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *GetAppResponse) GetAuthServer() *AuthServer {
	if x != nil {
		return x.AuthServer
	}
	return nil
}

var File_api_application_v1_application_proto protoreflect.FileDescriptor

var file_api_application_v1_application_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xb2, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x34,
	0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x03, 0x61, 0x70, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x13,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x78, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x61,
	0x70, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2c, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x1f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x38, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x81, 0x03, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x1d, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_application_v1_application_proto_rawDescOnce sync.Once
	file_api_application_v1_application_proto_rawDescData = file_api_application_v1_application_proto_rawDesc
)

func file_api_application_v1_application_proto_rawDescGZIP() []byte {
	file_api_application_v1_application_proto_rawDescOnce.Do(func() {
		file_api_application_v1_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_application_v1_application_proto_rawDescData)
	})
	return file_api_application_v1_application_proto_rawDescData
}

var file_api_application_v1_application_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_application_v1_application_proto_goTypes = []interface{}{
	(*Application)(nil),       // 0: application.Application
	(*AuthServer)(nil),        // 1: application.AuthServer
	(*Pod)(nil),               // 2: application.Pod
	(*CreateAppRequest)(nil),  // 3: application.CreateAppRequest
	(*CreateAppResponse)(nil), // 4: application.CreateAppResponse
	(*DeleteAppRequest)(nil),  // 5: application.DeleteAppRequest
	(*DeleteAppResponse)(nil), // 6: application.DeleteAppResponse
	(*UpdateAppRequest)(nil),  // 7: application.UpdateAppRequest
	(*UpdateAppResponse)(nil), // 8: application.UpdateAppResponse
	(*FindAppRequest)(nil),    // 9: application.FindAppRequest
	(*FindAppResponse)(nil),   // 10: application.FindAppResponse
	(*GetAppRequest)(nil),     // 11: application.GetAppRequest
	(*GetAppResponse)(nil),    // 12: application.GetAppResponse
}
var file_api_application_v1_application_proto_depIdxs = []int32{
	0,  // 0: application.CreateAppRequest.app:type_name -> application.Application
	1,  // 1: application.CreateAppRequest.auth_server:type_name -> application.AuthServer
	0,  // 2: application.UpdateAppRequest.app:type_name -> application.Application
	1,  // 3: application.UpdateAppRequest.auth_server:type_name -> application.AuthServer
	0,  // 4: application.FindAppResponse.apps:type_name -> application.Application
	0,  // 5: application.GetAppResponse.app:type_name -> application.Application
	1,  // 6: application.GetAppResponse.auth_server:type_name -> application.AuthServer
	3,  // 7: application.ApplicationService.CreateApp:input_type -> application.CreateAppRequest
	5,  // 8: application.ApplicationService.DeleteApp:input_type -> application.DeleteAppRequest
	7,  // 9: application.ApplicationService.UpdateApp:input_type -> application.UpdateAppRequest
	9,  // 10: application.ApplicationService.FindApp:input_type -> application.FindAppRequest
	11, // 11: application.ApplicationService.GetApp:input_type -> application.GetAppRequest
	4,  // 12: application.ApplicationService.CreateApp:output_type -> application.CreateAppResponse
	6,  // 13: application.ApplicationService.DeleteApp:output_type -> application.DeleteAppResponse
	8,  // 14: application.ApplicationService.UpdateApp:output_type -> application.UpdateAppResponse
	10, // 15: application.ApplicationService.FindApp:output_type -> application.FindAppResponse
	12, // 16: application.ApplicationService.GetApp:output_type -> application.GetAppResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_application_v1_application_proto_init() }
func file_api_application_v1_application_proto_init() {
	if File_api_application_v1_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_application_v1_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_api_application_v1_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthServer); i {
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
		file_api_application_v1_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
		file_api_application_v1_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRequest); i {
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
		file_api_application_v1_application_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppResponse); i {
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
		file_api_application_v1_application_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppRequest); i {
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
		file_api_application_v1_application_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppResponse); i {
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
		file_api_application_v1_application_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppRequest); i {
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
		file_api_application_v1_application_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppResponse); i {
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
		file_api_application_v1_application_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAppRequest); i {
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
		file_api_application_v1_application_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAppResponse); i {
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
		file_api_application_v1_application_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRequest); i {
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
		file_api_application_v1_application_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppResponse); i {
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
			RawDescriptor: file_api_application_v1_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_application_v1_application_proto_goTypes,
		DependencyIndexes: file_api_application_v1_application_proto_depIdxs,
		MessageInfos:      file_api_application_v1_application_proto_msgTypes,
	}.Build()
	File_api_application_v1_application_proto = out.File
	file_api_application_v1_application_proto_rawDesc = nil
	file_api_application_v1_application_proto_goTypes = nil
	file_api_application_v1_application_proto_depIdxs = nil
}