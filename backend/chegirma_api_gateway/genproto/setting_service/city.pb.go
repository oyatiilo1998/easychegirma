// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: city.proto

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

type CreatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatedResponse) Reset() {
	*x = CreatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedResponse) ProtoMessage() {}

func (x *CreatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedResponse.ProtoReflect.Descriptor instead.
func (*CreatedResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{0}
}

func (x *CreatedResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{1}
}

func (x *GetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RuName string `protobuf:"bytes,2,opt,name=ru_name,json=ruName,proto3" json:"ru_name,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Soato  uint32 `protobuf:"varint,4,opt,name=soato,proto3" json:"soato,omitempty"`
	Code   uint32 `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{2}
}

func (x *City) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *City) GetRuName() string {
	if x != nil {
		return x.RuName
	}
	return ""
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *City) GetSoato() uint32 {
	if x != nil {
		return x.Soato
	}
	return 0
}

func (x *City) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetAllCitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Soato uint32 `protobuf:"varint,2,opt,name=soato,proto3" json:"soato,omitempty"`
	Page  uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllCitiesRequest) Reset() {
	*x = GetAllCitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCitiesRequest) ProtoMessage() {}

func (x *GetAllCitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCitiesRequest.ProtoReflect.Descriptor instead.
func (*GetAllCitiesRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCitiesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllCitiesRequest) GetSoato() uint32 {
	if x != nil {
		return x.Soato
	}
	return 0
}

func (x *GetAllCitiesRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllCitiesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllCitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  uint32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Cities []*City `protobuf:"bytes,2,rep,name=cities,proto3" json:"cities,omitempty"`
}

func (x *GetAllCitiesResponse) Reset() {
	*x = GetAllCitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCitiesResponse) ProtoMessage() {}

func (x *GetAllCitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCitiesResponse.ProtoReflect.Descriptor instead.
func (*GetAllCitiesResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllCitiesResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllCitiesResponse) GetCities() []*City {
	if x != nil {
		return x.Cities
	}
	return nil
}

type SSExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SSExistsRequest) Reset() {
	*x = SSExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSExistsRequest) ProtoMessage() {}

func (x *SSExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSExistsRequest.ProtoReflect.Descriptor instead.
func (*SSExistsRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{5}
}

func (x *SSExistsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SSExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *SSExistsResponse) Reset() {
	*x = SSExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSExistsResponse) ProtoMessage() {}

func (x *SSExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSExistsResponse.ProtoReflect.Descriptor instead.
func (*SSExistsResponse) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{6}
}

func (x *SSExistsResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

type GetByCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetByCodeRequest) Reset() {
	*x = GetByCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByCodeRequest) ProtoMessage() {}

func (x *GetByCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByCodeRequest.ProtoReflect.Descriptor instead.
func (*GetByCodeRequest) Descriptor() ([]byte, []int) {
	return file_city_proto_rawDescGZIP(), []int{7}
}

func (x *GetByCodeRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_city_proto protoreflect.FileDescriptor

var file_city_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x61, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x6f, 0x61, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x69, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x6f, 0x61, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x6f,
	0x61, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x54, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x63,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x53, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x53, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_proto_rawDescOnce sync.Once
	file_city_proto_rawDescData = file_city_proto_rawDesc
)

func file_city_proto_rawDescGZIP() []byte {
	file_city_proto_rawDescOnce.Do(func() {
		file_city_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_proto_rawDescData)
	})
	return file_city_proto_rawDescData
}

var file_city_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_proto_goTypes = []interface{}{
	(*CreatedResponse)(nil),      // 0: genproto.CreatedResponse
	(*GetReq)(nil),               // 1: genproto.GetReq
	(*City)(nil),                 // 2: genproto.City
	(*GetAllCitiesRequest)(nil),  // 3: genproto.GetAllCitiesRequest
	(*GetAllCitiesResponse)(nil), // 4: genproto.GetAllCitiesResponse
	(*SSExistsRequest)(nil),      // 5: genproto.SSExistsRequest
	(*SSExistsResponse)(nil),     // 6: genproto.SSExistsResponse
	(*GetByCodeRequest)(nil),     // 7: genproto.GetByCodeRequest
}
var file_city_proto_depIdxs = []int32{
	2, // 0: genproto.GetAllCitiesResponse.cities:type_name -> genproto.City
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_proto_init() }
func file_city_proto_init() {
	if File_city_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedResponse); i {
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
		file_city_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_city_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*City); i {
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
		file_city_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCitiesRequest); i {
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
		file_city_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCitiesResponse); i {
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
		file_city_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSExistsRequest); i {
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
		file_city_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSExistsResponse); i {
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
		file_city_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByCodeRequest); i {
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
			RawDescriptor: file_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_proto_goTypes,
		DependencyIndexes: file_city_proto_depIdxs,
		MessageInfos:      file_city_proto_msgTypes,
	}.Build()
	File_city_proto = out.File
	file_city_proto_rawDesc = nil
	file_city_proto_goTypes = nil
	file_city_proto_depIdxs = nil
}
