// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: 38tao.proto

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

type GiveMeBTCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GiveMeBTCRequest) Reset() {
	*x = GiveMeBTCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__38tao_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveMeBTCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveMeBTCRequest) ProtoMessage() {}

func (x *GiveMeBTCRequest) ProtoReflect() protoreflect.Message {
	mi := &file__38tao_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveMeBTCRequest.ProtoReflect.Descriptor instead.
func (*GiveMeBTCRequest) Descriptor() ([]byte, []int) {
	return file__38tao_proto_rawDescGZIP(), []int{0}
}

func (x *GiveMeBTCRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GiveMeBTCRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GiveMeBTCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GiveMeBTCResponse) Reset() {
	*x = GiveMeBTCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__38tao_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveMeBTCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveMeBTCResponse) ProtoMessage() {}

func (x *GiveMeBTCResponse) ProtoReflect() protoreflect.Message {
	mi := &file__38tao_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveMeBTCResponse.ProtoReflect.Descriptor instead.
func (*GiveMeBTCResponse) Descriptor() ([]byte, []int) {
	return file__38tao_proto_rawDescGZIP(), []int{1}
}

func (x *GiveMeBTCResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GiveMeBTCResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File__38tao_proto protoreflect.FileDescriptor

var file__38tao_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x33, 0x38, 0x74, 0x61, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74,
	0x72, 0x65, 0x7a, 0x6f, 0x69, 0x74, 0x61, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x69, 0x76, 0x65,
	0x4d, 0x65, 0x42, 0x54, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x69, 0x76, 0x65, 0x4d, 0x65,
	0x42, 0x54, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x57, 0x0a,
	0x10, 0x54, 0x72, 0x65, 0x7a, 0x6f, 0x69, 0x74, 0x61, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x42, 0x54, 0x43, 0x12, 0x1b, 0x2e, 0x74, 0x72,
	0x65, 0x7a, 0x6f, 0x69, 0x74, 0x61, 0x6f, 0x2e, 0x47, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x42, 0x54,
	0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x65, 0x7a, 0x6f,
	0x69, 0x74, 0x61, 0x6f, 0x2e, 0x47, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x42, 0x54, 0x43, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x69,
	0x74, 0x61, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file__38tao_proto_rawDescOnce sync.Once
	file__38tao_proto_rawDescData = file__38tao_proto_rawDesc
)

func file__38tao_proto_rawDescGZIP() []byte {
	file__38tao_proto_rawDescOnce.Do(func() {
		file__38tao_proto_rawDescData = protoimpl.X.CompressGZIP(file__38tao_proto_rawDescData)
	})
	return file__38tao_proto_rawDescData
}

var file__38tao_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__38tao_proto_goTypes = []interface{}{
	(*GiveMeBTCRequest)(nil),  // 0: trezoitao.GiveMeBTCRequest
	(*GiveMeBTCResponse)(nil), // 1: trezoitao.GiveMeBTCResponse
}
var file__38tao_proto_depIdxs = []int32{
	0, // 0: trezoitao.TrezoitaoService.GetBTC:input_type -> trezoitao.GiveMeBTCRequest
	1, // 1: trezoitao.TrezoitaoService.GetBTC:output_type -> trezoitao.GiveMeBTCResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__38tao_proto_init() }
func file__38tao_proto_init() {
	if File__38tao_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__38tao_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiveMeBTCRequest); i {
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
		file__38tao_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiveMeBTCResponse); i {
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
			RawDescriptor: file__38tao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file__38tao_proto_goTypes,
		DependencyIndexes: file__38tao_proto_depIdxs,
		MessageInfos:      file__38tao_proto_msgTypes,
	}.Build()
	File__38tao_proto = out.File
	file__38tao_proto_rawDesc = nil
	file__38tao_proto_goTypes = nil
	file__38tao_proto_depIdxs = nil
}
