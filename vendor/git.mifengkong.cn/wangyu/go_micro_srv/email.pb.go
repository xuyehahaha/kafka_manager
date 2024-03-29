// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email.proto

package go_micro_srv_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SendRequest struct {
	EmailTo              string   `protobuf:"bytes,1,opt,name=email_to,json=emailTo,proto3" json:"email_to,omitempty"`
	EmailTitle           string   `protobuf:"bytes,2,opt,name=email_title,json=emailTitle,proto3" json:"email_title,omitempty"`
	EmailBody            string   `protobuf:"bytes,3,opt,name=email_body,json=emailBody,proto3" json:"email_body,omitempty"`
	EmailCc              string   `protobuf:"bytes,4,opt,name=email_cc,json=emailCc,proto3" json:"email_cc,omitempty"`
	EmailFile            []*File  `protobuf:"bytes,5,rep,name=email_file,json=emailFile,proto3" json:"email_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetEmailTo() string {
	if m != nil {
		return m.EmailTo
	}
	return ""
}

func (m *SendRequest) GetEmailTitle() string {
	if m != nil {
		return m.EmailTitle
	}
	return ""
}

func (m *SendRequest) GetEmailBody() string {
	if m != nil {
		return m.EmailBody
	}
	return ""
}

func (m *SendRequest) GetEmailCc() string {
	if m != nil {
		return m.EmailCc
	}
	return ""
}

func (m *SendRequest) GetEmailFile() []*File {
	if m != nil {
		return m.EmailFile
	}
	return nil
}

type File struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size                 uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{1}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *File) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendResponse struct {
	Code                 int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{2}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SendResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "email.SendRequest")
	proto.RegisterType((*File)(nil), "email.File")
	proto.RegisterType((*SendResponse)(nil), "email.SendResponse")
	proto.RegisterMapType((map[string]string)(nil), "email.SendResponse.DataEntry")
}

func init() { proto.RegisterFile("email.proto", fileDescriptor_6175298cb4ed6faa) }

var fileDescriptor_6175298cb4ed6faa = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0x4d, 0xa1, 0xbd, 0xc0, 0x57, 0xee, 0xcd, 0xcd, 0xc7, 0x5d, 0x94, 0xe6, 0x12, 0x49, 0x57,
	0x84, 0x05, 0x0d, 0xb8, 0xd0, 0xb0, 0x54, 0x71, 0xc1, 0xb2, 0xb8, 0x27, 0x43, 0x3b, 0x92, 0x89,
	0x65, 0x06, 0x99, 0x81, 0xa4, 0x2e, 0x7d, 0x05, 0x5f, 0xc2, 0x9d, 0x0f, 0xe3, 0x2b, 0xf8, 0x20,
	0x66, 0x66, 0x5a, 0xc4, 0xc4, 0x55, 0xcf, 0xf9, 0x7e, 0xce, 0x39, 0x5f, 0x3a, 0xe0, 0xd3, 0x0d,
	0x61, 0xf9, 0x68, 0xbb, 0x13, 0x4a, 0xa0, 0x67, 0x48, 0xf8, 0x7f, 0x2d, 0xc4, 0x3a, 0xa7, 0x31,
	0xd9, 0xb2, 0x98, 0x70, 0x2e, 0x14, 0x51, 0x4c, 0x70, 0x69, 0x87, 0xa2, 0x37, 0x07, 0xfc, 0x05,
	0xe5, 0x59, 0x42, 0x1f, 0xf7, 0x54, 0x2a, 0xec, 0x42, 0xd3, 0xac, 0x2d, 0x95, 0x08, 0x9c, 0xbe,
	0x33, 0x68, 0x25, 0x0d, 0xc3, 0xef, 0x04, 0x9e, 0x95, 0xf2, 0x4b, 0xc5, 0x54, 0x4e, 0x83, 0x9a,
	0xe9, 0x82, 0xed, 0xea, 0x0a, 0xf6, 0xc0, 0xb2, 0xe5, 0x4a, 0x64, 0x45, 0x50, 0x37, 0xfd, 0x96,
	0xa9, 0x5c, 0x89, 0xac, 0xf8, 0x92, 0x4e, 0xd3, 0xc0, 0x3d, 0x91, 0xbe, 0x4e, 0x71, 0x58, 0x6d,
	0xde, 0xb3, 0x9c, 0x06, 0x5e, 0xbf, 0x3e, 0xf0, 0x27, 0xfe, 0xc8, 0x1e, 0x73, 0xcb, 0x72, 0x5a,
	0xca, 0x68, 0x18, 0xcd, 0xc1, 0xd5, 0x5f, 0x0c, 0xa1, 0xa9, 0xa7, 0x39, 0xd9, 0xd0, 0x32, 0xe9,
	0x91, 0x23, 0x82, 0x2b, 0xd9, 0x93, 0xcd, 0xf8, 0x3b, 0x31, 0x58, 0xd7, 0x32, 0xa2, 0x88, 0xc9,
	0xd5, 0x4e, 0x0c, 0x8e, 0x5e, 0x1d, 0x68, 0xdb, 0xeb, 0xe5, 0x56, 0x70, 0x69, 0x86, 0x52, 0x91,
	0x59, 0x41, 0x2f, 0x31, 0x18, 0x03, 0x68, 0x6c, 0xa8, 0x94, 0x64, 0x5d, 0xdd, 0x5c, 0x51, 0x1c,
	0x1f, 0x25, 0x75, 0xe0, 0x5e, 0x19, 0xf8, 0x54, 0x70, 0x74, 0x43, 0x14, 0x99, 0x71, 0xb5, 0x2b,
	0xac, 0x63, 0x78, 0x01, 0xad, 0x63, 0x09, 0xff, 0x42, 0xfd, 0x81, 0x16, 0x65, 0x7a, 0x0d, 0xf1,
	0x1f, 0x78, 0x07, 0x92, 0xef, 0x2b, 0x27, 0x4b, 0xa6, 0xb5, 0x4b, 0x67, 0xb2, 0x00, 0x6f, 0xa6,
	0xe5, 0x71, 0x0e, 0xae, 0x76, 0x40, 0xfc, 0x66, 0x67, 0xfe, 0x5e, 0xd8, 0xf9, 0x21, 0x42, 0xd4,
	0x7d, 0x7e, 0xff, 0x78, 0xa9, 0x75, 0xa2, 0x3f, 0xf1, 0x61, 0x1c, 0x9b, 0x7e, 0x2c, 0x29, 0xcf,
	0xa6, 0xce, 0x70, 0xf5, 0xcb, 0x3c, 0x82, 0xf3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x05,
	0xd4, 0xb0, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type emailClient struct {
	cc *grpc.ClientConn
}

func NewEmailClient(cc *grpc.ClientConn) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/email.Email/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
type EmailServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
}

func RegisterEmailServer(s *grpc.Server, srv EmailServer) {
	s.RegisterService(&_Email_serviceDesc, srv)
}

func _Email_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Email_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email.Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Email_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
