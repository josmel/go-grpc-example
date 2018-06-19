// Code generated by protoc-gen-go. DO NOT EDIT.
// source: category.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type CQueryRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CQueryRequest) Reset()         { *m = CQueryRequest{} }
func (m *CQueryRequest) String() string { return proto.CompactTextString(m) }
func (*CQueryRequest) ProtoMessage()    {}
func (*CQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_8ece5b743b0640e2, []int{0}
}
func (m *CQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CQueryRequest.Unmarshal(m, b)
}
func (m *CQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CQueryRequest.Marshal(b, m, deterministic)
}
func (dst *CQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CQueryRequest.Merge(dst, src)
}
func (m *CQueryRequest) XXX_Size() int {
	return xxx_messageInfo_CQueryRequest.Size(m)
}
func (m *CQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CQueryRequest proto.InternalMessageInfo

func (m *CQueryRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CQueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CQueryRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

type CategoryInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ParentId             int64    `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Sort                 int64    `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryInfo) Reset()         { *m = CategoryInfo{} }
func (m *CategoryInfo) String() string { return proto.CompactTextString(m) }
func (*CategoryInfo) ProtoMessage()    {}
func (*CategoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_8ece5b743b0640e2, []int{1}
}
func (m *CategoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryInfo.Unmarshal(m, b)
}
func (m *CategoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryInfo.Marshal(b, m, deterministic)
}
func (dst *CategoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryInfo.Merge(dst, src)
}
func (m *CategoryInfo) XXX_Size() int {
	return xxx_messageInfo_CategoryInfo.Size(m)
}
func (m *CategoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryInfo proto.InternalMessageInfo

func (m *CategoryInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CategoryInfo) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *CategoryInfo) GetSort() int64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *CategoryInfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *CategoryInfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CategoryList struct {
	List                 []*CategoryInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CategoryList) Reset()         { *m = CategoryList{} }
func (m *CategoryList) String() string { return proto.CompactTextString(m) }
func (*CategoryList) ProtoMessage()    {}
func (*CategoryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_category_8ece5b743b0640e2, []int{2}
}
func (m *CategoryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryList.Unmarshal(m, b)
}
func (m *CategoryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryList.Marshal(b, m, deterministic)
}
func (dst *CategoryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryList.Merge(dst, src)
}
func (m *CategoryList) XXX_Size() int {
	return xxx_messageInfo_CategoryList.Size(m)
}
func (m *CategoryList) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryList.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryList proto.InternalMessageInfo

func (m *CategoryList) GetList() []*CategoryInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*CQueryRequest)(nil), "protos.CQueryRequest")
	proto.RegisterType((*CategoryInfo)(nil), "protos.CategoryInfo")
	proto.RegisterType((*CategoryList)(nil), "protos.CategoryList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CategoryClient is the client API for Category service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryClient interface {
	GetCategoryInfo(ctx context.Context, in *CQueryRequest, opts ...grpc.CallOption) (*CategoryInfo, error)
	GetCategoryList(ctx context.Context, in *CQueryRequest, opts ...grpc.CallOption) (*CategoryList, error)
}

type categoryClient struct {
	cc *grpc.ClientConn
}

func NewCategoryClient(cc *grpc.ClientConn) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) GetCategoryInfo(ctx context.Context, in *CQueryRequest, opts ...grpc.CallOption) (*CategoryInfo, error) {
	out := new(CategoryInfo)
	err := c.cc.Invoke(ctx, "/protos.Category/GetCategoryInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetCategoryList(ctx context.Context, in *CQueryRequest, opts ...grpc.CallOption) (*CategoryList, error) {
	out := new(CategoryList)
	err := c.cc.Invoke(ctx, "/protos.Category/GetCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
type CategoryServer interface {
	GetCategoryInfo(context.Context, *CQueryRequest) (*CategoryInfo, error)
	GetCategoryList(context.Context, *CQueryRequest) (*CategoryList, error)
}

func RegisterCategoryServer(s *grpc.Server, srv CategoryServer) {
	s.RegisterService(&_Category_serviceDesc, srv)
}

func _Category_GetCategoryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetCategoryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Category/GetCategoryInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetCategoryInfo(ctx, req.(*CQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Category/GetCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetCategoryList(ctx, req.(*CQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Category_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategoryInfo",
			Handler:    _Category_GetCategoryInfo_Handler,
		},
		{
			MethodName: "GetCategoryList",
			Handler:    _Category_GetCategoryList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}

func init() { proto.RegisterFile("category.proto", fileDescriptor_category_8ece5b743b0640e2) }

var fileDescriptor_category_8ece5b743b0640e2 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0x3b, 0x31,
	0x10, 0xc7, 0xc9, 0xf6, 0x0f, 0xed, 0xfc, 0x7e, 0x56, 0x0c, 0x55, 0x96, 0x5a, 0xa1, 0xec, 0xa9,
	0xf4, 0xd0, 0xc5, 0x7a, 0x91, 0xde, 0x4a, 0x0f, 0x52, 0xf0, 0xe2, 0x9e, 0x85, 0x12, 0x9b, 0x74,
	0x09, 0xac, 0xc9, 0x9a, 0x4c, 0x85, 0x5e, 0x7d, 0x05, 0x1f, 0xc1, 0x77, 0xf1, 0x05, 0x7c, 0x05,
	0x1f, 0x44, 0x36, 0xd9, 0xd5, 0x2a, 0x15, 0x3c, 0x65, 0xe6, 0x3b, 0x93, 0xcf, 0xfc, 0x49, 0xa0,
	0xb3, 0x62, 0x28, 0x52, 0x6d, 0xb6, 0xe3, 0xdc, 0x68, 0xd4, 0xb4, 0xe9, 0x0e, 0xdb, 0xeb, 0xa7,
	0x5a, 0xa7, 0x99, 0x88, 0x59, 0x2e, 0x63, 0xa6, 0x94, 0x46, 0x86, 0x52, 0x2b, 0xeb, 0xb3, 0xa2,
	0x05, 0x1c, 0xcc, 0x6f, 0x36, 0xc2, 0x6c, 0x13, 0xf1, 0xb0, 0x11, 0x16, 0x69, 0x07, 0x02, 0xc9,
	0x43, 0x32, 0x20, 0xc3, 0x5a, 0x12, 0x48, 0x4e, 0xbb, 0xd0, 0xc8, 0xe4, 0xbd, 0xc4, 0x30, 0x70,
	0x92, 0x77, 0x28, 0x85, 0x7a, 0xce, 0x52, 0x11, 0xd6, 0x9c, 0xe8, 0xec, 0xe8, 0x85, 0xc0, 0xff,
	0x79, 0xd9, 0xc3, 0x42, 0xad, 0xf5, 0x3e, 0x14, 0x4a, 0xcc, 0x84, 0x43, 0xb5, 0x13, 0xef, 0xd0,
	0x53, 0x68, 0xe7, 0xcc, 0x08, 0x85, 0x4b, 0xc9, 0x4b, 0x5e, 0xcb, 0x0b, 0x0b, 0x5e, 0xd4, 0xb1,
	0xda, 0x60, 0x58, 0xf7, 0x75, 0x0a, 0x9b, 0x9e, 0x01, 0xac, 0x8c, 0x60, 0x28, 0xf8, 0x92, 0x61,
	0xd8, 0x70, 0xac, 0x76, 0xa9, 0xcc, 0x5c, 0x78, 0x93, 0xf3, 0x2a, 0xdc, 0xf4, 0xe1, 0x52, 0x99,
	0x61, 0x74, 0xf9, 0xd5, 0xe4, 0xb5, 0xb4, 0x48, 0x87, 0x50, 0xcf, 0xa4, 0xc5, 0x90, 0x0c, 0x6a,
	0xc3, 0x7f, 0x93, 0xae, 0x5f, 0x8b, 0x1d, 0xef, 0x0e, 0x92, 0xb8, 0x8c, 0xc9, 0x2b, 0x81, 0x56,
	0x25, 0xd3, 0x5b, 0x38, 0xbc, 0x12, 0xf8, 0x6d, 0xdc, 0xe3, 0xcf, 0xbb, 0xbb, 0x0b, 0xed, 0xed,
	0x45, 0x46, 0xfd, 0xa7, 0xb7, 0xf7, 0xe7, 0xe0, 0x24, 0x3a, 0x8a, 0x1f, 0xcf, 0xe3, 0xea, 0xe5,
	0x62, 0xa9, 0xd6, 0x7a, 0x4a, 0x46, 0x3f, 0xe8, 0xae, 0xcf, 0xbf, 0xd2, 0x8b, 0xe4, 0x5f, 0xe8,
	0xc5, 0x14, 0x53, 0x32, 0xba, 0xf3, 0x3f, 0xe3, 0xe2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x71, 0xc7,
	0xa3, 0xe4, 0x32, 0x02, 0x00, 0x00,
}
