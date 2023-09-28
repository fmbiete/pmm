// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: managementpb/dump/dump.proto

package dumpv1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"

	_ "github.com/percona/pmm/api/inventorypb"
	_ "github.com/percona/pmm/api/managementpb/backup"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DumpStatus int32

const (
	DumpStatus_BACKUP_STATUS_INVALID     DumpStatus = 0
	DumpStatus_BACKUP_STATUS_IN_PROGRESS DumpStatus = 1
	DumpStatus_BACKUP_STATUS_SUCCESS     DumpStatus = 2
	DumpStatus_BACKUP_STATUS_ERROR       DumpStatus = 3
)

// Enum value maps for DumpStatus.
var (
	DumpStatus_name = map[int32]string{
		0: "BACKUP_STATUS_INVALID",
		1: "BACKUP_STATUS_IN_PROGRESS",
		2: "BACKUP_STATUS_SUCCESS",
		3: "BACKUP_STATUS_ERROR",
	}
	DumpStatus_value = map[string]int32{
		"BACKUP_STATUS_INVALID":     0,
		"BACKUP_STATUS_IN_PROGRESS": 1,
		"BACKUP_STATUS_SUCCESS":     2,
		"BACKUP_STATUS_ERROR":       3,
	}
)

func (x DumpStatus) Enum() *DumpStatus {
	p := new(DumpStatus)
	*p = x
	return p
}

func (x DumpStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DumpStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_dump_dump_proto_enumTypes[0].Descriptor()
}

func (DumpStatus) Type() protoreflect.EnumType {
	return &file_managementpb_dump_dump_proto_enumTypes[0]
}

func (x DumpStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DumpStatus.Descriptor instead.
func (DumpStatus) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{0}
}

type Dump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DumpId    string                 `protobuf:"bytes,1,opt,name=dump_id,json=dumpId,proto3" json:"dump_id,omitempty"`
	Status    DumpStatus             `protobuf:"varint,2,opt,name=status,proto3,enum=dump.v1beta1.DumpStatus" json:"status,omitempty"`
	NodeIds   []string               `protobuf:"bytes,3,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"` // TODO or node names?
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Dump) Reset() {
	*x = Dump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dump) ProtoMessage() {}

func (x *Dump) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dump.ProtoReflect.Descriptor instead.
func (*Dump) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{0}
}

func (x *Dump) GetDumpId() string {
	if x != nil {
		return x.DumpId
	}
	return ""
}

func (x *Dump) GetStatus() DumpStatus {
	if x != nil {
		return x.Status
	}
	return DumpStatus_BACKUP_STATUS_INVALID
}

func (x *Dump) GetNodeIds() []string {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

func (x *Dump) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Dump) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Dump) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type StartDumpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIds    []string               `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	ExportQan  bool                   `protobuf:"varint,4,opt,name=export_qan,json=exportQan,proto3" json:"export_qan,omitempty"`
	IgnoreLoad bool                   `protobuf:"varint,5,opt,name=ignore_load,json=ignoreLoad,proto3" json:"ignore_load,omitempty"`
}

func (x *StartDumpRequest) Reset() {
	*x = StartDumpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDumpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDumpRequest) ProtoMessage() {}

func (x *StartDumpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDumpRequest.ProtoReflect.Descriptor instead.
func (*StartDumpRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{1}
}

func (x *StartDumpRequest) GetNodeIds() []string {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

func (x *StartDumpRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *StartDumpRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *StartDumpRequest) GetExportQan() bool {
	if x != nil {
		return x.ExportQan
	}
	return false
}

func (x *StartDumpRequest) GetIgnoreLoad() bool {
	if x != nil {
		return x.IgnoreLoad
	}
	return false
}

type StartDumpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DumpId string `protobuf:"bytes,1,opt,name=dump_id,json=dumpId,proto3" json:"dump_id,omitempty"`
}

func (x *StartDumpResponse) Reset() {
	*x = StartDumpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDumpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDumpResponse) ProtoMessage() {}

func (x *StartDumpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDumpResponse.ProtoReflect.Descriptor instead.
func (*StartDumpResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{2}
}

func (x *StartDumpResponse) GetDumpId() string {
	if x != nil {
		return x.DumpId
	}
	return ""
}

type ListDumpsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDumpsRequest) Reset() {
	*x = ListDumpsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDumpsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDumpsRequest) ProtoMessage() {}

func (x *ListDumpsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDumpsRequest.ProtoReflect.Descriptor instead.
func (*ListDumpsRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{3}
}

type ListDumpsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dumps []*Dump `protobuf:"bytes,1,rep,name=dumps,proto3" json:"dumps,omitempty"`
}

func (x *ListDumpsResponse) Reset() {
	*x = ListDumpsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDumpsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDumpsResponse) ProtoMessage() {}

func (x *ListDumpsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDumpsResponse.ProtoReflect.Descriptor instead.
func (*ListDumpsResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{4}
}

func (x *ListDumpsResponse) GetDumps() []*Dump {
	if x != nil {
		return x.Dumps
	}
	return nil
}

type DeleteDumpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DumpIds []string `protobuf:"bytes,1,rep,name=dump_ids,json=dumpIds,proto3" json:"dump_ids,omitempty"`
}

func (x *DeleteDumpRequest) Reset() {
	*x = DeleteDumpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDumpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDumpRequest) ProtoMessage() {}

func (x *DeleteDumpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDumpRequest.ProtoReflect.Descriptor instead.
func (*DeleteDumpRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDumpRequest) GetDumpIds() []string {
	if x != nil {
		return x.DumpIds
	}
	return nil
}

type DeleteDumpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDumpResponse) Reset() {
	*x = DeleteDumpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDumpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDumpResponse) ProtoMessage() {}

func (x *DeleteDumpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDumpResponse.ProtoReflect.Descriptor instead.
func (*DeleteDumpResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{6}
}

type GetDumpLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DumpId    string `protobuf:"bytes,1,opt,name=dump_id,json=dumpId,proto3" json:"dump_id,omitempty"`
	Offset    uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	RestoreId string `protobuf:"bytes,4,opt,name=restore_id,json=restoreId,proto3" json:"restore_id,omitempty"`
}

func (x *GetDumpLogsRequest) Reset() {
	*x = GetDumpLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDumpLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDumpLogsRequest) ProtoMessage() {}

func (x *GetDumpLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDumpLogsRequest.ProtoReflect.Descriptor instead.
func (*GetDumpLogsRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{7}
}

func (x *GetDumpLogsRequest) GetDumpId() string {
	if x != nil {
		return x.DumpId
	}
	return ""
}

func (x *GetDumpLogsRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetDumpLogsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetDumpLogsRequest) GetRestoreId() string {
	if x != nil {
		return x.RestoreId
	}
	return ""
}

type GetDumpLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDumpLogsResponse) Reset() {
	*x = GetDumpLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dump_dump_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDumpLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDumpLogsResponse) ProtoMessage() {}

func (x *GetDumpLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dump_dump_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDumpLogsResponse.ProtoReflect.Descriptor instead.
func (*GetDumpLogsResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dump_dump_proto_rawDescGZIP(), []int{8}
}

var File_managementpb_dump_dump_proto protoreflect.FileDescriptor

var file_managementpb_dump_dump_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64,
	0x75, 0x6d, 0x70, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x04, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x75, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x75,
	0x6d, 0x70, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x75, 0x6d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdf,
	0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x71, 0x61, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x51, 0x61, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x61, 0x64,
	0x22, 0x2c, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x75, 0x6d, 0x70, 0x49, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x70,
	0x73, 0x22, 0x2e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x6d, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x6d, 0x70, 0x49, 0x64,
	0x73, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x75,
	0x6d, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x75, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x7a, 0x0a, 0x0a, 0x44, 0x75,
	0x6d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x43, 0x4b,
	0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x32, 0xed, 0x06, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x70, 0x73,
	0x12, 0xe5, 0x03, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1e,
	0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x96, 0x03, 0x92, 0x41, 0xe8, 0x02, 0x1a, 0xe5, 0x02, 0x43, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x20, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x20, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x3a, 0x0a,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x58, 0x54, 0x52, 0x41, 0x42,
	0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x45, 0x44, 0x20, 0x2d, 0x20, 0x78, 0x74, 0x72, 0x61, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65,
	0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x0a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x58, 0x54, 0x52, 0x41, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x20, 0x2d,
	0x20, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x78, 0x74, 0x72, 0x61, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x78, 0x62, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x0a, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x58, 0x54, 0x52, 0x41, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50,
	0x20, 0x2d, 0x20, 0x78, 0x74, 0x72, 0x61, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x20, 0x69, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x4d, 0x79, 0x53, 0x51, 0x4c, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x2f, 0x44, 0x75, 0x6d,
	0x70, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x77, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x75, 0x6d, 0x70, 0x73, 0x12, 0x1e, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01,
	0x2a, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x2f, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x2f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x7c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x6d, 0x70, 0x12,
	0x1f, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x75,
	0x6d, 0x70, 0x2f, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x84, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x20, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x64, 0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22,
	0x25, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x64, 0x75, 0x6d, 0x70, 0x2f, 0x44, 0x75, 0x6d, 0x70, 0x73, 0x2f, 0x47, 0x65, 0x74, 0x44, 0x75,
	0x6d, 0x70, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0xa8, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x75, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x44, 0x75, 0x6d,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x3b, 0x64, 0x75, 0x6d, 0x70, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x44, 0x75, 0x6d, 0x70, 0x2e,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0c, 0x44, 0x75, 0x6d, 0x70, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x18, 0x44, 0x75, 0x6d, 0x70, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0d, 0x44, 0x75, 0x6d, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_dump_dump_proto_rawDescOnce sync.Once
	file_managementpb_dump_dump_proto_rawDescData = file_managementpb_dump_dump_proto_rawDesc
)

func file_managementpb_dump_dump_proto_rawDescGZIP() []byte {
	file_managementpb_dump_dump_proto_rawDescOnce.Do(func() {
		file_managementpb_dump_dump_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_dump_dump_proto_rawDescData)
	})
	return file_managementpb_dump_dump_proto_rawDescData
}

var (
	file_managementpb_dump_dump_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_dump_dump_proto_msgTypes  = make([]protoimpl.MessageInfo, 9)
	file_managementpb_dump_dump_proto_goTypes   = []interface{}{
		(DumpStatus)(0),               // 0: dump.v1beta1.DumpStatus
		(*Dump)(nil),                  // 1: dump.v1beta1.Dump
		(*StartDumpRequest)(nil),      // 2: dump.v1beta1.StartDumpRequest
		(*StartDumpResponse)(nil),     // 3: dump.v1beta1.StartDumpResponse
		(*ListDumpsRequest)(nil),      // 4: dump.v1beta1.ListDumpsRequest
		(*ListDumpsResponse)(nil),     // 5: dump.v1beta1.ListDumpsResponse
		(*DeleteDumpRequest)(nil),     // 6: dump.v1beta1.DeleteDumpRequest
		(*DeleteDumpResponse)(nil),    // 7: dump.v1beta1.DeleteDumpResponse
		(*GetDumpLogsRequest)(nil),    // 8: dump.v1beta1.GetDumpLogsRequest
		(*GetDumpLogsResponse)(nil),   // 9: dump.v1beta1.GetDumpLogsResponse
		(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	}
)

var file_managementpb_dump_dump_proto_depIdxs = []int32{
	0,  // 0: dump.v1beta1.Dump.status:type_name -> dump.v1beta1.DumpStatus
	10, // 1: dump.v1beta1.Dump.start_time:type_name -> google.protobuf.Timestamp
	10, // 2: dump.v1beta1.Dump.end_time:type_name -> google.protobuf.Timestamp
	10, // 3: dump.v1beta1.Dump.created_at:type_name -> google.protobuf.Timestamp
	10, // 4: dump.v1beta1.StartDumpRequest.start_time:type_name -> google.protobuf.Timestamp
	10, // 5: dump.v1beta1.StartDumpRequest.end_time:type_name -> google.protobuf.Timestamp
	1,  // 6: dump.v1beta1.ListDumpsResponse.dumps:type_name -> dump.v1beta1.Dump
	2,  // 7: dump.v1beta1.Dumps.StartDump:input_type -> dump.v1beta1.StartDumpRequest
	4,  // 8: dump.v1beta1.Dumps.ListDumps:input_type -> dump.v1beta1.ListDumpsRequest
	6,  // 9: dump.v1beta1.Dumps.DeleteDump:input_type -> dump.v1beta1.DeleteDumpRequest
	8,  // 10: dump.v1beta1.Dumps.GetDumpLogs:input_type -> dump.v1beta1.GetDumpLogsRequest
	3,  // 11: dump.v1beta1.Dumps.StartDump:output_type -> dump.v1beta1.StartDumpResponse
	5,  // 12: dump.v1beta1.Dumps.ListDumps:output_type -> dump.v1beta1.ListDumpsResponse
	7,  // 13: dump.v1beta1.Dumps.DeleteDump:output_type -> dump.v1beta1.DeleteDumpResponse
	9,  // 14: dump.v1beta1.Dumps.GetDumpLogs:output_type -> dump.v1beta1.GetDumpLogsResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_managementpb_dump_dump_proto_init() }
func file_managementpb_dump_dump_proto_init() {
	if File_managementpb_dump_dump_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_dump_dump_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dump); i {
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
		file_managementpb_dump_dump_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDumpRequest); i {
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
		file_managementpb_dump_dump_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDumpResponse); i {
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
		file_managementpb_dump_dump_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDumpsRequest); i {
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
		file_managementpb_dump_dump_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDumpsResponse); i {
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
		file_managementpb_dump_dump_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDumpRequest); i {
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
		file_managementpb_dump_dump_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDumpResponse); i {
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
		file_managementpb_dump_dump_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDumpLogsRequest); i {
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
		file_managementpb_dump_dump_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDumpLogsResponse); i {
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
			RawDescriptor: file_managementpb_dump_dump_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_dump_dump_proto_goTypes,
		DependencyIndexes: file_managementpb_dump_dump_proto_depIdxs,
		EnumInfos:         file_managementpb_dump_dump_proto_enumTypes,
		MessageInfos:      file_managementpb_dump_dump_proto_msgTypes,
	}.Build()
	File_managementpb_dump_dump_proto = out.File
	file_managementpb_dump_dump_proto_rawDesc = nil
	file_managementpb_dump_dump_proto_goTypes = nil
	file_managementpb_dump_dump_proto_depIdxs = nil
}
