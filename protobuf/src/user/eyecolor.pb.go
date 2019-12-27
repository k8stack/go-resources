// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eyecolor.proto

package userpb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EyeColor int32

const (
	EyeColor_UNKNOWN_EYE_COLOR EyeColor = 0
	EyeColor_BLUE_EYE_COLOR    EyeColor = 1
	EyeColor_GREEN_EYE_COLOR   EyeColor = 2
	EyeColor_BROWN_EYE_COLOR   EyeColor = 3
	EyeColor_BLACK_EYE_COLOR   EyeColor = 4
)

var EyeColor_name = map[int32]string{
	0: "UNKNOWN_EYE_COLOR",
	1: "BLUE_EYE_COLOR",
	2: "GREEN_EYE_COLOR",
	3: "BROWN_EYE_COLOR",
	4: "BLACK_EYE_COLOR",
}

var EyeColor_value = map[string]int32{
	"UNKNOWN_EYE_COLOR": 0,
	"BLUE_EYE_COLOR":    1,
	"GREEN_EYE_COLOR":   2,
	"BROWN_EYE_COLOR":   3,
	"BLACK_EYE_COLOR":   4,
}

func (x EyeColor) String() string {
	return proto.EnumName(EyeColor_name, int32(x))
}

func (EyeColor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_983c48047a4195b5, []int{0}
}

func init() {
	proto.RegisterEnum("user.EyeColor", EyeColor_name, EyeColor_value)
}

func init() { proto.RegisterFile("eyecolor.proto", fileDescriptor_983c48047a4195b5) }

var fileDescriptor_983c48047a4195b5 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x4c, 0x4d,
	0xce, 0xcf, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2a, 0xe1, 0xe2, 0x70, 0xad, 0x4c, 0x75, 0x06, 0x89, 0x0b, 0x89, 0x72, 0x09, 0x86, 0xfa,
	0x79, 0xfb, 0xf9, 0x87, 0xfb, 0xc5, 0xbb, 0x46, 0xba, 0xc6, 0x3b, 0xfb, 0xfb, 0xf8, 0x07, 0x09,
	0x30, 0x08, 0x09, 0x71, 0xf1, 0x39, 0xf9, 0x84, 0xba, 0x22, 0x89, 0x31, 0x0a, 0x09, 0x73, 0xf1,
	0xbb, 0x07, 0xb9, 0xba, 0x22, 0x2b, 0x64, 0x02, 0x09, 0x3a, 0x05, 0xa1, 0xea, 0x66, 0x06, 0x0b,
	0xfa, 0x38, 0x3a, 0x7b, 0x23, 0x09, 0xb2, 0x38, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0x81, 0x5c, 0x53, 0x90,
	0x94, 0xc4, 0x06, 0x76, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xbe, 0xe2, 0x7a, 0xae,
	0x00, 0x00, 0x00,
}