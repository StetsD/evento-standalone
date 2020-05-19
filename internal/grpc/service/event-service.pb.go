// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.0
// source: evento-standalone/internal/proto/service/event-service.proto

package service

import (
	context "context"
	domain "evento-standalone/internal/grpc/domain"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddedEvent *domain.Event `protobuf:"bytes,1,opt,name=addedEvent,proto3" json:"addedEvent,omitempty"`
	Error      *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddEventResponse) Reset() {
	*x = AddEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evento_standalone_internal_proto_service_event_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventResponse) ProtoMessage() {}

func (x *AddEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evento_standalone_internal_proto_service_event_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventResponse.ProtoReflect.Descriptor instead.
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return file_evento_standalone_internal_proto_service_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddEventResponse) GetAddedEvent() *domain.Event {
	if x != nil {
		return x.AddedEvent
	}
	return nil
}

func (x *AddEventResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evento_standalone_internal_proto_service_event_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_evento_standalone_internal_proto_service_event_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_evento_standalone_internal_proto_service_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_evento_standalone_internal_proto_service_event_service_proto protoreflect.FileDescriptor

var file_evento_standalone_internal_proto_service_event_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x2d, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c,
	0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x33, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x2d,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3f, 0x0a, 0x0c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x03,
	0x61, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x2d, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x6f,
	0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evento_standalone_internal_proto_service_event_service_proto_rawDescOnce sync.Once
	file_evento_standalone_internal_proto_service_event_service_proto_rawDescData = file_evento_standalone_internal_proto_service_event_service_proto_rawDesc
)

func file_evento_standalone_internal_proto_service_event_service_proto_rawDescGZIP() []byte {
	file_evento_standalone_internal_proto_service_event_service_proto_rawDescOnce.Do(func() {
		file_evento_standalone_internal_proto_service_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_evento_standalone_internal_proto_service_event_service_proto_rawDescData)
	})
	return file_evento_standalone_internal_proto_service_event_service_proto_rawDescData
}

var file_evento_standalone_internal_proto_service_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_evento_standalone_internal_proto_service_event_service_proto_goTypes = []interface{}{
	(*AddEventResponse)(nil), // 0: service.AddEventResponse
	(*Error)(nil),            // 1: service.Error
	(*domain.Event)(nil),     // 2: domain.Event
}
var file_evento_standalone_internal_proto_service_event_service_proto_depIdxs = []int32{
	2, // 0: service.AddEventResponse.addedEvent:type_name -> domain.Event
	1, // 1: service.AddEventResponse.error:type_name -> service.Error
	2, // 2: service.EventService.add:input_type -> domain.Event
	0, // 3: service.EventService.add:output_type -> service.AddEventResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_evento_standalone_internal_proto_service_event_service_proto_init() }
func file_evento_standalone_internal_proto_service_event_service_proto_init() {
	if File_evento_standalone_internal_proto_service_event_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evento_standalone_internal_proto_service_event_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventResponse); i {
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
		file_evento_standalone_internal_proto_service_event_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_evento_standalone_internal_proto_service_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evento_standalone_internal_proto_service_event_service_proto_goTypes,
		DependencyIndexes: file_evento_standalone_internal_proto_service_event_service_proto_depIdxs,
		MessageInfos:      file_evento_standalone_internal_proto_service_event_service_proto_msgTypes,
	}.Build()
	File_evento_standalone_internal_proto_service_event_service_proto = out.File
	file_evento_standalone_internal_proto_service_event_service_proto_rawDesc = nil
	file_evento_standalone_internal_proto_service_event_service_proto_goTypes = nil
	file_evento_standalone_internal_proto_service_event_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	Add(ctx context.Context, in *domain.Event, opts ...grpc.CallOption) (*AddEventResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Add(ctx context.Context, in *domain.Event, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, "/service.EventService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	Add(context.Context, *domain.Event) (*AddEventResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) Add(context.Context, *domain.Event) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.EventService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Add(ctx, req.(*domain.Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _EventService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evento-standalone/internal/proto/service/event-service.proto",
}
