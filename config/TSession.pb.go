// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TSession.proto

package config

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

type TSessionSessionType int32

const (
	TSession_CONNECTOR TSessionSessionType = 0
	TSession_LOGIC     TSessionSessionType = 1
	TSession_CMD       TSessionSessionType = 2
)

var TSessionSessionType_name = map[int32]string{
	0: "CONNECTOR",
	1: "LOGIC",
	2: "CMD",
}

var TSessionSessionType_value = map[string]int32{
	"CONNECTOR": 0,
	"LOGIC":     1,
	"CMD":       2,
}

func (x TSessionSessionType) String() string {
	return proto.EnumName(TSessionSessionType_name, int32(x))
}

func (TSessionSessionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb2b1dadcb3d5c27, []int{0, 0}
}

type TSession struct {
	Cid                  string              `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	SrcSubRouter         string              `protobuf:"bytes,2,opt,name=srcSubRouter,proto3" json:"srcSubRouter,omitempty"`
	DstSubRouter         string              `protobuf:"bytes,3,opt,name=dstSubRouter,proto3" json:"dstSubRouter,omitempty"`
	Body                 []byte              `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	St                   TSessionSessionType `protobuf:"varint,5,opt,name=st,proto3,enum=config.TSessionSessionType" json:"st,omitempty"`
	ReplyToken           string              `protobuf:"bytes,6,opt,name=replyToken,proto3" json:"replyToken,omitempty"`
	Router               string              `protobuf:"bytes,7,opt,name=router,proto3" json:"router,omitempty"`
	LogicBindId          string              `protobuf:"bytes,8,opt,name=logicBindId,proto3" json:"logicBindId,omitempty"`
	RpcRespCode          int32               `protobuf:"varint,9,opt,name=rpcRespCode,proto3" json:"rpcRespCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TSession) Reset()         { *m = TSession{} }
func (m *TSession) String() string { return proto.CompactTextString(m) }
func (*TSession) ProtoMessage()    {}
func (*TSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb2b1dadcb3d5c27, []int{0}
}

func (m *TSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSession.Unmarshal(m, b)
}
func (m *TSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSession.Marshal(b, m, deterministic)
}
func (m *TSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSession.Merge(m, src)
}
func (m *TSession) XXX_Size() int {
	return xxx_messageInfo_TSession.Size(m)
}
func (m *TSession) XXX_DiscardUnknown() {
	xxx_messageInfo_TSession.DiscardUnknown(m)
}

var xxx_messageInfo_TSession proto.InternalMessageInfo

func (m *TSession) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *TSession) GetSrcSubRouter() string {
	if m != nil {
		return m.SrcSubRouter
	}
	return ""
}

func (m *TSession) GetDstSubRouter() string {
	if m != nil {
		return m.DstSubRouter
	}
	return ""
}

func (m *TSession) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *TSession) GetSt() TSessionSessionType {
	if m != nil {
		return m.St
	}
	return TSession_CONNECTOR
}

func (m *TSession) GetReplyToken() string {
	if m != nil {
		return m.ReplyToken
	}
	return ""
}

func (m *TSession) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *TSession) GetLogicBindId() string {
	if m != nil {
		return m.LogicBindId
	}
	return ""
}

func (m *TSession) GetRpcRespCode() int32 {
	if m != nil {
		return m.RpcRespCode
	}
	return 0
}

type CmdMsg struct {
	RunRoutineNum        int64    `protobuf:"varint,1,opt,name=runRoutineNum,proto3" json:"runRoutineNum,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdMsg) Reset()         { *m = CmdMsg{} }
func (m *CmdMsg) String() string { return proto.CompactTextString(m) }
func (*CmdMsg) ProtoMessage()    {}
func (*CmdMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb2b1dadcb3d5c27, []int{1}
}

func (m *CmdMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdMsg.Unmarshal(m, b)
}
func (m *CmdMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdMsg.Marshal(b, m, deterministic)
}
func (m *CmdMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdMsg.Merge(m, src)
}
func (m *CmdMsg) XXX_Size() int {
	return xxx_messageInfo_CmdMsg.Size(m)
}
func (m *CmdMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CmdMsg proto.InternalMessageInfo

func (m *CmdMsg) GetRunRoutineNum() int64 {
	if m != nil {
		return m.RunRoutineNum
	}
	return 0
}

func (m *CmdMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("config.TSessionSessionType", TSessionSessionType_name, TSessionSessionType_value)
	proto.RegisterType((*TSession)(nil), "config.TSession")
	proto.RegisterType((*CmdMsg)(nil), "config.CmdMsg")
}

func init() { proto.RegisterFile("TSession.proto", fileDescriptor_cb2b1dadcb3d5c27) }

var fileDescriptor_cb2b1dadcb3d5c27 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcb, 0x6a, 0xf3, 0x30,
	0x14, 0x84, 0x7f, 0xdb, 0x89, 0x13, 0x9f, 0x5c, 0x30, 0x67, 0xf1, 0xa3, 0x45, 0x29, 0xc6, 0x74,
	0xe1, 0x45, 0x09, 0xa5, 0x7d, 0x83, 0xa8, 0xa5, 0x04, 0x9a, 0x04, 0x14, 0xbf, 0x40, 0x62, 0xa9,
	0x41, 0x34, 0x91, 0x8c, 0x24, 0x2f, 0xf2, 0x9e, 0x7d, 0xa0, 0x62, 0xb9, 0x17, 0x67, 0xe5, 0xf1,
	0x37, 0x03, 0x67, 0x18, 0xc1, 0xbc, 0xdc, 0x09, 0x6b, 0xa5, 0x56, 0x8b, 0xda, 0x68, 0xa7, 0x31,
	0xae, 0xb4, 0x7a, 0x97, 0xc7, 0xfc, 0x33, 0x84, 0xf1, 0x8f, 0x85, 0x29, 0x44, 0x95, 0xe4, 0x24,
	0xc8, 0x82, 0x22, 0x61, 0xad, 0xc4, 0x1c, 0xa6, 0xd6, 0x54, 0xbb, 0xe6, 0xc0, 0x74, 0xe3, 0x84,
	0x21, 0xa1, 0xb7, 0xae, 0x58, 0x9b, 0xe1, 0xd6, 0xfd, 0x65, 0xa2, 0x2e, 0xd3, 0x67, 0x88, 0x30,
	0x38, 0x68, 0x7e, 0x21, 0x83, 0x2c, 0x28, 0xa6, 0xcc, 0x6b, 0xbc, 0x87, 0xd0, 0x3a, 0x32, 0xcc,
	0x82, 0x62, 0xfe, 0x78, 0xb3, 0xe8, 0xfa, 0x2c, 0x7e, 0x6b, 0xda, 0xee, 0x5b, 0x5e, 0x6a, 0xc1,
	0x42, 0xeb, 0xf0, 0x16, 0xc0, 0x88, 0xfa, 0x74, 0x29, 0xf5, 0x87, 0x50, 0x24, 0xf6, 0x37, 0x7a,
	0x04, 0xff, 0x43, 0x6c, 0xba, 0xfb, 0x23, 0xef, 0x7d, 0xff, 0x61, 0x06, 0x93, 0x93, 0x3e, 0xca,
	0x6a, 0x29, 0x15, 0x5f, 0x71, 0x32, 0xf6, 0x66, 0x1f, 0xb5, 0x09, 0x53, 0x57, 0x4c, 0xd8, 0x9a,
	0x6a, 0x2e, 0x48, 0x92, 0x05, 0xc5, 0x90, 0xf5, 0x51, 0xfe, 0x00, 0x93, 0x5e, 0x1d, 0x9c, 0x41,
	0x42, 0xb7, 0x9b, 0xcd, 0x0b, 0x2d, 0xb7, 0x2c, 0xfd, 0x87, 0x09, 0x0c, 0xdf, 0xb6, 0xaf, 0x2b,
	0x9a, 0x06, 0x38, 0x82, 0x88, 0xae, 0x9f, 0xd3, 0x30, 0x5f, 0x42, 0x4c, 0xcf, 0x7c, 0x6d, 0x8f,
	0x78, 0x07, 0x33, 0xd3, 0xa8, 0x76, 0x06, 0xa9, 0xc4, 0xa6, 0x39, 0xfb, 0x75, 0x23, 0x76, 0x0d,
	0xdb, 0x7d, 0xf8, 0xde, 0xed, 0xfd, 0xbe, 0x53, 0xe6, 0xf5, 0x21, 0xf6, 0x2f, 0xf5, 0xf4, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x5f, 0xc8, 0x6a, 0x4e, 0xbb, 0x01, 0x00, 0x00,
}
