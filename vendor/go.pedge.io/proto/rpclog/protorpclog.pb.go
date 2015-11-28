// Code generated by protoc-gen-go.
// source: rpclog/protorpclog.proto
// DO NOT EDIT!

package protorpclog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"

// discarding unused import protolog "go.pedge.io/protolog"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Call struct {
	Service  string                    `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Method   string                    `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	Request  string                    `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
	Response string                    `protobuf:"bytes,4,opt,name=response" json:"response,omitempty"`
	Error    string                    `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	Duration *google_protobuf.Duration `protobuf:"bytes,6,opt,name=duration" json:"duration,omitempty"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}

func (m *Call) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}
