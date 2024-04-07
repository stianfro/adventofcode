// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: proto/v2015/aoc.proto

package aoc2015

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

type SolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *SolutionRequest) Reset() {
	*x = SolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v2015_aoc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionRequest) ProtoMessage() {}

func (x *SolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2015_aoc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionRequest.ProtoReflect.Descriptor instead.
func (*SolutionRequest) Descriptor() ([]byte, []int) {
	return file_proto_v2015_aoc_proto_rawDescGZIP(), []int{0}
}

func (x *SolutionRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type SolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnswerPartOne int64 `protobuf:"varint,1,opt,name=answerPartOne,proto3" json:"answerPartOne,omitempty"`
	AnswerPartTwo int64 `protobuf:"varint,2,opt,name=answerPartTwo,proto3" json:"answerPartTwo,omitempty"`
}

func (x *SolutionResponse) Reset() {
	*x = SolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v2015_aoc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionResponse) ProtoMessage() {}

func (x *SolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2015_aoc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionResponse.ProtoReflect.Descriptor instead.
func (*SolutionResponse) Descriptor() ([]byte, []int) {
	return file_proto_v2015_aoc_proto_rawDescGZIP(), []int{1}
}

func (x *SolutionResponse) GetAnswerPartOne() int64 {
	if x != nil {
		return x.AnswerPartOne
	}
	return 0
}

func (x *SolutionResponse) GetAnswerPartTwo() int64 {
	if x != nil {
		return x.AnswerPartTwo
	}
	return 0
}

var File_proto_v2015_aoc_proto protoreflect.FileDescriptor

var file_proto_v2015_aoc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x30, 0x31, 0x35, 0x2f, 0x61, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x32, 0x30, 0x31, 0x35, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x5e, 0x0a,
	0x10, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x4f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x74, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x54, 0x77, 0x6f, 0x32, 0x56, 0x0a,
	0x0f, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x04, 0x44, 0x61, 0x79, 0x31, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x32, 0x30, 0x31, 0x35, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x32, 0x30, 0x31, 0x35, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x61, 0x6f, 0x63, 0x32, 0x30,
	0x31, 0x35, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v2015_aoc_proto_rawDescOnce sync.Once
	file_proto_v2015_aoc_proto_rawDescData = file_proto_v2015_aoc_proto_rawDesc
)

func file_proto_v2015_aoc_proto_rawDescGZIP() []byte {
	file_proto_v2015_aoc_proto_rawDescOnce.Do(func() {
		file_proto_v2015_aoc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v2015_aoc_proto_rawDescData)
	})
	return file_proto_v2015_aoc_proto_rawDescData
}

var file_proto_v2015_aoc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v2015_aoc_proto_goTypes = []interface{}{
	(*SolutionRequest)(nil),  // 0: proto.v2015.SolutionRequest
	(*SolutionResponse)(nil), // 1: proto.v2015.SolutionResponse
}
var file_proto_v2015_aoc_proto_depIdxs = []int32{
	0, // 0: proto.v2015.SolutionService.Day1:input_type -> proto.v2015.SolutionRequest
	1, // 1: proto.v2015.SolutionService.Day1:output_type -> proto.v2015.SolutionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v2015_aoc_proto_init() }
func file_proto_v2015_aoc_proto_init() {
	if File_proto_v2015_aoc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v2015_aoc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionRequest); i {
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
		file_proto_v2015_aoc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionResponse); i {
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
			RawDescriptor: file_proto_v2015_aoc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v2015_aoc_proto_goTypes,
		DependencyIndexes: file_proto_v2015_aoc_proto_depIdxs,
		MessageInfos:      file_proto_v2015_aoc_proto_msgTypes,
	}.Build()
	File_proto_v2015_aoc_proto = out.File
	file_proto_v2015_aoc_proto_rawDesc = nil
	file_proto_v2015_aoc_proto_goTypes = nil
	file_proto_v2015_aoc_proto_depIdxs = nil
}
