// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pug_sink.proto

package pug

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

type Request struct {
	Metrics              []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9e68b78a4ca3be, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Metric struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Value                float64  `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9e68b78a4ca3be, []int{1}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Metric) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Metric) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Metric) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9e68b78a4ca3be, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "pug.Request")
	proto.RegisterType((*Metric)(nil), "pug.Metric")
	proto.RegisterType((*Empty)(nil), "pug.Empty")
}

func init() { proto.RegisterFile("pug_sink.proto", fileDescriptor_6c9e68b78a4ca3be) }

var fileDescriptor_6c9e68b78a4ca3be = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0x89, 0xb9, 0xb9, 0xc5, 0xb9, 0xe2, 0x22, 0xb8, 0x08, 0xae, 0x42, 0x41, 0xc9, 0xaa,
	0x48, 0xf5, 0x15, 0x5c, 0xba, 0x49, 0x1f, 0x40, 0xaa, 0x0e, 0x21, 0xd4, 0xa4, 0xb1, 0x49, 0x84,
	0xbe, 0xbd, 0x74, 0x5a, 0x77, 0xdf, 0x7c, 0xf9, 0x99, 0x73, 0xe0, 0x36, 0x55, 0xf7, 0x9e, 0x7d,
	0x9c, 0xba, 0xb4, 0xcc, 0x65, 0x96, 0x3c, 0x55, 0xd7, 0x3e, 0x41, 0x63, 0xf1, 0xa7, 0x62, 0x2e,
	0xf2, 0x01, 0x9a, 0x80, 0x65, 0xf1, 0x9f, 0x59, 0x31, 0xcd, 0xcd, 0xa5, 0xbf, 0x74, 0xa9, 0xba,
	0xee, 0x8d, 0x9c, 0xfd, 0x3f, 0x6b, 0x23, 0x9c, 0x77, 0x25, 0x25, 0x9c, 0xe2, 0x18, 0x50, 0x31,
	0xcd, 0xcc, 0xb5, 0x25, 0xde, 0x5c, 0x59, 0x13, 0xaa, 0x2b, 0xcd, 0x8c, 0xb0, 0xc4, 0xe4, 0x7c,
	0x40, 0xc5, 0x35, 0x33, 0xdc, 0x12, 0xcb, 0x3b, 0x10, 0xbf, 0xe3, 0x77, 0x45, 0x75, 0xd2, 0xcc,
	0x30, 0xbb, 0x0f, 0x74, 0x73, 0x74, 0x59, 0x09, 0xcd, 0xb7, 0x1f, 0x37, 0x6e, 0x1b, 0x10, 0xaf,
	0x21, 0x95, 0xb5, 0x7f, 0x01, 0xd8, 0x17, 0x0f, 0x3e, 0x4e, 0xf2, 0x11, 0x60, 0xc0, 0xf8, 0x75,
	0x44, 0xb9, 0xa1, 0xa8, 0x47, 0x93, 0x7b, 0xa0, 0x89, 0x5e, 0x7d, 0x9c, 0xa9, 0xec, 0xf3, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xec, 0xf1, 0xa5, 0xfe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricSinkClient is the client API for MetricSink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricSinkClient interface {
	SendMetric(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
}

type metricSinkClient struct {
	cc *grpc.ClientConn
}

func NewMetricSinkClient(cc *grpc.ClientConn) MetricSinkClient {
	return &metricSinkClient{cc}
}

func (c *metricSinkClient) SendMetric(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pug.MetricSink/SendMetric", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricSinkServer is the server API for MetricSink service.
type MetricSinkServer interface {
	SendMetric(context.Context, *Request) (*Empty, error)
}

// UnimplementedMetricSinkServer can be embedded to have forward compatible implementations.
type UnimplementedMetricSinkServer struct {
}

func (*UnimplementedMetricSinkServer) SendMetric(ctx context.Context, req *Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMetric not implemented")
}

func RegisterMetricSinkServer(s *grpc.Server, srv MetricSinkServer) {
	s.RegisterService(&_MetricSink_serviceDesc, srv)
}

func _MetricSink_SendMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricSinkServer).SendMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pug.MetricSink/SendMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricSinkServer).SendMetric(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricSink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pug.MetricSink",
	HandlerType: (*MetricSinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMetric",
			Handler:    _MetricSink_SendMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pug_sink.proto",
}
