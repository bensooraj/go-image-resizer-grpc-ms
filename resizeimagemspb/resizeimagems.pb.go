// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resizeimagemspb/resizeimagems.proto

package resizeimagemspb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResizeImageRequest struct {
	ImageId              string   `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	ImageFilename        string   `protobuf:"bytes,2,opt,name=image_filename,json=imageFilename,proto3" json:"image_filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeImageRequest) Reset()         { *m = ResizeImageRequest{} }
func (m *ResizeImageRequest) String() string { return proto.CompactTextString(m) }
func (*ResizeImageRequest) ProtoMessage()    {}
func (*ResizeImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed71579964fc8148, []int{0}
}

func (m *ResizeImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeImageRequest.Unmarshal(m, b)
}
func (m *ResizeImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeImageRequest.Marshal(b, m, deterministic)
}
func (m *ResizeImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeImageRequest.Merge(m, src)
}
func (m *ResizeImageRequest) XXX_Size() int {
	return xxx_messageInfo_ResizeImageRequest.Size(m)
}
func (m *ResizeImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeImageRequest proto.InternalMessageInfo

func (m *ResizeImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *ResizeImageRequest) GetImageFilename() string {
	if m != nil {
		return m.ImageFilename
	}
	return ""
}

type ResizeImageResponse struct {
	ImagesResized        int32    `protobuf:"varint,1,opt,name=images_resized,json=imagesResized,proto3" json:"images_resized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeImageResponse) Reset()         { *m = ResizeImageResponse{} }
func (m *ResizeImageResponse) String() string { return proto.CompactTextString(m) }
func (*ResizeImageResponse) ProtoMessage()    {}
func (*ResizeImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed71579964fc8148, []int{1}
}

func (m *ResizeImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeImageResponse.Unmarshal(m, b)
}
func (m *ResizeImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeImageResponse.Marshal(b, m, deterministic)
}
func (m *ResizeImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeImageResponse.Merge(m, src)
}
func (m *ResizeImageResponse) XXX_Size() int {
	return xxx_messageInfo_ResizeImageResponse.Size(m)
}
func (m *ResizeImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeImageResponse proto.InternalMessageInfo

func (m *ResizeImageResponse) GetImagesResized() int32 {
	if m != nil {
		return m.ImagesResized
	}
	return 0
}

func init() {
	proto.RegisterType((*ResizeImageRequest)(nil), "resize_image_ms.ResizeImageRequest")
	proto.RegisterType((*ResizeImageResponse)(nil), "resize_image_ms.ResizeImageResponse")
}

func init() {
	proto.RegisterFile("resizeimagemspb/resizeimagems.proto", fileDescriptor_ed71579964fc8148)
}

var fileDescriptor_ed71579964fc8148 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x4a, 0x2d, 0xce,
	0xac, 0x4a, 0xcd, 0xcc, 0x4d, 0x4c, 0x4f, 0xcd, 0x2d, 0x2e, 0x48, 0xd2, 0x47, 0xe1, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0x43, 0x04, 0xe3, 0xc1, 0xa2, 0xf1, 0xb9, 0xc5, 0x4a, 0x61,
	0x5c, 0x42, 0x41, 0x60, 0x21, 0x4f, 0x90, 0x48, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
	0x24, 0x17, 0x07, 0x44, 0x45, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x3b, 0x98,
	0xef, 0x99, 0x22, 0xa4, 0xca, 0xc5, 0x07, 0x91, 0x4a, 0xcb, 0xcc, 0x49, 0xcd, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x02, 0x2b, 0xe0, 0x05, 0x8b, 0xba, 0x41, 0x05, 0x95, 0x6c, 0xb8, 0x84, 0x51, 0xcc,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x85, 0xeb, 0x2e, 0x8e, 0x87, 0x38, 0x04, 0x62, 0x3c, 0x2b,
	0x54, 0x77, 0x31, 0x44, 0x4b, 0x8a, 0x51, 0x29, 0x97, 0x38, 0x92, 0x6e, 0xdf, 0xcc, 0xe4, 0xa2,
	0xfc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xa1, 0x28, 0x2e, 0x6e, 0x24, 0x29, 0x21, 0x65,
	0x3d, 0x34, 0x1f, 0xe9, 0x61, 0x7a, 0x47, 0x4a, 0x05, 0xbf, 0x22, 0x88, 0xdb, 0x94, 0x18, 0x9c,
	0x04, 0xa3, 0xf8, 0xd1, 0x02, 0x31, 0x89, 0x0d, 0x1c, 0x6e, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x5b, 0xec, 0x2b, 0xcd, 0x5e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResizeImageMicroServiceClient is the client API for ResizeImageMicroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResizeImageMicroServiceClient interface {
	ResizeImage(ctx context.Context, in *ResizeImageRequest, opts ...grpc.CallOption) (*ResizeImageResponse, error)
}

type resizeImageMicroServiceClient struct {
	cc *grpc.ClientConn
}

func NewResizeImageMicroServiceClient(cc *grpc.ClientConn) ResizeImageMicroServiceClient {
	return &resizeImageMicroServiceClient{cc}
}

func (c *resizeImageMicroServiceClient) ResizeImage(ctx context.Context, in *ResizeImageRequest, opts ...grpc.CallOption) (*ResizeImageResponse, error) {
	out := new(ResizeImageResponse)
	err := c.cc.Invoke(ctx, "/resize_image_ms.ResizeImageMicroService/ResizeImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResizeImageMicroServiceServer is the server API for ResizeImageMicroService service.
type ResizeImageMicroServiceServer interface {
	ResizeImage(context.Context, *ResizeImageRequest) (*ResizeImageResponse, error)
}

// UnimplementedResizeImageMicroServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResizeImageMicroServiceServer struct {
}

func (*UnimplementedResizeImageMicroServiceServer) ResizeImage(ctx context.Context, req *ResizeImageRequest) (*ResizeImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeImage not implemented")
}

func RegisterResizeImageMicroServiceServer(s *grpc.Server, srv ResizeImageMicroServiceServer) {
	s.RegisterService(&_ResizeImageMicroService_serviceDesc, srv)
}

func _ResizeImageMicroService_ResizeImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResizeImageMicroServiceServer).ResizeImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resize_image_ms.ResizeImageMicroService/ResizeImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResizeImageMicroServiceServer).ResizeImage(ctx, req.(*ResizeImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResizeImageMicroService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resize_image_ms.ResizeImageMicroService",
	HandlerType: (*ResizeImageMicroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResizeImage",
			Handler:    _ResizeImageMicroService_ResizeImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resizeimagemspb/resizeimagems.proto",
}