// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/definitions.proto

package proto

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

type SingleOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *SingleOperation) Reset() {
	*x = SingleOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_definitions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleOperation) ProtoMessage() {}

func (x *SingleOperation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_definitions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleOperation.ProtoReflect.Descriptor instead.
func (*SingleOperation) Descriptor() ([]byte, []int) {
	return file_proto_definitions_proto_rawDescGZIP(), []int{0}
}

func (x *SingleOperation) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type ClientBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId string             `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	Requests []*SingleOperation `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
	Sender   int64              `protobuf:"varint,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *ClientBatch) Reset() {
	*x = ClientBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_definitions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientBatch) ProtoMessage() {}

func (x *ClientBatch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_definitions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientBatch.ProtoReflect.Descriptor instead.
func (*ClientBatch) Descriptor() ([]byte, []int) {
	return file_proto_definitions_proto_rawDescGZIP(), []int{1}
}

func (x *ClientBatch) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *ClientBatch) GetRequests() []*SingleOperation {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *ClientBatch) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

type ReplicaBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId string         `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	Requests []*ClientBatch `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
	Sender   int64          `protobuf:"varint,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *ReplicaBatch) Reset() {
	*x = ReplicaBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_definitions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicaBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicaBatch) ProtoMessage() {}

func (x *ReplicaBatch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_definitions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicaBatch.ProtoReflect.Descriptor instead.
func (*ReplicaBatch) Descriptor() ([]byte, []int) {
	return file_proto_definitions_proto_rawDescGZIP(), []int{2}
}

func (x *ReplicaBatch) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *ReplicaBatch) GetRequests() []*ClientBatch {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *ReplicaBatch) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"` // 1 for bootstrap, 2 for log print, 3 consensus start
	Note   string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	Sender int64  `protobuf:"varint,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_definitions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_definitions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_definitions_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Status) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Status) GetSender() int64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

type PaxosConsensus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender         int32                     `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver       int32                     `protobuf:"varint,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Type           int32                     `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"` // 1-prepare, 2-promise, 3-propose, 4-accept, 5 internal timeout
	Note           string                    `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	InstanceNumber int32                     `protobuf:"varint,5,opt,name=instance_number,json=instanceNumber,proto3" json:"instance_number,omitempty"` // initial id for prepare messages, instance id for propose message
	Ballot         int32                     `protobuf:"varint,6,opt,name=ballot,proto3" json:"ballot,omitempty"`
	View           int32                     `protobuf:"varint,7,opt,name=view,proto3" json:"view,omitempty"`
	PromiseReply   []*PaxosConsensusInstance `protobuf:"bytes,8,rep,name=promiseReply,proto3" json:"promiseReply,omitempty"`
	ProposeValue   *ReplicaBatch             `protobuf:"bytes,9,opt,name=propose_value,json=proposeValue,proto3" json:"propose_value,omitempty"`
	DecidedValues  []*PaxosConsensusInstance `protobuf:"bytes,10,rep,name=decided_values,json=decidedValues,proto3" json:"decided_values,omitempty"` // decided value for last instances
}

func (x *PaxosConsensus) Reset() {
	*x = PaxosConsensus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_definitions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaxosConsensus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaxosConsensus) ProtoMessage() {}

func (x *PaxosConsensus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_definitions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaxosConsensus.ProtoReflect.Descriptor instead.
func (*PaxosConsensus) Descriptor() ([]byte, []int) {
	return file_proto_definitions_proto_rawDescGZIP(), []int{4}
}

func (x *PaxosConsensus) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *PaxosConsensus) GetReceiver() int32 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *PaxosConsensus) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PaxosConsensus) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *PaxosConsensus) GetInstanceNumber() int32 {
	if x != nil {
		return x.InstanceNumber
	}
	return 0
}

func (x *PaxosConsensus) GetBallot() int32 {
	if x != nil {
		return x.Ballot
	}
	return 0
}

func (x *PaxosConsensus) GetView() int32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *PaxosConsensus) GetPromiseReply() []*PaxosConsensusInstance {
	if x != nil {
		return x.PromiseReply
	}
	return nil
}

func (x *PaxosConsensus) GetProposeValue() *ReplicaBatch {
	if x != nil {
		return x.ProposeValue
	}
	return nil
}

func (x *PaxosConsensus) GetDecidedValues() []*PaxosConsensusInstance {
	if x != nil {
		return x.DecidedValues
	}
	return nil
}

type PaxosConsensusInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32         `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Ballot int32         `protobuf:"varint,2,opt,name=ballot,proto3" json:"ballot,omitempty"`
	Value  *ReplicaBatch `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PaxosConsensusInstance) Reset() {
	*x = PaxosConsensusInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_definitions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaxosConsensusInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaxosConsensusInstance) ProtoMessage() {}

func (x *PaxosConsensusInstance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_definitions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaxosConsensusInstance.ProtoReflect.Descriptor instead.
func (*PaxosConsensusInstance) Descriptor() ([]byte, []int) {
	return file_proto_definitions_proto_rawDescGZIP(), []int{4, 0}
}

func (x *PaxosConsensusInstance) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PaxosConsensusInstance) GetBallot() int32 {
	if x != nil {
		return x.Ballot
	}
	return 0
}

func (x *PaxosConsensusInstance) GetValue() *ReplicaBatch {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_proto_definitions_proto protoreflect.FileDescriptor

var file_proto_definitions_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x70, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x6d, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x22, 0xd5, 0x03, 0x0a, 0x0e, 0x50, 0x61, 0x78, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6c,
	0x6c, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61, 0x6c, 0x6c, 0x6f,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x12, 0x3c, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x50, 0x61,
	0x78, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x64,
	0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x50, 0x61, 0x78, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0d, 0x64, 0x65, 0x63, 0x69, 0x64,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x5f, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x61,
	0x6c, 0x6c, 0x6f, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_definitions_proto_rawDescOnce sync.Once
	file_proto_definitions_proto_rawDescData = file_proto_definitions_proto_rawDesc
)

func file_proto_definitions_proto_rawDescGZIP() []byte {
	file_proto_definitions_proto_rawDescOnce.Do(func() {
		file_proto_definitions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_definitions_proto_rawDescData)
	})
	return file_proto_definitions_proto_rawDescData
}

var file_proto_definitions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_definitions_proto_goTypes = []interface{}{
	(*SingleOperation)(nil),        // 0: SingleOperation
	(*ClientBatch)(nil),            // 1: ClientBatch
	(*ReplicaBatch)(nil),           // 2: ReplicaBatch
	(*Status)(nil),                 // 3: Status
	(*PaxosConsensus)(nil),         // 4: PaxosConsensus
	(*PaxosConsensusInstance)(nil), // 5: PaxosConsensus.instance
}
var file_proto_definitions_proto_depIdxs = []int32{
	0, // 0: ClientBatch.requests:type_name -> SingleOperation
	1, // 1: ReplicaBatch.requests:type_name -> ClientBatch
	5, // 2: PaxosConsensus.promiseReply:type_name -> PaxosConsensus.instance
	2, // 3: PaxosConsensus.propose_value:type_name -> ReplicaBatch
	5, // 4: PaxosConsensus.decided_values:type_name -> PaxosConsensus.instance
	2, // 5: PaxosConsensus.instance.value:type_name -> ReplicaBatch
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_definitions_proto_init() }
func file_proto_definitions_proto_init() {
	if File_proto_definitions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_definitions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleOperation); i {
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
		file_proto_definitions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientBatch); i {
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
		file_proto_definitions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicaBatch); i {
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
		file_proto_definitions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_definitions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaxosConsensus); i {
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
		file_proto_definitions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaxosConsensusInstance); i {
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
			RawDescriptor: file_proto_definitions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_definitions_proto_goTypes,
		DependencyIndexes: file_proto_definitions_proto_depIdxs,
		MessageInfos:      file_proto_definitions_proto_msgTypes,
	}.Build()
	File_proto_definitions_proto = out.File
	file_proto_definitions_proto_rawDesc = nil
	file_proto_definitions_proto_goTypes = nil
	file_proto_definitions_proto_depIdxs = nil
}
