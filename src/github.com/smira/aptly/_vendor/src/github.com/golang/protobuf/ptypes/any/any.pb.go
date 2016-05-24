// Code generated by protoc-gen-go.
// source: github.com/golang/protobuf/ptypes/any/any.proto
// DO NOT EDIT!

/*
Package any is a generated protocol buffer package.

It is generated from these files:
	github.com/golang/protobuf/ptypes/any/any.proto

It has these top-level messages:
	Any
*/
package any

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// `Any` contains an arbitrary serialized message along with a URL
// that describes the type of the serialized message.
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
type Any struct {
	// A URL/resource name whose content describes the type of the
	// serialized message.
	//
	// For URLs which use the schema `http`, `https`, or no schema, the
	// following restrictions and interpretations apply:
	//
	// * If no schema is provided, `https` is assumed.
	// * The last segment of the URL's path must represent the fully
	//   qualified name of the type (as in `path/google.protobuf.Duration`).
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Schemas other than `http`, `https` (or the empty schema) might be
	// used with implementation specific semantics.
	//
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl" json:"type_url,omitempty"`
	// Must be valid serialized data of the above specified type.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Any) Reset()                    { *m = Any{} }
func (m *Any) String() string            { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()               {}
func (*Any) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (*Any) XXX_WellKnownType() string   { return "Any" }

func init() {
	proto.RegisterType((*Any)(nil), "google.protobuf.Any")
}

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xcc,
	0xab, 0x04, 0x61, 0x3d, 0xb0, 0xb8, 0x10, 0x7f, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x1e, 0x4c,
	0x95, 0x92, 0x19, 0x17, 0xb3, 0x63, 0x5e, 0xa5, 0x90, 0x24, 0x17, 0x07, 0x48, 0x79, 0x7c, 0x69,
	0x51, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x3b, 0x88, 0x1f, 0x5a, 0x94, 0x23, 0x24,
	0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x14, 0xe7, 0x09, 0x82, 0x70, 0x9c,
	0x8a, 0xb8, 0x84, 0x81, 0x96, 0xea, 0xa1, 0x19, 0xe7, 0xc4, 0x01, 0x34, 0x2c, 0x00, 0xc4, 0x09,
	0x60, 0x8c, 0x52, 0x25, 0xca, 0x71, 0x0b, 0x18, 0x19, 0x17, 0x31, 0x31, 0xbb, 0x07, 0x38, 0xad,
	0x62, 0x92, 0x73, 0x87, 0x98, 0x16, 0x00, 0x55, 0xa5, 0x17, 0x9e, 0x9a, 0x93, 0xe3, 0x9d, 0x97,
	0x5f, 0x9e, 0x17, 0x02, 0x52, 0x9d, 0xc4, 0x06, 0xd6, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xc6, 0x4d, 0x03, 0x23, 0xf6, 0x00, 0x00, 0x00,
}
