// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: netconfig.proto

package api

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

type UpdateHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewHostname string `protobuf:"bytes,1,opt,name=newHostname,proto3" json:"newHostname,omitempty"`
}

func (x *UpdateHostnameRequest) Reset() {
	*x = UpdateHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHostnameRequest) ProtoMessage() {}

func (x *UpdateHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_netconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHostnameRequest.ProtoReflect.Descriptor instead.
func (*UpdateHostnameRequest) Descriptor() ([]byte, []int) {
	return file_netconfig_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateHostnameRequest) GetNewHostname() string {
	if x != nil {
		return x.NewHostname
	}
	return ""
}

type AddDNSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewDns string `protobuf:"bytes,1,opt,name=newDns,proto3" json:"newDns,omitempty"`
}

func (x *AddDNSRequest) Reset() {
	*x = AddDNSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDNSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDNSRequest) ProtoMessage() {}

func (x *AddDNSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_netconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDNSRequest.ProtoReflect.Descriptor instead.
func (*AddDNSRequest) Descriptor() ([]byte, []int) {
	return file_netconfig_proto_rawDescGZIP(), []int{1}
}

func (x *AddDNSRequest) GetNewDns() string {
	if x != nil {
		return x.NewDns
	}
	return ""
}

type DeleteDNSRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsToRemove string `protobuf:"bytes,1,opt,name=dnsToRemove,proto3" json:"dnsToRemove,omitempty"`
}

func (x *DeleteDNSRequst) Reset() {
	*x = DeleteDNSRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDNSRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDNSRequst) ProtoMessage() {}

func (x *DeleteDNSRequst) ProtoReflect() protoreflect.Message {
	mi := &file_netconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDNSRequst.ProtoReflect.Descriptor instead.
func (*DeleteDNSRequst) Descriptor() ([]byte, []int) {
	return file_netconfig_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDNSRequst) GetDnsToRemove() string {
	if x != nil {
		return x.DnsToRemove
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_netconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_netconfig_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type DnsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *DnsListResponse) Reset() {
	*x = DnsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netconfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsListResponse) ProtoMessage() {}

func (x *DnsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_netconfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsListResponse.ProtoReflect.Descriptor instead.
func (*DnsListResponse) Descriptor() ([]byte, []int) {
	return file_netconfig_proto_rawDescGZIP(), []int{4}
}

func (x *DnsListResponse) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netconfig_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_netconfig_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_netconfig_proto_rawDescGZIP(), []int{5}
}

var File_netconfig_proto protoreflect.FileDescriptor

var file_netconfig_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x39, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x44, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x44, 0x6e, 0x73, 0x22, 0x33, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x6e, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x6e, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22,
	0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x2f, 0x0a, 0x0f, 0x44, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd5, 0x01, 0x0a, 0x09,
	0x4e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x06, 0x41, 0x64, 0x64,
	0x44, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x4e, 0x53,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x6e, 0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x6e, 0x73, 0x12,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x61, 0x64, 0x69, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x32, 0x31, 0x32, 0x37, 0x2f,
	0x6e, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_netconfig_proto_rawDescOnce sync.Once
	file_netconfig_proto_rawDescData = file_netconfig_proto_rawDesc
)

func file_netconfig_proto_rawDescGZIP() []byte {
	file_netconfig_proto_rawDescOnce.Do(func() {
		file_netconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_netconfig_proto_rawDescData)
	})
	return file_netconfig_proto_rawDescData
}

var file_netconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_netconfig_proto_goTypes = []interface{}{
	(*UpdateHostnameRequest)(nil), // 0: api.UpdateHostnameRequest
	(*AddDNSRequest)(nil),         // 1: api.AddDNSRequest
	(*DeleteDNSRequst)(nil),       // 2: api.DeleteDNSRequst
	(*Response)(nil),              // 3: api.Response
	(*DnsListResponse)(nil),       // 4: api.DnsListResponse
	(*Empty)(nil),                 // 5: api.Empty
}
var file_netconfig_proto_depIdxs = []int32{
	1, // 0: api.NetConfig.AddDns:input_type -> api.AddDNSRequest
	2, // 1: api.NetConfig.DeleteDns:input_type -> api.DeleteDNSRequst
	0, // 2: api.NetConfig.UpdateHostname:input_type -> api.UpdateHostnameRequest
	5, // 3: api.NetConfig.GetAllDns:input_type -> api.Empty
	5, // 4: api.NetConfig.AddDns:output_type -> api.Empty
	5, // 5: api.NetConfig.DeleteDns:output_type -> api.Empty
	5, // 6: api.NetConfig.UpdateHostname:output_type -> api.Empty
	4, // 7: api.NetConfig.GetAllDns:output_type -> api.DnsListResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_netconfig_proto_init() }
func file_netconfig_proto_init() {
	if File_netconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_netconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHostnameRequest); i {
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
		file_netconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDNSRequest); i {
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
		file_netconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDNSRequst); i {
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
		file_netconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_netconfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsListResponse); i {
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
		file_netconfig_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_netconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_netconfig_proto_goTypes,
		DependencyIndexes: file_netconfig_proto_depIdxs,
		MessageInfos:      file_netconfig_proto_msgTypes,
	}.Build()
	File_netconfig_proto = out.File
	file_netconfig_proto_rawDesc = nil
	file_netconfig_proto_goTypes = nil
	file_netconfig_proto_depIdxs = nil
}
