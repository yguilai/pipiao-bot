// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: wft/v1/wft.proto

package v1

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

type WmItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WmItemReq) Reset() {
	*x = WmItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wft_v1_wft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WmItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WmItemReq) ProtoMessage() {}

func (x *WmItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_wft_v1_wft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WmItemReq.ProtoReflect.Descriptor instead.
func (*WmItemReq) Descriptor() ([]byte, []int) {
	return file_wft_v1_wft_proto_rawDescGZIP(), []int{0}
}

func (x *WmItemReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WmItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*WmItemResp_WmItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *WmItemResp) Reset() {
	*x = WmItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wft_v1_wft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WmItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WmItemResp) ProtoMessage() {}

func (x *WmItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_wft_v1_wft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WmItemResp.ProtoReflect.Descriptor instead.
func (*WmItemResp) Descriptor() ([]byte, []int) {
	return file_wft_v1_wft_proto_rawDescGZIP(), []int{1}
}

func (x *WmItemResp) GetItems() []*WmItemResp_WmItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type WarframeItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueName string `protobuf:"bytes,1,opt,name=uniqueName,proto3" json:"uniqueName,omitempty"`
}

func (x *WarframeItemReq) Reset() {
	*x = WarframeItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wft_v1_wft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarframeItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarframeItemReq) ProtoMessage() {}

func (x *WarframeItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_wft_v1_wft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarframeItemReq.ProtoReflect.Descriptor instead.
func (*WarframeItemReq) Descriptor() ([]byte, []int) {
	return file_wft_v1_wft_proto_rawDescGZIP(), []int{2}
}

func (x *WarframeItemReq) GetUniqueName() string {
	if x != nil {
		return x.UniqueName
	}
	return ""
}

type WarframeItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *WarframeItemResp) Reset() {
	*x = WarframeItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wft_v1_wft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarframeItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarframeItemResp) ProtoMessage() {}

func (x *WarframeItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_wft_v1_wft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarframeItemResp.ProtoReflect.Descriptor instead.
func (*WarframeItemResp) Descriptor() ([]byte, []int) {
	return file_wft_v1_wft_proto_rawDescGZIP(), []int{3}
}

func (x *WarframeItemResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WarframeItemResp) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *WarframeItemResp) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type WmItemResp_WmItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Thumb    string `protobuf:"bytes,2,opt,name=thumb,proto3" json:"thumb,omitempty"`
	UrlName  string `protobuf:"bytes,3,opt,name=urlName,proto3" json:"urlName,omitempty"`
	ItemName string `protobuf:"bytes,4,opt,name=itemName,proto3" json:"itemName,omitempty"`
	Vaulted  bool   `protobuf:"varint,5,opt,name=vaulted,proto3" json:"vaulted,omitempty"`
}

func (x *WmItemResp_WmItem) Reset() {
	*x = WmItemResp_WmItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wft_v1_wft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WmItemResp_WmItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WmItemResp_WmItem) ProtoMessage() {}

func (x *WmItemResp_WmItem) ProtoReflect() protoreflect.Message {
	mi := &file_wft_v1_wft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WmItemResp_WmItem.ProtoReflect.Descriptor instead.
func (*WmItemResp_WmItem) Descriptor() ([]byte, []int) {
	return file_wft_v1_wft_proto_rawDescGZIP(), []int{1, 0}
}

func (x *WmItemResp_WmItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WmItemResp_WmItem) GetThumb() string {
	if x != nil {
		return x.Thumb
	}
	return ""
}

func (x *WmItemResp_WmItem) GetUrlName() string {
	if x != nil {
		return x.UrlName
	}
	return ""
}

func (x *WmItemResp_WmItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *WmItemResp_WmItem) GetVaulted() bool {
	if x != nil {
		return x.Vaulted
	}
	return false
}

var File_wft_v1_wft_proto protoreflect.FileDescriptor

var file_wft_v1_wft_proto_rawDesc = []byte{
	0x0a, 0x10, 0x77, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x77, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x1f, 0x0a, 0x09, 0x57, 0x6d,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0a,
	0x57, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x66, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x57, 0x6d,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x7e, 0x0a, 0x06, 0x57,
	0x6d, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x72, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x22, 0x31, 0x0a, 0x0f, 0x57,
	0x61, 0x72, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7b,
	0x0a, 0x10, 0x57, 0x61, 0x72, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x92, 0x01, 0x0a, 0x03,
	0x57, 0x66, 0x74, 0x12, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x11, 0x2e, 0x77,
	0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x77, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x2e,
	0x77, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x72, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x77, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x61, 0x72, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x1a, 0x5a, 0x18, 0x70, 0x69, 0x70, 0x69, 0x61, 0x6f, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x77, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wft_v1_wft_proto_rawDescOnce sync.Once
	file_wft_v1_wft_proto_rawDescData = file_wft_v1_wft_proto_rawDesc
)

func file_wft_v1_wft_proto_rawDescGZIP() []byte {
	file_wft_v1_wft_proto_rawDescOnce.Do(func() {
		file_wft_v1_wft_proto_rawDescData = protoimpl.X.CompressGZIP(file_wft_v1_wft_proto_rawDescData)
	})
	return file_wft_v1_wft_proto_rawDescData
}

var file_wft_v1_wft_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wft_v1_wft_proto_goTypes = []interface{}{
	(*WmItemReq)(nil),         // 0: wft.v1.WmItemReq
	(*WmItemResp)(nil),        // 1: wft.v1.WmItemResp
	(*WarframeItemReq)(nil),   // 2: wft.v1.WarframeItemReq
	(*WarframeItemResp)(nil),  // 3: wft.v1.WarframeItemResp
	(*WmItemResp_WmItem)(nil), // 4: wft.v1.WmItemResp.WmItem
}
var file_wft_v1_wft_proto_depIdxs = []int32{
	4, // 0: wft.v1.WmItemResp.items:type_name -> wft.v1.WmItemResp.WmItem
	0, // 1: wft.v1.Wft.GetWarframeMarketItem:input_type -> wft.v1.WmItemReq
	2, // 2: wft.v1.Wft.GetWarframeOfficalItem:input_type -> wft.v1.WarframeItemReq
	1, // 3: wft.v1.Wft.GetWarframeMarketItem:output_type -> wft.v1.WmItemResp
	3, // 4: wft.v1.Wft.GetWarframeOfficalItem:output_type -> wft.v1.WarframeItemResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wft_v1_wft_proto_init() }
func file_wft_v1_wft_proto_init() {
	if File_wft_v1_wft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wft_v1_wft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WmItemReq); i {
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
		file_wft_v1_wft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WmItemResp); i {
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
		file_wft_v1_wft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarframeItemReq); i {
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
		file_wft_v1_wft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarframeItemResp); i {
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
		file_wft_v1_wft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WmItemResp_WmItem); i {
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
	file_wft_v1_wft_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wft_v1_wft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wft_v1_wft_proto_goTypes,
		DependencyIndexes: file_wft_v1_wft_proto_depIdxs,
		MessageInfos:      file_wft_v1_wft_proto_msgTypes,
	}.Build()
	File_wft_v1_wft_proto = out.File
	file_wft_v1_wft_proto_rawDesc = nil
	file_wft_v1_wft_proto_goTypes = nil
	file_wft_v1_wft_proto_depIdxs = nil
}
