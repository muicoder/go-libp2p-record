// Code generated by protoc-gen-gogo.
// source: record.proto
// DO NOT EDIT!

/*
Package record_pb is a generated protocol buffer package.

It is generated from these files:
	record.proto

It has these top-level messages:
	Record
*/
package record_pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// Record represents a dht record that contains a value
// for a key value pair
type Record struct {
	// The key that references this record
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The actual value this record is storing
	Value []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	// hash of the authors public key
	Author *string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	// A PKI signature for the key+value+author
	Signature []byte `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	// Time the record was received, set by receiver
	TimeReceived     *string `protobuf:"bytes,5,opt,name=timeReceived" json:"timeReceived,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorRecord, []int{0} }

func (m *Record) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetAuthor() string {
	if m != nil && m.Author != nil {
		return *m.Author
	}
	return ""
}

func (m *Record) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Record) GetTimeReceived() string {
	if m != nil && m.TimeReceived != nil {
		return *m.TimeReceived
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "record.pb.Record")
}

var fileDescriptorRecord = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xf1, 0x92, 0x94, 0xba, 0x18,
	0xb9, 0xd8, 0x82, 0xc0, 0x3c, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09,
	0x28, 0xc6, 0x13, 0x04, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0x25, 0x96, 0x96, 0x64, 0xe4, 0x17, 0x49,
	0x30, 0x83, 0x95, 0x42, 0x79, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0x99, 0xe9, 0x79, 0x89, 0x25, 0xa5,
	0x45, 0xa9, 0x12, 0x2c, 0x60, 0x1d, 0x08, 0x01, 0x21, 0x25, 0x2e, 0x9e, 0x92, 0xcc, 0xdc, 0x54,
	0xa0, 0x5d, 0xa9, 0x99, 0x65, 0xa9, 0x29, 0x12, 0xac, 0x60, 0xbd, 0x28, 0x62, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x87, 0x59, 0xbc, 0x53, 0xa6, 0x00, 0x00, 0x00,
}