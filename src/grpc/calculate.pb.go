// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.1
// source: calculate.proto

package grpc

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

type Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digit []float64 `protobuf:"fixed64,1,rep,packed,name=digit,proto3" json:"digit,omitempty"`
}

func (x *Array) Reset() {
	*x = Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Array) ProtoMessage() {}

func (x *Array) ProtoReflect() protoreflect.Message {
	mi := &file_calculate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Array.ProtoReflect.Descriptor instead.
func (*Array) Descriptor() ([]byte, []int) {
	return file_calculate_proto_rawDescGZIP(), []int{0}
}

func (x *Array) GetDigit() []float64 {
	if x != nil {
		return x.Digit
	}
	return nil
}

type MatrixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matrix1 []*Array `protobuf:"bytes,1,rep,name=matrix1,proto3" json:"matrix1,omitempty"`
	Matrix2 []*Array `protobuf:"bytes,2,rep,name=matrix2,proto3" json:"matrix2,omitempty"`
}

func (x *MatrixRequest) Reset() {
	*x = MatrixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixRequest) ProtoMessage() {}

func (x *MatrixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixRequest.ProtoReflect.Descriptor instead.
func (*MatrixRequest) Descriptor() ([]byte, []int) {
	return file_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *MatrixRequest) GetMatrix1() []*Array {
	if x != nil {
		return x.Matrix1
	}
	return nil
}

func (x *MatrixRequest) GetMatrix2() []*Array {
	if x != nil {
		return x.Matrix2
	}
	return nil
}

type MatrixResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matrix []*Array `protobuf:"bytes,1,rep,name=matrix,proto3" json:"matrix,omitempty"`
}

func (x *MatrixResponse) Reset() {
	*x = MatrixResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixResponse) ProtoMessage() {}

func (x *MatrixResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixResponse.ProtoReflect.Descriptor instead.
func (*MatrixResponse) Descriptor() ([]byte, []int) {
	return file_calculate_proto_rawDescGZIP(), []int{2}
}

func (x *MatrixResponse) GetMatrix() []*Array {
	if x != nil {
		return x.Matrix
	}
	return nil
}

var File_calculate_proto protoreflect.FileDescriptor

var file_calculate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x1d, 0x0a, 0x05,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x67, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x05, 0x64, 0x69, 0x67, 0x69, 0x74, 0x22, 0x67, 0x0a, 0x0d, 0x4d,
	0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07,
	0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x31, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52,
	0x07, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x31, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x32, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x32, 0x22, 0x3a, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x32, 0x99, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x12, 0x42, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x53, 0x75,
	0x6d, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x4d, 0x75, 0x6c, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculate_proto_rawDescOnce sync.Once
	file_calculate_proto_rawDescData = file_calculate_proto_rawDesc
)

func file_calculate_proto_rawDescGZIP() []byte {
	file_calculate_proto_rawDescOnce.Do(func() {
		file_calculate_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculate_proto_rawDescData)
	})
	return file_calculate_proto_rawDescData
}

var file_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_calculate_proto_goTypes = []interface{}{
	(*Array)(nil),          // 0: calculate.Array
	(*MatrixRequest)(nil),  // 1: calculate.MatrixRequest
	(*MatrixResponse)(nil), // 2: calculate.MatrixResponse
}
var file_calculate_proto_depIdxs = []int32{
	0, // 0: calculate.MatrixRequest.matrix1:type_name -> calculate.Array
	0, // 1: calculate.MatrixRequest.matrix2:type_name -> calculate.Array
	0, // 2: calculate.MatrixResponse.matrix:type_name -> calculate.Array
	1, // 3: calculate.CalculateMatrix.MatrixSum:input_type -> calculate.MatrixRequest
	1, // 4: calculate.CalculateMatrix.MatrixMul:input_type -> calculate.MatrixRequest
	2, // 5: calculate.CalculateMatrix.MatrixSum:output_type -> calculate.MatrixResponse
	2, // 6: calculate.CalculateMatrix.MatrixMul:output_type -> calculate.MatrixResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_calculate_proto_init() }
func file_calculate_proto_init() {
	if File_calculate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Array); i {
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
		file_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixRequest); i {
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
		file_calculate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixResponse); i {
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
			RawDescriptor: file_calculate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculate_proto_goTypes,
		DependencyIndexes: file_calculate_proto_depIdxs,
		MessageInfos:      file_calculate_proto_msgTypes,
	}.Build()
	File_calculate_proto = out.File
	file_calculate_proto_rawDesc = nil
	file_calculate_proto_goTypes = nil
	file_calculate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateMatrixClient is the client API for CalculateMatrix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateMatrixClient interface {
	MatrixSum(ctx context.Context, in *MatrixRequest, opts ...grpc.CallOption) (*MatrixResponse, error)
	MatrixMul(ctx context.Context, in *MatrixRequest, opts ...grpc.CallOption) (*MatrixResponse, error)
}

type calculateMatrixClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateMatrixClient(cc grpc.ClientConnInterface) CalculateMatrixClient {
	return &calculateMatrixClient{cc}
}

func (c *calculateMatrixClient) MatrixSum(ctx context.Context, in *MatrixRequest, opts ...grpc.CallOption) (*MatrixResponse, error) {
	out := new(MatrixResponse)
	err := c.cc.Invoke(ctx, "/calculate.CalculateMatrix/MatrixSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateMatrixClient) MatrixMul(ctx context.Context, in *MatrixRequest, opts ...grpc.CallOption) (*MatrixResponse, error) {
	out := new(MatrixResponse)
	err := c.cc.Invoke(ctx, "/calculate.CalculateMatrix/MatrixMul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateMatrixServer is the server API for CalculateMatrix service.
type CalculateMatrixServer interface {
	MatrixSum(context.Context, *MatrixRequest) (*MatrixResponse, error)
	MatrixMul(context.Context, *MatrixRequest) (*MatrixResponse, error)
}

// UnimplementedCalculateMatrixServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateMatrixServer struct {
}

func (*UnimplementedCalculateMatrixServer) MatrixSum(context.Context, *MatrixRequest) (*MatrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatrixSum not implemented")
}
func (*UnimplementedCalculateMatrixServer) MatrixMul(context.Context, *MatrixRequest) (*MatrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatrixMul not implemented")
}

func RegisterCalculateMatrixServer(s *grpc.Server, srv CalculateMatrixServer) {
	s.RegisterService(&_CalculateMatrix_serviceDesc, srv)
}

func _CalculateMatrix_MatrixSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateMatrixServer).MatrixSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.CalculateMatrix/MatrixSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateMatrixServer).MatrixSum(ctx, req.(*MatrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateMatrix_MatrixMul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateMatrixServer).MatrixMul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.CalculateMatrix/MatrixMul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateMatrixServer).MatrixMul(ctx, req.(*MatrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateMatrix_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculate.CalculateMatrix",
	HandlerType: (*CalculateMatrixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MatrixSum",
			Handler:    _CalculateMatrix_MatrixSum_Handler,
		},
		{
			MethodName: "MatrixMul",
			Handler:    _CalculateMatrix_MatrixMul_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculate.proto",
}
