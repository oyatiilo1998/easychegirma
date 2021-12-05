// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: global_entity.proto

package setting_service

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

type CreateOrUpdateGlobalEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Code        int64  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreateOrUpdateGlobalEntity) Reset() {
	*x = CreateOrUpdateGlobalEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateGlobalEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateGlobalEntity) ProtoMessage() {}

func (x *CreateOrUpdateGlobalEntity) ProtoReflect() protoreflect.Message {
	mi := &file_global_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateGlobalEntity.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateGlobalEntity) Descriptor() ([]byte, []int) {
	return file_global_entity_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateGlobalEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrUpdateGlobalEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrUpdateGlobalEntity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateOrUpdateGlobalEntity) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GlobalEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Code        int64  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GlobalEntity) Reset() {
	*x = GlobalEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalEntity) ProtoMessage() {}

func (x *GlobalEntity) ProtoReflect() protoreflect.Message {
	mi := &file_global_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalEntity.ProtoReflect.Descriptor instead.
func (*GlobalEntity) Descriptor() ([]byte, []int) {
	return file_global_entity_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GlobalEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GlobalEntity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GlobalEntity) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GlobalEntity) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GlobalEntity) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetAllGlobalEntitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetAllGlobalEntitiesRequest) Reset() {
	*x = GetAllGlobalEntitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGlobalEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGlobalEntitiesRequest) ProtoMessage() {}

func (x *GetAllGlobalEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_global_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGlobalEntitiesRequest.ProtoReflect.Descriptor instead.
func (*GetAllGlobalEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_global_entity_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllGlobalEntitiesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllGlobalEntitiesRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetAllGlobalEntitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalEntities []*GlobalEntity `protobuf:"bytes,1,rep,name=global_entities,json=globalEntities,proto3" json:"global_entities,omitempty"`
	Count          int64           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllGlobalEntitiesResponse) Reset() {
	*x = GetAllGlobalEntitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGlobalEntitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGlobalEntitiesResponse) ProtoMessage() {}

func (x *GetAllGlobalEntitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGlobalEntitiesResponse.ProtoReflect.Descriptor instead.
func (*GetAllGlobalEntitiesResponse) Descriptor() ([]byte, []int) {
	return file_global_entity_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllGlobalEntitiesResponse) GetGlobalEntities() []*GlobalEntity {
	if x != nil {
		return x.GlobalEntities
	}
	return nil
}

func (x *GetAllGlobalEntitiesResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_global_entity_proto protoreflect.FileDescriptor

var file_global_entity_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x76, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x47, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x75, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_global_entity_proto_rawDescOnce sync.Once
	file_global_entity_proto_rawDescData = file_global_entity_proto_rawDesc
)

func file_global_entity_proto_rawDescGZIP() []byte {
	file_global_entity_proto_rawDescOnce.Do(func() {
		file_global_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_entity_proto_rawDescData)
	})
	return file_global_entity_proto_rawDescData
}

var file_global_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_global_entity_proto_goTypes = []interface{}{
	(*CreateOrUpdateGlobalEntity)(nil),   // 0: genproto.CreateOrUpdateGlobalEntity
	(*GlobalEntity)(nil),                 // 1: genproto.GlobalEntity
	(*GetAllGlobalEntitiesRequest)(nil),  // 2: genproto.GetAllGlobalEntitiesRequest
	(*GetAllGlobalEntitiesResponse)(nil), // 3: genproto.GetAllGlobalEntitiesResponse
}
var file_global_entity_proto_depIdxs = []int32{
	1, // 0: genproto.GetAllGlobalEntitiesResponse.global_entities:type_name -> genproto.GlobalEntity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_global_entity_proto_init() }
func file_global_entity_proto_init() {
	if File_global_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_global_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateGlobalEntity); i {
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
		file_global_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalEntity); i {
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
		file_global_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGlobalEntitiesRequest); i {
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
		file_global_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGlobalEntitiesResponse); i {
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
			RawDescriptor: file_global_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_entity_proto_goTypes,
		DependencyIndexes: file_global_entity_proto_depIdxs,
		MessageInfos:      file_global_entity_proto_msgTypes,
	}.Build()
	File_global_entity_proto = out.File
	file_global_entity_proto_rawDesc = nil
	file_global_entity_proto_goTypes = nil
	file_global_entity_proto_depIdxs = nil
}