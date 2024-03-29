// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Common.proto

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

type RequestHeader struct {
	Chn                  string   `protobuf:"bytes,1,opt,name=chn,proto3" json:"chn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetChn() string {
	if m != nil {
		return m.Chn
	}
	return ""
}

type ResponseHeader struct {
	ErrorNo              uint32   `protobuf:"varint,1,opt,name=errorNo,proto3" json:"errorNo,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{1}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetErrorNo() uint32 {
	if m != nil {
		return m.ErrorNo
	}
	return 0
}

func (m *ResponseHeader) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type CommentInfo struct {
	CommentId            uint64   `protobuf:"varint,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
	UserId               uint64   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	CommentText          string   `protobuf:"bytes,4,opt,name=commentText,proto3" json:"commentText,omitempty"`
	Stars                uint64   `protobuf:"varint,5,opt,name=stars,proto3" json:"stars,omitempty"`
	Fans                 string   `protobuf:"bytes,6,opt,name=fans,proto3" json:"fans,omitempty"`
	AddTime              string   `protobuf:"bytes,7,opt,name=addTime,proto3" json:"addTime,omitempty"`
	LastModifyTime       string   `protobuf:"bytes,8,opt,name=lastModifyTime,proto3" json:"lastModifyTime,omitempty"`
	Chn                  string   `protobuf:"bytes,9,opt,name=chn,proto3" json:"chn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentInfo) Reset()         { *m = CommentInfo{} }
func (m *CommentInfo) String() string { return proto.CompactTextString(m) }
func (*CommentInfo) ProtoMessage()    {}
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{2}
}

func (m *CommentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentInfo.Unmarshal(m, b)
}
func (m *CommentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentInfo.Marshal(b, m, deterministic)
}
func (m *CommentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentInfo.Merge(m, src)
}
func (m *CommentInfo) XXX_Size() int {
	return xxx_messageInfo_CommentInfo.Size(m)
}
func (m *CommentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommentInfo proto.InternalMessageInfo

func (m *CommentInfo) GetCommentId() uint64 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

func (m *CommentInfo) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CommentInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CommentInfo) GetCommentText() string {
	if m != nil {
		return m.CommentText
	}
	return ""
}

func (m *CommentInfo) GetStars() uint64 {
	if m != nil {
		return m.Stars
	}
	return 0
}

func (m *CommentInfo) GetFans() string {
	if m != nil {
		return m.Fans
	}
	return ""
}

func (m *CommentInfo) GetAddTime() string {
	if m != nil {
		return m.AddTime
	}
	return ""
}

func (m *CommentInfo) GetLastModifyTime() string {
	if m != nil {
		return m.LastModifyTime
	}
	return ""
}

func (m *CommentInfo) GetChn() string {
	if m != nil {
		return m.Chn
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "terry.proto.comment.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "terry.proto.comment.ResponseHeader")
	proto.RegisterType((*CommentInfo)(nil), "terry.proto.comment.CommentInfo")
}

func init() { proto.RegisterFile("Common.proto", fileDescriptor_ee72d9a89737215c) }

var fileDescriptor_ee72d9a89737215c = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4f, 0x3a, 0x31,
	0x10, 0xc5, 0xc3, 0x9f, 0x65, 0x61, 0x87, 0x3f, 0xc4, 0x8c, 0xc6, 0x34, 0xc6, 0x03, 0xee, 0xc1,
	0x78, 0xe2, 0xe2, 0x47, 0x20, 0x31, 0xee, 0x01, 0x0e, 0x0d, 0x5f, 0xa0, 0xd2, 0x59, 0x25, 0xb1,
	0x1d, 0x6c, 0x4b, 0x22, 0x9f, 0x5e, 0xb3, 0xb3, 0xbb, 0x68, 0xbc, 0xbd, 0xdf, 0xeb, 0xeb, 0x9b,
	0x76, 0xe0, 0xff, 0x8a, 0x9d, 0x63, 0xbf, 0x3c, 0x04, 0x4e, 0x8c, 0x97, 0x89, 0x42, 0x38, 0xb5,
	0xb0, 0xdc, 0xb1, 0x73, 0xe4, 0x53, 0x79, 0x07, 0x33, 0x4d, 0x1f, 0x47, 0x8a, 0xe9, 0x99, 0x8c,
	0xa5, 0x80, 0x17, 0x30, 0xdc, 0xbd, 0x79, 0x35, 0x58, 0x0c, 0x1e, 0x0a, 0xdd, 0xc8, 0xf2, 0x09,
	0xe6, 0x9a, 0xe2, 0x81, 0x7d, 0xa4, 0x2e, 0xa3, 0x60, 0x4c, 0x21, 0x70, 0xd8, 0xb0, 0xe4, 0x66,
	0xba, 0x47, 0xbc, 0x81, 0x89, 0xc8, 0x75, 0x7c, 0x55, 0xff, 0xa4, 0xe2, 0xcc, 0xe5, 0xd7, 0x00,
	0xa6, 0xab, 0x76, 0x6c, 0xe5, 0x6b, 0xc6, 0x5b, 0x28, 0xba, 0x57, 0x54, 0x56, 0x7a, 0x32, 0xfd,
	0x63, 0xe0, 0x35, 0xe4, 0xc7, 0x48, 0xa1, 0xb2, 0xd2, 0x93, 0xe9, 0x8e, 0x9a, 0x09, 0x8d, 0xda,
	0x18, 0x47, 0x6a, 0xd8, 0x4e, 0xe8, 0x19, 0x17, 0x30, 0xed, 0x0a, 0xb6, 0xf4, 0x99, 0x54, 0x26,
	0xc7, 0xbf, 0x2d, 0xbc, 0x82, 0x51, 0x4c, 0x26, 0x44, 0x35, 0x92, 0xd2, 0x16, 0x10, 0x21, 0xab,
	0x8d, 0x8f, 0x2a, 0x97, 0x0b, 0xa2, 0x9b, 0x3f, 0x1a, 0x6b, 0xb7, 0x7b, 0x47, 0x6a, 0x2c, 0x76,
	0x8f, 0x78, 0x0f, 0xf3, 0x77, 0x13, 0xd3, 0x9a, 0xed, 0xbe, 0x3e, 0x49, 0x60, 0x22, 0x81, 0x3f,
	0x6e, 0xbf, 0xc9, 0xe2, 0xbc, 0xc9, 0x97, 0x5c, 0x76, 0xff, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x88, 0xce, 0xd5, 0x9c, 0x98, 0x01, 0x00, 0x00,
}
