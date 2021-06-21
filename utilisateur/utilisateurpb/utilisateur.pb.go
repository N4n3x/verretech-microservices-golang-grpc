// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/utilisateur.proto

package utilisateurpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	localisationpb "verretech-microservices/generic/localisationpb"
	pointRetraitpb "verretech-microservices/generic/pointRetraitpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Preferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Localisation *localisationpb.Localisation `protobuf:"bytes,1,opt,name=localisation,proto3" json:"localisation,omitempty"`
	PointRetrait *pointRetraitpb.PointRetrait `protobuf:"bytes,2,opt,name=pointRetrait,proto3" json:"pointRetrait,omitempty"`
}

func (x *Preferences) Reset() {
	*x = Preferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preferences) ProtoMessage() {}

func (x *Preferences) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preferences.ProtoReflect.Descriptor instead.
func (*Preferences) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{0}
}

func (x *Preferences) GetLocalisation() *localisationpb.Localisation {
	if x != nil {
		return x.Localisation
	}
	return nil
}

func (x *Preferences) GetPointRetrait() *pointRetraitpb.PointRetrait {
	if x != nil {
		return x.PointRetrait
	}
	return nil
}

type Utilisateur struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string       `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Nom            string       `protobuf:"bytes,2,opt,name=nom,proto3" json:"nom,omitempty"`
	Prenom         string       `protobuf:"bytes,3,opt,name=prenom,proto3" json:"prenom,omitempty"`
	Mail           string       `protobuf:"bytes,4,opt,name=mail,proto3" json:"mail,omitempty"`
	HashMotDePasse string       `protobuf:"bytes,5,opt,name=hashMotDePasse,proto3" json:"hashMotDePasse,omitempty"`
	Preferences    *Preferences `protobuf:"bytes,6,opt,name=preferences,proto3" json:"preferences,omitempty"`
	Permission     []string     `protobuf:"bytes,7,rep,name=permission,proto3" json:"permission,omitempty"`
}

func (x *Utilisateur) Reset() {
	*x = Utilisateur{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utilisateur) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utilisateur) ProtoMessage() {}

func (x *Utilisateur) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utilisateur.ProtoReflect.Descriptor instead.
func (*Utilisateur) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{1}
}

func (x *Utilisateur) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Utilisateur) GetNom() string {
	if x != nil {
		return x.Nom
	}
	return ""
}

func (x *Utilisateur) GetPrenom() string {
	if x != nil {
		return x.Prenom
	}
	return ""
}

func (x *Utilisateur) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *Utilisateur) GetHashMotDePasse() string {
	if x != nil {
		return x.HashMotDePasse
	}
	return ""
}

func (x *Utilisateur) GetPreferences() *Preferences {
	if x != nil {
		return x.Preferences
	}
	return nil
}

func (x *Utilisateur) GetPermission() []string {
	if x != nil {
		return x.Permission
	}
	return nil
}

type UtilisateurRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Utilisateur *Utilisateur `protobuf:"bytes,1,opt,name=utilisateur,proto3" json:"utilisateur,omitempty"`
}

func (x *UtilisateurRequest) Reset() {
	*x = UtilisateurRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilisateurRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilisateurRequest) ProtoMessage() {}

func (x *UtilisateurRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilisateurRequest.ProtoReflect.Descriptor instead.
func (*UtilisateurRequest) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{2}
}

func (x *UtilisateurRequest) GetUtilisateur() *Utilisateur {
	if x != nil {
		return x.Utilisateur
	}
	return nil
}

type UtilisateursRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UtilisateursRequest) Reset() {
	*x = UtilisateursRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilisateursRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilisateursRequest) ProtoMessage() {}

func (x *UtilisateursRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilisateursRequest.ProtoReflect.Descriptor instead.
func (*UtilisateursRequest) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{3}
}

type UtilisateurResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Utilisateur *Utilisateur `protobuf:"bytes,1,opt,name=utilisateur,proto3" json:"utilisateur,omitempty"`
}

func (x *UtilisateurResponse) Reset() {
	*x = UtilisateurResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilisateurResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilisateurResponse) ProtoMessage() {}

func (x *UtilisateurResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilisateurResponse.ProtoReflect.Descriptor instead.
func (*UtilisateurResponse) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{4}
}

func (x *UtilisateurResponse) GetUtilisateur() *Utilisateur {
	if x != nil {
		return x.Utilisateur
	}
	return nil
}

type UtilisateursResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Utilisateur []*Utilisateur `protobuf:"bytes,1,rep,name=utilisateur,proto3" json:"utilisateur,omitempty"`
}

func (x *UtilisateursResponse) Reset() {
	*x = UtilisateursResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilisateursResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilisateursResponse) ProtoMessage() {}

func (x *UtilisateursResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilisateursResponse.ProtoReflect.Descriptor instead.
func (*UtilisateursResponse) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{5}
}

func (x *UtilisateursResponse) GetUtilisateur() []*Utilisateur {
	if x != nil {
		return x.Utilisateur
	}
	return nil
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{6}
}

func (x *BoolResponse) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State       bool         `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Utilisateur *Utilisateur `protobuf:"bytes,2,opt,name=utilisateur,proto3" json:"utilisateur,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_utilisateur_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_utilisateur_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_utilisateur_proto_rawDescGZIP(), []int{7}
}

func (x *AuthResponse) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

func (x *AuthResponse) GetUtilisateur() *Utilisateur {
	if x != nil {
		return x.Utilisateur
	}
	return nil
}

var File_proto_utilisateur_proto protoreflect.FileDescriptor

var file_proto_utilisateur_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74,
	0x65, 0x75, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69,
	0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69, 0x74, 0x52, 0x0c, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x72, 0x61, 0x69, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x55,
	0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68,
	0x4d, 0x6f, 0x74, 0x44, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x4d, 0x6f, 0x74, 0x44, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74,
	0x65, 0x75, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x12,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73,
	0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75,
	0x72, 0x52, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x22, 0x15,
	0x0a, 0x13, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x13, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61,
	0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x0b, 0x75, 0x74, 0x69,
	0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x22, 0x52, 0x0a, 0x14, 0x55, 0x74, 0x69, 0x6c,
	0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74,
	0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52,
	0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x22, 0x24, 0x0a, 0x0c,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x60, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c,
	0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c,
	0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61,
	0x74, 0x65, 0x75, 0x72, 0x32, 0xb2, 0x03, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x12, 0x53, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x12, 0x1f, 0x2e,
	0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c,
	0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69,
	0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73,
	0x61, 0x74, 0x65, 0x75, 0x72, 0x12, 0x1f, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74,
	0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61,
	0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x12, 0x1f, 0x2e, 0x75, 0x74, 0x69,
	0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61,
	0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73,
	0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x73,
	0x12, 0x20, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55,
	0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72,
	0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1f, 0x2e,
	0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x55, 0x74, 0x69, 0x6c,
	0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x76, 0x65, 0x72,
	0x72, 0x65, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72,
	0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x73, 0x61, 0x74, 0x65, 0x75, 0x72, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_utilisateur_proto_rawDescOnce sync.Once
	file_proto_utilisateur_proto_rawDescData = file_proto_utilisateur_proto_rawDesc
)

func file_proto_utilisateur_proto_rawDescGZIP() []byte {
	file_proto_utilisateur_proto_rawDescOnce.Do(func() {
		file_proto_utilisateur_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_utilisateur_proto_rawDescData)
	})
	return file_proto_utilisateur_proto_rawDescData
}

var file_proto_utilisateur_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_utilisateur_proto_goTypes = []interface{}{
	(*Preferences)(nil),                 // 0: utilisateur.Preferences
	(*Utilisateur)(nil),                 // 1: utilisateur.Utilisateur
	(*UtilisateurRequest)(nil),          // 2: utilisateur.UtilisateurRequest
	(*UtilisateursRequest)(nil),         // 3: utilisateur.UtilisateursRequest
	(*UtilisateurResponse)(nil),         // 4: utilisateur.UtilisateurResponse
	(*UtilisateursResponse)(nil),        // 5: utilisateur.UtilisateursResponse
	(*BoolResponse)(nil),                // 6: utilisateur.BoolResponse
	(*AuthResponse)(nil),                // 7: utilisateur.AuthResponse
	(*localisationpb.Localisation)(nil), // 8: localisation.Localisation
	(*pointRetraitpb.PointRetrait)(nil), // 9: pointRetrait.PointRetrait
}
var file_proto_utilisateur_proto_depIdxs = []int32{
	8,  // 0: utilisateur.Preferences.localisation:type_name -> localisation.Localisation
	9,  // 1: utilisateur.Preferences.pointRetrait:type_name -> pointRetrait.PointRetrait
	0,  // 2: utilisateur.Utilisateur.preferences:type_name -> utilisateur.Preferences
	1,  // 3: utilisateur.UtilisateurRequest.utilisateur:type_name -> utilisateur.Utilisateur
	1,  // 4: utilisateur.UtilisateurResponse.utilisateur:type_name -> utilisateur.Utilisateur
	1,  // 5: utilisateur.UtilisateursResponse.utilisateur:type_name -> utilisateur.Utilisateur
	1,  // 6: utilisateur.AuthResponse.utilisateur:type_name -> utilisateur.Utilisateur
	2,  // 7: utilisateur.ServiceUtilisateur.AddUtilisateur:input_type -> utilisateur.UtilisateurRequest
	2,  // 8: utilisateur.ServiceUtilisateur.UpdateUtilisateur:input_type -> utilisateur.UtilisateurRequest
	2,  // 9: utilisateur.ServiceUtilisateur.GetUtilisateur:input_type -> utilisateur.UtilisateurRequest
	3,  // 10: utilisateur.ServiceUtilisateur.GetUtilisateurs:input_type -> utilisateur.UtilisateursRequest
	2,  // 11: utilisateur.ServiceUtilisateur.Auth:input_type -> utilisateur.UtilisateurRequest
	4,  // 12: utilisateur.ServiceUtilisateur.AddUtilisateur:output_type -> utilisateur.UtilisateurResponse
	4,  // 13: utilisateur.ServiceUtilisateur.UpdateUtilisateur:output_type -> utilisateur.UtilisateurResponse
	4,  // 14: utilisateur.ServiceUtilisateur.GetUtilisateur:output_type -> utilisateur.UtilisateurResponse
	5,  // 15: utilisateur.ServiceUtilisateur.GetUtilisateurs:output_type -> utilisateur.UtilisateursResponse
	7,  // 16: utilisateur.ServiceUtilisateur.Auth:output_type -> utilisateur.AuthResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_utilisateur_proto_init() }
func file_proto_utilisateur_proto_init() {
	if File_proto_utilisateur_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_utilisateur_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Preferences); i {
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
		file_proto_utilisateur_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utilisateur); i {
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
		file_proto_utilisateur_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilisateurRequest); i {
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
		file_proto_utilisateur_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilisateursRequest); i {
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
		file_proto_utilisateur_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilisateurResponse); i {
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
		file_proto_utilisateur_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilisateursResponse); i {
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
		file_proto_utilisateur_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
		file_proto_utilisateur_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
			RawDescriptor: file_proto_utilisateur_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_utilisateur_proto_goTypes,
		DependencyIndexes: file_proto_utilisateur_proto_depIdxs,
		MessageInfos:      file_proto_utilisateur_proto_msgTypes,
	}.Build()
	File_proto_utilisateur_proto = out.File
	file_proto_utilisateur_proto_rawDesc = nil
	file_proto_utilisateur_proto_goTypes = nil
	file_proto_utilisateur_proto_depIdxs = nil
}
