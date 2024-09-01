// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: grpcapi/implant.proto

package grpcapi

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

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_implant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_implant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_grpcapi_implant_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In  string `protobuf:"bytes,1,opt,name=In,proto3" json:"In,omitempty"`
	Out string `protobuf:"bytes,2,opt,name=Out,proto3" json:"Out,omitempty"`
	Id  string `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_implant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_implant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_grpcapi_implant_proto_rawDescGZIP(), []int{1}
}

func (x *Command) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *Command) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

func (x *Command) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Implants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Implants []*Identity `protobuf:"bytes,1,rep,name=Implants,proto3" json:"Implants,omitempty"`
}

func (x *Implants) Reset() {
	*x = Implants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_implant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Implants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Implants) ProtoMessage() {}

func (x *Implants) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_implant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Implants.ProtoReflect.Descriptor instead.
func (*Implants) Descriptor() ([]byte, []int) {
	return file_grpcapi_implant_proto_rawDescGZIP(), []int{2}
}

func (x *Implants) GetImplants() []*Identity {
	if x != nil {
		return x.Implants
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
		mi := &file_grpcapi_implant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_implant_proto_msgTypes[3]
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
	return file_grpcapi_implant_proto_rawDescGZIP(), []int{3}
}

var File_grpcapi_implant_proto protoreflect.FileDescriptor

var file_grpcapi_implant_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69,
	0x22, 0x2e, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x39, 0x0a,
	0x08, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x49, 0x6d, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0xa7, 0x01, 0x0a, 0x07, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x33, 0x0a,
	0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x11, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x37, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6d,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0x6b, 0x0a, 0x05, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e,
	0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x72, 0x74, 0x65, 0x67, 0x6f, 0x34, 0x32,
	0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x61, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcapi_implant_proto_rawDescOnce sync.Once
	file_grpcapi_implant_proto_rawDescData = file_grpcapi_implant_proto_rawDesc
)

func file_grpcapi_implant_proto_rawDescGZIP() []byte {
	file_grpcapi_implant_proto_rawDescOnce.Do(func() {
		file_grpcapi_implant_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcapi_implant_proto_rawDescData)
	})
	return file_grpcapi_implant_proto_rawDescData
}

var file_grpcapi_implant_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpcapi_implant_proto_goTypes = []interface{}{
	(*Identity)(nil), // 0: grpcapi.Identity
	(*Command)(nil),  // 1: grpcapi.Command
	(*Implants)(nil), // 2: grpcapi.Implants
	(*Empty)(nil),    // 3: grpcapi.Empty
}
var file_grpcapi_implant_proto_depIdxs = []int32{
	0, // 0: grpcapi.Implants.Implants:type_name -> grpcapi.Identity
	0, // 1: grpcapi.Implant.FetchCommand:input_type -> grpcapi.Identity
	1, // 2: grpcapi.Implant.SendOutput:input_type -> grpcapi.Command
	0, // 3: grpcapi.Implant.RegisterImplant:input_type -> grpcapi.Identity
	1, // 4: grpcapi.Admin.RunCommand:input_type -> grpcapi.Command
	3, // 5: grpcapi.Admin.GetImplants:input_type -> grpcapi.Empty
	1, // 6: grpcapi.Implant.FetchCommand:output_type -> grpcapi.Command
	3, // 7: grpcapi.Implant.SendOutput:output_type -> grpcapi.Empty
	0, // 8: grpcapi.Implant.RegisterImplant:output_type -> grpcapi.Identity
	1, // 9: grpcapi.Admin.RunCommand:output_type -> grpcapi.Command
	2, // 10: grpcapi.Admin.GetImplants:output_type -> grpcapi.Implants
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpcapi_implant_proto_init() }
func file_grpcapi_implant_proto_init() {
	if File_grpcapi_implant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcapi_implant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_grpcapi_implant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_grpcapi_implant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Implants); i {
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
		file_grpcapi_implant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_grpcapi_implant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpcapi_implant_proto_goTypes,
		DependencyIndexes: file_grpcapi_implant_proto_depIdxs,
		MessageInfos:      file_grpcapi_implant_proto_msgTypes,
	}.Build()
	File_grpcapi_implant_proto = out.File
	file_grpcapi_implant_proto_rawDesc = nil
	file_grpcapi_implant_proto_goTypes = nil
	file_grpcapi_implant_proto_depIdxs = nil
}
