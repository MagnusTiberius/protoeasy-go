// Code generated by protoc-gen-gogo.
// source: google/logging/type/http_request.proto
// DO NOT EDIT!

package google_logging_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_protobuf1 "go.pedge.io/pb/gogo/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A common proto for logging HTTP requests. Only contains semantics
// defined by the HTTP specification. Product-specific logging
// information MUST be defined in a separate message.
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl,proto3" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize,proto3" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// The IP address (IPv4 or IPv6) of the origin server that the request was
	// sent to.
	ServerIp string `protobuf:"bytes,13,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer,proto3" json:"referer,omitempty"`
	// The request processing latency on the server, from the time the request was
	// received until the response was sent.
	Latency *google_protobuf1.Duration `protobuf:"bytes,14,opt,name=latency" json:"latency,omitempty"`
	// Whether or not a cache lookup was attempted.
	CacheLookup bool `protobuf:"varint,11,opt,name=cache_lookup,json=cacheLookup,proto3" json:"cache_lookup,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit,proto3" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	CacheValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=cache_validated_with_origin_server,json=cacheValidatedWithOriginServer,proto3" json:"cache_validated_with_origin_server,omitempty"`
	// The number of HTTP response bytes inserted into cache. Set only when a
	// cache fill was attempted.
	CacheFillBytes int64 `protobuf:"varint,12,opt,name=cache_fill_bytes,json=cacheFillBytes,proto3" json:"cache_fill_bytes,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptorHttpRequest, []int{0} }

func (m *HttpRequest) GetLatency() *google_protobuf1.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}

func init() { proto.RegisterFile("google/logging/type/http_request.proto", fileDescriptorHttpRequest) }

var fileDescriptorHttpRequest = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdb, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0x1e, 0xf6, 0xe0, 0x3d, 0xa8, 0x32, 0x12, 0xb8, 0x05, 0xca, 0x52, 0x04, 0xca,
	0x55, 0x22, 0xd1, 0x27, 0x60, 0x41, 0xd0, 0x45, 0x45, 0x54, 0x29, 0x87, 0xcb, 0x28, 0xbb, 0x99,
	0x4d, 0x2c, 0xbc, 0xb1, 0xb1, 0x27, 0x45, 0xdb, 0xd7, 0xe0, 0x2d, 0x78, 0x4a, 0x94, 0xb1, 0x23,
	0x71, 0xc1, 0xa5, 0xbf, 0xff, 0xfb, 0x77, 0xbc, 0x13, 0xb3, 0x57, 0x95, 0xd6, 0x95, 0x82, 0x54,
	0xe9, 0xaa, 0x92, 0x4d, 0x95, 0xe2, 0xde, 0x40, 0x5a, 0x23, 0x9a, 0xdc, 0xc2, 0xcf, 0x16, 0x1c,
	0x26, 0xc6, 0x6a, 0xd4, 0xfc, 0x81, 0xf7, 0x92, 0xe0, 0x25, 0x9d, 0x77, 0xf6, 0x24, 0x94, 0x0b,
	0x23, 0xd3, 0xa2, 0x69, 0x34, 0x16, 0x28, 0x75, 0xe3, 0x7c, 0xe5, 0xec, 0x3c, 0xa4, 0x74, 0x5a,
	0xb7, 0xdb, 0xb4, 0x6c, 0x2d, 0x09, 0x3e, 0xbf, 0xf8, 0x7d, 0xc4, 0x26, 0x57, 0x88, 0x26, 0xf3,
	0x83, 0xf8, 0x4b, 0x36, 0x0f, 0x33, 0xf3, 0x1d, 0x60, 0xad, 0x4b, 0x11, 0x2d, 0xa2, 0x78, 0x9c,
	0xcd, 0x02, 0xfd, 0x44, 0x90, 0x3f, 0x63, 0x93, 0x5e, 0x6b, 0xad, 0x12, 0x07, 0xe4, 0xb0, 0x80,
	0xbe, 0x5a, 0xc5, 0x9f, 0xb3, 0x69, 0x2f, 0x38, 0x79, 0x0f, 0xe2, 0x70, 0x11, 0xc5, 0x87, 0x59,
	0x5f, 0xba, 0x95, 0xf7, 0xc0, 0x1f, 0xb2, 0x81, 0xc3, 0x02, 0x5b, 0x27, 0x8e, 0x16, 0x51, 0x7c,
	0x9c, 0x85, 0x13, 0x7f, 0xc1, 0x66, 0x16, 0x9c, 0xd1, 0x8d, 0x03, 0xdf, 0x3d, 0xa6, 0xee, 0xb4,
	0x87, 0x54, 0x7e, 0xca, 0x58, 0xeb, 0xc0, 0xe6, 0x45, 0x05, 0x0d, 0x8a, 0x01, 0xcd, 0x1f, 0x77,
	0xe4, 0x4d, 0x07, 0xf8, 0x63, 0x36, 0xb6, 0xb0, 0xd3, 0x08, 0xb9, 0x34, 0x62, 0x48, 0xe9, 0xc8,
	0x83, 0x95, 0xe9, 0x42, 0x07, 0xf6, 0x0e, 0x6c, 0x17, 0xce, 0x7c, 0xe8, 0xc1, 0xca, 0x70, 0xc1,
	0x86, 0x16, 0xb6, 0x60, 0xc1, 0x8a, 0x11, 0x45, 0xfd, 0x91, 0x5f, 0xb2, 0xa1, 0x2a, 0x10, 0x9a,
	0xcd, 0x5e, 0xcc, 0x17, 0x51, 0x3c, 0x79, 0x7d, 0x9a, 0x84, 0xef, 0xd1, 0x2f, 0x37, 0x79, 0x17,
	0x96, 0x9b, 0xf5, 0x66, 0xb7, 0x87, 0x4d, 0xb1, 0xa9, 0x21, 0x57, 0x5a, 0xff, 0x68, 0x8d, 0x98,
	0x2c, 0xa2, 0x78, 0x94, 0x4d, 0x88, 0x5d, 0x13, 0xea, 0xae, 0xe3, 0x95, 0x5a, 0xa2, 0x18, 0x53,
	0x3e, 0x22, 0x70, 0x25, 0x91, 0x7f, 0x64, 0x17, 0x3e, 0xbc, 0x2b, 0x94, 0x2c, 0x0b, 0x84, 0x32,
	0xff, 0x25, 0xb1, 0xce, 0xb5, 0x95, 0x95, 0x6c, 0x72, 0x7f, 0x6d, 0xc1, 0xa8, 0x75, 0x4e, 0xe6,
	0xb7, 0x5e, 0xfc, 0x2e, 0xb1, 0xfe, 0x4c, 0xda, 0x2d, 0x59, 0x3c, 0x66, 0x27, 0xfe, 0xb7, 0xb6,
	0x52, 0xa9, 0x7c, 0xbd, 0x47, 0x70, 0x62, 0x4a, 0xbb, 0x9d, 0x13, 0x7f, 0x2f, 0x95, 0x5a, 0x76,
	0x74, 0xb9, 0x62, 0x8f, 0x36, 0x7a, 0x97, 0xfc, 0xe7, 0xb9, 0x2d, 0x4f, 0xfe, 0x79, 0x2d, 0x37,
	0xdd, 0xff, 0xbe, 0x89, 0xfe, 0x1c, 0x9c, 0x7e, 0xf0, 0xe6, 0x5b, 0xa5, 0xdb, 0x32, 0xb9, 0x0e,
	0xfe, 0x97, 0xbd, 0x81, 0xf5, 0x80, 0x96, 0x73, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xf7,
	0x98, 0xab, 0xe4, 0x02, 0x00, 0x00,
}
