// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: api/device/v1/device.proto

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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeviceId  string `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Brand     string `protobuf:"bytes,6,opt,name=brand,proto3" json:"brand,omitempty"`
	Type      string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Protocol  string `protobuf:"bytes,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip        string `protobuf:"bytes,9,opt,name=ip,proto3" json:"ip,omitempty"`
	Port      string `protobuf:"bytes,10,opt,name=port,proto3" json:"port,omitempty"`
	Region    uint64 `protobuf:"varint,11,opt,name=region,proto3" json:"region,omitempty"`
	Attrs     []byte `protobuf:"bytes,12,opt,name=attrs,proto3" json:"attrs,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Device) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Device) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Device) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Device) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Device) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Device) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Device) GetRegion() uint64 {
	if x != nil {
		return x.Region
	}
	return 0
}

func (x *Device) GetAttrs() []byte {
	if x != nil {
		return x.Attrs
	}
	return nil
}

type CreateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table  string  `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Device *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateDeviceRequest) Reset() {
	*x = CreateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceRequest) ProtoMessage() {}

func (x *CreateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeviceRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *CreateDeviceRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type CreateDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDeviceResponse) Reset() {
	*x = CreateDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceResponse) ProtoMessage() {}

func (x *CreateDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceResponse.ProtoReflect.Descriptor instead.
func (*CreateDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{2}
}

type DeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Id    uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDeviceRequest) Reset() {
	*x = DeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceRequest) ProtoMessage() {}

func (x *DeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDeviceRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *DeleteDeviceRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDeviceResponse) Reset() {
	*x = DeleteDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceResponse) ProtoMessage() {}

func (x *DeleteDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{4}
}

type UpdateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table  string  `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Device *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *UpdateDeviceRequest) Reset() {
	*x = UpdateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceRequest) ProtoMessage() {}

func (x *UpdateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDeviceRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *UpdateDeviceRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type UpdateDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDeviceResponse) Reset() {
	*x = UpdateDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceResponse) ProtoMessage() {}

func (x *UpdateDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{6}
}

type SelectDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table   string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Index   uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Size    uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Keyword string `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Region  uint64 `protobuf:"varint,5,opt,name=region,proto3" json:"region,omitempty"`
	Cons    []byte `protobuf:"bytes,6,opt,name=cons,proto3" json:"cons,omitempty"`
}

func (x *SelectDeviceRequest) Reset() {
	*x = SelectDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectDeviceRequest) ProtoMessage() {}

func (x *SelectDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectDeviceRequest.ProtoReflect.Descriptor instead.
func (*SelectDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{7}
}

func (x *SelectDeviceRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *SelectDeviceRequest) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SelectDeviceRequest) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SelectDeviceRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SelectDeviceRequest) GetRegion() uint64 {
	if x != nil {
		return x.Region
	}
	return 0
}

func (x *SelectDeviceRequest) GetCons() []byte {
	if x != nil {
		return x.Cons
	}
	return nil
}

type SelectDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   uint64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Page    uint64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size    uint64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Total   uint64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	Devices []*Device `protobuf:"bytes,5,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *SelectDeviceResponse) Reset() {
	*x = SelectDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectDeviceResponse) ProtoMessage() {}

func (x *SelectDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectDeviceResponse.ProtoReflect.Descriptor instead.
func (*SelectDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{8}
}

func (x *SelectDeviceResponse) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SelectDeviceResponse) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SelectDeviceResponse) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SelectDeviceResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SelectDeviceResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type GetDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Id    uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeviceRequest) Reset() {
	*x = GetDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceRequest) ProtoMessage() {}

func (x *GetDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{9}
}

func (x *GetDeviceRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *GetDeviceRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *GetDeviceResponse) Reset() {
	*x = GetDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceResponse) ProtoMessage() {}

func (x *GetDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{10}
}

func (x *GetDeviceResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type BatchDeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Id    []uint64 `protobuf:"varint,2,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *BatchDeleteDeviceRequest) Reset() {
	*x = BatchDeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteDeviceRequest) ProtoMessage() {}

func (x *BatchDeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*BatchDeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{11}
}

func (x *BatchDeleteDeviceRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *BatchDeleteDeviceRequest) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type BatchDeleteDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BatchDeleteDeviceResponse) Reset() {
	*x = BatchDeleteDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_v1_device_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteDeviceResponse) ProtoMessage() {}

func (x *BatchDeleteDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_v1_device_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteDeviceResponse.ProtoReflect.Descriptor instead.
func (*BatchDeleteDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_v1_device_proto_rawDescGZIP(), []int{12}
}

var File_api_device_v1_device_proto protoreflect.FileDescriptor

var file_api_device_v1_device_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x22, 0x53, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x16, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63,
	0x6f, 0x6e, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x28, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x40, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xd7, 0x03, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_device_v1_device_proto_rawDescOnce sync.Once
	file_api_device_v1_device_proto_rawDescData = file_api_device_v1_device_proto_rawDesc
)

func file_api_device_v1_device_proto_rawDescGZIP() []byte {
	file_api_device_v1_device_proto_rawDescOnce.Do(func() {
		file_api_device_v1_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_device_v1_device_proto_rawDescData)
	})
	return file_api_device_v1_device_proto_rawDescData
}

var file_api_device_v1_device_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_device_v1_device_proto_goTypes = []interface{}{
	(*Device)(nil),                    // 0: device.Device
	(*CreateDeviceRequest)(nil),       // 1: device.CreateDeviceRequest
	(*CreateDeviceResponse)(nil),      // 2: device.CreateDeviceResponse
	(*DeleteDeviceRequest)(nil),       // 3: device.DeleteDeviceRequest
	(*DeleteDeviceResponse)(nil),      // 4: device.DeleteDeviceResponse
	(*UpdateDeviceRequest)(nil),       // 5: device.UpdateDeviceRequest
	(*UpdateDeviceResponse)(nil),      // 6: device.UpdateDeviceResponse
	(*SelectDeviceRequest)(nil),       // 7: device.SelectDeviceRequest
	(*SelectDeviceResponse)(nil),      // 8: device.SelectDeviceResponse
	(*GetDeviceRequest)(nil),          // 9: device.GetDeviceRequest
	(*GetDeviceResponse)(nil),         // 10: device.GetDeviceResponse
	(*BatchDeleteDeviceRequest)(nil),  // 11: device.BatchDeleteDeviceRequest
	(*BatchDeleteDeviceResponse)(nil), // 12: device.BatchDeleteDeviceResponse
}
var file_api_device_v1_device_proto_depIdxs = []int32{
	0,  // 0: device.CreateDeviceRequest.device:type_name -> device.Device
	0,  // 1: device.UpdateDeviceRequest.device:type_name -> device.Device
	0,  // 2: device.SelectDeviceResponse.devices:type_name -> device.Device
	0,  // 3: device.GetDeviceResponse.device:type_name -> device.Device
	1,  // 4: device.DeviceService.CreateDevice:input_type -> device.CreateDeviceRequest
	3,  // 5: device.DeviceService.DeleteDevice:input_type -> device.DeleteDeviceRequest
	5,  // 6: device.DeviceService.UpdateDevice:input_type -> device.UpdateDeviceRequest
	7,  // 7: device.DeviceService.SelectDevice:input_type -> device.SelectDeviceRequest
	9,  // 8: device.DeviceService.GetDevice:input_type -> device.GetDeviceRequest
	11, // 9: device.DeviceService.BatchDeleteDevice:input_type -> device.BatchDeleteDeviceRequest
	2,  // 10: device.DeviceService.CreateDevice:output_type -> device.CreateDeviceResponse
	4,  // 11: device.DeviceService.DeleteDevice:output_type -> device.DeleteDeviceResponse
	6,  // 12: device.DeviceService.UpdateDevice:output_type -> device.UpdateDeviceResponse
	8,  // 13: device.DeviceService.SelectDevice:output_type -> device.SelectDeviceResponse
	10, // 14: device.DeviceService.GetDevice:output_type -> device.GetDeviceResponse
	12, // 15: device.DeviceService.BatchDeleteDevice:output_type -> device.BatchDeleteDeviceResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_device_v1_device_proto_init() }
func file_api_device_v1_device_proto_init() {
	if File_api_device_v1_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_device_v1_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_api_device_v1_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceRequest); i {
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
		file_api_device_v1_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceResponse); i {
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
		file_api_device_v1_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceRequest); i {
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
		file_api_device_v1_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceResponse); i {
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
		file_api_device_v1_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceRequest); i {
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
		file_api_device_v1_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceResponse); i {
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
		file_api_device_v1_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectDeviceRequest); i {
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
		file_api_device_v1_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectDeviceResponse); i {
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
		file_api_device_v1_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceRequest); i {
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
		file_api_device_v1_device_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceResponse); i {
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
		file_api_device_v1_device_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteDeviceRequest); i {
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
		file_api_device_v1_device_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteDeviceResponse); i {
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
			RawDescriptor: file_api_device_v1_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_device_v1_device_proto_goTypes,
		DependencyIndexes: file_api_device_v1_device_proto_depIdxs,
		MessageInfos:      file_api_device_v1_device_proto_msgTypes,
	}.Build()
	File_api_device_v1_device_proto = out.File
	file_api_device_v1_device_proto_rawDesc = nil
	file_api_device_v1_device_proto_goTypes = nil
	file_api_device_v1_device_proto_depIdxs = nil
}
