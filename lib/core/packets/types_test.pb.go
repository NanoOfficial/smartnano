package packets

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TestPacketDataPayload struct {
}

func (m *TestPacketDataPayload) Reset()         { *m = TestPacketDataPayload{} }
func (m *TestPacketDataPayload) String() string { return proto.CompactTextString(m) }
func (*TestPacketDataPayload) ProtoMessage()    {}
func (*TestPacketDataPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa36fb394f866bb, []int{0}
}
func (m *TestPacketDataPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestPacketDataPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestPacketDataPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestPacketDataPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestPacketDataPayload.Merge(m, src)
}
func (m *TestPacketDataPayload) XXX_Size() int {
	return m.Size()
}
func (m *TestPacketDataPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TestPacketDataPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TestPacketDataPayload proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestPacketDataPayload)(nil), "nano.packets.TestPacketDataPayload")
}

func init() { proto.RegisterFile("nano/packets/types_test.proto", fileDescriptor_6fa36fb394f866bb) }

var fileDescriptor_6fa36fb394f866bb = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2e, 0xca, 0x2f,
	0x2e, 0xd6, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x29, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x8e,
	0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0xcb, 0xeb, 0x41,
	0xe5, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x32, 0xfa, 0x20, 0x16, 0x44, 0x91, 0x92, 0x38,
	0x97, 0x68, 0x48, 0x6a, 0x71, 0x49, 0x00, 0x58, 0x91, 0x4b, 0x62, 0x49, 0x62, 0x40, 0x62, 0x65,
	0x4e, 0x7e, 0x62, 0x8a, 0x93, 0xeb, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9e, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x9f, 0x92, 0x58, 0x92, 0x98, 0x9c, 0x91, 0x98, 0x99, 0x97, 0x93, 0x98, 0xa4, 0x0f, 0x71, 0x53,
	0x05, 0xcc, 0x55, 0x49, 0x6c, 0x60, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x29,
	0xf7, 0x96, 0xad, 0x00, 0x00, 0x00,
}

func (m *TestPacketDataPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestPacketDataPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestPacketDataPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTypesTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestPacketDataPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTypesTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesTest(x uint64) (n int) {
	return sovTypesTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestPacketDataPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestPacketDataPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestPacketDataPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypesTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesTest
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
func skipTypesTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesTest
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
					return 0, ErrIntOverflowTypesTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypesTest
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
			if length < 0 {
				return 0, ErrInvalidLengthTypesTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesTest = fmt.Errorf("proto: unexpected end of group")
)
