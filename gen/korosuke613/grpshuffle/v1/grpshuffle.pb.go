// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: korosuke613/grpshuffle/v1/grpshuffle.proto

package grpshufflev1

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

type Combination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []string `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"` // Target to be shuffled.
}

func (x *Combination) Reset() {
	*x = Combination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Combination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Combination) ProtoMessage() {}

func (x *Combination) ProtoReflect() protoreflect.Message {
	mi := &file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Combination.ProtoReflect.Descriptor instead.
func (*Combination) Descriptor() ([]byte, []int) {
	return file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescGZIP(), []int{0}
}

func (x *Combination) GetTargets() []string {
	if x != nil {
		return x.Targets
	}
	return nil
}

type ShuffleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets    []string `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`        // Target to be shuffled.
	Divide     uint64   `protobuf:"varint,4,opt,name=divide,proto3" json:"divide,omitempty"`         // The number of groups to divide into.
	Sequential bool     `protobuf:"varint,3,opt,name=sequential,proto3" json:"sequential,omitempty"` // If true, do not shuffle.
}

func (x *ShuffleRequest) Reset() {
	*x = ShuffleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShuffleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShuffleRequest) ProtoMessage() {}

func (x *ShuffleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShuffleRequest.ProtoReflect.Descriptor instead.
func (*ShuffleRequest) Descriptor() ([]byte, []int) {
	return file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescGZIP(), []int{1}
}

func (x *ShuffleRequest) GetTargets() []string {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *ShuffleRequest) GetDivide() uint64 {
	if x != nil {
		return x.Divide
	}
	return 0
}

func (x *ShuffleRequest) GetSequential() bool {
	if x != nil {
		return x.Sequential
	}
	return false
}

type ShuffleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Combinations []*Combination `protobuf:"bytes,1,rep,name=combinations,proto3" json:"combinations,omitempty"` // Set of targets.
}

func (x *ShuffleResponse) Reset() {
	*x = ShuffleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShuffleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShuffleResponse) ProtoMessage() {}

func (x *ShuffleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShuffleResponse.ProtoReflect.Descriptor instead.
func (*ShuffleResponse) Descriptor() ([]byte, []int) {
	return file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescGZIP(), []int{2}
}

func (x *ShuffleResponse) GetCombinations() []*Combination {
	if x != nil {
		return x.Combinations
	}
	return nil
}

var File_korosuke613_grpshuffle_v1_grpshuffle_proto protoreflect.FileDescriptor

var file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2f, 0x67, 0x72,
	0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x73,
	0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6b, 0x6f,
	0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2e, 0x67, 0x72, 0x70, 0x73, 0x68, 0x75,
	0x66, 0x66, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x22, 0x73, 0x0a, 0x0e, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x0f, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2e, 0x67, 0x72, 0x70,
	0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0x72, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x07, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c,
	0x65, 0x12, 0x29, 0x2e, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2e,
	0x67, 0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x75, 0x66, 0x66, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6b,
	0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2e, 0x67, 0x72, 0x70, 0x73, 0x68,
	0x75, 0x66, 0x66, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x84, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2e, 0x67, 0x72, 0x70,
	0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x47, 0x72, 0x70, 0x73,
	0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x75,
	0x6b, 0x65, 0x36, 0x31, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33,
	0x2f, 0x67, 0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67,
	0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4b, 0x47,
	0x58, 0xaa, 0x02, 0x19, 0x4b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x2e,
	0x47, 0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19,
	0x4b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x5c, 0x47, 0x72, 0x70, 0x73,
	0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x25, 0x4b, 0x6f, 0x72, 0x6f,
	0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x5c, 0x47, 0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66,
	0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1b, 0x4b, 0x6f, 0x72, 0x6f, 0x73, 0x75, 0x6b, 0x65, 0x36, 0x31, 0x33, 0x3a,
	0x3a, 0x47, 0x72, 0x70, 0x73, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescOnce sync.Once
	file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescData = file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDesc
)

func file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescGZIP() []byte {
	file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescOnce.Do(func() {
		file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescData = protoimpl.X.CompressGZIP(file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescData)
	})
	return file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDescData
}

var file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_korosuke613_grpshuffle_v1_grpshuffle_proto_goTypes = []interface{}{
	(*Combination)(nil),     // 0: korosuke613.grpshuffle.v1.Combination
	(*ShuffleRequest)(nil),  // 1: korosuke613.grpshuffle.v1.ShuffleRequest
	(*ShuffleResponse)(nil), // 2: korosuke613.grpshuffle.v1.ShuffleResponse
}
var file_korosuke613_grpshuffle_v1_grpshuffle_proto_depIdxs = []int32{
	0, // 0: korosuke613.grpshuffle.v1.ShuffleResponse.combinations:type_name -> korosuke613.grpshuffle.v1.Combination
	1, // 1: korosuke613.grpshuffle.v1.ComputeService.Shuffle:input_type -> korosuke613.grpshuffle.v1.ShuffleRequest
	2, // 2: korosuke613.grpshuffle.v1.ComputeService.Shuffle:output_type -> korosuke613.grpshuffle.v1.ShuffleResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_korosuke613_grpshuffle_v1_grpshuffle_proto_init() }
func file_korosuke613_grpshuffle_v1_grpshuffle_proto_init() {
	if File_korosuke613_grpshuffle_v1_grpshuffle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Combination); i {
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
		file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShuffleRequest); i {
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
		file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShuffleResponse); i {
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
			RawDescriptor: file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_korosuke613_grpshuffle_v1_grpshuffle_proto_goTypes,
		DependencyIndexes: file_korosuke613_grpshuffle_v1_grpshuffle_proto_depIdxs,
		MessageInfos:      file_korosuke613_grpshuffle_v1_grpshuffle_proto_msgTypes,
	}.Build()
	File_korosuke613_grpshuffle_v1_grpshuffle_proto = out.File
	file_korosuke613_grpshuffle_v1_grpshuffle_proto_rawDesc = nil
	file_korosuke613_grpshuffle_v1_grpshuffle_proto_goTypes = nil
	file_korosuke613_grpshuffle_v1_grpshuffle_proto_depIdxs = nil
}
