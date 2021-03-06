// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: calculator/calculatepb/calculate.proto

package calculatepb

import (
	context "context"
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

type FindMaxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *FindMaxRequest) Reset() {
	*x = FindMaxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatepb_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxRequest) ProtoMessage() {}

func (x *FindMaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatepb_calculate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxRequest.ProtoReflect.Descriptor instead.
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatepb_calculate_proto_rawDescGZIP(), []int{0}
}

func (x *FindMaxRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type FindMaxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *FindMaxResponse) Reset() {
	*x = FindMaxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatepb_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxResponse) ProtoMessage() {}

func (x *FindMaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatepb_calculate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxResponse.ProtoReflect.Descriptor instead.
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatepb_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *FindMaxResponse) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

type SquareRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *SquareRootRequest) Reset() {
	*x = SquareRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatepb_calculate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootRequest) ProtoMessage() {}

func (x *SquareRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatepb_calculate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootRequest.ProtoReflect.Descriptor instead.
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatepb_calculate_proto_rawDescGZIP(), []int{2}
}

func (x *SquareRootRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type SquareRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberRoot float64 `protobuf:"fixed64,2,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
}

func (x *SquareRootResponse) Reset() {
	*x = SquareRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatepb_calculate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootResponse) ProtoMessage() {}

func (x *SquareRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatepb_calculate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootResponse.ProtoReflect.Descriptor instead.
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatepb_calculate_proto_rawDescGZIP(), []int{3}
}

func (x *SquareRootResponse) GetNumberRoot() float64 {
	if x != nil {
		return x.NumberRoot
	}
	return 0
}

var File_calculator_calculatepb_calculate_proto protoreflect.FileDescriptor

var file_calculator_calculatepb_calculate_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2b, 0x0a, 0x0f,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x22, 0x35, 0x0a, 0x12, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x74, 0x32, 0x86, 0x01, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x0f,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0a, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calculatepb_calculate_proto_rawDescOnce sync.Once
	file_calculator_calculatepb_calculate_proto_rawDescData = file_calculator_calculatepb_calculate_proto_rawDesc
)

func file_calculator_calculatepb_calculate_proto_rawDescGZIP() []byte {
	file_calculator_calculatepb_calculate_proto_rawDescOnce.Do(func() {
		file_calculator_calculatepb_calculate_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculatepb_calculate_proto_rawDescData)
	})
	return file_calculator_calculatepb_calculate_proto_rawDescData
}

var file_calculator_calculatepb_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calculator_calculatepb_calculate_proto_goTypes = []interface{}{
	(*FindMaxRequest)(nil),     // 0: FindMaxRequest
	(*FindMaxResponse)(nil),    // 1: FindMaxResponse
	(*SquareRootRequest)(nil),  // 2: SquareRootRequest
	(*SquareRootResponse)(nil), // 3: SquareRootResponse
}
var file_calculator_calculatepb_calculate_proto_depIdxs = []int32{
	0, // 0: CalculateMaxService.FindMaximum:input_type -> FindMaxRequest
	2, // 1: CalculateMaxService.SquareRoot:input_type -> SquareRootRequest
	1, // 2: CalculateMaxService.FindMaximum:output_type -> FindMaxResponse
	3, // 3: CalculateMaxService.SquareRoot:output_type -> SquareRootResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_calculatepb_calculate_proto_init() }
func file_calculator_calculatepb_calculate_proto_init() {
	if File_calculator_calculatepb_calculate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculatepb_calculate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxRequest); i {
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
		file_calculator_calculatepb_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxResponse); i {
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
		file_calculator_calculatepb_calculate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootRequest); i {
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
		file_calculator_calculatepb_calculate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootResponse); i {
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
			RawDescriptor: file_calculator_calculatepb_calculate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculatepb_calculate_proto_goTypes,
		DependencyIndexes: file_calculator_calculatepb_calculate_proto_depIdxs,
		MessageInfos:      file_calculator_calculatepb_calculate_proto_msgTypes,
	}.Build()
	File_calculator_calculatepb_calculate_proto = out.File
	file_calculator_calculatepb_calculate_proto_rawDesc = nil
	file_calculator_calculatepb_calculate_proto_goTypes = nil
	file_calculator_calculatepb_calculate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateMaxServiceClient is the client API for CalculateMaxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateMaxServiceClient interface {
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculateMaxService_FindMaximumClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculateMaxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateMaxServiceClient(cc grpc.ClientConnInterface) CalculateMaxServiceClient {
	return &calculateMaxServiceClient{cc}
}

func (c *calculateMaxServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculateMaxService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateMaxService_serviceDesc.Streams[0], "/CalculateMaxService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateMaxServiceFindMaximumClient{stream}
	return x, nil
}

type CalculateMaxService_FindMaximumClient interface {
	Send(*FindMaxRequest) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type calculateMaxServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculateMaxServiceFindMaximumClient) Send(m *FindMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateMaxServiceFindMaximumClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateMaxServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/CalculateMaxService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateMaxServiceServer is the server API for CalculateMaxService service.
type CalculateMaxServiceServer interface {
	FindMaximum(CalculateMaxService_FindMaximumServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculateMaxServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateMaxServiceServer struct {
}

func (*UnimplementedCalculateMaxServiceServer) FindMaximum(CalculateMaxService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (*UnimplementedCalculateMaxServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculateMaxServiceServer(s *grpc.Server, srv CalculateMaxServiceServer) {
	s.RegisterService(&_CalculateMaxService_serviceDesc, srv)
}

func _CalculateMaxService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateMaxServiceServer).FindMaximum(&calculateMaxServiceFindMaximumServer{stream})
}

type CalculateMaxService_FindMaximumServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequest, error)
	grpc.ServerStream
}

type calculateMaxServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculateMaxServiceFindMaximumServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateMaxServiceFindMaximumServer) Recv() (*FindMaxRequest, error) {
	m := new(FindMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateMaxService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateMaxServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalculateMaxService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateMaxServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateMaxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CalculateMaxService",
	HandlerType: (*CalculateMaxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateMaxService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculateMaxService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatepb/calculate.proto",
}
