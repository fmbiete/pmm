// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: inventorypb/agent_status.proto

package inventorypb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AgentStatus represents actual Agent status.
type AgentStatus int32

const (
	AgentStatus_AGENT_STATUS_INVALID AgentStatus = 0
	// Agent is starting.
	AgentStatus_STARTING AgentStatus = 1
	// Agent is running.
	AgentStatus_RUNNING AgentStatus = 2
	// Agent encountered error and will be restarted automatically soon.
	AgentStatus_WAITING AgentStatus = 3
	// Agent is stopping.
	AgentStatus_STOPPING AgentStatus = 4
	// Agent finished.
	AgentStatus_DONE AgentStatus = 5
	// Agent is not connected, we don't know anything about it's state.
	AgentStatus_UNKNOWN AgentStatus = 6
)

// Enum value maps for AgentStatus.
var (
	AgentStatus_name = map[int32]string{
		0: "AGENT_STATUS_INVALID",
		1: "STARTING",
		2: "RUNNING",
		3: "WAITING",
		4: "STOPPING",
		5: "DONE",
		6: "UNKNOWN",
	}
	AgentStatus_value = map[string]int32{
		"AGENT_STATUS_INVALID": 0,
		"STARTING":             1,
		"RUNNING":              2,
		"WAITING":              3,
		"STOPPING":             4,
		"DONE":                 5,
		"UNKNOWN":              6,
	}
)

func (x AgentStatus) Enum() *AgentStatus {
	p := new(AgentStatus)
	*p = x
	return p
}

func (x AgentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_inventorypb_agent_status_proto_enumTypes[0].Descriptor()
}

func (AgentStatus) Type() protoreflect.EnumType {
	return &file_inventorypb_agent_status_proto_enumTypes[0]
}

func (x AgentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentStatus.Descriptor instead.
func (AgentStatus) EnumDescriptor() ([]byte, []int) {
	return file_inventorypb_agent_status_proto_rawDescGZIP(), []int{0}
}

var File_inventorypb_agent_status_proto protoreflect.FileDescriptor

var file_inventorypb_agent_status_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2a, 0x74, 0x0a, 0x0b, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x06, 0x42, 0x8d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x42, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0xa2,
	0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0xca, 0x02, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0xe2, 0x02, 0x15,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventorypb_agent_status_proto_rawDescOnce sync.Once
	file_inventorypb_agent_status_proto_rawDescData = file_inventorypb_agent_status_proto_rawDesc
)

func file_inventorypb_agent_status_proto_rawDescGZIP() []byte {
	file_inventorypb_agent_status_proto_rawDescOnce.Do(func() {
		file_inventorypb_agent_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventorypb_agent_status_proto_rawDescData)
	})
	return file_inventorypb_agent_status_proto_rawDescData
}

var (
	file_inventorypb_agent_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_inventorypb_agent_status_proto_goTypes   = []interface{}{
		(AgentStatus)(0), // 0: inventory.AgentStatus
	}
)

var file_inventorypb_agent_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventorypb_agent_status_proto_init() }
func file_inventorypb_agent_status_proto_init() {
	if File_inventorypb_agent_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventorypb_agent_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventorypb_agent_status_proto_goTypes,
		DependencyIndexes: file_inventorypb_agent_status_proto_depIdxs,
		EnumInfos:         file_inventorypb_agent_status_proto_enumTypes,
	}.Build()
	File_inventorypb_agent_status_proto = out.File
	file_inventorypb_agent_status_proto_rawDesc = nil
	file_inventorypb_agent_status_proto_goTypes = nil
	file_inventorypb_agent_status_proto_depIdxs = nil
}
