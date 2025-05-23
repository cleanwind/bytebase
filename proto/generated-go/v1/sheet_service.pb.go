// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: v1/sheet_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of the SheetPayload.
type SheetPayload_Type int32

const (
	SheetPayload_TYPE_UNSPECIFIED SheetPayload_Type = 0
	SheetPayload_SCHEMA_DESIGN    SheetPayload_Type = 1
)

// Enum value maps for SheetPayload_Type.
var (
	SheetPayload_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "SCHEMA_DESIGN",
	}
	SheetPayload_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"SCHEMA_DESIGN":    1,
	}
)

func (x SheetPayload_Type) Enum() *SheetPayload_Type {
	p := new(SheetPayload_Type)
	*p = x
	return p
}

func (x SheetPayload_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SheetPayload_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_sheet_service_proto_enumTypes[0].Descriptor()
}

func (SheetPayload_Type) Type() protoreflect.EnumType {
	return &file_v1_sheet_service_proto_enumTypes[0]
}

func (x SheetPayload_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SheetPayload_Type.Descriptor instead.
func (SheetPayload_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{6, 0}
}

type CreateSheetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The parent resource where this sheet will be created.
	// Format: projects/{project}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The sheet to create.
	Sheet         *Sheet `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSheetRequest) Reset() {
	*x = CreateSheetRequest{}
	mi := &file_v1_sheet_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSheetRequest) ProtoMessage() {}

func (x *CreateSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSheetRequest.ProtoReflect.Descriptor instead.
func (*CreateSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSheetRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateSheetRequest) GetSheet() *Sheet {
	if x != nil {
		return x.Sheet
	}
	return nil
}

type BatchCreateSheetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The parent resource where all sheets will be created.
	// Format: projects/{project}
	Parent        string                `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Requests      []*CreateSheetRequest `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchCreateSheetRequest) Reset() {
	*x = BatchCreateSheetRequest{}
	mi := &file_v1_sheet_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchCreateSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateSheetRequest) ProtoMessage() {}

func (x *BatchCreateSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateSheetRequest.ProtoReflect.Descriptor instead.
func (*BatchCreateSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{1}
}

func (x *BatchCreateSheetRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *BatchCreateSheetRequest) GetRequests() []*CreateSheetRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type BatchCreateSheetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sheets        []*Sheet               `protobuf:"bytes,1,rep,name=sheets,proto3" json:"sheets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchCreateSheetResponse) Reset() {
	*x = BatchCreateSheetResponse{}
	mi := &file_v1_sheet_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchCreateSheetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateSheetResponse) ProtoMessage() {}

func (x *BatchCreateSheetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateSheetResponse.ProtoReflect.Descriptor instead.
func (*BatchCreateSheetResponse) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{2}
}

func (x *BatchCreateSheetResponse) GetSheets() []*Sheet {
	if x != nil {
		return x.Sheets
	}
	return nil
}

type GetSheetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the sheet to retrieve.
	// Format: projects/{project}/sheets/{sheet}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// By default, the content of the sheet is cut off, set the `raw` to true to retrieve the full content.
	Raw           bool `protobuf:"varint,2,opt,name=raw,proto3" json:"raw,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSheetRequest) Reset() {
	*x = GetSheetRequest{}
	mi := &file_v1_sheet_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSheetRequest) ProtoMessage() {}

func (x *GetSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSheetRequest.ProtoReflect.Descriptor instead.
func (*GetSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetSheetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSheetRequest) GetRaw() bool {
	if x != nil {
		return x.Raw
	}
	return false
}

type UpdateSheetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The sheet to update.
	//
	// The sheet's `name` field is used to identify the sheet to update.
	// Format: projects/{project}/sheets/{sheet}
	Sheet *Sheet `protobuf:"bytes,1,opt,name=sheet,proto3" json:"sheet,omitempty"`
	// The list of fields to be updated.
	// Fields are specified relative to the sheet.
	// (e.g. `title`, `statement`; *not* `sheet.title` or `sheet.statement`)
	// Only support update the following fields for now:
	// - `title`
	// - `statement`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSheetRequest) Reset() {
	*x = UpdateSheetRequest{}
	mi := &file_v1_sheet_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSheetRequest) ProtoMessage() {}

func (x *UpdateSheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSheetRequest.ProtoReflect.Descriptor instead.
func (*UpdateSheetRequest) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSheetRequest) GetSheet() *Sheet {
	if x != nil {
		return x.Sheet
	}
	return nil
}

func (x *UpdateSheetRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Sheet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the sheet resource, generated by the server.
	// Canonical parent is project.
	// Format: projects/{project}/sheets/{sheet}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The title of the sheet.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// The creator of the Sheet.
	// Format: users/{email}
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	// The create time of the sheet.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the sheet.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The content of the sheet.
	// By default, it will be cut off, if it doesn't match the `content_size`, you can
	// set the `raw` to true in GetSheet request to retrieve the full content.
	Content []byte `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	// content_size is the full size of the content, may not match the size of the `content` field.
	ContentSize int64         `protobuf:"varint,8,opt,name=content_size,json=contentSize,proto3" json:"content_size,omitempty"`
	Payload     *SheetPayload `protobuf:"bytes,13,opt,name=payload,proto3" json:"payload,omitempty"`
	// The SQL dialect.
	Engine        Engine `protobuf:"varint,14,opt,name=engine,proto3,enum=bytebase.v1.Engine" json:"engine,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sheet) Reset() {
	*x = Sheet{}
	mi := &file_v1_sheet_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sheet) ProtoMessage() {}

func (x *Sheet) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sheet.ProtoReflect.Descriptor instead.
func (*Sheet) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{5}
}

func (x *Sheet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sheet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Sheet) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Sheet) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Sheet) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Sheet) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Sheet) GetContentSize() int64 {
	if x != nil {
		return x.ContentSize
	}
	return 0
}

func (x *Sheet) GetPayload() *SheetPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Sheet) GetEngine() Engine {
	if x != nil {
		return x.Engine
	}
	return Engine_ENGINE_UNSPECIFIED
}

type SheetPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Type  SheetPayload_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.v1.SheetPayload_Type" json:"type,omitempty"`
	// The start and end position of each command in the sheet statement.
	Commands      []*SheetCommand `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SheetPayload) Reset() {
	*x = SheetPayload{}
	mi := &file_v1_sheet_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SheetPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetPayload) ProtoMessage() {}

func (x *SheetPayload) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetPayload.ProtoReflect.Descriptor instead.
func (*SheetPayload) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{6}
}

func (x *SheetPayload) GetType() SheetPayload_Type {
	if x != nil {
		return x.Type
	}
	return SheetPayload_TYPE_UNSPECIFIED
}

func (x *SheetPayload) GetCommands() []*SheetCommand {
	if x != nil {
		return x.Commands
	}
	return nil
}

type SheetCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         int32                  `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End           int32                  `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SheetCommand) Reset() {
	*x = SheetCommand{}
	mi := &file_v1_sheet_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SheetCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetCommand) ProtoMessage() {}

func (x *SheetCommand) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sheet_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetCommand.ProtoReflect.Descriptor instead.
func (*SheetCommand) Descriptor() ([]byte, []int) {
	return file_v1_sheet_service_proto_rawDescGZIP(), []int{7}
}

func (x *SheetCommand) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SheetCommand) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_v1_sheet_service_proto protoreflect.FileDescriptor

var file_v1_sheet_service_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x16,
	0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2e,
	0x0a, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65,
	0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x93,
	0x01, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xe2, 0x41, 0x01, 0x02,
	0xfa, 0x41, 0x16, 0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x52, 0x06, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xe2,
	0x41, 0x01, 0x02, 0xfa, 0x41, 0x14, 0x0a, 0x12, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72,
	0x61, 0x77, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x02, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xd1, 0x03, 0x0a, 0x05, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05,
	0xe2, 0x41, 0x02, 0x02, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x02, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a,
	0xea, 0x41, 0x37, 0x0a, 0x12, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x7d, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x44,
	0x45, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x01, 0x22, 0x36, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x32,
	0x83, 0x05, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x98, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x12, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x22, 0x54, 0xda, 0x41, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x2c, 0x73, 0x68, 0x65, 0x65, 0x74, 0x8a, 0xea, 0x30, 0x10, 0x62, 0x62, 0x2e, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x3a, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x1e, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x12, 0xae, 0x01, 0x0a, 0x10,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x12, 0x24, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x8a,
	0xea, 0x30, 0x10, 0x62, 0x62, 0x2e, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a,
	0x22, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73,
	0x3a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x80, 0x01, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x22, 0x42, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x8a, 0xea, 0x30, 0x0d, 0x62, 0x62, 0x2e, 0x73, 0x68, 0x65, 0x65, 0x74,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x90, 0xea, 0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12,
	0xa3, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12,
	0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x22, 0x5f, 0xda, 0x41, 0x11, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x8a, 0xea, 0x30, 0x10, 0x62, 0x62,
	0x2e, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x90, 0xea,
	0x30, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x32,
	0x24, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1_sheet_service_proto_rawDescOnce sync.Once
	file_v1_sheet_service_proto_rawDescData []byte
)

func file_v1_sheet_service_proto_rawDescGZIP() []byte {
	file_v1_sheet_service_proto_rawDescOnce.Do(func() {
		file_v1_sheet_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_sheet_service_proto_rawDesc), len(file_v1_sheet_service_proto_rawDesc)))
	})
	return file_v1_sheet_service_proto_rawDescData
}

var file_v1_sheet_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_sheet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_sheet_service_proto_goTypes = []any{
	(SheetPayload_Type)(0),           // 0: bytebase.v1.SheetPayload.Type
	(*CreateSheetRequest)(nil),       // 1: bytebase.v1.CreateSheetRequest
	(*BatchCreateSheetRequest)(nil),  // 2: bytebase.v1.BatchCreateSheetRequest
	(*BatchCreateSheetResponse)(nil), // 3: bytebase.v1.BatchCreateSheetResponse
	(*GetSheetRequest)(nil),          // 4: bytebase.v1.GetSheetRequest
	(*UpdateSheetRequest)(nil),       // 5: bytebase.v1.UpdateSheetRequest
	(*Sheet)(nil),                    // 6: bytebase.v1.Sheet
	(*SheetPayload)(nil),             // 7: bytebase.v1.SheetPayload
	(*SheetCommand)(nil),             // 8: bytebase.v1.SheetCommand
	(*fieldmaskpb.FieldMask)(nil),    // 9: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(Engine)(0),                      // 11: bytebase.v1.Engine
}
var file_v1_sheet_service_proto_depIdxs = []int32{
	6,  // 0: bytebase.v1.CreateSheetRequest.sheet:type_name -> bytebase.v1.Sheet
	1,  // 1: bytebase.v1.BatchCreateSheetRequest.requests:type_name -> bytebase.v1.CreateSheetRequest
	6,  // 2: bytebase.v1.BatchCreateSheetResponse.sheets:type_name -> bytebase.v1.Sheet
	6,  // 3: bytebase.v1.UpdateSheetRequest.sheet:type_name -> bytebase.v1.Sheet
	9,  // 4: bytebase.v1.UpdateSheetRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 5: bytebase.v1.Sheet.create_time:type_name -> google.protobuf.Timestamp
	10, // 6: bytebase.v1.Sheet.update_time:type_name -> google.protobuf.Timestamp
	7,  // 7: bytebase.v1.Sheet.payload:type_name -> bytebase.v1.SheetPayload
	11, // 8: bytebase.v1.Sheet.engine:type_name -> bytebase.v1.Engine
	0,  // 9: bytebase.v1.SheetPayload.type:type_name -> bytebase.v1.SheetPayload.Type
	8,  // 10: bytebase.v1.SheetPayload.commands:type_name -> bytebase.v1.SheetCommand
	1,  // 11: bytebase.v1.SheetService.CreateSheet:input_type -> bytebase.v1.CreateSheetRequest
	2,  // 12: bytebase.v1.SheetService.BatchCreateSheet:input_type -> bytebase.v1.BatchCreateSheetRequest
	4,  // 13: bytebase.v1.SheetService.GetSheet:input_type -> bytebase.v1.GetSheetRequest
	5,  // 14: bytebase.v1.SheetService.UpdateSheet:input_type -> bytebase.v1.UpdateSheetRequest
	6,  // 15: bytebase.v1.SheetService.CreateSheet:output_type -> bytebase.v1.Sheet
	3,  // 16: bytebase.v1.SheetService.BatchCreateSheet:output_type -> bytebase.v1.BatchCreateSheetResponse
	6,  // 17: bytebase.v1.SheetService.GetSheet:output_type -> bytebase.v1.Sheet
	6,  // 18: bytebase.v1.SheetService.UpdateSheet:output_type -> bytebase.v1.Sheet
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_v1_sheet_service_proto_init() }
func file_v1_sheet_service_proto_init() {
	if File_v1_sheet_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_sheet_service_proto_rawDesc), len(file_v1_sheet_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_sheet_service_proto_goTypes,
		DependencyIndexes: file_v1_sheet_service_proto_depIdxs,
		EnumInfos:         file_v1_sheet_service_proto_enumTypes,
		MessageInfos:      file_v1_sheet_service_proto_msgTypes,
	}.Build()
	File_v1_sheet_service_proto = out.File
	file_v1_sheet_service_proto_goTypes = nil
	file_v1_sheet_service_proto_depIdxs = nil
}
