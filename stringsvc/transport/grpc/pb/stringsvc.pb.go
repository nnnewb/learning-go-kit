// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: stringsvc/transport/pb/stringsvc.proto

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

type UppercaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *UppercaseRequest) Reset() {
	*x = UppercaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UppercaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UppercaseRequest) ProtoMessage() {}

func (x *UppercaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UppercaseRequest.ProtoReflect.Descriptor instead.
func (*UppercaseRequest) Descriptor() ([]byte, []int) {
	return file_stringsvc_transport_pb_stringsvc_proto_rawDescGZIP(), []int{0}
}

func (x *UppercaseRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type UppercaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V string `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *UppercaseResponse) Reset() {
	*x = UppercaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UppercaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UppercaseResponse) ProtoMessage() {}

func (x *UppercaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UppercaseResponse.ProtoReflect.Descriptor instead.
func (*UppercaseResponse) Descriptor() ([]byte, []int) {
	return file_stringsvc_transport_pb_stringsvc_proto_rawDescGZIP(), []int{1}
}

func (x *UppercaseResponse) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

type CountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_stringsvc_transport_pb_stringsvc_proto_rawDescGZIP(), []int{2}
}

func (x *CountRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

type CountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *CountResponse) Reset() {
	*x = CountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResponse) ProtoMessage() {}

func (x *CountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stringsvc_transport_pb_stringsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResponse.ProtoReflect.Descriptor instead.
func (*CountResponse) Descriptor() ([]byte, []int) {
	return file_stringsvc_transport_pb_stringsvc_proto_rawDescGZIP(), []int{3}
}

func (x *CountResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

var File_stringsvc_transport_pb_stringsvc_proto protoreflect.FileDescriptor

var file_stringsvc_transport_pb_stringsvc_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x76, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x10, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22, 0x21, 0x0a, 0x11, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x22, 0x1c, 0x0a,
	0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x22, 0x1d, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e, 0x32, 0x6b, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x55,
	0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x55, 0x70, 0x70, 0x65, 0x72,
	0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x70, 0x6c, 0x61, 0x79, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x76, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stringsvc_transport_pb_stringsvc_proto_rawDescOnce sync.Once
	file_stringsvc_transport_pb_stringsvc_proto_rawDescData = file_stringsvc_transport_pb_stringsvc_proto_rawDesc
)

func file_stringsvc_transport_pb_stringsvc_proto_rawDescGZIP() []byte {
	file_stringsvc_transport_pb_stringsvc_proto_rawDescOnce.Do(func() {
		file_stringsvc_transport_pb_stringsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_stringsvc_transport_pb_stringsvc_proto_rawDescData)
	})
	return file_stringsvc_transport_pb_stringsvc_proto_rawDescData
}

var file_stringsvc_transport_pb_stringsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stringsvc_transport_pb_stringsvc_proto_goTypes = []interface{}{
	(*UppercaseRequest)(nil),  // 0: UppercaseRequest
	(*UppercaseResponse)(nil), // 1: UppercaseResponse
	(*CountRequest)(nil),      // 2: CountRequest
	(*CountResponse)(nil),     // 3: CountResponse
}
var file_stringsvc_transport_pb_stringsvc_proto_depIdxs = []int32{
	0, // 0: StringService.Uppercase:input_type -> UppercaseRequest
	2, // 1: StringService.Count:input_type -> CountRequest
	1, // 2: StringService.Uppercase:output_type -> UppercaseResponse
	3, // 3: StringService.Count:output_type -> CountResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stringsvc_transport_pb_stringsvc_proto_init() }
func file_stringsvc_transport_pb_stringsvc_proto_init() {
	if File_stringsvc_transport_pb_stringsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stringsvc_transport_pb_stringsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UppercaseRequest); i {
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
		file_stringsvc_transport_pb_stringsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UppercaseResponse); i {
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
		file_stringsvc_transport_pb_stringsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRequest); i {
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
		file_stringsvc_transport_pb_stringsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountResponse); i {
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
			RawDescriptor: file_stringsvc_transport_pb_stringsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stringsvc_transport_pb_stringsvc_proto_goTypes,
		DependencyIndexes: file_stringsvc_transport_pb_stringsvc_proto_depIdxs,
		MessageInfos:      file_stringsvc_transport_pb_stringsvc_proto_msgTypes,
	}.Build()
	File_stringsvc_transport_pb_stringsvc_proto = out.File
	file_stringsvc_transport_pb_stringsvc_proto_rawDesc = nil
	file_stringsvc_transport_pb_stringsvc_proto_goTypes = nil
	file_stringsvc_transport_pb_stringsvc_proto_depIdxs = nil
}
