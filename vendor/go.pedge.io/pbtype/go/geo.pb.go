// Code generated by protoc-gen-go.
// source: geo.proto
// DO NOT EDIT!

package pbtype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LatLng struct {
	LatPicos  int64 `protobuf:"varint,1,opt,name=lat_picos" json:"lat_picos,omitempty"`
	LongPicos int64 `protobuf:"varint,2,opt,name=long_picos" json:"long_picos,omitempty"`
}

func (m *LatLng) Reset()                    { *m = LatLng{} }
func (m *LatLng) String() string            { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()               {}
func (*LatLng) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*LatLng)(nil), "pb.type.LatLng")
}

var fileDescriptor3 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4f, 0xcd, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x48, 0xd2, 0x2b, 0xa9, 0x2c, 0x48, 0x55, 0xd2,
	0xe7, 0x62, 0xf3, 0x49, 0x2c, 0xf1, 0xc9, 0x4b, 0x17, 0x12, 0xe4, 0xe2, 0xcc, 0x49, 0x2c, 0x89,
	0x2f, 0xc8, 0x4c, 0xce, 0x2f, 0x96, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x16, 0x12, 0xe2, 0xe2, 0xca,
	0xc9, 0xcf, 0x4b, 0x87, 0x8a, 0x31, 0x81, 0xc4, 0x9c, 0x38, 0xa2, 0xd8, 0x0a, 0x92, 0x40, 0x5a,
	0x93, 0xd8, 0xc0, 0x46, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa3, 0xa6, 0xf7, 0x57,
	0x00, 0x00, 0x00,
}