// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coreum/dex/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// MsgIssue defines message to issue new fungible token.
type MsgCreateLimitOrder struct {
	Owner         string        `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	OfferedAmount types.Coin    `protobuf:"bytes,2,opt,name=offered_amount,json=offeredAmount,proto3" json:"offered_amount"`
	SellPrice     types.DecCoin `protobuf:"bytes,3,opt,name=sell_price,json=sellPrice,proto3" json:"sell_price"`
}

func (m *MsgCreateLimitOrder) Reset()         { *m = MsgCreateLimitOrder{} }
func (m *MsgCreateLimitOrder) String() string { return proto.CompactTextString(m) }
func (*MsgCreateLimitOrder) ProtoMessage()    {}
func (*MsgCreateLimitOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3181ef84525da2, []int{0}
}
func (m *MsgCreateLimitOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateLimitOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateLimitOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateLimitOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateLimitOrder.Merge(m, src)
}
func (m *MsgCreateLimitOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateLimitOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateLimitOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateLimitOrder proto.InternalMessageInfo

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3181ef84525da2, []int{1}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return m.Size()
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateLimitOrder)(nil), "coreum.dex.v1.MsgCreateLimitOrder")
	proto.RegisterType((*EmptyResponse)(nil), "coreum.dex.v1.EmptyResponse")
}

func init() { proto.RegisterFile("coreum/dex/v1/tx.proto", fileDescriptor_6b3181ef84525da2) }

var fileDescriptor_6b3181ef84525da2 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xdb, 0xcb, 0xbd, 0x37, 0x61, 0x6e, 0xb8, 0x9a, 0x4a, 0x0c, 0x12, 0x32, 0x12, 0x56,
	0xac, 0x66, 0x52, 0xf4, 0x05, 0x00, 0x65, 0x25, 0xc1, 0x10, 0x57, 0xba, 0x20, 0xfd, 0x73, 0xa8,
	0x4d, 0xe8, 0x9c, 0x66, 0x66, 0x5a, 0xcb, 0x5b, 0xf8, 0x34, 0x3e, 0x03, 0x4b, 0x96, 0xae, 0x8c,
	0xc2, 0x8b, 0x98, 0xb6, 0x2c, 0x44, 0xd9, 0x9d, 0x73, 0x72, 0xbe, 0xdf, 0x97, 0x73, 0x3e, 0x72,
	0xea, 0xa1, 0x84, 0x24, 0xe2, 0x3e, 0x64, 0x3c, 0xb5, 0xb9, 0xce, 0x58, 0x2c, 0x51, 0xa3, 0x55,
	0x2b, 0xe7, 0xcc, 0x87, 0x8c, 0xa5, 0x76, 0x93, 0x7a, 0xa8, 0x22, 0x54, 0xdc, 0x75, 0x14, 0xf0,
	0xd4, 0x76, 0x41, 0x3b, 0x36, 0xf7, 0x30, 0x14, 0xe5, 0x7a, 0xb3, 0x1e, 0x60, 0x80, 0x45, 0xc9,
	0xf3, 0xaa, 0x9c, 0x76, 0x5e, 0x4c, 0x72, 0x32, 0x56, 0xc1, 0x50, 0x82, 0xa3, 0xe1, 0x26, 0x8c,
	0x42, 0x3d, 0x91, 0x3e, 0x48, 0xab, 0x4e, 0xfe, 0xe0, 0x93, 0x00, 0xd9, 0x30, 0xdb, 0x66, 0xb7,
	0x3a, 0x2d, 0x1b, 0x6b, 0x44, 0xfe, 0xe3, 0x7c, 0x0e, 0x12, 0xfc, 0x99, 0x13, 0x61, 0x22, 0x74,
	0xe3, 0x57, 0xdb, 0xec, 0xfe, 0xeb, 0x9d, 0xb1, 0xd2, 0x9c, 0xe5, 0xe6, 0x6c, 0x67, 0xce, 0x86,
	0x18, 0x8a, 0xc1, 0xef, 0xd5, 0xdb, 0xb9, 0x31, 0xad, 0xed, 0x64, 0xfd, 0x42, 0x65, 0xf5, 0x09,
	0x51, 0xb0, 0x58, 0xcc, 0x62, 0x19, 0x7a, 0xd0, 0xa8, 0x14, 0x8c, 0xd6, 0x41, 0xc6, 0x15, 0x78,
	0x5f, 0x30, 0xd5, 0x5c, 0x75, 0x9b, 0x8b, 0x3a, 0x47, 0xa4, 0x76, 0x1d, 0xc5, 0x7a, 0x39, 0x05,
	0x15, 0xa3, 0x50, 0xd0, 0x7b, 0x20, 0x95, 0xb1, 0x0a, 0xac, 0x3b, 0x72, 0xfc, 0xe3, 0x98, 0x0e,
	0xdb, 0x7b, 0x15, 0x3b, 0x70, 0x70, 0xb3, 0xf5, 0x6d, 0x67, 0x0f, 0x3e, 0x98, 0xac, 0x3e, 0xa8,
	0xb1, 0xda, 0x50, 0x73, 0xbd, 0xa1, 0xe6, 0xfb, 0x86, 0x9a, 0xcf, 0x5b, 0x6a, 0xac, 0xb7, 0xd4,
	0x78, 0xdd, 0x52, 0xe3, 0xde, 0x0e, 0x42, 0xfd, 0x98, 0xb8, 0xcc, 0xc3, 0x88, 0x0f, 0x0b, 0xca,
	0x08, 0x13, 0xe1, 0x3b, 0x3a, 0x44, 0xc1, 0x77, 0xe9, 0xa5, 0x97, 0x3c, 0x2b, 0x22, 0xd4, 0xcb,
	0x18, 0x94, 0xfb, 0xb7, 0x78, 0xff, 0xc5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0xfe, 0x9a,
	0x0e, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Issue defines a method to issue a new fungible token.
	CreateLimitOrder(ctx context.Context, in *MsgCreateLimitOrder, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateLimitOrder(ctx context.Context, in *MsgCreateLimitOrder, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/coreum.dex.v1.Msg/CreateLimitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Issue defines a method to issue a new fungible token.
	CreateLimitOrder(context.Context, *MsgCreateLimitOrder) (*EmptyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateLimitOrder(ctx context.Context, req *MsgCreateLimitOrder) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimitOrder not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateLimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateLimitOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateLimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreum.dex.v1.Msg/CreateLimitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateLimitOrder(ctx, req.(*MsgCreateLimitOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coreum.dex.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLimitOrder",
			Handler:    _Msg_CreateLimitOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coreum/dex/v1/tx.proto",
}

func (m *MsgCreateLimitOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateLimitOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateLimitOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SellPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.OfferedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EmptyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmptyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateLimitOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.OfferedAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.SellPrice.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *EmptyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateLimitOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateLimitOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateLimitOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SellPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *EmptyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
