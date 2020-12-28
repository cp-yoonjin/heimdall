// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: clerk/v1beta/clerk.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventRecord struct {
	Id         uint64                                        `protobuf:"varint,1,opt,name=id,proto3" json:"id" yaml:"id"`
	Contract   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=contract,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"contract,omitempty"`
	Data       []byte                                        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	RecordTime time.Time                                     `protobuf:"bytes,4,opt,name=record_time,json=recordTime,proto3,stdtime" json:"record_time" yaml:"record_time"`
	LogIndex   uint64                                        `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	TxHash     []byte                                        `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	ChainId    string                                        `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *EventRecord) Reset()         { *m = EventRecord{} }
func (m *EventRecord) String() string { return proto.CompactTextString(m) }
func (*EventRecord) ProtoMessage()    {}
func (*EventRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1edb6cdcc8eb1288, []int{0}
}
func (m *EventRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRecord.Merge(m, src)
}
func (m *EventRecord) XXX_Size() int {
	return m.Size()
}
func (m *EventRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRecord.DiscardUnknown(m)
}

var xxx_messageInfo_EventRecord proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRecord)(nil), "heimdall.clerk.v1beta1.EventRecord")
}

func init() { proto.RegisterFile("clerk/v1beta/clerk.proto", fileDescriptor_1edb6cdcc8eb1288) }

var fileDescriptor_1edb6cdcc8eb1288 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x50, 0xfa, 0xc7, 0x65, 0x32, 0x08, 0x4c, 0x91, 0xe2, 0xaa, 0x2c, 0x5d, 0xce,
	0x56, 0x61, 0xbb, 0xed, 0x2a, 0x21, 0xdd, 0x0d, 0x2c, 0x11, 0x13, 0x0c, 0x95, 0x63, 0x9b, 0xc4,
	0x6a, 0x12, 0x57, 0xb1, 0xef, 0xc8, 0x7d, 0x03, 0xc6, 0xfb, 0x08, 0x7c, 0x9c, 0x63, 0xbb, 0x91,
	0x29, 0xa0, 0x76, 0x63, 0xbc, 0x91, 0x09, 0xc5, 0xa6, 0xe8, 0x26, 0xbf, 0xcf, 0xab, 0xe7, 0xb5,
	0x9e, 0xdf, 0xfb, 0x42, 0x2c, 0x4a, 0xd5, 0x6c, 0xd9, 0xd5, 0x2a, 0x53, 0x8e, 0x33, 0x2f, 0xe8,
	0xae, 0x31, 0xce, 0xa0, 0xe7, 0x85, 0xd2, 0x95, 0xe4, 0x65, 0x49, 0x43, 0x37, 0x58, 0x56, 0xb3,
	0x67, 0xb9, 0xc9, 0x8d, 0xb7, 0xb0, 0xbe, 0x0a, 0xee, 0x19, 0xc9, 0x8d, 0xc9, 0x4b, 0xc5, 0xbc,
	0xca, 0x2e, 0x3f, 0x33, 0xa7, 0x2b, 0x65, 0x1d, 0xaf, 0x76, 0xc1, 0xb0, 0xf8, 0x1e, 0xc3, 0xe9,
	0xbb, 0x2b, 0x55, 0xbb, 0x54, 0x09, 0xd3, 0x48, 0xf4, 0x1a, 0xc6, 0x5a, 0x62, 0x30, 0x07, 0xcb,
	0xc1, 0xfa, 0xe9, 0xef, 0x8e, 0xc4, 0x5a, 0xde, 0x77, 0x64, 0x72, 0xcd, 0xab, 0xf2, 0x74, 0xa1,
	0xe5, 0x22, 0x8d, 0xb5, 0x44, 0xef, 0xe1, 0x58, 0x98, 0xda, 0x35, 0x5c, 0x38, 0x1c, 0xcf, 0xc1,
	0x72, 0xb2, 0x5e, 0xfd, 0xe9, 0xc8, 0x49, 0xae, 0x5d, 0x71, 0x99, 0x51, 0x61, 0x2a, 0x26, 0x8c,
	0xad, 0x8c, 0xfd, 0xf7, 0x9c, 0x58, 0xb9, 0x65, 0xee, 0x7a, 0xa7, 0x2c, 0x3d, 0x13, 0xe2, 0x4c,
	0xca, 0x46, 0x59, 0x9b, 0xfe, 0xff, 0x02, 0x21, 0x38, 0x90, 0xdc, 0x71, 0xfc, 0x68, 0x0e, 0x96,
	0x4f, 0x52, 0x5f, 0xa3, 0x4f, 0x70, 0xda, 0xf8, 0x44, 0x9b, 0x3e, 0x31, 0x1e, 0xcc, 0xc1, 0x72,
	0xfa, 0x66, 0x46, 0x03, 0x0e, 0x3d, 0xe2, 0xd0, 0x0f, 0x47, 0x9c, 0x75, 0x72, 0xdb, 0x91, 0xe8,
	0xbe, 0x23, 0x28, 0x44, 0x7d, 0x30, 0xbc, 0xb8, 0xf9, 0x49, 0x40, 0x0a, 0x43, 0xa7, 0x1f, 0x40,
	0xaf, 0xe0, 0xa4, 0x34, 0xf9, 0x46, 0xd7, 0x52, 0xb5, 0xf8, 0x71, 0xcf, 0x9a, 0x8e, 0x4b, 0x93,
	0x5f, 0xf4, 0x1a, 0xbd, 0x80, 0x23, 0xd7, 0x6e, 0x0a, 0x6e, 0x0b, 0x3c, 0xf4, 0x81, 0x86, 0xae,
	0x3d, 0xe7, 0xb6, 0x40, 0x2f, 0xe1, 0x58, 0x14, 0x5c, 0xd7, 0x1b, 0x2d, 0xf1, 0xa8, 0xa7, 0x4e,
	0x47, 0x5e, 0x5f, 0xc8, 0xd3, 0xc1, 0xd7, 0x6f, 0x24, 0x5a, 0x9f, 0xdf, 0xee, 0x13, 0x70, 0xb7,
	0x4f, 0xc0, 0xaf, 0x7d, 0x02, 0x6e, 0x0e, 0x49, 0x74, 0x77, 0x48, 0xa2, 0x1f, 0x87, 0x24, 0xfa,
	0x48, 0x1f, 0xac, 0xa6, 0xe2, 0x4e, 0x8b, 0x5a, 0xb9, 0x2f, 0xa6, 0xd9, 0xb2, 0xe3, 0x31, 0x59,
	0x1b, 0x8e, 0x1c, 0xd6, 0x94, 0x0d, 0x3d, 0xe0, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x8c, 0xca, 0x3a, 0x07, 0x02, 0x00, 0x00,
}

func (m *EventRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.LogIndex != 0 {
		i = encodeVarintClerk(dAtA, i, uint64(m.LogIndex))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RecordTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RecordTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClerk(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintClerk(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintClerk(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClerk(dAtA []byte, offset int, v uint64) int {
	offset -= sovClerk(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClerk(uint64(m.Id))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RecordTime)
	n += 1 + l + sovClerk(uint64(l))
	if m.LogIndex != 0 {
		n += 1 + sovClerk(uint64(m.LogIndex))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovClerk(uint64(l))
	}
	return n
}

func sovClerk(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClerk(x uint64) (n int) {
	return sovClerk(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClerk
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
			return fmt.Errorf("proto: EventRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RecordTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogIndex", wireType)
			}
			m.LogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClerk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClerk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClerk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClerk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClerk
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClerk
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
func skipClerk(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClerk
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
					return 0, ErrIntOverflowClerk
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
					return 0, ErrIntOverflowClerk
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
				return 0, ErrInvalidLengthClerk
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClerk
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClerk
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClerk        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClerk          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClerk = fmt.Errorf("proto: unexpected end of group")
)