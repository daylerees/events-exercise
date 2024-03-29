// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: accounts.proto

package accounts

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

// BadgeColour is an enum containing possible colours for badges.
type BadgeColour int32

const (
	BadgeColour_RED   BadgeColour = 0
	BadgeColour_GREEN BadgeColour = 1
	BadgeColour_BLUE  BadgeColour = 2
)

// Enum value maps for BadgeColour.
var (
	BadgeColour_name = map[int32]string{
		0: "RED",
		1: "GREEN",
		2: "BLUE",
	}
	BadgeColour_value = map[string]int32{
		"RED":   0,
		"GREEN": 1,
		"BLUE":  2,
	}
)

func (x BadgeColour) Enum() *BadgeColour {
	p := new(BadgeColour)
	*p = x
	return p
}

func (x BadgeColour) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BadgeColour) Descriptor() protoreflect.EnumDescriptor {
	return file_accounts_proto_enumTypes[0].Descriptor()
}

func (BadgeColour) Type() protoreflect.EnumType {
	return &file_accounts_proto_enumTypes[0]
}

func (x BadgeColour) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BadgeColour.Descriptor instead.
func (BadgeColour) EnumDescriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

// UserAccountCreatedEvent indicates that a new user account has been created.
type UserAccountCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserAccountCreatedEvent) Reset() {
	*x = UserAccountCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccountCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccountCreatedEvent) ProtoMessage() {}

func (x *UserAccountCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccountCreatedEvent.ProtoReflect.Descriptor instead.
func (*UserAccountCreatedEvent) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *UserAccountCreatedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// UserAccountUpdatedEvent updates a user account providing or updating additional user information.
type UserAccountUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserAccountUpdatedEvent) Reset() {
	*x = UserAccountUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccountUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccountUpdatedEvent) ProtoMessage() {}

func (x *UserAccountUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccountUpdatedEvent.ProtoReflect.Descriptor instead.
func (*UserAccountUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *UserAccountUpdatedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAccountUpdatedEvent) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserAccountUpdatedEvent) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// UserGainedBadgeEvent indicates that a user has gained a badge of a certain colour.
type UserGainedBadgeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BadgeColour BadgeColour `protobuf:"varint,2,opt,name=badge_colour,json=badgeColour,proto3,enum=accounts.BadgeColour" json:"badge_colour,omitempty"`
}

func (x *UserGainedBadgeEvent) Reset() {
	*x = UserGainedBadgeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGainedBadgeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGainedBadgeEvent) ProtoMessage() {}

func (x *UserGainedBadgeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGainedBadgeEvent.ProtoReflect.Descriptor instead.
func (*UserGainedBadgeEvent) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *UserGainedBadgeEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserGainedBadgeEvent) GetBadgeColour() BadgeColour {
	if x != nil {
		return x.BadgeColour
	}
	return BadgeColour_RED
}

// UserLostBadgeEvent indicates that a user has lost a badge of a certain colour.
type UserLostBadgeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BadgeColour BadgeColour `protobuf:"varint,2,opt,name=badge_colour,json=badgeColour,proto3,enum=accounts.BadgeColour" json:"badge_colour,omitempty"`
}

func (x *UserLostBadgeEvent) Reset() {
	*x = UserLostBadgeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLostBadgeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLostBadgeEvent) ProtoMessage() {}

func (x *UserLostBadgeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLostBadgeEvent.ProtoReflect.Descriptor instead.
func (*UserLostBadgeEvent) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *UserLostBadgeEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLostBadgeEvent) GetBadgeColour() BadgeColour {
	if x != nil {
		return x.BadgeColour
	}
	return BadgeColour_RED
}

var File_accounts_proto protoreflect.FileDescriptor

var file_accounts_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x17, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65,
	0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x69, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x42, 0x61, 0x64, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x61, 0x64, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6c,
	0x6f, 0x75, 0x72, 0x52, 0x0b, 0x62, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72,
	0x22, 0x67, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x64, 0x67,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x0c, 0x62, 0x61, 0x64, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x52, 0x0b, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x2a, 0x2b, 0x0a, 0x0b, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_accounts_proto_rawDescOnce sync.Once
	file_accounts_proto_rawDescData = file_accounts_proto_rawDesc
)

func file_accounts_proto_rawDescGZIP() []byte {
	file_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_proto_rawDescData)
	})
	return file_accounts_proto_rawDescData
}

var file_accounts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accounts_proto_goTypes = []interface{}{
	(BadgeColour)(0),                // 0: accounts.BadgeColour
	(*UserAccountCreatedEvent)(nil), // 1: accounts.UserAccountCreatedEvent
	(*UserAccountUpdatedEvent)(nil), // 2: accounts.UserAccountUpdatedEvent
	(*UserGainedBadgeEvent)(nil),    // 3: accounts.UserGainedBadgeEvent
	(*UserLostBadgeEvent)(nil),      // 4: accounts.UserLostBadgeEvent
}
var file_accounts_proto_depIdxs = []int32{
	0, // 0: accounts.UserGainedBadgeEvent.badge_colour:type_name -> accounts.BadgeColour
	0, // 1: accounts.UserLostBadgeEvent.badge_colour:type_name -> accounts.BadgeColour
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_accounts_proto_init() }
func file_accounts_proto_init() {
	if File_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccountCreatedEvent); i {
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
		file_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccountUpdatedEvent); i {
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
		file_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGainedBadgeEvent); i {
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
		file_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLostBadgeEvent); i {
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
			RawDescriptor: file_accounts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_proto_depIdxs,
		EnumInfos:         file_accounts_proto_enumTypes,
		MessageInfos:      file_accounts_proto_msgTypes,
	}.Build()
	File_accounts_proto = out.File
	file_accounts_proto_rawDesc = nil
	file_accounts_proto_goTypes = nil
	file_accounts_proto_depIdxs = nil
}
