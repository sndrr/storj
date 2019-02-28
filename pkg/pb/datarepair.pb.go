// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: datarepair.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// InjuredSegment is the queue item used for the data repair queue
type InjuredSegment struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LostPieces           []int32  `protobuf:"varint,2,rep,packed,name=lost_pieces,json=lostPieces" json:"lost_pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjuredSegment) Reset()         { *m = InjuredSegment{} }
func (m *InjuredSegment) String() string { return proto.CompactTextString(m) }
func (*InjuredSegment) ProtoMessage()    {}
func (*InjuredSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_datarepair_6f9ccd8b18c20e5e, []int{0}
}
func (m *InjuredSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjuredSegment.Unmarshal(m, b)
}
func (m *InjuredSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjuredSegment.Marshal(b, m, deterministic)
}
func (dst *InjuredSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjuredSegment.Merge(dst, src)
}
func (m *InjuredSegment) XXX_Size() int {
	return xxx_messageInfo_InjuredSegment.Size(m)
}
func (m *InjuredSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_InjuredSegment.DiscardUnknown(m)
}

var xxx_messageInfo_InjuredSegment proto.InternalMessageInfo

func (m *InjuredSegment) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *InjuredSegment) GetLostPieces() []int32 {
	if m != nil {
		return m.LostPieces
	}
	return nil
}

func init() {
	proto.RegisterType((*InjuredSegment)(nil), "storjv3_0_0.InjuredSegment")
}

func init() { proto.RegisterFile("datarepair.proto", fileDescriptor_datarepair_6f9ccd8b18c20e5e) }

var fileDescriptor_datarepair_6f9ccd8b18c20e5e = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2c, 0x49,
	0x2c, 0x4a, 0x2d, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x2e,
	0xc9, 0x2f, 0xca, 0x2a, 0x33, 0x8e, 0x37, 0x88, 0x37, 0x50, 0x72, 0xe5, 0xe2, 0xf3, 0xcc, 0xcb,
	0x2a, 0x2d, 0x4a, 0x4d, 0x09, 0x4e, 0x4d, 0xcf, 0x4d, 0xcd, 0x2b, 0x11, 0x12, 0xe2, 0x62, 0x29,
	0x48, 0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xe4, 0xb9, 0xb8,
	0x73, 0xf2, 0x8b, 0x4b, 0xe2, 0x0b, 0x32, 0x53, 0x93, 0x53, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x35,
	0x58, 0x83, 0xb8, 0x40, 0x42, 0x01, 0x60, 0x11, 0x27, 0x96, 0x28, 0xa6, 0x82, 0xa4, 0x24, 0x36,
	0xb0, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6b, 0x22, 0x6f, 0x74, 0x00, 0x00,
	0x00,
}
