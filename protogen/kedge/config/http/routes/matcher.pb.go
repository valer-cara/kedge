// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kedge/config/http/routes/matcher.proto

package kedge_config_http_routes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProxyMode int32

const (
	ProxyMode_ANY ProxyMode = 0
	// / Reverse Proxy is when the FE serves an authority (Host) publicly and clients connect to that authority
	// / directly. This is used to expose publicly DNS-resolvable names.
	ProxyMode_REVERSE_PROXY ProxyMode = 1
	// / Forward Proxy is when the FE serves as an HTTP_PROXY for a browser or an application. The resolution of the
	// / backend is done by the FE itself, so non-public names can be addressed.
	// / This may be from the 90s, but it still is very useful.
	// /
	// / IMPORTANT: If you have a PAC file configured in Firefox, the HTTPS rule behaves differently than in Chrome. The
	// / proxied requests are not FORWARD_PROXY requests but REVERSE_PROXY_REQUESTS.
	ProxyMode_FORWARD_PROXY ProxyMode = 2
)

var ProxyMode_name = map[int32]string{
	0: "ANY",
	1: "REVERSE_PROXY",
	2: "FORWARD_PROXY",
}
var ProxyMode_value = map[string]int32{
	"ANY":           0,
	"REVERSE_PROXY": 1,
	"FORWARD_PROXY": 2,
}

func (x ProxyMode) String() string {
	return proto.EnumName(ProxyMode_name, int32(x))
}
func (ProxyMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// / Matcher describes a rule that matches any HTTP route.
type Matcher struct {
	// / path_rules is a globbing expression that matches a URL path of the request.
	// / See: https://cloud.google.com/compute/docs/load-balancing/http/url-map
	// / If not present, '/*' is default.
	PathRules []string `protobuf:"bytes,1,rep,name=path_rules,json=pathRules" json:"path_rules,omitempty"`
	// / host matches on the ':authority' header (a.k.a. Host header) enabling Virtual Host-like proxying.
	// / The matching is done through lower-case string-equality.
	// / If none are present, the route skips ':authority' checks.
	// / Accepts regexp RE2 expression.
	Host string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	// / header matches any HTTP inbound request Headers.
	// / Eeach key provided must find a match for the route to match.
	// / The matching is done through lower-case key match and explicit string-equality of values.
	// / If none are present, the route skips metadata checks.
	Header map[string]string `protobuf:"bytes,4,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// / proxy_mode controlls what kind of inbound requests this route matches. See
	ProxyMode ProxyMode `protobuf:"varint,5,opt,name=proxy_mode,json=proxyMode,enum=kedge.config.http.routes.ProxyMode" json:"proxy_mode,omitempty"`
	// / Optional port matcher. If 0 route will ignore port.
	Port uint32 `protobuf:"varint,6,opt,name=port" json:"port,omitempty"`
}

func (m *Matcher) Reset()                    { *m = Matcher{} }
func (m *Matcher) String() string            { return proto.CompactTextString(m) }
func (*Matcher) ProtoMessage()               {}
func (*Matcher) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Matcher) GetPathRules() []string {
	if m != nil {
		return m.PathRules
	}
	return nil
}

func (m *Matcher) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Matcher) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Matcher) GetProxyMode() ProxyMode {
	if m != nil {
		return m.ProxyMode
	}
	return ProxyMode_ANY
}

func (m *Matcher) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Matcher)(nil), "kedge.config.http.routes.Matcher")
	proto.RegisterEnum("kedge.config.http.routes.ProxyMode", ProxyMode_name, ProxyMode_value)
}

func init() { proto.RegisterFile("kedge/config/http/routes/matcher.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x86, 0x6d, 0xbb, 0x0f, 0x7a, 0xc6, 0x64, 0x06, 0x2f, 0xc2, 0x40, 0x28, 0x0a, 0x52, 0x84,
	0x25, 0x30, 0x41, 0x74, 0x77, 0x13, 0x2b, 0xde, 0xcc, 0x8d, 0x08, 0xea, 0xae, 0x46, 0xb7, 0xc6,
	0xb6, 0x6c, 0x5d, 0x4a, 0x9a, 0x6e, 0xee, 0x8f, 0xf8, 0x7b, 0x25, 0xe9, 0x1c, 0xde, 0xec, 0xee,
	0xed, 0xc3, 0x73, 0xce, 0x7b, 0x1a, 0xb8, 0x5e, 0xf2, 0x28, 0xe6, 0x74, 0x21, 0xd6, 0x5f, 0x69,
	0x4c, 0x13, 0xa5, 0x72, 0x2a, 0x45, 0xa9, 0x78, 0x41, 0xb3, 0x50, 0x2d, 0x12, 0x2e, 0x49, 0x2e,
	0x85, 0x12, 0x08, 0x1b, 0x8f, 0x54, 0x1e, 0xd1, 0x1e, 0xa9, 0xbc, 0xee, 0x5d, 0x9c, 0xaa, 0xa4,
	0x9c, 0x93, 0x85, 0xc8, 0x68, 0xb6, 0x4d, 0xd5, 0x52, 0x6c, 0x69, 0x2c, 0x7a, 0x66, 0xac, 0xb7,
	0x09, 0x57, 0x69, 0x14, 0x2a, 0x21, 0x0b, 0x7a, 0x88, 0xd5, 0xc6, 0xcb, 0x1f, 0x1b, 0x9a, 0xa3,
	0xaa, 0x03, 0x5d, 0x00, 0xe4, 0xa1, 0x4a, 0x66, 0xb2, 0x5c, 0xf1, 0x02, 0x5b, 0x9e, 0xe3, 0xbb,
	0xcc, 0xd5, 0x84, 0x69, 0x80, 0x10, 0xd4, 0x12, 0x51, 0x28, 0xec, 0x78, 0x96, 0xef, 0x32, 0x93,
	0x51, 0x00, 0x8d, 0x84, 0x87, 0x11, 0x97, 0xb8, 0xe6, 0x39, 0x7e, 0xab, 0xdf, 0x23, 0xc7, 0x2e,
	0x24, 0xfb, 0x16, 0xf2, 0x62, 0xfc, 0x60, 0xad, 0xe4, 0x8e, 0xed, 0x87, 0xd1, 0x23, 0x40, 0x2e,
	0xc5, 0xf7, 0x6e, 0x96, 0x89, 0x88, 0xe3, 0xba, 0x67, 0xf9, 0xa7, 0xfd, 0xab, 0xe3, 0xab, 0x26,
	0xda, 0x1d, 0x89, 0x88, 0x33, 0x37, 0xff, 0x8b, 0xfa, 0xbc, 0x5c, 0x48, 0x85, 0x1b, 0x9e, 0xe5,
	0xb7, 0x99, 0xc9, 0xdd, 0x07, 0x68, 0xfd, 0xab, 0x43, 0x1d, 0x70, 0x96, 0x7c, 0x87, 0x2d, 0xf3,
	0x03, 0x3a, 0xa2, 0x73, 0xa8, 0x6f, 0xc2, 0x55, 0xc9, 0xb1, 0x6d, 0x58, 0xf5, 0x31, 0xb0, 0xef,
	0xad, 0x9b, 0x01, 0xb8, 0x87, 0x1a, 0xd4, 0x04, 0x67, 0xf8, 0x3a, 0xed, 0x9c, 0xa0, 0x33, 0x68,
	0xb3, 0xe0, 0x3d, 0x60, 0x6f, 0xc1, 0x6c, 0xc2, 0xc6, 0x9f, 0xd3, 0x8e, 0xa5, 0xd1, 0xf3, 0x98,
	0x7d, 0x0c, 0xd9, 0xd3, 0x1e, 0xd9, 0xf3, 0x86, 0x79, 0xdb, 0xdb, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xca, 0xa4, 0x88, 0x8d, 0xd7, 0x01, 0x00, 0x00,
}