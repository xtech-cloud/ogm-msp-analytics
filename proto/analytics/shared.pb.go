// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/analytics/shared.proto

package analytics

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

// 状态
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 状态信息
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_analytics_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 空白回复
type BlankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
}

func (x *BlankResponse) Reset() {
	*x = BlankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankResponse) ProtoMessage() {}

func (x *BlankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankResponse.ProtoReflect.Descriptor instead.
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return file_proto_analytics_shared_proto_rawDescGZIP(), []int{1}
}

func (x *BlankResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 个体
type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber    string `protobuf:"bytes,1,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`       // 序列号
	SoftwareFamily  string `protobuf:"bytes,2,opt,name=softwareFamily,proto3" json:"softwareFamily,omitempty"`   // 软件家族
	SoftwareVersion string `protobuf:"bytes,3,opt,name=softwareVersion,proto3" json:"softwareVersion,omitempty"` // 软件版本
	SystemFamily    string `protobuf:"bytes,4,opt,name=systemFamily,proto3" json:"systemFamily,omitempty"`       // 系统家族
	SystemVersion   string `protobuf:"bytes,5,opt,name=systemVersion,proto3" json:"systemVersion,omitempty"`     // 系统版本
	DeviceModel     string `protobuf:"bytes,6,opt,name=deviceModel,proto3" json:"deviceModel,omitempty"`         // 设备模型
	DeviceType      string `protobuf:"bytes,7,opt,name=deviceType,proto3" json:"deviceType,omitempty"`           // 设备类型
	Profile         string `protobuf:"bytes,8,opt,name=profile,proto3" json:"profile,omitempty"`                 // 简要资料
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_proto_analytics_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Agent) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Agent) GetSoftwareFamily() string {
	if x != nil {
		return x.SoftwareFamily
	}
	return ""
}

func (x *Agent) GetSoftwareVersion() string {
	if x != nil {
		return x.SoftwareVersion
	}
	return ""
}

func (x *Agent) GetSystemFamily() string {
	if x != nil {
		return x.SystemFamily
	}
	return ""
}

func (x *Agent) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *Agent) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *Agent) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Agent) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

// 参数
type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // 键
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // 值
	Type  string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`   // 类型
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_proto_analytics_shared_proto_rawDescGZIP(), []int{3}
}

func (x *Parameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Parameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Parameter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// 事件
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // 事件唯一ID
	Alias     string       `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`         // 事件显示名
	Parameter []*Parameter `protobuf:"bytes,3,rep,name=parameter,proto3" json:"parameter,omitempty"` // 事件参数
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_analytics_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_analytics_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_analytics_shared_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Event) GetParameter() []*Parameter {
	if x != nil {
		return x.Parameter
	}
	return nil
}

var File_proto_analytics_shared_proto protoreflect.FileDescriptor

var file_proto_analytics_shared_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x0d, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa3, 0x02, 0x0a, 0x05, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61,
	0x72, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x28,
	0x0a, 0x0f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x47,
	0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x28, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x3b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_analytics_shared_proto_rawDescOnce sync.Once
	file_proto_analytics_shared_proto_rawDescData = file_proto_analytics_shared_proto_rawDesc
)

func file_proto_analytics_shared_proto_rawDescGZIP() []byte {
	file_proto_analytics_shared_proto_rawDescOnce.Do(func() {
		file_proto_analytics_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_analytics_shared_proto_rawDescData)
	})
	return file_proto_analytics_shared_proto_rawDescData
}

var file_proto_analytics_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_analytics_shared_proto_goTypes = []interface{}{
	(*Status)(nil),        // 0: Status
	(*BlankResponse)(nil), // 1: BlankResponse
	(*Agent)(nil),         // 2: Agent
	(*Parameter)(nil),     // 3: Parameter
	(*Event)(nil),         // 4: Event
}
var file_proto_analytics_shared_proto_depIdxs = []int32{
	0, // 0: BlankResponse.status:type_name -> Status
	3, // 1: Event.parameter:type_name -> Parameter
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_analytics_shared_proto_init() }
func file_proto_analytics_shared_proto_init() {
	if File_proto_analytics_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_analytics_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_analytics_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankResponse); i {
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
		file_proto_analytics_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_proto_analytics_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
		file_proto_analytics_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_proto_analytics_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_analytics_shared_proto_goTypes,
		DependencyIndexes: file_proto_analytics_shared_proto_depIdxs,
		MessageInfos:      file_proto_analytics_shared_proto_msgTypes,
	}.Build()
	File_proto_analytics_shared_proto = out.File
	file_proto_analytics_shared_proto_rawDesc = nil
	file_proto_analytics_shared_proto_goTypes = nil
	file_proto_analytics_shared_proto_depIdxs = nil
}
