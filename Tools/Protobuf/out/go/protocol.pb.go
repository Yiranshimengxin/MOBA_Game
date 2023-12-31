// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.1
// source: protocol.proto

package pb

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

type MsgID int32

const (
	MsgID_INVALID             MsgID = 0
	MsgID_LOGIN_REQ           MsgID = 1000 //客户端往服务器发送的登录id
	MsgID_LOGIN_RSP           MsgID = 1001 //服务器往客户端返回的登录id
	MsgID_MATCH_REQ           MsgID = 1002
	MsgID_MATCH_RSP           MsgID = 1003
	MsgID_NOTIFY_CONFIRM      MsgID = 1004
	MsgID_CONFIRM_REQ         MsgID = 1005
	MsgID_NOTIFY_SELECT       MsgID = 1006
	MsgID_SELECT_REQ          MsgID = 1007
	MsgID_NOTIFY_LOAD         MsgID = 1008
	MsgID_LOAD_PROGRESS_REQ   MsgID = 1009
	MsgID_NOTIFY_PROGRESS     MsgID = 1010
	MsgID_BATTLE_START_REQ    MsgID = 1011
	MsgID_NOTIFY_BATTLE_START MsgID = 1012
	MsgID_QPERATE_REQ         MsgID = 1013
	MsgID_NOTIFY_OPERATE      MsgID = 1014
	MsgID_BATTLE_END_REQ      MsgID = 1015
	MsgID_NOTIFY_BATTLE_END   MsgID = 1016
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0:    "INVALID",
		1000: "LOGIN_REQ",
		1001: "LOGIN_RSP",
		1002: "MATCH_REQ",
		1003: "MATCH_RSP",
		1004: "NOTIFY_CONFIRM",
		1005: "CONFIRM_REQ",
		1006: "NOTIFY_SELECT",
		1007: "SELECT_REQ",
		1008: "NOTIFY_LOAD",
		1009: "LOAD_PROGRESS_REQ",
		1010: "NOTIFY_PROGRESS",
		1011: "BATTLE_START_REQ",
		1012: "NOTIFY_BATTLE_START",
		1013: "QPERATE_REQ",
		1014: "NOTIFY_OPERATE",
		1015: "BATTLE_END_REQ",
		1016: "NOTIFY_BATTLE_END",
	}
	MsgID_value = map[string]int32{
		"INVALID":             0,
		"LOGIN_REQ":           1000,
		"LOGIN_RSP":           1001,
		"MATCH_REQ":           1002,
		"MATCH_RSP":           1003,
		"NOTIFY_CONFIRM":      1004,
		"CONFIRM_REQ":         1005,
		"NOTIFY_SELECT":       1006,
		"SELECT_REQ":          1007,
		"NOTIFY_LOAD":         1008,
		"LOAD_PROGRESS_REQ":   1009,
		"NOTIFY_PROGRESS":     1010,
		"BATTLE_START_REQ":    1011,
		"NOTIFY_BATTLE_START": 1012,
		"QPERATE_REQ":         1013,
		"NOTIFY_OPERATE":      1014,
		"BATTLE_END_REQ":      1015,
		"NOTIFY_BATTLE_END":   1016,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_protocol_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

type PBMatchReqPvpType int32

const (
	PBMatchReq_PVP_1V1 PBMatchReqPvpType = 0
	PBMatchReq_PVP_2V2 PBMatchReqPvpType = 1
	PBMatchReq_PVP_5V5 PBMatchReqPvpType = 2
)

// Enum value maps for PBMatchReqPvpType.
var (
	PBMatchReqPvpType_name = map[int32]string{
		0: "PVP_1V1",
		1: "PVP_2V2",
		2: "PVP_5V5",
	}
	PBMatchReqPvpType_value = map[string]int32{
		"PVP_1V1": 0,
		"PVP_2V2": 1,
		"PVP_5V5": 2,
	}
)

func (x PBMatchReqPvpType) Enum() *PBMatchReqPvpType {
	p := new(PBMatchReqPvpType)
	*p = x
	return p
}

func (x PBMatchReqPvpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PBMatchReqPvpType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_enumTypes[1].Descriptor()
}

func (PBMatchReqPvpType) Type() protoreflect.EnumType {
	return &file_protocol_proto_enumTypes[1]
}

func (x PBMatchReqPvpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PBMatchReqPvpType.Descriptor instead.
func (PBMatchReqPvpType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{3, 0}
}

type Head struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    //协议id
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` //协议里的内容
}

func (x *Head) Reset() {
	*x = Head{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Head) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Head) ProtoMessage() {}

func (x *Head) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Head.ProtoReflect.Descriptor instead.
func (*Head) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Head) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Head) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PBLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *PBLoginReq) Reset() {
	*x = PBLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBLoginReq) ProtoMessage() {}

func (x *PBLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBLoginReq.ProtoReflect.Descriptor instead.
func (*PBLoginReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *PBLoginReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type PBLoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Uid     uint64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *PBLoginRsp) Reset() {
	*x = PBLoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBLoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBLoginRsp) ProtoMessage() {}

func (x *PBLoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBLoginRsp.ProtoReflect.Descriptor instead.
func (*PBLoginRsp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *PBLoginRsp) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PBLoginRsp) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type PBMatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PBMatchReqPvpType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PBMatchReqPvpType" json:"type,omitempty"`
}

func (x *PBMatchReq) Reset() {
	*x = PBMatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBMatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBMatchReq) ProtoMessage() {}

func (x *PBMatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBMatchReq.ProtoReflect.Descriptor instead.
func (*PBMatchReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *PBMatchReq) GetType() PBMatchReqPvpType {
	if x != nil {
		return x.Type
	}
	return PBMatchReq_PVP_1V1
}

type PBMtachRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreTime int32 `protobuf:"varint,1,opt,name=preTime,proto3" json:"preTime,omitempty"`
}

func (x *PBMtachRsp) Reset() {
	*x = PBMtachRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBMtachRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBMtachRsp) ProtoMessage() {}

func (x *PBMtachRsp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBMtachRsp.ProtoReflect.Descriptor instead.
func (*PBMtachRsp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *PBMtachRsp) GetPreTime() int32 {
	if x != nil {
		return x.PreTime
	}
	return 0
}

type PBConfirmReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PBConfirmReq) Reset() {
	*x = PBConfirmReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBConfirmReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBConfirmReq) ProtoMessage() {}

func (x *PBConfirmReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBConfirmReq.ProtoReflect.Descriptor instead.
func (*PBConfirmReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{5}
}

type PBConfirmData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IconIndex   int32 `protobuf:"varint,1,opt,name=iconIndex,proto3" json:"iconIndex,omitempty"`
	ConfirmDone bool  `protobuf:"varint,2,opt,name=confirmDone,proto3" json:"confirmDone,omitempty"`
}

func (x *PBConfirmData) Reset() {
	*x = PBConfirmData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBConfirmData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBConfirmData) ProtoMessage() {}

func (x *PBConfirmData) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBConfirmData.ProtoReflect.Descriptor instead.
func (*PBConfirmData) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *PBConfirmData) GetIconIndex() int32 {
	if x != nil {
		return x.IconIndex
	}
	return 0
}

func (x *PBConfirmData) GetConfirmDone() bool {
	if x != nil {
		return x.ConfirmDone
	}
	return false
}

type PBnotifyConfirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId     int32            `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Dismiss    bool             `protobuf:"varint,2,opt,name=dismiss,proto3" json:"dismiss,omitempty"`
	ConfirmArr []*PBConfirmData `protobuf:"bytes,3,rep,name=confirmArr,proto3" json:"confirmArr,omitempty"`
}

func (x *PBnotifyConfirm) Reset() {
	*x = PBnotifyConfirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBnotifyConfirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBnotifyConfirm) ProtoMessage() {}

func (x *PBnotifyConfirm) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBnotifyConfirm.ProtoReflect.Descriptor instead.
func (*PBnotifyConfirm) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *PBnotifyConfirm) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PBnotifyConfirm) GetDismiss() bool {
	if x != nil {
		return x.Dismiss
	}
	return false
}

func (x *PBnotifyConfirm) GetConfirmArr() []*PBConfirmData {
	if x != nil {
		return x.ConfirmArr
	}
	return nil
}

type PBSelectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int32 `protobuf:"varint,1,opt,name=heroId,proto3" json:"heroId,omitempty"`
}

func (x *PBSelectReq) Reset() {
	*x = PBSelectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBSelectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBSelectReq) ProtoMessage() {}

func (x *PBSelectReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBSelectReq.ProtoReflect.Descriptor instead.
func (*PBSelectReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{8}
}

func (x *PBSelectReq) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x2a, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x1e, 0x0a, 0x0a, 0x50, 0x42, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0x38, 0x0a, 0x0a, 0x50, 0x42, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x0a, 0x50, 0x42,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x42, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x76, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x30, 0x0a, 0x07, 0x70, 0x76, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x56, 0x50, 0x5f, 0x31, 0x56, 0x31, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x56, 0x50, 0x5f, 0x32, 0x56, 0x32, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x56, 0x50,
	0x5f, 0x35, 0x56, 0x35, 0x10, 0x02, 0x22, 0x26, 0x0a, 0x0a, 0x50, 0x42, 0x4d, 0x74, 0x61, 0x63,
	0x68, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x50, 0x42, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x22, 0x4f,
	0x0a, 0x0d, 0x50, 0x42, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x6f, 0x6e, 0x65, 0x22,
	0x76, 0x0a, 0x0f, 0x50, 0x42, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73,
	0x6d, 0x69, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41,
	0x72, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x42,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x41, 0x72, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x42, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x2a, 0xe5,
	0x02, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52,
	0x45, 0x51, 0x10, 0xe8, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52,
	0x53, 0x50, 0x10, 0xe9, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52,
	0x45, 0x51, 0x10, 0xea, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52,
	0x53, 0x50, 0x10, 0xeb, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x10, 0xec, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x52, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xed, 0x07, 0x12, 0x12, 0x0a, 0x0d,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x10, 0xee, 0x07,
	0x12, 0x0f, 0x0a, 0x0a, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xef,
	0x07, 0x12, 0x10, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4c, 0x4f, 0x41, 0x44,
	0x10, 0xf0, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xf1, 0x07, 0x12, 0x14, 0x0a, 0x0f, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0xf2,
	0x07, 0x12, 0x15, 0x0a, 0x10, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xf3, 0x07, 0x12, 0x18, 0x0a, 0x13, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x59, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10,
	0xf4, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x51, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x10, 0xf5, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10, 0xf6, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x42, 0x41, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xf7, 0x07, 0x12, 0x16,
	0x0a, 0x11, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f,
	0x45, 0x4e, 0x44, 0x10, 0xf8, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protocol_proto_goTypes = []interface{}{
	(MsgID)(0),              // 0: pb.MsgID
	(PBMatchReqPvpType)(0),  // 1: pb.PBMatchReq.pvpType
	(*Head)(nil),            // 2: pb.Head
	(*PBLoginReq)(nil),      // 3: pb.PBLoginReq
	(*PBLoginRsp)(nil),      // 4: pb.PBLoginRsp
	(*PBMatchReq)(nil),      // 5: pb.PBMatchReq
	(*PBMtachRsp)(nil),      // 6: pb.PBMtachRsp
	(*PBConfirmReq)(nil),    // 7: pb.PBConfirmReq
	(*PBConfirmData)(nil),   // 8: pb.PBConfirmData
	(*PBnotifyConfirm)(nil), // 9: pb.PBnotifyConfirm
	(*PBSelectReq)(nil),     // 10: pb.PBSelectReq
}
var file_protocol_proto_depIdxs = []int32{
	1, // 0: pb.PBMatchReq.type:type_name -> pb.PBMatchReq.pvpType
	8, // 1: pb.PBnotifyConfirm.confirmArr:type_name -> pb.PBConfirmData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Head); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBLoginReq); i {
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
		file_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBLoginRsp); i {
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
		file_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBMatchReq); i {
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
		file_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBMtachRsp); i {
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
		file_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBConfirmReq); i {
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
		file_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBConfirmData); i {
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
		file_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBnotifyConfirm); i {
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
		file_protocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBSelectReq); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		EnumInfos:         file_protocol_proto_enumTypes,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}
