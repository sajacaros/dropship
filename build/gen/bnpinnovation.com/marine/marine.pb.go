// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: marine/proto/marine.proto

// https://developers.google.com/protocol-buffers/docs/gotutorial

package marine

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProjectStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Uptime  string `protobuf:"bytes,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Pid     int32  `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ProjectStatus) Reset() {
	*x = ProjectStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marine_proto_marine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectStatus) ProtoMessage() {}

func (x *ProjectStatus) ProtoReflect() protoreflect.Message {
	mi := &file_marine_proto_marine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectStatus.ProtoReflect.Descriptor instead.
func (*ProjectStatus) Descriptor() ([]byte, []int) {
	return file_marine_proto_marine_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectStatus) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ProjectStatus) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *ProjectStatus) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type StatusSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*ProjectStatus `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *StatusSummary) Reset() {
	*x = StatusSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marine_proto_marine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSummary) ProtoMessage() {}

func (x *StatusSummary) ProtoReflect() protoreflect.Message {
	mi := &file_marine_proto_marine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSummary.ProtoReflect.Descriptor instead.
func (*StatusSummary) Descriptor() ([]byte, []int) {
	return file_marine_proto_marine_proto_rawDescGZIP(), []int{1}
}

func (x *StatusSummary) GetProjects() []*ProjectStatus {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ProjectIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectIdentity) Reset() {
	*x = ProjectIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marine_proto_marine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectIdentity) ProtoMessage() {}

func (x *ProjectIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_marine_proto_marine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectIdentity.ProtoReflect.Descriptor instead.
func (*ProjectIdentity) Descriptor() ([]byte, []int) {
	return file_marine_proto_marine_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectIdentity) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

var File_marine_proto_marine_proto protoreflect.FileDescriptor

var file_marine_proto_marine_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x72,
	0x69, 0x6e, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x2b, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0xba, 0x04, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5a, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x63, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x15,
	0x2e, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f,
	0x64, 0x72, 0x6f, 0x70, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x62, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x69,
	0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x22, 0x20, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x60, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x17, 0x2e, 0x6d,
	0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x73, 0x68, 0x69, 0x70,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x64, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x64, 0x72, 0x6f, 0x70,
	0x73, 0x68, 0x69, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x07,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x6e, 0x70,
	0x69, 0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x72, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_marine_proto_marine_proto_rawDescOnce sync.Once
	file_marine_proto_marine_proto_rawDescData = file_marine_proto_marine_proto_rawDesc
)

func file_marine_proto_marine_proto_rawDescGZIP() []byte {
	file_marine_proto_marine_proto_rawDescOnce.Do(func() {
		file_marine_proto_marine_proto_rawDescData = protoimpl.X.CompressGZIP(file_marine_proto_marine_proto_rawDescData)
	})
	return file_marine_proto_marine_proto_rawDescData
}

var file_marine_proto_marine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_marine_proto_marine_proto_goTypes = []interface{}{
	(*ProjectStatus)(nil),   // 0: marine.ProjectStatus
	(*StatusSummary)(nil),   // 1: marine.StatusSummary
	(*ProjectIdentity)(nil), // 2: marine.ProjectIdentity
	(*empty.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_marine_proto_marine_proto_depIdxs = []int32{
	0, // 0: marine.StatusSummary.projects:type_name -> marine.ProjectStatus
	3, // 1: marine.ProjectService.Summary:input_type -> google.protobuf.Empty
	2, // 2: marine.ProjectService.Status:input_type -> marine.ProjectIdentity
	2, // 3: marine.ProjectService.Start:input_type -> marine.ProjectIdentity
	2, // 4: marine.ProjectService.Stop:input_type -> marine.ProjectIdentity
	2, // 5: marine.ProjectService.Update:input_type -> marine.ProjectIdentity
	3, // 6: marine.ProjectService.Install:input_type -> google.protobuf.Empty
	1, // 7: marine.ProjectService.Summary:output_type -> marine.StatusSummary
	0, // 8: marine.ProjectService.Status:output_type -> marine.ProjectStatus
	3, // 9: marine.ProjectService.Start:output_type -> google.protobuf.Empty
	3, // 10: marine.ProjectService.Stop:output_type -> google.protobuf.Empty
	3, // 11: marine.ProjectService.Update:output_type -> google.protobuf.Empty
	3, // 12: marine.ProjectService.Install:output_type -> google.protobuf.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_marine_proto_marine_proto_init() }
func file_marine_proto_marine_proto_init() {
	if File_marine_proto_marine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_marine_proto_marine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectStatus); i {
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
		file_marine_proto_marine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusSummary); i {
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
		file_marine_proto_marine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectIdentity); i {
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
			RawDescriptor: file_marine_proto_marine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_marine_proto_marine_proto_goTypes,
		DependencyIndexes: file_marine_proto_marine_proto_depIdxs,
		MessageInfos:      file_marine_proto_marine_proto_msgTypes,
	}.Build()
	File_marine_proto_marine_proto = out.File
	file_marine_proto_marine_proto_rawDesc = nil
	file_marine_proto_marine_proto_goTypes = nil
	file_marine_proto_marine_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	//   Unary
	Summary(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusSummary, error)
	Status(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*ProjectStatus, error)
	Start(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*empty.Empty, error)
	Stop(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*empty.Empty, error)
	Install(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Summary(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusSummary, error) {
	out := new(StatusSummary)
	err := c.cc.Invoke(ctx, "/marine.ProjectService/Summary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Status(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*ProjectStatus, error) {
	out := new(ProjectStatus)
	err := c.cc.Invoke(ctx, "/marine.ProjectService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Start(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/marine.ProjectService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Stop(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/marine.ProjectService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Update(ctx context.Context, in *ProjectIdentity, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/marine.ProjectService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Install(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/marine.ProjectService/Install", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
type ProjectServiceServer interface {
	//   Unary
	Summary(context.Context, *empty.Empty) (*StatusSummary, error)
	Status(context.Context, *ProjectIdentity) (*ProjectStatus, error)
	Start(context.Context, *ProjectIdentity) (*empty.Empty, error)
	Stop(context.Context, *ProjectIdentity) (*empty.Empty, error)
	Update(context.Context, *ProjectIdentity) (*empty.Empty, error)
	Install(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedProjectServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (*UnimplementedProjectServiceServer) Summary(context.Context, *empty.Empty) (*StatusSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Summary not implemented")
}
func (*UnimplementedProjectServiceServer) Status(context.Context, *ProjectIdentity) (*ProjectStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedProjectServiceServer) Start(context.Context, *ProjectIdentity) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedProjectServiceServer) Stop(context.Context, *ProjectIdentity) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedProjectServiceServer) Update(context.Context, *ProjectIdentity) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProjectServiceServer) Install(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}

func RegisterProjectServiceServer(s *grpc.Server, srv ProjectServiceServer) {
	s.RegisterService(&_ProjectService_serviceDesc, srv)
}

func _ProjectService_Summary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Summary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marine.ProjectService/Summary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Summary(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marine.ProjectService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Status(ctx, req.(*ProjectIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marine.ProjectService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Start(ctx, req.(*ProjectIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marine.ProjectService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Stop(ctx, req.(*ProjectIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marine.ProjectService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Update(ctx, req.(*ProjectIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marine.ProjectService/Install",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Install(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marine.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Summary",
			Handler:    _ProjectService_Summary_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ProjectService_Status_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ProjectService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ProjectService_Stop_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProjectService_Update_Handler,
		},
		{
			MethodName: "Install",
			Handler:    _ProjectService_Install_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marine/proto/marine.proto",
}
