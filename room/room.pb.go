// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.proto

package room

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

type CreateRoomRequest struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	MaxUser              int32    `protobuf:"varint,4,opt,name=max_user,json=maxUser,proto3" json:"max_user,omitempty"`
	Token                []byte   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
}

func (m *CreateRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomRequest.Unmarshal(m, b)
}
func (m *CreateRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRequest.Merge(m, src)
}
func (m *CreateRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoomRequest.Size(m)
}
func (m *CreateRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRequest proto.InternalMessageInfo

func (m *CreateRoomRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *CreateRoomRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CreateRoomRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRoomRequest) GetMaxUser() int32 {
	if m != nil {
		return m.MaxUser
	}
	return 0
}

func (m *CreateRoomRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type CreateRoomResponse struct {
	Room                 *Room    `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomResponse) Reset()         { *m = CreateRoomResponse{} }
func (m *CreateRoomResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoomResponse) ProtoMessage()    {}
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
}

func (m *CreateRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomResponse.Unmarshal(m, b)
}
func (m *CreateRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomResponse.Marshal(b, m, deterministic)
}
func (m *CreateRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomResponse.Merge(m, src)
}
func (m *CreateRoomResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoomResponse.Size(m)
}
func (m *CreateRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomResponse proto.InternalMessageInfo

func (m *CreateRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type Room struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RequirePassword      bool     `protobuf:"varint,2,opt,name=require_password,json=requirePassword,proto3" json:"require_password,omitempty"`
	MaxUser              int32    `protobuf:"varint,3,opt,name=max_user,json=maxUser,proto3" json:"max_user,omitempty"`
	ConnectedUser        int32    `protobuf:"varint,4,opt,name=connected_user,json=connectedUser,proto3" json:"connected_user,omitempty"`
	Server               *Server  `protobuf:"bytes,5,opt,name=server,proto3" json:"server,omitempty"`
	ApplicationName      string   `protobuf:"bytes,6,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	Version              string   `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{2}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *Room) GetRequirePassword() bool {
	if m != nil {
		return m.RequirePassword
	}
	return false
}

func (m *Room) GetMaxUser() int32 {
	if m != nil {
		return m.MaxUser
	}
	return 0
}

func (m *Room) GetConnectedUser() int32 {
	if m != nil {
		return m.ConnectedUser
	}
	return 0
}

func (m *Room) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *Room) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *Room) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Server struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ServerId             int32    `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Token                []byte   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{3}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Server) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Server) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *Server) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoomRequest)(nil), "CreateRoomRequest")
	proto.RegisterType((*CreateRoomResponse)(nil), "CreateRoomResponse")
	proto.RegisterType((*Room)(nil), "Room")
	proto.RegisterType((*Server)(nil), "Server")
}

func init() { proto.RegisterFile("room.proto", fileDescriptor_c5fd27dd97284ef4) }

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0x36, 0x7f, 0x9d, 0xaa, 0xd5, 0x55, 0x30, 0xad, 0x07, 0x4b, 0x40, 0xa8, 0x97, 0x08,
	0x15, 0x9f, 0x40, 0x2f, 0xbd, 0x88, 0xac, 0x78, 0x0e, 0x6b, 0x32, 0x60, 0xd0, 0x64, 0xd2, 0xdd,
	0x6d, 0xed, 0x0b, 0xf9, 0x82, 0x3e, 0x81, 0xec, 0xa6, 0xb6, 0x29, 0x2d, 0x78, 0xdb, 0xef, 0x9b,
	0x21, 0xf9, 0x7e, 0x06, 0x40, 0x12, 0x95, 0x49, 0x2d, 0x49, 0x53, 0xfc, 0xed, 0xc0, 0xe9, 0x83,
	0x44, 0xa1, 0x91, 0x13, 0x95, 0x1c, 0x67, 0x73, 0x54, 0x9a, 0xdd, 0xc0, 0x89, 0xa8, 0xeb, 0xcf,
	0x22, 0x13, 0xba, 0xa0, 0x2a, 0xad, 0x44, 0x89, 0x91, 0x33, 0x72, 0xc6, 0x5d, 0xde, 0x6f, 0xf1,
	0x4f, 0xa2, 0x44, 0x16, 0x41, 0xb0, 0x40, 0xa9, 0x0a, 0xaa, 0xa2, 0x03, 0xbb, 0xf1, 0x07, 0xd9,
	0x10, 0xc2, 0x5a, 0x28, 0xf5, 0x45, 0x32, 0x8f, 0x3a, 0x76, 0xb4, 0xc6, 0x6c, 0x00, 0x61, 0x29,
	0x96, 0xe9, 0x5c, 0xa1, 0x8c, 0xdc, 0x91, 0x33, 0xf6, 0x78, 0x50, 0x8a, 0xe5, 0xab, 0x42, 0xc9,
	0xce, 0xc1, 0xd3, 0xf4, 0x81, 0x55, 0xe4, 0x8d, 0x9c, 0xf1, 0x21, 0x6f, 0x40, 0x7c, 0x0b, 0xac,
	0x2d, 0x53, 0xd5, 0x54, 0x29, 0x64, 0x03, 0x70, 0x8d, 0x17, 0xab, 0xad, 0x37, 0xf1, 0x12, 0x3b,
	0xb4, 0x54, 0xfc, 0xe3, 0x80, 0x6b, 0x20, 0xbb, 0x80, 0xc0, 0x10, 0x69, 0x91, 0xdb, 0x35, 0x8f,
	0xfb, 0x06, 0x4e, 0x73, 0x63, 0x52, 0xe2, 0x6c, 0x5e, 0x48, 0x4c, 0xd7, 0x3a, 0x8d, 0x85, 0x90,
	0xf7, 0x57, 0xfc, 0xf3, 0x3e, 0xb9, 0x9d, 0x6d, 0xb9, 0xd7, 0x70, 0x9c, 0x51, 0x55, 0x61, 0xa6,
	0x31, 0x6f, 0xfb, 0x39, 0x5a, 0xb3, 0x76, 0xed, 0x0a, 0x7c, 0x85, 0x72, 0x81, 0xd2, 0xda, 0xea,
	0x4d, 0x82, 0xe4, 0xc5, 0x42, 0xbe, 0xa2, 0xf7, 0x46, 0xee, 0xff, 0x1b, 0x79, 0xb0, 0x15, 0x79,
	0x9c, 0x81, 0xdf, 0x7c, 0x96, 0x31, 0x70, 0xdf, 0x49, 0xe9, 0x55, 0x6b, 0xf6, 0x6d, 0xb8, 0x9a,
	0xa4, 0xb6, 0x26, 0x3d, 0x6e, 0xdf, 0xec, 0x12, 0xba, 0x8d, 0x00, 0x93, 0x4f, 0x63, 0x2d, 0x6c,
	0x88, 0x69, 0xbe, 0xa9, 0xc2, 0x6d, 0x55, 0x31, 0x79, 0x84, 0x9e, 0x09, 0xd6, 0xfc, 0xa8, 0xc8,
	0x90, 0xdd, 0x03, 0x6c, 0x9a, 0x61, 0x2c, 0xd9, 0xb9, 0xa6, 0xe1, 0x59, 0xb2, 0x5b, 0xdd, 0x9b,
	0x6f, 0xef, 0xef, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x76, 0x15, 0x0f, 0xb8, 0x8d, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomServiceClient interface {
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
}

type roomServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoomServiceClient(cc *grpc.ClientConn) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
type RoomServiceServer interface {
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
}

// UnimplementedRoomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (*UnimplementedRoomServiceServer) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room.proto",
}
