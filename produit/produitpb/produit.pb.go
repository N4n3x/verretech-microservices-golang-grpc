// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: proto/produit.proto

package produitpb

import (
	proto "./proto"
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

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointRetrait *proto.PointRetrait `protobuf:"bytes,1,opt,name=pointRetrait,proto3" json:"pointRetrait,omitempty"`
	Qte          int32               `protobuf:"varint,2,opt,name=qte,proto3" json:"qte,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{0}
}

func (x *Stock) GetPointRetrait() *proto.PointRetrait {
	if x != nil {
		return x.PointRetrait
	}
	return nil
}

func (x *Stock) GetQte() int32 {
	if x != nil {
		return x.Qte
	}
	return 0
}

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{1}
}

func (x *Photo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Produit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref         string   `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Prix        float32  `protobuf:"fixed32,3,opt,name=prix,proto3" json:"prix,omitempty"`
	Photos      []*Photo `protobuf:"bytes,4,rep,name=photos,proto3" json:"photos,omitempty"`
	Stocks      []*Stock `protobuf:"bytes,5,rep,name=stocks,proto3" json:"stocks,omitempty"`
	Tags        []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Produit) Reset() {
	*x = Produit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Produit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Produit) ProtoMessage() {}

func (x *Produit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Produit.ProtoReflect.Descriptor instead.
func (*Produit) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{2}
}

func (x *Produit) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Produit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Produit) GetPrix() float32 {
	if x != nil {
		return x.Prix
	}
	return 0
}

func (x *Produit) GetPhotos() []*Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Produit) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

func (x *Produit) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type ListProduits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Produits []*Produit `protobuf:"bytes,1,rep,name=produits,proto3" json:"produits,omitempty"`
}

func (x *ListProduits) Reset() {
	*x = ListProduits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProduits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProduits) ProtoMessage() {}

func (x *ListProduits) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProduits.ProtoReflect.Descriptor instead.
func (*ListProduits) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{3}
}

func (x *ListProduits) GetProduits() []*Produit {
	if x != nil {
		return x.Produits
	}
	return nil
}

type ProduitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listproduits *ListProduits `protobuf:"bytes,1,opt,name=listproduits,proto3" json:"listproduits,omitempty"`
}

func (x *ProduitsRequest) Reset() {
	*x = ProduitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduitsRequest) ProtoMessage() {}

func (x *ProduitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduitsRequest.ProtoReflect.Descriptor instead.
func (*ProduitsRequest) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{4}
}

func (x *ProduitsRequest) GetListproduits() *ListProduits {
	if x != nil {
		return x.Listproduits
	}
	return nil
}

type GetAllProduitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllProduitsRequest) Reset() {
	*x = GetAllProduitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProduitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProduitsRequest) ProtoMessage() {}

func (x *GetAllProduitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProduitsRequest.ProtoReflect.Descriptor instead.
func (*GetAllProduitsRequest) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{5}
}

type ProduitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Produit *Produit `protobuf:"bytes,1,opt,name=produit,proto3" json:"produit,omitempty"`
}

func (x *ProduitResponse) Reset() {
	*x = ProduitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduitResponse) ProtoMessage() {}

func (x *ProduitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduitResponse.ProtoReflect.Descriptor instead.
func (*ProduitResponse) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{6}
}

func (x *ProduitResponse) GetProduit() *Produit {
	if x != nil {
		return x.Produit
	}
	return nil
}

type ProduitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listproduits *ListProduits `protobuf:"bytes,1,opt,name=listproduits,proto3" json:"listproduits,omitempty"`
}

func (x *ProduitsResponse) Reset() {
	*x = ProduitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_produit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduitsResponse) ProtoMessage() {}

func (x *ProduitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_produit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduitsResponse.ProtoReflect.Descriptor instead.
func (*ProduitsResponse) Descriptor() ([]byte, []int) {
	return file_proto_produit_proto_rawDescGZIP(), []int{7}
}

func (x *ProduitsResponse) GetListproduits() *ListProduits {
	if x != nil {
		return x.Listproduits
	}
	return nil
}

var File_proto_produit_proto protoreflect.FileDescriptor

var file_proto_produit_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x1a, 0x18,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x52, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x71, 0x74, 0x65, 0x22, 0x19, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xb5,
	0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x72, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x72,
	0x69, 0x78, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x3c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x69, 0x74, 0x73, 0x22, 0x4c, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x69, 0x74, 0x73, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69,
	0x74, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x22, 0x4d, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x52, 0x0c, 0x6c, 0x69, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x32, 0xe9, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x12, 0x45, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x42, 0x79,
	0x52, 0x65, 0x66, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x69, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x69, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_produit_proto_rawDescOnce sync.Once
	file_proto_produit_proto_rawDescData = file_proto_produit_proto_rawDesc
)

func file_proto_produit_proto_rawDescGZIP() []byte {
	file_proto_produit_proto_rawDescOnce.Do(func() {
		file_proto_produit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_produit_proto_rawDescData)
	})
	return file_proto_produit_proto_rawDescData
}

var file_proto_produit_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_produit_proto_goTypes = []interface{}{
	(*Stock)(nil),                 // 0: produit.Stock
	(*Photo)(nil),                 // 1: produit.Photo
	(*Produit)(nil),               // 2: produit.Produit
	(*ListProduits)(nil),          // 3: produit.ListProduits
	(*ProduitsRequest)(nil),       // 4: produit.ProduitsRequest
	(*GetAllProduitsRequest)(nil), // 5: produit.GetAllProduitsRequest
	(*ProduitResponse)(nil),       // 6: produit.ProduitResponse
	(*ProduitsResponse)(nil),      // 7: produit.ProduitsResponse
	(*proto.PointRetrait)(nil),    // 8: pointRetrait.PointRetrait
}
var file_proto_produit_proto_depIdxs = []int32{
	8,  // 0: produit.Stock.pointRetrait:type_name -> pointRetrait.PointRetrait
	1,  // 1: produit.Produit.photos:type_name -> produit.Photo
	0,  // 2: produit.Produit.stocks:type_name -> produit.Stock
	2,  // 3: produit.ListProduits.produits:type_name -> produit.Produit
	3,  // 4: produit.ProduitsRequest.listproduits:type_name -> produit.ListProduits
	2,  // 5: produit.ProduitResponse.produit:type_name -> produit.Produit
	3,  // 6: produit.ProduitsResponse.listproduits:type_name -> produit.ListProduits
	4,  // 7: produit.ServiceProduit.UpdateProduits:input_type -> produit.ProduitsRequest
	5,  // 8: produit.ServiceProduit.GetAllProduits:input_type -> produit.GetAllProduitsRequest
	5,  // 9: produit.ServiceProduit.GetProduitByRef:input_type -> produit.GetAllProduitsRequest
	7,  // 10: produit.ServiceProduit.UpdateProduits:output_type -> produit.ProduitsResponse
	7,  // 11: produit.ServiceProduit.GetAllProduits:output_type -> produit.ProduitsResponse
	2,  // 12: produit.ServiceProduit.GetProduitByRef:output_type -> produit.Produit
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_produit_proto_init() }
func file_proto_produit_proto_init() {
	if File_proto_produit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_produit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_proto_produit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
		file_proto_produit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Produit); i {
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
		file_proto_produit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProduits); i {
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
		file_proto_produit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProduitsRequest); i {
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
		file_proto_produit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProduitsRequest); i {
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
		file_proto_produit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProduitResponse); i {
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
		file_proto_produit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProduitsResponse); i {
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
			RawDescriptor: file_proto_produit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_produit_proto_goTypes,
		DependencyIndexes: file_proto_produit_proto_depIdxs,
		MessageInfos:      file_proto_produit_proto_msgTypes,
	}.Build()
	File_proto_produit_proto = out.File
	file_proto_produit_proto_rawDesc = nil
	file_proto_produit_proto_goTypes = nil
	file_proto_produit_proto_depIdxs = nil
}
