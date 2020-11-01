// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package arhatgopb

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CmdType int32

const (
	_INVALID_CMD                   CmdType = 0
	CMD_PERIPHERAL_CONNECT         CmdType = 11
	CMD_PERIPHERAL_OPERATE         CmdType = 12
	CMD_PERIPHERAL_COLLECT_METRICS CmdType = 13
	CMD_PERIPHERAL_CLOSE           CmdType = 14
)

var CmdType_name = map[int32]string{
	0:  "_INVALID_CMD",
	11: "CMD_PERIPHERAL_CONNECT",
	12: "CMD_PERIPHERAL_OPERATE",
	13: "CMD_PERIPHERAL_COLLECT_METRICS",
	14: "CMD_PERIPHERAL_CLOSE",
}

var CmdType_value = map[string]int32{
	"_INVALID_CMD":                   0,
	"CMD_PERIPHERAL_CONNECT":         11,
	"CMD_PERIPHERAL_OPERATE":         12,
	"CMD_PERIPHERAL_COLLECT_METRICS": 13,
	"CMD_PERIPHERAL_CLOSE":           14,
}

func (CmdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

type MsgType int32

const (
	_INVALID_MSG                    MsgType = 0
	MSG_ERROR                       MsgType = 1
	MSG_DONE                        MsgType = 2
	MSG_REGISTER                    MsgType = 3
	MSG_PERIPHERAL_OPERATION_RESULT MsgType = 11
	MSG_PERIPHERAL_METRICS          MsgType = 12
	MSG_PERIPHERAL_EVENTS           MsgType = 13
)

var MsgType_name = map[int32]string{
	0:  "_INVALID_MSG",
	1:  "MSG_ERROR",
	2:  "MSG_DONE",
	3:  "MSG_REGISTER",
	11: "MSG_PERIPHERAL_OPERATION_RESULT",
	12: "MSG_PERIPHERAL_METRICS",
	13: "MSG_PERIPHERAL_EVENTS",
}

var MsgType_value = map[string]int32{
	"_INVALID_MSG":                    0,
	"MSG_ERROR":                       1,
	"MSG_DONE":                        2,
	"MSG_REGISTER":                    3,
	"MSG_PERIPHERAL_OPERATION_RESULT": 11,
	"MSG_PERIPHERAL_METRICS":          12,
	"MSG_PERIPHERAL_EVENTS":           13,
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

type CodecType int32

const (
	_INVALID_CODEC CodecType = 0
	CODEC_PROTOBUF CodecType = 1
	CODEC_JSON     CodecType = 2
)

var CodecType_name = map[int32]string{
	0: "_INVALID_CODEC",
	1: "CODEC_PROTOBUF",
	2: "CODEC_JSON",
}

var CodecType_value = map[string]int32{
	"_INVALID_CODEC": 0,
	"CODEC_PROTOBUF": 1,
	"CODEC_JSON":     2,
}

func (CodecType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}

type ExtensionType int32

const (
	_INVALID_EXTENSION   ExtensionType = 0
	EXTENSION_PERIPHERAL ExtensionType = 1
)

var ExtensionType_name = map[int32]string{
	0: "_INVALID_EXTENSION",
	1: "EXTENSION_PERIPHERAL",
}

var ExtensionType_value = map[string]int32{
	"_INVALID_EXTENSION":   0,
	"EXTENSION_PERIPHERAL": 1,
}

func (ExtensionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{3}
}

type Cmd struct {
	Kind CmdType `protobuf:"varint,1,opt,name=kind,proto3,enum=arhat.CmdType" json:"kind,omitempty"`
	// id assigned by the arhat extension manager
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Seq     uint64 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Cmd) Reset()      { *m = Cmd{} }
func (*Cmd) ProtoMessage() {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}
func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(m, src)
}
func (m *Cmd) XXX_Size() int {
	return m.Size()
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetKind() CmdType {
	if m != nil {
		return m.Kind
	}
	return _INVALID_CMD
}

func (m *Cmd) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cmd) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Cmd) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Msg struct {
	Kind MsgType `protobuf:"varint,1,opt,name=kind,proto3,enum=arhat.MsgType" json:"kind,omitempty"`
	// id assigned by the arhat extension manager
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// (optional) if this is an answer for a Cmd, set ack to seq in Cmd
	Ack     uint64 `protobuf:"varint,3,opt,name=ack,proto3" json:"ack,omitempty"`
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Msg) Reset()      { *m = Msg{} }
func (*Msg) ProtoMessage() {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetKind() MsgType {
	if m != nil {
		return m.Kind
	}
	return _INVALID_MSG
}

func (m *Msg) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Msg) GetAck() uint64 {
	if m != nil {
		return m.Ack
	}
	return 0
}

func (m *Msg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DoneMsg struct {
}

func (m *DoneMsg) Reset()      { *m = DoneMsg{} }
func (*DoneMsg) ProtoMessage() {}
func (*DoneMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}
func (m *DoneMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DoneMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DoneMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DoneMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoneMsg.Merge(m, src)
}
func (m *DoneMsg) XXX_Size() int {
	return m.Size()
}
func (m *DoneMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DoneMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DoneMsg proto.InternalMessageInfo

type ErrorMsg struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ErrorMsg) Reset()      { *m = ErrorMsg{} }
func (*ErrorMsg) ProtoMessage() {}
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{3}
}
func (m *ErrorMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMsg.Merge(m, src)
}
func (m *ErrorMsg) XXX_Size() int {
	return m.Size()
}
func (m *ErrorMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMsg proto.InternalMessageInfo

func (m *ErrorMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// RegisterMsg is always sent in json codec, so the extension server
// can decoded it and determine real codec used in the stream
type RegisterMsg struct {
	Name          string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Codec         CodecType     `protobuf:"varint,2,opt,name=codec,proto3,enum=arhat.CodecType" json:"codec,omitempty"`
	ExtensionType ExtensionType `protobuf:"varint,3,opt,name=extension_type,json=extensionType,proto3,enum=arhat.ExtensionType" json:"extension_type,omitempty"`
}

func (m *RegisterMsg) Reset()      { *m = RegisterMsg{} }
func (*RegisterMsg) ProtoMessage() {}
func (*RegisterMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{4}
}
func (m *RegisterMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterMsg.Merge(m, src)
}
func (m *RegisterMsg) XXX_Size() int {
	return m.Size()
}
func (m *RegisterMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterMsg proto.InternalMessageInfo

func (m *RegisterMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterMsg) GetCodec() CodecType {
	if m != nil {
		return m.Codec
	}
	return _INVALID_CODEC
}

func (m *RegisterMsg) GetExtensionType() ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return _INVALID_EXTENSION
}

func init() {
	proto.RegisterEnum("arhat.CmdType", CmdType_name, CmdType_value)
	proto.RegisterEnum("arhat.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("arhat.CodecType", CodecType_name, CodecType_value)
	proto.RegisterEnum("arhat.ExtensionType", ExtensionType_name, ExtensionType_value)
	proto.RegisterType((*Cmd)(nil), "arhat.Cmd")
	proto.RegisterType((*Msg)(nil), "arhat.Msg")
	proto.RegisterType((*DoneMsg)(nil), "arhat.DoneMsg")
	proto.RegisterType((*ErrorMsg)(nil), "arhat.ErrorMsg")
	proto.RegisterType((*RegisterMsg)(nil), "arhat.RegisterMsg")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xf5, 0x26, 0x29, 0x69, 0x26, 0x89, 0xb5, 0x5a, 0x95, 0xca, 0x70, 0xd8, 0x46, 0x41, 0x42,
	0x55, 0x04, 0x45, 0x2a, 0x47, 0x4e, 0xa9, 0xbd, 0x14, 0x23, 0x7f, 0x44, 0x6b, 0xb7, 0x42, 0x5c,
	0x2c, 0x37, 0x5e, 0x05, 0xab, 0xd4, 0x36, 0xb6, 0x85, 0xe8, 0x8d, 0x13, 0x67, 0xf8, 0x13, 0x88,
	0x9f, 0xc2, 0xb1, 0xc7, 0x1e, 0xa9, 0x7b, 0xe1, 0xd8, 0x9f, 0x80, 0xbc, 0x49, 0x4a, 0x3f, 0xa0,
	0x97, 0x68, 0xde, 0x7b, 0x33, 0x79, 0xcf, 0x33, 0x36, 0x74, 0xb3, 0x3c, 0x2d, 0xd3, 0x2d, 0xf9,
	0x4b, 0x56, 0xc2, 0xfc, 0x5d, 0x58, 0x0e, 0x05, 0x34, 0xf5, 0xa3, 0x88, 0x0c, 0xa1, 0x75, 0x18,
	0x27, 0x91, 0x86, 0x06, 0x68, 0x53, 0xdd, 0x56, 0xb7, 0xa4, 0xb8, 0xa5, 0x1f, 0x45, 0xfe, 0x71,
	0x26, 0xb8, 0xd4, 0x88, 0x0a, 0x8d, 0x38, 0xd2, 0x1a, 0x03, 0xb4, 0xd9, 0xe2, 0x8d, 0x38, 0x22,
	0x18, 0x9a, 0x85, 0xf8, 0xa0, 0x35, 0x25, 0x51, 0x97, 0x44, 0x83, 0x76, 0x16, 0x1e, 0xbf, 0x4f,
	0xc3, 0x48, 0x6b, 0x0d, 0xd0, 0x66, 0x8f, 0x2f, 0x61, 0x6d, 0x63, 0x17, 0xb3, 0xff, 0xd8, 0xd8,
	0xc5, 0xec, 0x6e, 0x9b, 0x70, 0x7a, 0xb8, 0xb4, 0x09, 0xa7, 0x87, 0x77, 0xd8, 0x74, 0xa0, 0x6d,
	0xa4, 0x89, 0xb0, 0x8b, 0xd9, 0xf0, 0x09, 0xac, 0xb2, 0x3c, 0x4f, 0xf3, 0xda, 0x76, 0x00, 0xdd,
	0x48, 0x14, 0xd3, 0x3c, 0xce, 0xca, 0x38, 0x4d, 0xa4, 0x7b, 0x87, 0x5f, 0xa5, 0x86, 0x5f, 0x10,
	0x74, 0xb9, 0x98, 0xc5, 0x45, 0x29, 0xe4, 0x04, 0x81, 0x56, 0x12, 0x1e, 0x89, 0x45, 0xab, 0xac,
	0xc9, 0x63, 0x58, 0x99, 0xa6, 0x91, 0x98, 0xca, 0x6c, 0xea, 0x36, 0x5e, 0x2e, 0xa9, 0xe6, 0x64,
	0xfe, 0xb9, 0x4c, 0x5e, 0x80, 0x2a, 0x3e, 0x95, 0x22, 0x29, 0xe2, 0x34, 0x09, 0xca, 0xe3, 0x4c,
	0xc8, 0xec, 0xea, 0xf6, 0xda, 0x62, 0x80, 0x2d, 0x45, 0x39, 0xd4, 0x17, 0x57, 0xe1, 0xe8, 0x1b,
	0x82, 0xf6, 0x62, 0xed, 0x04, 0x43, 0x2f, 0x30, 0x9d, 0xfd, 0xb1, 0x65, 0x1a, 0x81, 0x6e, 0x1b,
	0x58, 0x21, 0x0f, 0x61, 0x5d, 0xb7, 0x8d, 0x60, 0xc2, 0xb8, 0x39, 0x79, 0xc5, 0xf8, 0xd8, 0x0a,
	0x74, 0xd7, 0x71, 0x98, 0xee, 0xe3, 0xee, 0x3f, 0x34, 0x77, 0xc2, 0xf8, 0xd8, 0x67, 0xb8, 0x47,
	0x86, 0x40, 0x6f, 0xcd, 0x59, 0x16, 0xd3, 0xfd, 0xc0, 0x66, 0x3e, 0x37, 0x75, 0x0f, 0xf7, 0x89,
	0x06, 0x6b, 0x37, 0x7b, 0x2c, 0xd7, 0x63, 0x58, 0x1d, 0x7d, 0x47, 0xd0, 0x5e, 0xdc, 0xe8, 0x5a,
	0x26, 0xdb, 0xdb, 0xc5, 0x0a, 0xe9, 0x43, 0xc7, 0xf6, 0x76, 0x03, 0xc6, 0xb9, 0xcb, 0x31, 0x22,
	0x3d, 0x58, 0xad, 0xa1, 0xe1, 0x3a, 0x0c, 0x37, 0xea, 0xf6, 0x1a, 0x71, 0xb6, 0x6b, 0x7a, 0x3e,
	0xe3, 0xb8, 0x49, 0x1e, 0xc1, 0x46, 0xcd, 0xdc, 0x8a, 0x69, 0xba, 0x4e, 0xc0, 0x99, 0xb7, 0x67,
	0x2d, 0x9e, 0xe5, 0x46, 0xd3, 0x32, 0x67, 0x8f, 0x3c, 0x80, 0xfb, 0x37, 0x34, 0xb6, 0xcf, 0x1c,
	0xdf, 0xc3, 0xfd, 0x91, 0x0e, 0x9d, 0xcb, 0x6b, 0x10, 0x02, 0xea, 0xdf, 0xed, 0xb9, 0x06, 0xd3,
	0xb1, 0x52, 0x73, 0xb2, 0x0c, 0x26, 0xdc, 0xf5, 0xdd, 0x9d, 0xbd, 0x97, 0x18, 0x11, 0x15, 0x60,
	0xce, 0xbd, 0xf6, 0x5c, 0x07, 0x37, 0x46, 0x63, 0xe8, 0x5f, 0xbb, 0x10, 0x59, 0x07, 0x72, 0xf9,
	0x47, 0xec, 0x8d, 0xcf, 0x1c, 0xcf, 0x74, 0x1d, 0xac, 0xd4, 0x0b, 0xbb, 0x84, 0x57, 0xe2, 0x60,
	0xb4, 0xb3, 0x77, 0x72, 0x46, 0x95, 0xd3, 0x33, 0xaa, 0x5c, 0x9c, 0x51, 0xf4, 0xb9, 0xa2, 0xe8,
	0x47, 0x45, 0xd1, 0xcf, 0x8a, 0xa2, 0x93, 0x8a, 0xa2, 0x5f, 0x15, 0x45, 0xbf, 0x2b, 0xaa, 0x5c,
	0x54, 0x14, 0x7d, 0x3d, 0xa7, 0xca, 0xc9, 0x39, 0x55, 0x4e, 0xcf, 0xa9, 0xf2, 0x76, 0x63, 0xfe,
	0x7a, 0x44, 0xe2, 0xe3, 0x33, 0x59, 0x3d, 0x95, 0x1f, 0xea, 0xbc, 0x9e, 0xa5, 0xd9, 0xc1, 0xc1,
	0x3d, 0x49, 0x3c, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x57, 0x05, 0xe7, 0x54, 0xc8, 0x03, 0x00,
	0x00,
}

func (x CmdType) String() string {
	s, ok := CmdType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x MsgType) String() string {
	s, ok := MsgType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CodecType) String() string {
	s, ok := CodecType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ExtensionType) String() string {
	s, ok := ExtensionType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Cmd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cmd)
	if !ok {
		that2, ok := that.(Cmd)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Seq != that1.Seq {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	return true
}
func (this *Msg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Msg)
	if !ok {
		that2, ok := that.(Msg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Ack != that1.Ack {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	return true
}
func (this *DoneMsg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DoneMsg)
	if !ok {
		that2, ok := that.(DoneMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *ErrorMsg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ErrorMsg)
	if !ok {
		that2, ok := that.(ErrorMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	return true
}
func (this *RegisterMsg) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisterMsg)
	if !ok {
		that2, ok := that.(RegisterMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Codec != that1.Codec {
		return false
	}
	if this.ExtensionType != that1.ExtensionType {
		return false
	}
	return true
}
func (this *Cmd) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&arhatgopb.Cmd{")
	s = append(s, "Kind: "+fmt.Sprintf("%#v", this.Kind)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Seq: "+fmt.Sprintf("%#v", this.Seq)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Msg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&arhatgopb.Msg{")
	s = append(s, "Kind: "+fmt.Sprintf("%#v", this.Kind)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Ack: "+fmt.Sprintf("%#v", this.Ack)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DoneMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&arhatgopb.DoneMsg{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ErrorMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&arhatgopb.ErrorMsg{")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RegisterMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&arhatgopb.RegisterMsg{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Codec: "+fmt.Sprintf("%#v", this.Codec)+",\n")
	s = append(s, "ExtensionType: "+fmt.Sprintf("%#v", this.ExtensionType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProto(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Cmd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cmd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cmd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if m.Seq != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Kind != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Msg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if m.Ack != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Ack))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Kind != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DoneMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DoneMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DoneMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ErrorMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtensionType != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.ExtensionType))
		i--
		dAtA[i] = 0x18
	}
	if m.Codec != 0 {
		i = encodeVarintProto(dAtA, i, uint64(m.Codec))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProto(dAtA []byte, offset int, v uint64) int {
	offset -= sovProto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cmd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovProto(uint64(m.Kind))
	}
	if m.Id != 0 {
		n += 1 + sovProto(uint64(m.Id))
	}
	if m.Seq != 0 {
		n += 1 + sovProto(uint64(m.Seq))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovProto(uint64(m.Kind))
	}
	if m.Id != 0 {
		n += 1 + sovProto(uint64(m.Id))
	}
	if m.Ack != 0 {
		n += 1 + sovProto(uint64(m.Ack))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *DoneMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ErrorMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *RegisterMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	if m.Codec != 0 {
		n += 1 + sovProto(uint64(m.Codec))
	}
	if m.ExtensionType != 0 {
		n += 1 + sovProto(uint64(m.ExtensionType))
	}
	return n
}

func sovProto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProto(x uint64) (n int) {
	return sovProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Cmd) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Cmd{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Seq:` + fmt.Sprintf("%v", this.Seq) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Msg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Msg{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Ack:` + fmt.Sprintf("%v", this.Ack) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DoneMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DoneMsg{`,
		`}`,
	}, "")
	return s
}
func (this *ErrorMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ErrorMsg{`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RegisterMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterMsg{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Codec:` + fmt.Sprintf("%v", this.Codec) + `,`,
		`ExtensionType:` + fmt.Sprintf("%v", this.ExtensionType) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProto(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Cmd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cmd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cmd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= CmdType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= MsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ack", wireType)
			}
			m.Ack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ack |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DoneMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DoneMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DoneMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrorMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			m.Codec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Codec |= CodecType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtensionType", wireType)
			}
			m.ExtensionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtensionType |= ExtensionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProto = fmt.Errorf("proto: unexpected end of group")
)
