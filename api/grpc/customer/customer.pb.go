// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

/*
Package grpc_customer is a generated protocol buffer package.

It is generated from these files:
	customer.proto

It has these top-level messages:
	CustomerRequest
	CustomerResponse
	CustomerFilter
*/
package grpc_customer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import grpc_example "NarasimhaReddy12/golang_examples/api/grpc/example"
import dnsin "Infoblox-CTO/cdc.manager/cdc-app/pkg/pb/dnsin"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Request message for creating a new customer
type CustomerRequest struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email     string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone     string                     `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Addresses []*CustomerRequest_Address `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
	Info      *grpc_example.Information  `protobuf:"bytes,6,opt,name=info" json:"info,omitempty"`
	Status    *dnsin.GetStatusRequest    `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
}

func (m *CustomerRequest) Reset()                    { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()               {}
func (*CustomerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CustomerRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CustomerRequest) GetAddresses() []*CustomerRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *CustomerRequest) GetInfo() *grpc_example.Information {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *CustomerRequest) GetStatus() *dnsin.GetStatusRequest {
	if m != nil {
		return m.Status
	}
	return nil
}

type CustomerRequest_Address struct {
	Street            string `protobuf:"bytes,1,opt,name=street" json:"street,omitempty"`
	City              string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	State             string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Zip               string `protobuf:"bytes,4,opt,name=zip" json:"zip,omitempty"`
	IsShippingAddress bool   `protobuf:"varint,5,opt,name=isShippingAddress" json:"isShippingAddress,omitempty"`
}

func (m *CustomerRequest_Address) Reset()                    { *m = CustomerRequest_Address{} }
func (m *CustomerRequest_Address) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest_Address) ProtoMessage()               {}
func (*CustomerRequest_Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CustomerRequest_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *CustomerRequest_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CustomerRequest_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CustomerRequest_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *CustomerRequest_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

type CustomerResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *CustomerResponse) Reset()                    { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string            { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()               {}
func (*CustomerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CustomerResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CustomerFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *CustomerFilter) Reset()                    { *m = CustomerFilter{} }
func (m *CustomerFilter) String() string            { return proto.CompactTextString(m) }
func (*CustomerFilter) ProtoMessage()               {}
func (*CustomerFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomerFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomerRequest)(nil), "grpc.customer.CustomerRequest")
	proto.RegisterType((*CustomerRequest_Address)(nil), "grpc.customer.CustomerRequest.Address")
	proto.RegisterType((*CustomerResponse)(nil), "grpc.customer.CustomerResponse")
	proto.RegisterType((*CustomerFilter)(nil), "grpc.customer.CustomerFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Customer service

type CustomerClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error)
	// Create a new Customer - A simple RPC
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Customer_serviceDesc.Streams[0], c.cc, "/grpc.customer.Customer/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_GetCustomersClient interface {
	Recv() (*CustomerRequest, error)
	grpc.ClientStream
}

type customerGetCustomersClient struct {
	grpc.ClientStream
}

func (x *customerGetCustomersClient) Recv() (*CustomerRequest, error) {
	m := new(CustomerRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := grpc.Invoke(ctx, "/grpc.customer.Customer/CreateCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Customer service

type CustomerServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(*CustomerFilter, Customer_GetCustomersServer) error
	// Create a new Customer - A simple RPC
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).GetCustomers(m, &customerGetCustomersServer{stream})
}

type Customer_GetCustomersServer interface {
	Send(*CustomerRequest) error
	grpc.ServerStream
}

type customerGetCustomersServer struct {
	grpc.ServerStream
}

func (x *customerGetCustomersServer) Send(m *CustomerRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.customer.Customer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _Customer_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "customer.proto",
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0xe4, 0xb7, 0x71, 0xbf, 0x2f, 0x04, 0x83, 0xc0, 0x44, 0x50, 0xa2, 0x59, 0xa0, 0x51,
	0x45, 0xc6, 0x34, 0xac, 0x40, 0x48, 0x08, 0x05, 0x51, 0xb1, 0x01, 0x69, 0xca, 0x1e, 0x39, 0x33,
	0xb7, 0x13, 0xab, 0x33, 0xb6, 0x6b, 0x3b, 0xd0, 0xb0, 0x64, 0xcb, 0x12, 0xf1, 0x64, 0xbc, 0x02,
	0x0b, 0x1e, 0x03, 0xd9, 0x33, 0xa6, 0x2a, 0x55, 0x61, 0x93, 0xdc, 0x73, 0x7d, 0xee, 0x39, 0xf6,
	0xb9, 0x83, 0xc6, 0xf9, 0xc6, 0x58, 0x59, 0x83, 0x4e, 0x95, 0x96, 0x56, 0xe2, 0xff, 0x4b, 0xad,
	0xf2, 0x34, 0x34, 0xa7, 0x77, 0x4b, 0x29, 0xcb, 0x0a, 0x28, 0x53, 0x9c, 0x32, 0x21, 0xa4, 0x65,
	0x96, 0x4b, 0x61, 0x1a, 0xf2, 0xf4, 0xf9, 0x1b, 0xa6, 0x99, 0xe1, 0xf5, 0x9a, 0x65, 0x50, 0x14,
	0xdb, 0x83, 0x05, 0x2d, 0x65, 0xc5, 0x44, 0xf9, 0x1e, 0xce, 0x58, 0xad, 0x2a, 0x30, 0x7e, 0xce,
	0x29, 0xd2, 0xb6, 0x13, 0xfe, 0x5b, 0x81, 0x27, 0xaf, 0xc5, 0xb1, 0x5c, 0x55, 0xf2, 0x6c, 0xbe,
	0x7c, 0xf7, 0x96, 0xe6, 0x45, 0x9e, 0xd6, 0x4c, 0xb0, 0x12, 0xb4, 0xab, 0xe7, 0x4c, 0x29, 0xaa,
	0x4e, 0x4a, 0xaa, 0x56, 0xb4, 0x10, 0x86, 0x8b, 0xe6, 0xb7, 0x19, 0x8d, 0xbf, 0x75, 0xd1, 0xb5,
	0x65, 0x7b, 0xcd, 0x0c, 0x4e, 0x37, 0x60, 0x2c, 0x1e, 0xa3, 0x0e, 0x2f, 0x48, 0x34, 0x8b, 0x92,
	0x7e, 0xd6, 0xe1, 0x05, 0xc6, 0xa8, 0x27, 0x58, 0x0d, 0xa4, 0x33, 0x8b, 0x92, 0x51, 0xe6, 0x6b,
	0x7c, 0x13, 0xf5, 0xa1, 0x66, 0xbc, 0x22, 0x5d, 0xdf, 0x6c, 0x80, 0xeb, 0xaa, 0xb5, 0x14, 0x40,
	0x7a, 0x4d, 0xd7, 0x03, 0xfc, 0x12, 0x8d, 0x58, 0x51, 0x68, 0x30, 0x06, 0x0c, 0xe9, 0xcf, 0xba,
	0xc9, 0xee, 0xe2, 0x41, 0x7a, 0x21, 0xa0, 0xf4, 0x8f, 0x2b, 0xa4, 0x2f, 0x1a, 0x7e, 0x76, 0x3e,
	0x88, 0xe7, 0xa8, 0xc7, 0xc5, 0xb1, 0x24, 0x83, 0x59, 0x94, 0xec, 0x2e, 0xee, 0x34, 0x02, 0x21,
	0x07, 0x17, 0x80, 0xae, 0x7d, 0xaa, 0x99, 0xa7, 0x61, 0x8a, 0x06, 0xc6, 0x32, 0xbb, 0x31, 0x64,
	0xe8, 0x07, 0x6e, 0xa7, 0xcd, 0xb3, 0x0f, 0xc1, 0x1e, 0xf9, 0x7e, 0x6b, 0x95, 0xb5, 0xb4, 0xe9,
	0x97, 0x08, 0x0d, 0x5b, 0x5b, 0x7c, 0xcb, 0x0d, 0x6b, 0x00, 0xeb, 0x53, 0x18, 0x65, 0x2d, 0x72,
	0x49, 0xe4, 0xdc, 0x6e, 0x43, 0x12, 0xae, 0x76, 0x6f, 0x76, 0x0a, 0x10, 0x92, 0xf0, 0x00, 0x4f,
	0x50, 0xf7, 0x13, 0x57, 0x6d, 0x0e, 0xae, 0xc4, 0x0f, 0xd1, 0x75, 0x6e, 0x8e, 0xd6, 0x5c, 0x29,
	0x2e, 0xca, 0xd6, 0x88, 0xf4, 0x67, 0x51, 0xb2, 0x93, 0x5d, 0x3e, 0x88, 0x9f, 0xa1, 0xc9, 0x79,
	0x26, 0x46, 0x49, 0x61, 0xe0, 0xd2, 0x5e, 0x08, 0x1a, 0x9a, 0x4d, 0x9e, 0x3b, 0x9d, 0x8e, 0xd7,
	0x09, 0x30, 0xde, 0x47, 0xe3, 0x30, 0xfd, 0x8a, 0x57, 0x16, 0xb4, 0xe3, 0x9e, 0xc0, 0xf6, 0xa3,
	0xd4, 0x45, 0xfb, 0xa4, 0x00, 0x17, 0x3f, 0x23, 0xb4, 0x13, 0xc8, 0x98, 0xa3, 0xff, 0x0e, 0xc1,
	0x06, 0x68, 0xf0, 0xbd, 0x2b, 0xf6, 0xd4, 0xa8, 0x4e, 0xf7, 0xfe, 0xbe, 0xc6, 0x98, 0x7c, 0xfe,
	0xfe, 0xe3, 0x6b, 0x07, 0xe3, 0x09, 0xfd, 0x70, 0x40, 0x03, 0x8b, 0x96, 0x60, 0x1f, 0x45, 0xf8,
	0x14, 0x8d, 0x97, 0x1a, 0x98, 0x85, 0xdf, 0xe6, 0xff, 0x50, 0x9b, 0xde, 0xbf, 0xf2, 0xbc, 0x09,
	0x28, 0xde, 0xf3, 0x76, 0x24, 0xbe, 0x71, 0xc1, 0x2e, 0xf7, 0x2e, 0x4f, 0xa3, 0xfd, 0xd5, 0xc0,
	0x7f, 0xf3, 0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x46, 0x20, 0x35, 0xae, 0x03, 0x00,
	0x00,
}