// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CommentReadAo.proto

package terry_proto_comment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetCommentRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	CommentId            uint64         `protobuf:"varint,2,opt,name=commentId,proto3" json:"commentId,omitempty"`
	UserId               uint64         `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             uint64         `protobuf:"varint,4,opt,name=userName,proto3" json:"userName,omitempty"`
	CommentText          string         `protobuf:"bytes,5,opt,name=commentText,proto3" json:"commentText,omitempty"`
	AddTime              string         `protobuf:"bytes,6,opt,name=addTime,proto3" json:"addTime,omitempty"`
	LastModifyTime       string         `protobuf:"bytes,7,opt,name=lastModifyTime,proto3" json:"lastModifyTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetCommentRequest) Reset()         { *m = GetCommentRequest{} }
func (m *GetCommentRequest) String() string { return proto.CompactTextString(m) }
func (*GetCommentRequest) ProtoMessage()    {}
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff0a69c68c4ec5e, []int{0}
}

func (m *GetCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommentRequest.Unmarshal(m, b)
}
func (m *GetCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommentRequest.Marshal(b, m, deterministic)
}
func (m *GetCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommentRequest.Merge(m, src)
}
func (m *GetCommentRequest) XXX_Size() int {
	return xxx_messageInfo_GetCommentRequest.Size(m)
}
func (m *GetCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommentRequest proto.InternalMessageInfo

func (m *GetCommentRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetCommentRequest) GetCommentId() uint64 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

func (m *GetCommentRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetCommentRequest) GetUserName() uint64 {
	if m != nil {
		return m.UserName
	}
	return 0
}

func (m *GetCommentRequest) GetCommentText() string {
	if m != nil {
		return m.CommentText
	}
	return ""
}

func (m *GetCommentRequest) GetAddTime() string {
	if m != nil {
		return m.AddTime
	}
	return ""
}

func (m *GetCommentRequest) GetLastModifyTime() string {
	if m != nil {
		return m.LastModifyTime
	}
	return ""
}

type GetCommentResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	CommentInfo          *CommentInfo    `protobuf:"bytes,2,opt,name=commentInfo,proto3" json:"commentInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetCommentResponse) Reset()         { *m = GetCommentResponse{} }
func (m *GetCommentResponse) String() string { return proto.CompactTextString(m) }
func (*GetCommentResponse) ProtoMessage()    {}
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff0a69c68c4ec5e, []int{1}
}

func (m *GetCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommentResponse.Unmarshal(m, b)
}
func (m *GetCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommentResponse.Marshal(b, m, deterministic)
}
func (m *GetCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommentResponse.Merge(m, src)
}
func (m *GetCommentResponse) XXX_Size() int {
	return xxx_messageInfo_GetCommentResponse.Size(m)
}
func (m *GetCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommentResponse proto.InternalMessageInfo

func (m *GetCommentResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetCommentResponse) GetCommentInfo() *CommentInfo {
	if m != nil {
		return m.CommentInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCommentRequest)(nil), "terry.proto.comment.GetCommentRequest")
	proto.RegisterType((*GetCommentResponse)(nil), "terry.proto.comment.GetCommentResponse")
}

func init() { proto.RegisterFile("CommentReadAo.proto", fileDescriptor_2ff0a69c68c4ec5e) }

var fileDescriptor_2ff0a69c68c4ec5e = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0xa5, 0x56, 0x6a, 0x07, 0x6d, 0xe2, 0x34, 0x31, 0x1b, 0xe2, 0x81, 0x60, 0x52, 0x39,
	0x71, 0xc0, 0x9b, 0x9e, 0xb4, 0x07, 0xe5, 0xa0, 0x07, 0xd2, 0x17, 0xc0, 0xee, 0x34, 0x36, 0x11,
	0xb6, 0xc2, 0x36, 0xb1, 0x0f, 0xe0, 0x23, 0xf8, 0xbe, 0x86, 0xdd, 0x15, 0x10, 0x31, 0xbd, 0xf1,
	0xff, 0xf3, 0xcd, 0x9f, 0xe5, 0x1f, 0x98, 0xce, 0x45, 0x96, 0x51, 0x2e, 0x13, 0x4a, 0xf9, 0x9d,
	0x08, 0x37, 0x85, 0x90, 0x02, 0xa7, 0x92, 0x8a, 0x62, 0xa7, 0x45, 0xb8, 0xd4, 0x80, 0x7b, 0x52,
	0x91, 0x22, 0xd7, 0xae, 0xff, 0x39, 0x80, 0xb3, 0x07, 0x92, 0xf5, 0xf6, 0xfb, 0x96, 0x4a, 0x89,
	0x37, 0x60, 0xbf, 0x52, 0xca, 0xa9, 0x60, 0x96, 0x67, 0x05, 0x4e, 0xe4, 0x87, 0x3d, 0x49, 0xa1,
	0xa1, 0x1f, 0x15, 0x99, 0x98, 0x0d, 0xbc, 0x80, 0xb1, 0x01, 0x62, 0xce, 0x06, 0x9e, 0x15, 0x0c,
	0x93, 0xc6, 0xc0, 0x73, 0xb0, 0xb7, 0x25, 0x15, 0x31, 0x67, 0x87, 0x6a, 0x64, 0x14, 0xba, 0x70,
	0x5c, 0x7d, 0x3d, 0xa7, 0x19, 0xb1, 0xa1, 0x9a, 0xd4, 0x1a, 0x3d, 0x70, 0x4c, 0xc0, 0x82, 0x3e,
	0x24, 0x3b, 0xf2, 0xac, 0x60, 0x9c, 0xb4, 0x2d, 0x64, 0x30, 0x4a, 0x39, 0x5f, 0xac, 0x33, 0x62,
	0xb6, 0x9a, 0xfe, 0x48, 0x9c, 0xc1, 0xe4, 0x2d, 0x2d, 0xe5, 0x93, 0xe0, 0xeb, 0xd5, 0x4e, 0x01,
	0x23, 0x05, 0x74, 0x5c, 0xff, 0xcb, 0x02, 0x6c, 0xf7, 0x50, 0x6e, 0x44, 0x5e, 0x12, 0xde, 0x76,
	0x8a, 0xb8, 0xfc, 0xa7, 0x08, 0x8d, 0x77, 0x9a, 0xb8, 0xaf, 0xdf, 0x1d, 0xe7, 0x2b, 0xa1, 0xba,
	0x70, 0x22, 0xaf, 0x37, 0x61, 0xde, 0x70, 0x49, 0x7b, 0x29, 0x92, 0x70, 0xfa, 0xeb, 0xb2, 0xb8,
	0x84, 0x49, 0xf3, 0xce, 0x0a, 0xc1, 0x59, 0x6f, 0xe2, 0x9f, 0xa3, 0xba, 0x57, 0x7b, 0x39, 0xfd,
	0x17, 0xfe, 0xc1, 0x8b, 0xad, 0x98, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x7e, 0xf4,
	0x5a, 0x56, 0x02, 0x00, 0x00,
}
