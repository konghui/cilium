// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/regex.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A regex matcher designed for safety when used with untrusted input.
type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	// The regex match string. The string must be supported by the configured engine.
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{0}
}

func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

// Google's `RE2 <https://github.com/google/re2>`_ regex engine. The regex string must adhere to
// the documented `syntax <https://github.com/google/re2/wiki/Syntax>`_. The engine is designed
// to complete execution in linear time as well as limit the amount of memory used.
type RegexMatcher_GoogleRE2 struct {
	// This field controls the RE2 "program size" which is a rough estimate of how complex a
	// compiled regex is to evaluate. A regex that has a program size greater than the configured
	// value will fail to compile. In this case, the configured max program size can be increased
	// or the regex can be simplified. If not specified, the default is 100.
	MaxProgramSize       *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_program_size,json=maxProgramSize,proto3" json:"max_program_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{0, 0}
}

func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

func (m *RegexMatcher_GoogleRE2) GetMaxProgramSize() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxProgramSize
	}
	return nil
}

// Describes how to match a string and then produce a new string using a regular
// expression and a substitution string.
type RegexMatchAndSubstitute struct {
	// The regular expression used to find portions of a string (hereafter called
	// the "subject string") that should be replaced. When a new string is
	// produced during the substitution operation, the new string is initially
	// the same as the subject string, but then all matches in the subject string
	// are replaced by the substitution string. If replacing all matches isn't
	// desired, regular expression anchors can be used to ensure a single match,
	// so as to replace just one occurrence of a pattern. Capture groups can be
	// used in the pattern to extract portions of the subject string, and then
	// referenced in the substitution string.
	Pattern *RegexMatcher `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// The string that should be substituted into matching portions of the
	// subject string during a substitution operation to produce a new string.
	// Capture groups in the pattern can be referenced in the substitution
	// string. Note, however, that the syntax for referring to capture groups is
	// defined by the chosen regular expression engine. Google's `RE2
	// <https://github.com/google/re2>`_ regular expression engine uses a
	// backslash followed by the capture group number to denote a numbered
	// capture group. E.g., ``\1`` refers to capture group 1, and ``\2`` refers
	// to capture group 2.
	Substitution         string   `protobuf:"bytes,2,opt,name=substitution,proto3" json:"substitution,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatchAndSubstitute) Reset()         { *m = RegexMatchAndSubstitute{} }
func (m *RegexMatchAndSubstitute) String() string { return proto.CompactTextString(m) }
func (*RegexMatchAndSubstitute) ProtoMessage()    {}
func (*RegexMatchAndSubstitute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{1}
}

func (m *RegexMatchAndSubstitute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatchAndSubstitute.Unmarshal(m, b)
}
func (m *RegexMatchAndSubstitute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatchAndSubstitute.Marshal(b, m, deterministic)
}
func (m *RegexMatchAndSubstitute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatchAndSubstitute.Merge(m, src)
}
func (m *RegexMatchAndSubstitute) XXX_Size() int {
	return xxx_messageInfo_RegexMatchAndSubstitute.Size(m)
}
func (m *RegexMatchAndSubstitute) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatchAndSubstitute.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatchAndSubstitute proto.InternalMessageInfo

func (m *RegexMatchAndSubstitute) GetPattern() *RegexMatcher {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *RegexMatchAndSubstitute) GetSubstitution() string {
	if m != nil {
		return m.Substitution
	}
	return ""
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.v3.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.v3.RegexMatcher.GoogleRE2")
	proto.RegisterType((*RegexMatchAndSubstitute)(nil), "envoy.type.matcher.v3.RegexMatchAndSubstitute")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/regex.proto", fileDescriptor_7de4f215ebc85482) }

var fileDescriptor_7de4f215ebc85482 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0xcd, 0xe8, 0xee, 0xda, 0x77, 0x17, 0x29, 0x03, 0xb2, 0x4b, 0x71, 0xd7, 0xb6, 0x0b,
	0xb2, 0xfe, 0x4b, 0x74, 0x7a, 0xab, 0x78, 0x30, 0xe0, 0xbf, 0x83, 0x50, 0xb2, 0x28, 0xde, 0x4a,
	0x6a, 0x5f, 0xc7, 0x40, 0x9b, 0x84, 0x4c, 0x66, 0x6c, 0xf7, 0xe4, 0x45, 0x10, 0x3f, 0x82, 0x5f,
	0xc2, 0xbb, 0x77, 0x41, 0xf0, 0xe4, 0xb7, 0x91, 0x3d, 0xc9, 0x4c, 0x3a, 0x6a, 0x71, 0x50, 0x6f,
	0xe1, 0x7d, 0x9f, 0xe7, 0xe1, 0xf7, 0x3e, 0x81, 0x1e, 0xea, 0xc2, 0x2c, 0x99, 0x5f, 0x5a, 0x64,
	0x73, 0xe9, 0x5f, 0xbc, 0x42, 0xc7, 0x8a, 0x01, 0x73, 0x98, 0xe2, 0x82, 0x5a, 0x67, 0xbc, 0x89,
	0x2f, 0x56, 0x12, 0x5a, 0x4a, 0xe8, 0x4a, 0x42, 0x8b, 0x41, 0xe7, 0x20, 0x35, 0x26, 0x9d, 0x21,
	0xab, 0x44, 0x93, 0xfc, 0x25, 0x7b, 0xed, 0xa4, 0xb5, 0xe8, 0xb2, 0x60, 0xeb, 0xec, 0xe7, 0x53,
	0x2b, 0x99, 0xd4, 0xda, 0x78, 0xe9, 0x95, 0xd1, 0x19, 0xcb, 0xbc, 0xf4, 0x79, 0xbd, 0xee, 0xfd,
	0xb1, 0x2e, 0xd0, 0x65, 0xca, 0x68, 0xa5, 0xd3, 0x95, 0x64, 0xb7, 0x90, 0x33, 0x35, 0x95, 0x1e,
	0x59, 0xfd, 0x08, 0x8b, 0xfe, 0xd7, 0x08, 0x76, 0x44, 0x49, 0xf8, 0x24, 0xe0, 0xc4, 0xcf, 0x01,
	0x02, 0xcd, 0xd8, 0x61, 0xb2, 0x47, 0xba, 0xe4, 0x68, 0x3b, 0xb9, 0x49, 0x1b, 0xb9, 0xe9, 0xef,
	0x46, 0xfa, 0xb0, 0x72, 0x89, 0xfb, 0x09, 0x3f, 0x7f, 0xca, 0x37, 0xde, 0x93, 0xa8, 0x4d, 0x1e,
	0x9d, 0x11, 0xad, 0x10, 0x26, 0x30, 0x89, 0xf7, 0x61, 0xa3, 0xea, 0x62, 0x2f, 0xea, 0x92, 0xa3,
	0x16, 0xdf, 0x3a, 0xe5, 0xe7, 0x5c, 0xd4, 0x25, 0x22, 0x4c, 0x3b, 0x6f, 0x09, 0xb4, 0x7e, 0x66,
	0xc4, 0x0f, 0xa0, 0x3d, 0x97, 0x8b, 0xb1, 0x75, 0x26, 0x75, 0x72, 0x3e, 0xce, 0xd4, 0x09, 0xae,
	0x60, 0x2e, 0xd1, 0x10, 0x49, 0xeb, 0xb6, 0xe8, 0xd3, 0xc7, 0xda, 0x0f, 0x92, 0x67, 0x72, 0x96,
	0xa3, 0xb8, 0x30, 0x97, 0x8b, 0x51, 0x30, 0x1d, 0xab, 0x13, 0x1c, 0xde, 0xfa, 0xf0, 0xf9, 0xdd,
	0xc1, 0x75, 0xb8, 0xda, 0x70, 0x40, 0x33, 0xfd, 0xf0, 0x4a, 0xe9, 0xe8, 0xc1, 0xe5, 0x7f, 0x38,
	0x78, 0x0c, 0xdb, 0xa8, 0x53, 0xa5, 0x71, 0x5c, 0x6a, 0xe2, 0xb3, 0xdf, 0x39, 0xe9, 0x7f, 0x24,
	0xb0, 0xfb, 0x4b, 0x74, 0x4f, 0x4f, 0x8f, 0xf3, 0x49, 0xe6, 0x95, 0xcf, 0x3d, 0xc6, 0x77, 0x61,
	0xcb, 0x4a, 0xef, 0xd1, 0xe9, 0xd5, 0x21, 0x87, 0xff, 0xd1, 0xaa, 0xa8, 0x3d, 0x71, 0x1f, 0x76,
	0xb2, 0x3a, 0x4c, 0x19, 0x1d, 0x4a, 0x14, 0x6b, 0xb3, 0xe1, 0xed, 0x12, 0xfd, 0x06, 0x5c, 0xfb,
	0x2b, 0xfa, 0x1a, 0x15, 0xbf, 0xf3, 0xe9, 0xcd, 0x97, 0x6f, 0x9b, 0x51, 0x3b, 0x82, 0x43, 0x65,
	0x02, 0x90, 0x75, 0x66, 0xb1, 0x6c, 0x66, 0xe3, 0x50, 0xe5, 0x8c, 0xca, 0xe6, 0x47, 0x64, 0xb2,
	0x59, 0x7d, 0xc1, 0xe0, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x5c, 0x3f, 0x10, 0xfa, 0x02,
	0x00, 0x00,
}
