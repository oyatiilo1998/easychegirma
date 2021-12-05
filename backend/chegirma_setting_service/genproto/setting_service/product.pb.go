// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: product.proto

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url            string    `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name           string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl       string    `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	DiscountAmount uint32    `protobuf:"varint,5,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	PriceBefore    uint64    `protobuf:"varint,6,opt,name=price_before,json=priceBefore,proto3" json:"price_before,omitempty"`
	PriceAfter     uint64    `protobuf:"varint,7,opt,name=price_after,json=priceAfter,proto3" json:"price_after,omitempty"`
	Category       *Category `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	FromTime       string    `protobuf:"bytes,9,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime         string    `protobuf:"bytes,10,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	Owner          *User     `protobuf:"bytes,11,opt,name=owner,proto3" json:"owner,omitempty"`
	Description    string    `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Product) GetDiscountAmount() uint32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *Product) GetPriceBefore() uint64 {
	if x != nil {
		return x.PriceBefore
	}
	return 0
}

func (x *Product) GetPriceAfter() uint64 {
	if x != nil {
		return x.PriceAfter
	}
	return 0
}

func (x *Product) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Product) GetFromTime() string {
	if x != nil {
		return x.FromTime
	}
	return ""
}

func (x *Product) GetToTime() string {
	if x != nil {
		return x.ToTime
	}
	return ""
}

func (x *Product) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateUpdateProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url            string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl       string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	DiscountAmount uint32 `protobuf:"varint,5,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	PriceBefore    uint64 `protobuf:"varint,6,opt,name=price_before,json=priceBefore,proto3" json:"price_before,omitempty"`
	PriceAfter     uint64 `protobuf:"varint,7,opt,name=price_after,json=priceAfter,proto3" json:"price_after,omitempty"`
	CategoryId     string `protobuf:"bytes,8,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	FromTime       string `protobuf:"bytes,9,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime         string `protobuf:"bytes,10,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	OwnerId        string `protobuf:"bytes,11,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Description    string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateUpdateProduct) Reset() {
	*x = CreateUpdateProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUpdateProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUpdateProduct) ProtoMessage() {}

func (x *CreateUpdateProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUpdateProduct.ProtoReflect.Descriptor instead.
func (*CreateUpdateProduct) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUpdateProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateUpdateProduct) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateUpdateProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUpdateProduct) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateUpdateProduct) GetDiscountAmount() uint32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *CreateUpdateProduct) GetPriceBefore() uint64 {
	if x != nil {
		return x.PriceBefore
	}
	return 0
}

func (x *CreateUpdateProduct) GetPriceAfter() uint64 {
	if x != nil {
		return x.PriceAfter
	}
	return 0
}

func (x *CreateUpdateProduct) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CreateUpdateProduct) GetFromTime() string {
	if x != nil {
		return x.FromTime
	}
	return ""
}

func (x *CreateUpdateProduct) GetToTime() string {
	if x != nil {
		return x.ToTime
	}
	return ""
}

func (x *CreateUpdateProduct) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateUpdateProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Page       uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit      uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	OwnerId    string `protobuf:"bytes,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *GetAllProductsRequest) Reset() {
	*x = GetAllProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductsRequest) ProtoMessage() {}

func (x *GetAllProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductsRequest.ProtoReflect.Descriptor instead.
func (*GetAllProductsRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllProductsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllProductsRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *GetAllProductsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllProductsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllProductsRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetAllProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    uint32     `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetAllProductsResponse) Reset() {
	*x = GetAllProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductsResponse) ProtoMessage() {}

func (x *GetAllProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductsResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductsResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllProductsResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetAllByCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAllByCategoryRequest) Reset() {
	*x = GetAllByCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByCategoryRequest) ProtoMessage() {}

func (x *GetAllByCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetAllByCategoryRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllByCategoryRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *GetAllByCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xe9, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f,
	0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x5d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x4e,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_proto_goTypes = []interface{}{
	(*Product)(nil),                 // 0: genproto.Product
	(*CreateUpdateProduct)(nil),     // 1: genproto.CreateUpdateProduct
	(*GetAllProductsRequest)(nil),   // 2: genproto.GetAllProductsRequest
	(*GetAllProductsResponse)(nil),  // 3: genproto.GetAllProductsResponse
	(*GetAllByCategoryRequest)(nil), // 4: genproto.GetAllByCategoryRequest
	(*DeleteProductRequest)(nil),    // 5: genproto.DeleteProductRequest
	(*Category)(nil),                // 6: genproto.Category
	(*User)(nil),                    // 7: genproto.User
}
var file_product_proto_depIdxs = []int32{
	6, // 0: genproto.Product.category:type_name -> genproto.Category
	7, // 1: genproto.Product.owner:type_name -> genproto.User
	0, // 2: genproto.GetAllProductsResponse.products:type_name -> genproto.Product
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	file_category_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUpdateProduct); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductsRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductsResponse); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByCategoryRequest); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductRequest); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}