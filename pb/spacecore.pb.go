// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: pb/spacecore.proto

package pb

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

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Cid     string `protobuf:"bytes,3,opt,name=cid,proto3" json:"cid,omitempty"`
	Path    string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugin) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *Plugin) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type RegisterPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Plugin  string `protobuf:"bytes,3,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *RegisterPluginRequest) Reset() {
	*x = RegisterPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPluginRequest) ProtoMessage() {}

func (x *RegisterPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPluginRequest.ProtoReflect.Descriptor instead.
func (*RegisterPluginRequest) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterPluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterPluginRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegisterPluginRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

type RegisterPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Cid     string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *RegisterPluginResponse) Reset() {
	*x = RegisterPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPluginResponse) ProtoMessage() {}

func (x *RegisterPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPluginResponse.ProtoReflect.Descriptor instead.
func (*RegisterPluginResponse) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterPluginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterPluginResponse) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type DiscoverPluginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Cid  *string `protobuf:"bytes,2,opt,name=cid,proto3,oneof" json:"cid,omitempty"`
}

func (x *DiscoverPluginsRequest) Reset() {
	*x = DiscoverPluginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverPluginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverPluginsRequest) ProtoMessage() {}

func (x *DiscoverPluginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverPluginsRequest.ProtoReflect.Descriptor instead.
func (*DiscoverPluginsRequest) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoverPluginsRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *DiscoverPluginsRequest) GetCid() string {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return ""
}

type DiscoverPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins []*Plugin `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *DiscoverPluginsResponse) Reset() {
	*x = DiscoverPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverPluginsResponse) ProtoMessage() {}

func (x *DiscoverPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverPluginsResponse.ProtoReflect.Descriptor instead.
func (*DiscoverPluginsResponse) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{4}
}

func (x *DiscoverPluginsResponse) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

type GetPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetPluginRequest) Reset() {
	*x = GetPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginRequest) ProtoMessage() {}

func (x *GetPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginRequest.ProtoReflect.Descriptor instead.
func (*GetPluginRequest) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{5}
}

func (x *GetPluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPluginRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin *Plugin `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *GetPluginResponse) Reset() {
	*x = GetPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginResponse) ProtoMessage() {}

func (x *GetPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginResponse.ProtoReflect.Descriptor instead.
func (*GetPluginResponse) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{6}
}

func (x *GetPluginResponse) GetPlugin() *Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type DownloadPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *DownloadPluginRequest) Reset() {
	*x = DownloadPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPluginRequest) ProtoMessage() {}

func (x *DownloadPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPluginRequest.ProtoReflect.Descriptor instead.
func (*DownloadPluginRequest) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadPluginRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type DownloadPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *DownloadPluginResponse) Reset() {
	*x = DownloadPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_spacecore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPluginResponse) ProtoMessage() {}

func (x *DownloadPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_spacecore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPluginResponse.ProtoReflect.Descriptor instead.
func (*DownloadPluginResponse) Descriptor() ([]byte, []int) {
	return file_pb_spacecore_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadPluginResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_pb_spacecore_proto protoreflect.FileDescriptor

var file_pb_spacecore_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5c, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x5d, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x44, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x16, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15,
	0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x63, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x07,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x22, 0x29, 0x0a, 0x15, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x32, 0x0a,
	0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x32, 0xdf, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a,
	0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_spacecore_proto_rawDescOnce sync.Once
	file_pb_spacecore_proto_rawDescData = file_pb_spacecore_proto_rawDesc
)

func file_pb_spacecore_proto_rawDescGZIP() []byte {
	file_pb_spacecore_proto_rawDescOnce.Do(func() {
		file_pb_spacecore_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_spacecore_proto_rawDescData)
	})
	return file_pb_spacecore_proto_rawDescData
}

var file_pb_spacecore_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_spacecore_proto_goTypes = []interface{}{
	(*Plugin)(nil),                  // 0: pb.Plugin
	(*RegisterPluginRequest)(nil),   // 1: pb.RegisterPluginRequest
	(*RegisterPluginResponse)(nil),  // 2: pb.RegisterPluginResponse
	(*DiscoverPluginsRequest)(nil),  // 3: pb.DiscoverPluginsRequest
	(*DiscoverPluginsResponse)(nil), // 4: pb.DiscoverPluginsResponse
	(*GetPluginRequest)(nil),        // 5: pb.GetPluginRequest
	(*GetPluginResponse)(nil),       // 6: pb.GetPluginResponse
	(*DownloadPluginRequest)(nil),   // 7: pb.DownloadPluginRequest
	(*DownloadPluginResponse)(nil),  // 8: pb.DownloadPluginResponse
}
var file_pb_spacecore_proto_depIdxs = []int32{
	0, // 0: pb.DiscoverPluginsResponse.plugins:type_name -> pb.Plugin
	0, // 1: pb.GetPluginResponse.plugin:type_name -> pb.Plugin
	1, // 2: pb.PluginRegistry.RegisterPlugin:input_type -> pb.RegisterPluginRequest
	3, // 3: pb.PluginRegistry.DiscoverPlugins:input_type -> pb.DiscoverPluginsRequest
	5, // 4: pb.PluginRegistry.GetPlugin:input_type -> pb.GetPluginRequest
	2, // 5: pb.PluginRegistry.RegisterPlugin:output_type -> pb.RegisterPluginResponse
	4, // 6: pb.PluginRegistry.DiscoverPlugins:output_type -> pb.DiscoverPluginsResponse
	6, // 7: pb.PluginRegistry.GetPlugin:output_type -> pb.GetPluginResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_spacecore_proto_init() }
func file_pb_spacecore_proto_init() {
	if File_pb_spacecore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_spacecore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_pb_spacecore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPluginRequest); i {
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
		file_pb_spacecore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPluginResponse); i {
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
		file_pb_spacecore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverPluginsRequest); i {
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
		file_pb_spacecore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverPluginsResponse); i {
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
		file_pb_spacecore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginRequest); i {
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
		file_pb_spacecore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginResponse); i {
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
		file_pb_spacecore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadPluginRequest); i {
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
		file_pb_spacecore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadPluginResponse); i {
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
	file_pb_spacecore_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_spacecore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_spacecore_proto_goTypes,
		DependencyIndexes: file_pb_spacecore_proto_depIdxs,
		MessageInfos:      file_pb_spacecore_proto_msgTypes,
	}.Build()
	File_pb_spacecore_proto = out.File
	file_pb_spacecore_proto_rawDesc = nil
	file_pb_spacecore_proto_goTypes = nil
	file_pb_spacecore_proto_depIdxs = nil
}
