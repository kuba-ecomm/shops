// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service.proto

package shops

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

type ShopEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShopEmpty) Reset() {
	*x = ShopEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopEmpty) ProtoMessage() {}

func (x *ShopEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopEmpty.ProtoReflect.Descriptor instead.
func (*ShopEmpty) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

// UserGetter is
type ShopGetter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Getter:
	//
	//	*ShopGetter_Uuid
	//	*ShopGetter_Code
	//	*ShopGetter_Name
	Getter isShopGetter_Getter `protobuf_oneof:"getter"`
}

func (x *ShopGetter) Reset() {
	*x = ShopGetter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopGetter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopGetter) ProtoMessage() {}

func (x *ShopGetter) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopGetter.ProtoReflect.Descriptor instead.
func (*ShopGetter) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (m *ShopGetter) GetGetter() isShopGetter_Getter {
	if m != nil {
		return m.Getter
	}
	return nil
}

func (x *ShopGetter) GetUuid() []byte {
	if x, ok := x.GetGetter().(*ShopGetter_Uuid); ok {
		return x.Uuid
	}
	return nil
}

func (x *ShopGetter) GetCode() string {
	if x, ok := x.GetGetter().(*ShopGetter_Code); ok {
		return x.Code
	}
	return ""
}

func (x *ShopGetter) GetName() string {
	if x, ok := x.GetGetter().(*ShopGetter_Name); ok {
		return x.Name
	}
	return ""
}

type isShopGetter_Getter interface {
	isShopGetter_Getter()
}

type ShopGetter_Uuid struct {
	Uuid []byte `protobuf:"bytes,1,opt,name=uuid,proto3,oneof"`
}

type ShopGetter_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

type ShopGetter_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

func (*ShopGetter_Uuid) isShopGetter_Getter() {}

func (*ShopGetter_Code) isShopGetter_Getter() {}

func (*ShopGetter_Name) isShopGetter_Getter() {}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x70, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x58, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x74, 0x65, 0x72, 0x32, 0x3a, 0x0a, 0x0c,
	0x53, 0x68, 0x6f, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_proto_goTypes = []interface{}{
	(*ShopEmpty)(nil),  // 0: models.ShopEmpty
	(*ShopGetter)(nil), // 1: models.ShopGetter
	(*Shop)(nil),       // 2: models.Shop
}
var file_service_proto_depIdxs = []int32{
	1, // 0: models.ShopsService.ShopBy:input_type -> models.ShopGetter
	2, // 1: models.ShopsService.ShopBy:output_type -> models.Shop
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopEmpty); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopGetter); i {
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
	file_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ShopGetter_Uuid)(nil),
		(*ShopGetter_Code)(nil),
		(*ShopGetter_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}