// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: catalog-common/entities.proto

package grpc_catalog_common_go

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

// OpStatus enumeration defining the possible status of an operation.
type OpStatus int32

const (
	// Unknown operation status indicates an error in the processing of the request as all operations should
	// have a determined status.
	OpStatus_UNKNOWN OpStatus = 0
	// Success indicates that the operation was completed successfully.
	OpStatus_SUCCESS OpStatus = 1
)

// Enum value maps for OpStatus.
var (
	OpStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
	}
	OpStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
	}
)

func (x OpStatus) Enum() *OpStatus {
	p := new(OpStatus)
	*p = x
	return p
}

func (x OpStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_catalog_common_entities_proto_enumTypes[0].Descriptor()
}

func (OpStatus) Type() protoreflect.EnumType {
	return &file_catalog_common_entities_proto_enumTypes[0]
}

func (x OpStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpStatus.Descriptor instead.
func (OpStatus) EnumDescriptor() ([]byte, []int) {
	return file_catalog_common_entities_proto_rawDescGZIP(), []int{0}
}

// Empty message as required by gRPC for services not requiring any parameters.
type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_common_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_common_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_catalog_common_entities_proto_rawDescGZIP(), []int{0}
}

// OpResponse with a common definition for the result of executing an operation through gRPC. Notice that
// in case of an error, the call must fail unless the operation is executed asynchronously and this
// message contains the result of a past operation.
type OpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the operation.
	Status OpStatus `protobuf:"varint,1,opt,name=status,proto3,enum=catalog_common.OpStatus" json:"status,omitempty"`
	// StatusName with the string representation of the status.
	StatusName string `protobuf:"bytes,2,opt,name=status_name,json=statusName,proto3" json:"status_name,omitempty"`
	// UserInfo with additional information to show to the user.
	UserInfo string `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *OpResponse) Reset() {
	*x = OpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_common_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpResponse) ProtoMessage() {}

func (x *OpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_common_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpResponse.ProtoReflect.Descriptor instead.
func (*OpResponse) Descriptor() ([]byte, []int) {
	return file_catalog_common_entities_proto_rawDescGZIP(), []int{1}
}

func (x *OpResponse) GetStatus() OpStatus {
	if x != nil {
		return x.Status
	}
	return OpStatus_UNKNOWN
}

func (x *OpResponse) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

func (x *OpResponse) GetUserInfo() string {
	if x != nil {
		return x.UserInfo
	}
	return ""
}

var File_catalog_common_entities_proto protoreflect.FileDescriptor

var file_catalog_common_entities_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22,
	0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x7c, 0x0a, 0x0a, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x24, 0x0a,
	0x08, 0x4f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x61, 0x70, 0x70, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x67,
	0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_common_entities_proto_rawDescOnce sync.Once
	file_catalog_common_entities_proto_rawDescData = file_catalog_common_entities_proto_rawDesc
)

func file_catalog_common_entities_proto_rawDescGZIP() []byte {
	file_catalog_common_entities_proto_rawDescOnce.Do(func() {
		file_catalog_common_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_common_entities_proto_rawDescData)
	})
	return file_catalog_common_entities_proto_rawDescData
}

var file_catalog_common_entities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_catalog_common_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_catalog_common_entities_proto_goTypes = []interface{}{
	(OpStatus)(0),        // 0: catalog_common.OpStatus
	(*EmptyRequest)(nil), // 1: catalog_common.EmptyRequest
	(*OpResponse)(nil),   // 2: catalog_common.OpResponse
}
var file_catalog_common_entities_proto_depIdxs = []int32{
	0, // 0: catalog_common.OpResponse.status:type_name -> catalog_common.OpStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_catalog_common_entities_proto_init() }
func file_catalog_common_entities_proto_init() {
	if File_catalog_common_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_common_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_catalog_common_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpResponse); i {
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
			RawDescriptor: file_catalog_common_entities_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_catalog_common_entities_proto_goTypes,
		DependencyIndexes: file_catalog_common_entities_proto_depIdxs,
		EnumInfos:         file_catalog_common_entities_proto_enumTypes,
		MessageInfos:      file_catalog_common_entities_proto_msgTypes,
	}.Build()
	File_catalog_common_entities_proto = out.File
	file_catalog_common_entities_proto_rawDesc = nil
	file_catalog_common_entities_proto_goTypes = nil
	file_catalog_common_entities_proto_depIdxs = nil
}
