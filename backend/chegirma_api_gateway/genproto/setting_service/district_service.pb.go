// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: district_service.proto

package setting_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_district_service_proto protoreflect.FileDescriptor

var file_district_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe1, 0x04, 0x0a, 0x0f,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x79, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79,
	0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x53, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x53, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x54, 0x61, 0x78, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x54, 0x61, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_district_service_proto_goTypes = []interface{}{
	(*CreateUpdateDistrict)(nil),      // 0: genproto.CreateUpdateDistrict
	(*GetReq)(nil),                    // 1: genproto.GetReq
	(*GetAllDistrictsRequest)(nil),    // 2: genproto.GetAllDistrictsRequest
	(*GetAllByCityRegionRequest)(nil), // 3: genproto.GetAllByCityRegionRequest
	(*SSExistsRequest)(nil),           // 4: genproto.SSExistsRequest
	(*DistrictCostRequest)(nil),       // 5: genproto.DistrictCostRequest
	(*DistrictTaxRequest)(nil),        // 6: genproto.DistrictTaxRequest
	(*CreatedResponse)(nil),           // 7: genproto.CreatedResponse
	(*District)(nil),                  // 8: genproto.District
	(*GetAllDistrictsResponse)(nil),   // 9: genproto.GetAllDistrictsResponse
	(*empty.Empty)(nil),               // 10: google.protobuf.Empty
	(*SSExistsResponse)(nil),          // 11: genproto.SSExistsResponse
}
var file_district_service_proto_depIdxs = []int32{
	0,  // 0: genproto.DistrictService.Create:input_type -> genproto.CreateUpdateDistrict
	1,  // 1: genproto.DistrictService.Get:input_type -> genproto.GetReq
	2,  // 2: genproto.DistrictService.GetAll:input_type -> genproto.GetAllDistrictsRequest
	3,  // 3: genproto.DistrictService.GetAllByCityRegion:input_type -> genproto.GetAllByCityRegionRequest
	0,  // 4: genproto.DistrictService.Update:input_type -> genproto.CreateUpdateDistrict
	4,  // 5: genproto.DistrictService.DistrictExists:input_type -> genproto.SSExistsRequest
	5,  // 6: genproto.DistrictService.UpdateDistrictCost:input_type -> genproto.DistrictCostRequest
	6,  // 7: genproto.DistrictService.UpdateDistrictTax:input_type -> genproto.DistrictTaxRequest
	7,  // 8: genproto.DistrictService.Create:output_type -> genproto.CreatedResponse
	8,  // 9: genproto.DistrictService.Get:output_type -> genproto.District
	9,  // 10: genproto.DistrictService.GetAll:output_type -> genproto.GetAllDistrictsResponse
	9,  // 11: genproto.DistrictService.GetAllByCityRegion:output_type -> genproto.GetAllDistrictsResponse
	10, // 12: genproto.DistrictService.Update:output_type -> google.protobuf.Empty
	11, // 13: genproto.DistrictService.DistrictExists:output_type -> genproto.SSExistsResponse
	10, // 14: genproto.DistrictService.UpdateDistrictCost:output_type -> google.protobuf.Empty
	10, // 15: genproto.DistrictService.UpdateDistrictTax:output_type -> google.protobuf.Empty
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_district_service_proto_init() }
func file_district_service_proto_init() {
	if File_district_service_proto != nil {
		return
	}
	file_district_proto_init()
	file_city_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_district_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_district_service_proto_goTypes,
		DependencyIndexes: file_district_service_proto_depIdxs,
	}.Build()
	File_district_service_proto = out.File
	file_district_service_proto_rawDesc = nil
	file_district_service_proto_goTypes = nil
	file_district_service_proto_depIdxs = nil
}
