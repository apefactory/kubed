// Code generated by protoc-gen-go.
// source: metadata.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerInfoResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Provider  string                  `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	ServerUrl string                  `protobuf:"bytes,3,opt,name=server_url,json=serverUrl" json:"server_url,omitempty"`
	CaCert    string                  `protobuf:"bytes,4,opt,name=ca_cert,json=caCert" json:"ca_cert,omitempty"`
}

func (m *ServerInfoResponse) Reset()                    { *m = ServerInfoResponse{} }
func (m *ServerInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerInfoResponse) ProtoMessage()               {}
func (*ServerInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServerInfoResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ServerInfoResponse) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ServerInfoResponse) GetServerUrl() string {
	if m != nil {
		return m.ServerUrl
	}
	return ""
}

func (m *ServerInfoResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerInfoResponse)(nil), "appscode.ci.v1beta1.ServerInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	ServerInfo(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ServerInfo(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	out := new(ServerInfoResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Metadata/ServerInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	ServerInfo(context.Context, *appscode_dtypes.VoidRequest) (*ServerInfoResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_ServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Metadata/ServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ServerInfo(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerInfo",
			Handler:    _Metadata_ServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x53, 0x34, 0xfc, 0xa9, 0xd1, 0x43, 0x3d, 0xb0, 0x59, 0x51, 0x09, 0x17, 0xb9, 0xb8,
	0x0d, 0x10, 0xe3, 0x59, 0x3c, 0x79, 0x30, 0x21, 0x4b, 0xf4, 0xe0, 0x85, 0x94, 0xee, 0x40, 0x6a,
	0xa0, 0x53, 0xdb, 0x42, 0xe2, 0x95, 0x57, 0x30, 0x31, 0xf1, 0xec, 0x33, 0xf8, 0x24, 0xbe, 0x82,
	0x0f, 0x62, 0xdc, 0x5d, 0x40, 0x82, 0x5e, 0x9a, 0xb4, 0xbf, 0x99, 0x7e, 0xdf, 0x7c, 0x43, 0x0f,
	0xa6, 0xe0, 0x45, 0x22, 0xbc, 0x88, 0x8c, 0x45, 0x8f, 0xec, 0x50, 0x18, 0xe3, 0x24, 0x26, 0x10,
	0x49, 0x15, 0xcd, 0x5b, 0x43, 0xf0, 0xa2, 0x15, 0xd6, 0xc6, 0x88, 0xe3, 0x09, 0x70, 0x61, 0x14,
	0x17, 0x5a, 0xa3, 0x17, 0x5e, 0xa1, 0x76, 0x59, 0x4b, 0x78, 0xb2, 0x6c, 0xf9, 0x87, 0x9f, 0x6e,
	0xf0, 0xc4, 0x3f, 0x1b, 0x70, 0x3c, 0x3d, 0xb3, 0x82, 0xc6, 0x1b, 0xa1, 0xac, 0x0f, 0x76, 0x0e,
	0xf6, 0x46, 0x8f, 0x30, 0x06, 0x67, 0x50, 0x3b, 0x60, 0x9c, 0x16, 0x9d, 0x17, 0x7e, 0xe6, 0x02,
	0x52, 0x27, 0xcd, 0xbd, 0x76, 0x35, 0x5a, 0x79, 0xcb, 0x3e, 0x89, 0xfa, 0x29, 0x8e, 0xf3, 0x32,
	0x16, 0xd2, 0xb2, 0xb1, 0x38, 0x57, 0x09, 0xd8, 0xa0, 0x50, 0x27, 0xcd, 0x4a, 0xbc, 0xba, 0xb3,
	0x63, 0x4a, 0x5d, 0x2a, 0x31, 0x98, 0xd9, 0x49, 0xb0, 0x93, 0xd2, 0x4a, 0xf6, 0x72, 0x67, 0x27,
	0xac, 0x4a, 0x4b, 0x52, 0x0c, 0x24, 0x58, 0x1f, 0xec, 0xa6, 0xac, 0x28, 0xc5, 0x35, 0x58, 0xdf,
	0x7e, 0x27, 0xb4, 0x7c, 0x9b, 0x47, 0xc4, 0x5e, 0x09, 0xa5, 0x6b, 0xa3, 0xac, 0xb6, 0x65, 0xe8,
	0x1e, 0x55, 0x12, 0xc3, 0xd3, 0x0c, 0x9c, 0x0f, 0xcf, 0xa2, 0x3f, 0xa2, 0x8c, 0xb6, 0xe7, 0x6c,
	0x5c, 0x2d, 0x3e, 0x82, 0x42, 0x99, 0x2c, 0x3e, 0xbf, 0x5e, 0x0a, 0x17, 0xac, 0xc3, 0x07, 0x1b,
	0x79, 0x49, 0xc5, 0xf3, 0x5e, 0xbe, 0x5c, 0x15, 0xcf, 0x7c, 0x9f, 0x2b, 0x3d, 0x42, 0xfe, 0xe8,
	0x50, 0x77, 0x2f, 0xe9, 0x91, 0xc4, 0xe9, 0x5a, 0x50, 0x18, 0xf5, 0x4b, 0xb4, 0xbb, 0xbf, 0x9c,
	0xa0, 0xf7, 0x93, 0x77, 0x8f, 0x3c, 0x94, 0x72, 0x32, 0x2c, 0xa6, 0x1b, 0xe8, 0x7c, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x64, 0x50, 0x7c, 0xba, 0x07, 0x02, 0x00, 0x00,
}
