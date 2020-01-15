// Code generated by protoc-gen-go. DO NOT EDIT.
// source: employee.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Employee struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Position             string   `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Grade                uint64   `protobuf:"varint,5,opt,name=grade,proto3" json:"grade,omitempty"`
	TuMembership         bool     `protobuf:"varint,6,opt,name=tu_membership,json=tuMembership,proto3" json:"tu_membership,omitempty"`
	Salary               float64  `protobuf:"fixed64,7,opt,name=salary,proto3" json:"salary,omitempty"`
	Accruals             float64  `protobuf:"fixed64,8,opt,name=accruals,proto3" json:"accruals,omitempty"`
	Deduction            float64  `protobuf:"fixed64,9,opt,name=deduction,proto3" json:"deduction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{0}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Employee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Employee) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Employee) GetGrade() uint64 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *Employee) GetTuMembership() bool {
	if m != nil {
		return m.TuMembership
	}
	return false
}

func (m *Employee) GetSalary() float64 {
	if m != nil {
		return m.Salary
	}
	return 0
}

func (m *Employee) GetAccruals() float64 {
	if m != nil {
		return m.Accruals
	}
	return 0
}

func (m *Employee) GetDeduction() float64 {
	if m != nil {
		return m.Deduction
	}
	return 0
}

type ID struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{1}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Date struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{2}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type IDTime struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDTime) Reset()         { *m = IDTime{} }
func (m *IDTime) String() string { return proto.CompactTextString(m) }
func (*IDTime) ProtoMessage()    {}
func (*IDTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{3}
}

func (m *IDTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDTime.Unmarshal(m, b)
}
func (m *IDTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDTime.Marshal(b, m, deterministic)
}
func (m *IDTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDTime.Merge(m, src)
}
func (m *IDTime) XXX_Size() int {
	return xxx_messageInfo_IDTime.Size(m)
}
func (m *IDTime) XXX_DiscardUnknown() {
	xxx_messageInfo_IDTime.DiscardUnknown(m)
}

var xxx_messageInfo_IDTime proto.InternalMessageInfo

func (m *IDTime) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IDTime) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*Employee)(nil), "api.Employee")
	proto.RegisterType((*ID)(nil), "api.ID")
	proto.RegisterType((*Date)(nil), "api.Date")
	proto.RegisterType((*IDTime)(nil), "api.IDTime")
}

func init() { proto.RegisterFile("employee.proto", fileDescriptor_eb50a19aa79a6eac) }

var fileDescriptor_eb50a19aa79a6eac = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xdd, 0x8a, 0x9b, 0x40,
	0x18, 0x65, 0x8c, 0x31, 0xfa, 0xe5, 0xa7, 0x30, 0x0d, 0x61, 0xb0, 0xbd, 0x10, 0x4b, 0xc1, 0x8b,
	0xd6, 0xb4, 0x0d, 0x7d, 0x80, 0x82, 0xa1, 0x04, 0x5a, 0x28, 0xd2, 0xfb, 0x65, 0x92, 0xf9, 0x36,
	0x3b, 0xa0, 0x51, 0x74, 0xbc, 0xf0, 0x89, 0xf6, 0xe1, 0xf6, 0x25, 0x16, 0xc7, 0x9f, 0x24, 0xbb,
	0x04, 0xf6, 0x6e, 0xce, 0xf9, 0xce, 0xf7, 0x73, 0xce, 0xc0, 0x02, 0xd3, 0x3c, 0xc9, 0x6a, 0xc4,
	0x30, 0x2f, 0x32, 0x95, 0xd1, 0x11, 0xcf, 0xa5, 0xfb, 0xe1, 0x98, 0x65, 0xc7, 0x04, 0xd7, 0x9a,
	0xda, 0x57, 0xf7, 0x6b, 0x4c, 0x73, 0x55, 0xb7, 0x0a, 0xff, 0x89, 0x80, 0xbd, 0xed, 0x9a, 0xe8,
	0x02, 0x0c, 0x29, 0x18, 0xf1, 0x48, 0x60, 0xc6, 0x86, 0x14, 0x94, 0x82, 0x29, 0xb8, 0x42, 0x66,
	0x78, 0x24, 0x70, 0x62, 0xfd, 0x6e, 0xb8, 0x13, 0x4f, 0x91, 0x8d, 0x5a, 0xae, 0x79, 0x53, 0x17,
	0xec, 0x3c, 0x2b, 0xa5, 0x92, 0xd9, 0x89, 0x99, 0x9a, 0x1f, 0x30, 0x5d, 0xc2, 0xf8, 0x58, 0x70,
	0x81, 0x6c, 0xac, 0xc7, 0xb6, 0x80, 0x7e, 0x82, 0xb9, 0xaa, 0xee, 0x52, 0x4c, 0xf7, 0x58, 0x94,
	0x0f, 0x32, 0x67, 0x96, 0x47, 0x02, 0x3b, 0x9e, 0xa9, 0xea, 0xef, 0xc0, 0xd1, 0x15, 0x58, 0x25,
	0x4f, 0x78, 0x51, 0xb3, 0x89, 0x47, 0x02, 0x12, 0x77, 0xa8, 0x59, 0xc7, 0x0f, 0x87, 0xa2, 0xe2,
	0x49, 0xc9, 0x6c, 0x5d, 0x19, 0x30, 0xfd, 0x08, 0x8e, 0x40, 0x51, 0x1d, 0xf4, 0x2d, 0x8e, 0x2e,
	0x9e, 0x09, 0x7f, 0x09, 0xc6, 0x2e, 0x7a, 0x69, 0xd3, 0x77, 0xc1, 0x8c, 0x3a, 0x6b, 0xda, 0x2e,
	0x39, 0xdb, 0xf5, 0xbf, 0x80, 0xb5, 0x8b, 0xfe, 0xcb, 0xf4, 0x4d, 0xe1, 0xfc, 0x78, 0x34, 0xc0,
	0xe9, 0xd3, 0x2c, 0xe9, 0x67, 0x98, 0xfe, 0x12, 0x62, 0x48, 0x77, 0x1e, 0xf2, 0x5c, 0x86, 0x3d,
	0x74, 0x27, 0x1a, 0xee, 0x22, 0xfa, 0x13, 0x66, 0x5b, 0x21, 0xd5, 0x2d, 0xdd, 0x2a, 0x6c, 0xff,
	0x2f, 0xec, 0xff, 0xaf, 0x29, 0xa9, 0x9a, 0x7e, 0x87, 0x45, 0x84, 0x09, 0x2a, 0x1c, 0x1a, 0xfb,
	0x89, 0x37, 0x5b, 0xbe, 0xc2, 0xbb, 0xdf, 0x38, 0x2c, 0xfa, 0x23, 0x4b, 0x45, 0x1d, 0xdd, 0xd3,
	0xd8, 0x77, 0xaf, 0xf7, 0x7e, 0x23, 0xcd, 0xfd, 0x17, 0xf2, 0xf3, 0xf8, 0x6b, 0x21, 0xdd, 0xc0,
	0xfb, 0x0b, 0xd9, 0x3f, 0x5e, 0xa7, 0x78, 0x52, 0x25, 0x9d, 0x76, 0xf2, 0x26, 0xbc, 0x57, 0xb3,
	0xf7, 0x96, 0x3e, 0x6d, 0xf3, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x0c, 0xec, 0x20, 0xb2, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmployeesClient is the client API for Employees service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeesClient interface {
	AddEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ID, error)
	EditEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*empty.Empty, error)
	GetEmployeeList(ctx context.Context, in *Date, opts ...grpc.CallOption) (Employees_GetEmployeeListClient, error)
	GetEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error)
	GetEmployeePayments(ctx context.Context, in *IDTime, opts ...grpc.CallOption) (Employees_GetEmployeePaymentsClient, error)
}

type employeesClient struct {
	cc *grpc.ClientConn
}

func NewEmployeesClient(cc *grpc.ClientConn) EmployeesClient {
	return &employeesClient{cc}
}

func (c *employeesClient) AddEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/api.Employees/AddEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeesClient) EditEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Employees/EditEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeesClient) DeleteEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Employees/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeesClient) GetEmployeeList(ctx context.Context, in *Date, opts ...grpc.CallOption) (Employees_GetEmployeeListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Employees_serviceDesc.Streams[0], "/api.Employees/GetEmployeeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeesGetEmployeeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Employees_GetEmployeeListClient interface {
	Recv() (*Employee, error)
	grpc.ClientStream
}

type employeesGetEmployeeListClient struct {
	grpc.ClientStream
}

func (x *employeesGetEmployeeListClient) Recv() (*Employee, error) {
	m := new(Employee)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeesClient) GetEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/api.Employees/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeesClient) GetEmployeePayments(ctx context.Context, in *IDTime, opts ...grpc.CallOption) (Employees_GetEmployeePaymentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Employees_serviceDesc.Streams[1], "/api.Employees/GetEmployeePayments", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeesGetEmployeePaymentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Employees_GetEmployeePaymentsClient interface {
	Recv() (*Employee, error)
	grpc.ClientStream
}

type employeesGetEmployeePaymentsClient struct {
	grpc.ClientStream
}

func (x *employeesGetEmployeePaymentsClient) Recv() (*Employee, error) {
	m := new(Employee)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeesServer is the server API for Employees service.
type EmployeesServer interface {
	AddEmployee(context.Context, *Employee) (*ID, error)
	EditEmployee(context.Context, *Employee) (*empty.Empty, error)
	DeleteEmployee(context.Context, *ID) (*empty.Empty, error)
	GetEmployeeList(*Date, Employees_GetEmployeeListServer) error
	GetEmployee(context.Context, *ID) (*Employee, error)
	GetEmployeePayments(*IDTime, Employees_GetEmployeePaymentsServer) error
}

// UnimplementedEmployeesServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeesServer struct {
}

func (*UnimplementedEmployeesServer) AddEmployee(ctx context.Context, req *Employee) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (*UnimplementedEmployeesServer) EditEmployee(ctx context.Context, req *Employee) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEmployee not implemented")
}
func (*UnimplementedEmployeesServer) DeleteEmployee(ctx context.Context, req *ID) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (*UnimplementedEmployeesServer) GetEmployeeList(req *Date, srv Employees_GetEmployeeListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEmployeeList not implemented")
}
func (*UnimplementedEmployeesServer) GetEmployee(ctx context.Context, req *ID) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (*UnimplementedEmployeesServer) GetEmployeePayments(req *IDTime, srv Employees_GetEmployeePaymentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEmployeePayments not implemented")
}

func RegisterEmployeesServer(s *grpc.Server, srv EmployeesServer) {
	s.RegisterService(&_Employees_serviceDesc, srv)
}

func _Employees_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeesServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Employees/AddEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeesServer).AddEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employees_EditEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeesServer).EditEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Employees/EditEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeesServer).EditEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employees_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeesServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Employees/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeesServer).DeleteEmployee(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employees_GetEmployeeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Date)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeesServer).GetEmployeeList(m, &employeesGetEmployeeListServer{stream})
}

type Employees_GetEmployeeListServer interface {
	Send(*Employee) error
	grpc.ServerStream
}

type employeesGetEmployeeListServer struct {
	grpc.ServerStream
}

func (x *employeesGetEmployeeListServer) Send(m *Employee) error {
	return x.ServerStream.SendMsg(m)
}

func _Employees_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeesServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Employees/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeesServer).GetEmployee(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employees_GetEmployeePayments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IDTime)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeesServer).GetEmployeePayments(m, &employeesGetEmployeePaymentsServer{stream})
}

type Employees_GetEmployeePaymentsServer interface {
	Send(*Employee) error
	grpc.ServerStream
}

type employeesGetEmployeePaymentsServer struct {
	grpc.ServerStream
}

func (x *employeesGetEmployeePaymentsServer) Send(m *Employee) error {
	return x.ServerStream.SendMsg(m)
}

var _Employees_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Employees",
	HandlerType: (*EmployeesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEmployee",
			Handler:    _Employees_AddEmployee_Handler,
		},
		{
			MethodName: "EditEmployee",
			Handler:    _Employees_EditEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _Employees_DeleteEmployee_Handler,
		},
		{
			MethodName: "GetEmployee",
			Handler:    _Employees_GetEmployee_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEmployeeList",
			Handler:       _Employees_GetEmployeeList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetEmployeePayments",
			Handler:       _Employees_GetEmployeePayments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "employee.proto",
}