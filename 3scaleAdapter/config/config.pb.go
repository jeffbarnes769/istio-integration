// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/3scaleAdapter/config/config.proto

/*
	Package config is a generated protocol buffer package.

	config for threescale

	It is generated from these files:
		mixer/adapter/3scaleAdapter/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// config for threescale
type Params struct {
	ServiceId   string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	SystemUrl   string `protobuf:"bytes,2,opt,name=system_url,json=systemUrl,proto3" json:"system_url,omitempty"`
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Params) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Params) GetSystemUrl() string {
	if m != nil {
		return m.SystemUrl
	}
	return ""
}

func (m *Params) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "adapter.threescale.config.Params")
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.ServiceId != that1.ServiceId {
		return false
	}
	if this.SystemUrl != that1.SystemUrl {
		return false
	}
	if this.AccessToken != that1.AccessToken {
		return false
	}
	return true
}
func (this *Params) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&config.Params{")
	s = append(s, "ServiceId: "+fmt.Sprintf("%#v", this.ServiceId)+",\n")
	s = append(s, "SystemUrl: "+fmt.Sprintf("%#v", this.SystemUrl)+",\n")
	s = append(s, "AccessToken: "+fmt.Sprintf("%#v", this.AccessToken)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ServiceId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ServiceId)))
		i += copy(dAtA[i:], m.ServiceId)
	}
	if len(m.SystemUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SystemUrl)))
		i += copy(dAtA[i:], m.SystemUrl)
	}
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.SystemUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`ServiceId:` + fmt.Sprintf("%v", this.ServiceId) + `,`,
		`SystemUrl:` + fmt.Sprintf("%v", this.SystemUrl) + `,`,
		`AccessToken:` + fmt.Sprintf("%v", this.AccessToken) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/adapter/3scaleAdapter/config/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcf, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x4f, 0x4c, 0x49, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0x37, 0x2e, 0x4e, 0x4e, 0xcc, 0x49,
	0x75, 0x84, 0xf2, 0x92, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0xa1, 0x94, 0x5e, 0x41, 0x51, 0x7e, 0x49,
	0xbe, 0x90, 0x24, 0x54, 0xa9, 0x5e, 0x49, 0x46, 0x51, 0x6a, 0x2a, 0x58, 0xb9, 0x1e, 0x44, 0x81,
	0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x95, 0x3e, 0x88, 0x05, 0xd1, 0xa0, 0x94, 0xc9, 0xc5,
	0x16, 0x90, 0x58, 0x94, 0x98, 0x5b, 0x2c, 0x24, 0xcb, 0xc5, 0x55, 0x9c, 0x5a, 0x54, 0x96, 0x99,
	0x9c, 0x1a, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x09, 0x15, 0xf1, 0x4c,
	0x01, 0x4b, 0x57, 0x16, 0x97, 0xa4, 0xe6, 0xc6, 0x97, 0x16, 0xe5, 0x48, 0x30, 0x41, 0xa5, 0xc1,
	0x22, 0xa1, 0x45, 0x39, 0x42, 0x8a, 0x5c, 0x3c, 0x89, 0xc9, 0xc9, 0xa9, 0xc5, 0xc5, 0xf1, 0x25,
	0xf9, 0xd9, 0xa9, 0x79, 0x12, 0xcc, 0x60, 0x05, 0xdc, 0x10, 0xb1, 0x10, 0x90, 0x90, 0x93, 0xc9,
	0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e,
	0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0x88, 0x62,
	0x83, 0x38, 0x3b, 0x89, 0x0d, 0xec, 0x4e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x5a,
	0x48, 0xaf, 0x0b, 0x01, 0x00, 0x00,
}
