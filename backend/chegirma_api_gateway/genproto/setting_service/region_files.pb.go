// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: region_files.proto

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

type RegionFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RegionId string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Url      string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Comment  string `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	User     string `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Type     string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *RegionFiles) Reset() {
	*x = RegionFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionFiles) ProtoMessage() {}

func (x *RegionFiles) ProtoReflect() protoreflect.Message {
	mi := &file_region_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionFiles.ProtoReflect.Descriptor instead.
func (*RegionFiles) Descriptor() ([]byte, []int) {
	return file_region_files_proto_rawDescGZIP(), []int{0}
}

func (x *RegionFiles) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegionFiles) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *RegionFiles) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *RegionFiles) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RegionFiles) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *RegionFiles) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RegionFiles) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetRegionFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetRegionFilesRequest) Reset() {
	*x = GetRegionFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionFilesRequest) ProtoMessage() {}

func (x *GetRegionFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionFilesRequest.ProtoReflect.Descriptor instead.
func (*GetRegionFilesRequest) Descriptor() ([]byte, []int) {
	return file_region_files_proto_rawDescGZIP(), []int{1}
}

func (x *GetRegionFilesRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *GetRegionFilesRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetAllRegionFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Search   string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Page     uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit    uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Type     string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetAllRegionFilesRequest) Reset() {
	*x = GetAllRegionFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_files_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRegionFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRegionFilesRequest) ProtoMessage() {}

func (x *GetAllRegionFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_files_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRegionFilesRequest.ProtoReflect.Descriptor instead.
func (*GetAllRegionFilesRequest) Descriptor() ([]byte, []int) {
	return file_region_files_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllRegionFilesRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *GetAllRegionFilesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllRegionFilesRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllRegionFilesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllRegionFilesRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetAllRegionFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionFiles []*RegionFiles `protobuf:"bytes,1,rep,name=region_files,json=regionFiles,proto3" json:"region_files,omitempty"`
	Count       uint32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllRegionFilesResponse) Reset() {
	*x = GetAllRegionFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_files_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRegionFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRegionFilesResponse) ProtoMessage() {}

func (x *GetAllRegionFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_region_files_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRegionFilesResponse.ProtoReflect.Descriptor instead.
func (*GetAllRegionFilesResponse) Descriptor() ([]byte, []int) {
	return file_region_files_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllRegionFilesResponse) GetRegionFiles() []*RegionFiles {
	if x != nil {
		return x.RegionFiles
	}
	return nil
}

func (x *GetAllRegionFilesResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_region_files_proto protoreflect.FileDescriptor

var file_region_files_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x01, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_region_files_proto_rawDescOnce sync.Once
	file_region_files_proto_rawDescData = file_region_files_proto_rawDesc
)

func file_region_files_proto_rawDescGZIP() []byte {
	file_region_files_proto_rawDescOnce.Do(func() {
		file_region_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_region_files_proto_rawDescData)
	})
	return file_region_files_proto_rawDescData
}

var file_region_files_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_region_files_proto_goTypes = []interface{}{
	(*RegionFiles)(nil),               // 0: genproto.RegionFiles
	(*GetRegionFilesRequest)(nil),     // 1: genproto.GetRegionFilesRequest
	(*GetAllRegionFilesRequest)(nil),  // 2: genproto.GetAllRegionFilesRequest
	(*GetAllRegionFilesResponse)(nil), // 3: genproto.GetAllRegionFilesResponse
}
var file_region_files_proto_depIdxs = []int32{
	0, // 0: genproto.GetAllRegionFilesResponse.region_files:type_name -> genproto.RegionFiles
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_region_files_proto_init() }
func file_region_files_proto_init() {
	if File_region_files_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_region_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionFiles); i {
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
		file_region_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionFilesRequest); i {
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
		file_region_files_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRegionFilesRequest); i {
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
		file_region_files_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRegionFilesResponse); i {
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
			RawDescriptor: file_region_files_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_region_files_proto_goTypes,
		DependencyIndexes: file_region_files_proto_depIdxs,
		MessageInfos:      file_region_files_proto_msgTypes,
	}.Build()
	File_region_files_proto = out.File
	file_region_files_proto_rawDesc = nil
	file_region_files_proto_goTypes = nil
	file_region_files_proto_depIdxs = nil
}