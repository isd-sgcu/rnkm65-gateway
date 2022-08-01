// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: baan.proto

package proto

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

type BaanSize int32

const (
	BaanSize_S   BaanSize = 0
	BaanSize_M   BaanSize = 1
	BaanSize_L   BaanSize = 2
	BaanSize_XL  BaanSize = 3
	BaanSize_XXL BaanSize = 4
)

// Enum value maps for BaanSize.
var (
	BaanSize_name = map[int32]string{
		0: "S",
		1: "M",
		2: "L",
		3: "XL",
		4: "XXL",
	}
	BaanSize_value = map[string]int32{
		"S":   0,
		"M":   1,
		"L":   2,
		"XL":  3,
		"XXL": 4,
	}
)

func (x BaanSize) Enum() *BaanSize {
	p := new(BaanSize)
	*p = x
	return p
}

func (x BaanSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaanSize) Descriptor() protoreflect.EnumDescriptor {
	return file_baan_proto_enumTypes[0].Descriptor()
}

func (BaanSize) Type() protoreflect.EnumType {
	return &file_baan_proto_enumTypes[0]
}

func (x BaanSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaanSize.Descriptor instead.
func (BaanSize) EnumDescriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{0}
}

type Baan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NameTH        string   `protobuf:"bytes,2,opt,name=nameTH,proto3" json:"nameTH,omitempty"`
	DescriptionTH string   `protobuf:"bytes,3,opt,name=descriptionTH,proto3" json:"descriptionTH,omitempty"`
	NameEN        string   `protobuf:"bytes,4,opt,name=nameEN,proto3" json:"nameEN,omitempty"`
	DescriptionEN string   `protobuf:"bytes,5,opt,name=descriptionEN,proto3" json:"descriptionEN,omitempty"`
	Size          BaanSize `protobuf:"varint,6,opt,name=size,proto3,enum=baan.BaanSize" json:"size,omitempty"`
	Facebook      string   `protobuf:"bytes,7,opt,name=facebook,proto3" json:"facebook,omitempty"`
	FacebookUrl   string   `protobuf:"bytes,8,opt,name=facebookUrl,proto3" json:"facebookUrl,omitempty"`
	Instagram     string   `protobuf:"bytes,9,opt,name=instagram,proto3" json:"instagram,omitempty"`
	InstagramUrl  string   `protobuf:"bytes,10,opt,name=instagramUrl,proto3" json:"instagramUrl,omitempty"`
	Line          string   `protobuf:"bytes,11,opt,name=line,proto3" json:"line,omitempty"`
	LineUrl       string   `protobuf:"bytes,12,opt,name=lineUrl,proto3" json:"lineUrl,omitempty"`
	ImageUrl      string   `protobuf:"bytes,13,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
}

func (x *Baan) Reset() {
	*x = Baan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Baan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Baan) ProtoMessage() {}

func (x *Baan) ProtoReflect() protoreflect.Message {
	mi := &file_baan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Baan.ProtoReflect.Descriptor instead.
func (*Baan) Descriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{0}
}

func (x *Baan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Baan) GetNameTH() string {
	if x != nil {
		return x.NameTH
	}
	return ""
}

func (x *Baan) GetDescriptionTH() string {
	if x != nil {
		return x.DescriptionTH
	}
	return ""
}

func (x *Baan) GetNameEN() string {
	if x != nil {
		return x.NameEN
	}
	return ""
}

func (x *Baan) GetDescriptionEN() string {
	if x != nil {
		return x.DescriptionEN
	}
	return ""
}

func (x *Baan) GetSize() BaanSize {
	if x != nil {
		return x.Size
	}
	return BaanSize_S
}

func (x *Baan) GetFacebook() string {
	if x != nil {
		return x.Facebook
	}
	return ""
}

func (x *Baan) GetFacebookUrl() string {
	if x != nil {
		return x.FacebookUrl
	}
	return ""
}

func (x *Baan) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *Baan) GetInstagramUrl() string {
	if x != nil {
		return x.InstagramUrl
	}
	return ""
}

func (x *Baan) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

func (x *Baan) GetLineUrl() string {
	if x != nil {
		return x.LineUrl
	}
	return ""
}

func (x *Baan) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type BaanInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NameTH   string `protobuf:"bytes,2,opt,name=nameTH,proto3" json:"nameTH,omitempty"`
	NameEN   string `protobuf:"bytes,3,opt,name=nameEN,proto3" json:"nameEN,omitempty"`
	ImageUrl string `protobuf:"bytes,4,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
}

func (x *BaanInfo) Reset() {
	*x = BaanInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaanInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaanInfo) ProtoMessage() {}

func (x *BaanInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaanInfo.ProtoReflect.Descriptor instead.
func (*BaanInfo) Descriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{1}
}

func (x *BaanInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BaanInfo) GetNameTH() string {
	if x != nil {
		return x.NameTH
	}
	return ""
}

func (x *BaanInfo) GetNameEN() string {
	if x != nil {
		return x.NameEN
	}
	return ""
}

func (x *BaanInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type FindAllBaanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllBaanRequest) Reset() {
	*x = FindAllBaanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllBaanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllBaanRequest) ProtoMessage() {}

func (x *FindAllBaanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllBaanRequest.ProtoReflect.Descriptor instead.
func (*FindAllBaanRequest) Descriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{2}
}

type FindAllBaanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baans []*Baan `protobuf:"bytes,1,rep,name=baans,proto3" json:"baans,omitempty"`
}

func (x *FindAllBaanResponse) Reset() {
	*x = FindAllBaanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllBaanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllBaanResponse) ProtoMessage() {}

func (x *FindAllBaanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllBaanResponse.ProtoReflect.Descriptor instead.
func (*FindAllBaanResponse) Descriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllBaanResponse) GetBaans() []*Baan {
	if x != nil {
		return x.Baans
	}
	return nil
}

type FindOneBaanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneBaanRequest) Reset() {
	*x = FindOneBaanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneBaanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneBaanRequest) ProtoMessage() {}

func (x *FindOneBaanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneBaanRequest.ProtoReflect.Descriptor instead.
func (*FindOneBaanRequest) Descriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneBaanRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneBaanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baan *Baan `protobuf:"bytes,1,opt,name=baan,proto3" json:"baan,omitempty"`
}

func (x *FindOneBaanResponse) Reset() {
	*x = FindOneBaanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneBaanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneBaanResponse) ProtoMessage() {}

func (x *FindOneBaanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneBaanResponse.ProtoReflect.Descriptor instead.
func (*FindOneBaanResponse) Descriptor() ([]byte, []int) {
	return file_baan_proto_rawDescGZIP(), []int{5}
}

func (x *FindOneBaanResponse) GetBaan() *Baan {
	if x != nil {
		return x.Baan
	}
	return nil
}

var File_baan_proto protoreflect.FileDescriptor

var file_baan_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x61, 0x6e, 0x22, 0x80, 0x03, 0x0a, 0x04, 0x42, 0x61, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x61, 0x6d, 0x65, 0x54, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d,
	0x65, 0x54, 0x48, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x48, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d,
	0x65, 0x45, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45,
	0x4e, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x4e, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x61, 0x6e, 0x2e, 0x42, 0x61, 0x61,
	0x6e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x61, 0x63, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61,
	0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x67, 0x72, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x66, 0x0a, 0x08, 0x42, 0x61, 0x61, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x48, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d,
	0x65, 0x45, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45,
	0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x14, 0x0a,
	0x12, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x61,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x61,
	0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x61, 0x6e,
	0x2e, 0x42, 0x61, 0x61, 0x6e, 0x52, 0x05, 0x62, 0x61, 0x61, 0x6e, 0x73, 0x22, 0x24, 0x0a, 0x12,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x61, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x35, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x61, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x61, 0x61,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x61, 0x6e, 0x2e, 0x42,
	0x61, 0x61, 0x6e, 0x52, 0x04, 0x62, 0x61, 0x61, 0x6e, 0x2a, 0x30, 0x0a, 0x08, 0x42, 0x61, 0x61,
	0x6e, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x53, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01,
	0x4d, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x4c, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x58, 0x4c,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x58, 0x4c, 0x10, 0x04, 0x32, 0x99, 0x01, 0x0a, 0x0b,
	0x42, 0x61, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x61, 0x6e, 0x12, 0x18, 0x2e, 0x62, 0x61, 0x61,
	0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x61, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x42, 0x61, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x61, 0x61, 0x6e,
	0x12, 0x18, 0x2e, 0x62, 0x61, 0x61, 0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42,
	0x61, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x61,
	0x6e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x61, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x73, 0x72, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baan_proto_rawDescOnce sync.Once
	file_baan_proto_rawDescData = file_baan_proto_rawDesc
)

func file_baan_proto_rawDescGZIP() []byte {
	file_baan_proto_rawDescOnce.Do(func() {
		file_baan_proto_rawDescData = protoimpl.X.CompressGZIP(file_baan_proto_rawDescData)
	})
	return file_baan_proto_rawDescData
}

var file_baan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_baan_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_baan_proto_goTypes = []interface{}{
	(BaanSize)(0),               // 0: baan.BaanSize
	(*Baan)(nil),                // 1: baan.Baan
	(*BaanInfo)(nil),            // 2: baan.BaanInfo
	(*FindAllBaanRequest)(nil),  // 3: baan.FindAllBaanRequest
	(*FindAllBaanResponse)(nil), // 4: baan.FindAllBaanResponse
	(*FindOneBaanRequest)(nil),  // 5: baan.FindOneBaanRequest
	(*FindOneBaanResponse)(nil), // 6: baan.FindOneBaanResponse
}
var file_baan_proto_depIdxs = []int32{
	0, // 0: baan.Baan.size:type_name -> baan.BaanSize
	1, // 1: baan.FindAllBaanResponse.baans:type_name -> baan.Baan
	1, // 2: baan.FindOneBaanResponse.baan:type_name -> baan.Baan
	3, // 3: baan.BaanService.FindAllBaan:input_type -> baan.FindAllBaanRequest
	5, // 4: baan.BaanService.FindOneBaan:input_type -> baan.FindOneBaanRequest
	4, // 5: baan.BaanService.FindAllBaan:output_type -> baan.FindAllBaanResponse
	6, // 6: baan.BaanService.FindOneBaan:output_type -> baan.FindOneBaanResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_baan_proto_init() }
func file_baan_proto_init() {
	if File_baan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Baan); i {
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
		file_baan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaanInfo); i {
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
		file_baan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllBaanRequest); i {
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
		file_baan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllBaanResponse); i {
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
		file_baan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneBaanRequest); i {
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
		file_baan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneBaanResponse); i {
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
			RawDescriptor: file_baan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baan_proto_goTypes,
		DependencyIndexes: file_baan_proto_depIdxs,
		EnumInfos:         file_baan_proto_enumTypes,
		MessageInfos:      file_baan_proto_msgTypes,
	}.Build()
	File_baan_proto = out.File
	file_baan_proto_rawDesc = nil
	file_baan_proto_goTypes = nil
	file_baan_proto_depIdxs = nil
}
