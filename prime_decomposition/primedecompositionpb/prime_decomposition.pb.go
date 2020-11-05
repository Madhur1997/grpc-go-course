// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: primedecompositionpb/prime_decomposition.proto

package primedecompositionpb

import (
	context "context"
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

type ReqNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ReqNumber) Reset() {
	*x = ReqNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primedecompositionpb_prime_decomposition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqNumber) ProtoMessage() {}

func (x *ReqNumber) ProtoReflect() protoreflect.Message {
	mi := &file_primedecompositionpb_prime_decomposition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqNumber.ProtoReflect.Descriptor instead.
func (*ReqNumber) Descriptor() ([]byte, []int) {
	return file_primedecompositionpb_prime_decomposition_proto_rawDescGZIP(), []int{0}
}

func (x *ReqNumber) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeFactor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeFactor) Reset() {
	*x = PrimeFactor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primedecompositionpb_prime_decomposition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeFactor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactor) ProtoMessage() {}

func (x *PrimeFactor) ProtoReflect() protoreflect.Message {
	mi := &file_primedecompositionpb_prime_decomposition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactor.ProtoReflect.Descriptor instead.
func (*PrimeFactor) Descriptor() ([]byte, []int) {
	return file_primedecompositionpb_prime_decomposition_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeFactor) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_primedecompositionpb_prime_decomposition_proto protoreflect.FileDescriptor

var file_primedecompositionpb_prime_decomposition_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x22, 0x1d, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0x6d, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x1f, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x1a, 0x21, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x64, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_primedecompositionpb_prime_decomposition_proto_rawDescOnce sync.Once
	file_primedecompositionpb_prime_decomposition_proto_rawDescData = file_primedecompositionpb_prime_decomposition_proto_rawDesc
)

func file_primedecompositionpb_prime_decomposition_proto_rawDescGZIP() []byte {
	file_primedecompositionpb_prime_decomposition_proto_rawDescOnce.Do(func() {
		file_primedecompositionpb_prime_decomposition_proto_rawDescData = protoimpl.X.CompressGZIP(file_primedecompositionpb_prime_decomposition_proto_rawDescData)
	})
	return file_primedecompositionpb_prime_decomposition_proto_rawDescData
}

var file_primedecompositionpb_prime_decomposition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_primedecompositionpb_prime_decomposition_proto_goTypes = []interface{}{
	(*ReqNumber)(nil),   // 0: primedecompositionpb.ReqNumber
	(*PrimeFactor)(nil), // 1: primedecompositionpb.PrimeFactor
}
var file_primedecompositionpb_prime_decomposition_proto_depIdxs = []int32{
	0, // 0: primedecompositionpb.PrimeDecomposition.getPrimeFactors:input_type -> primedecompositionpb.ReqNumber
	1, // 1: primedecompositionpb.PrimeDecomposition.getPrimeFactors:output_type -> primedecompositionpb.PrimeFactor
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_primedecompositionpb_prime_decomposition_proto_init() }
func file_primedecompositionpb_prime_decomposition_proto_init() {
	if File_primedecompositionpb_prime_decomposition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_primedecompositionpb_prime_decomposition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqNumber); i {
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
		file_primedecompositionpb_prime_decomposition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeFactor); i {
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
			RawDescriptor: file_primedecompositionpb_prime_decomposition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_primedecompositionpb_prime_decomposition_proto_goTypes,
		DependencyIndexes: file_primedecompositionpb_prime_decomposition_proto_depIdxs,
		MessageInfos:      file_primedecompositionpb_prime_decomposition_proto_msgTypes,
	}.Build()
	File_primedecompositionpb_prime_decomposition_proto = out.File
	file_primedecompositionpb_prime_decomposition_proto_rawDesc = nil
	file_primedecompositionpb_prime_decomposition_proto_goTypes = nil
	file_primedecompositionpb_prime_decomposition_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PrimeDecompositionClient is the client API for PrimeDecomposition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimeDecompositionClient interface {
	GetPrimeFactors(ctx context.Context, in *ReqNumber, opts ...grpc.CallOption) (PrimeDecomposition_GetPrimeFactorsClient, error)
}

type primeDecompositionClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeDecompositionClient(cc grpc.ClientConnInterface) PrimeDecompositionClient {
	return &primeDecompositionClient{cc}
}

func (c *primeDecompositionClient) GetPrimeFactors(ctx context.Context, in *ReqNumber, opts ...grpc.CallOption) (PrimeDecomposition_GetPrimeFactorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeDecomposition_serviceDesc.Streams[0], "/primedecompositionpb.PrimeDecomposition/getPrimeFactors", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeDecompositionGetPrimeFactorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeDecomposition_GetPrimeFactorsClient interface {
	Recv() (*PrimeFactor, error)
	grpc.ClientStream
}

type primeDecompositionGetPrimeFactorsClient struct {
	grpc.ClientStream
}

func (x *primeDecompositionGetPrimeFactorsClient) Recv() (*PrimeFactor, error) {
	m := new(PrimeFactor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeDecompositionServer is the server API for PrimeDecomposition service.
type PrimeDecompositionServer interface {
	GetPrimeFactors(*ReqNumber, PrimeDecomposition_GetPrimeFactorsServer) error
}

// UnimplementedPrimeDecompositionServer can be embedded to have forward compatible implementations.
type UnimplementedPrimeDecompositionServer struct {
}

func (*UnimplementedPrimeDecompositionServer) GetPrimeFactors(*ReqNumber, PrimeDecomposition_GetPrimeFactorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPrimeFactors not implemented")
}

func RegisterPrimeDecompositionServer(s *grpc.Server, srv PrimeDecompositionServer) {
	s.RegisterService(&_PrimeDecomposition_serviceDesc, srv)
}

func _PrimeDecomposition_GetPrimeFactors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqNumber)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeDecompositionServer).GetPrimeFactors(m, &primeDecompositionGetPrimeFactorsServer{stream})
}

type PrimeDecomposition_GetPrimeFactorsServer interface {
	Send(*PrimeFactor) error
	grpc.ServerStream
}

type primeDecompositionGetPrimeFactorsServer struct {
	grpc.ServerStream
}

func (x *primeDecompositionGetPrimeFactorsServer) Send(m *PrimeFactor) error {
	return x.ServerStream.SendMsg(m)
}

var _PrimeDecomposition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "primedecompositionpb.PrimeDecomposition",
	HandlerType: (*PrimeDecompositionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getPrimeFactors",
			Handler:       _PrimeDecomposition_GetPrimeFactors_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "primedecompositionpb/prime_decomposition.proto",
}