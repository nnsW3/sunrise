// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquidstaking/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aae18fe0b821bcc, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aae18fe0b821bcc, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgLiquidStake struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgLiquidStake) Reset()         { *m = MsgLiquidStake{} }
func (m *MsgLiquidStake) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidStake) ProtoMessage()    {}
func (*MsgLiquidStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aae18fe0b821bcc, []int{2}
}
func (m *MsgLiquidStake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidStake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidStake.Merge(m, src)
}
func (m *MsgLiquidStake) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidStake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidStake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidStake proto.InternalMessageInfo

func (m *MsgLiquidStake) GetDelegatorAddress() string {
	if m != nil {
		return m.DelegatorAddress
	}
	return ""
}

func (m *MsgLiquidStake) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgLiquidStakeResponse struct {
}

func (m *MsgLiquidStakeResponse) Reset()         { *m = MsgLiquidStakeResponse{} }
func (m *MsgLiquidStakeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidStakeResponse) ProtoMessage()    {}
func (*MsgLiquidStakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aae18fe0b821bcc, []int{3}
}
func (m *MsgLiquidStakeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidStakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidStakeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidStakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidStakeResponse.Merge(m, src)
}
func (m *MsgLiquidStakeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidStakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidStakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidStakeResponse proto.InternalMessageInfo

type MsgLiquidUnstake struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgLiquidUnstake) Reset()         { *m = MsgLiquidUnstake{} }
func (m *MsgLiquidUnstake) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidUnstake) ProtoMessage()    {}
func (*MsgLiquidUnstake) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aae18fe0b821bcc, []int{4}
}
func (m *MsgLiquidUnstake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidUnstake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidUnstake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidUnstake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidUnstake.Merge(m, src)
}
func (m *MsgLiquidUnstake) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidUnstake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidUnstake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidUnstake proto.InternalMessageInfo

func (m *MsgLiquidUnstake) GetDelegatorAddress() string {
	if m != nil {
		return m.DelegatorAddress
	}
	return ""
}

func (m *MsgLiquidUnstake) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgLiquidUnstakeResponse struct {
	CompletionTime time.Time `protobuf:"bytes,1,opt,name=completion_time,json=completionTime,proto3,stdtime" json:"completion_time"`
}

func (m *MsgLiquidUnstakeResponse) Reset()         { *m = MsgLiquidUnstakeResponse{} }
func (m *MsgLiquidUnstakeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidUnstakeResponse) ProtoMessage()    {}
func (*MsgLiquidUnstakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aae18fe0b821bcc, []int{5}
}
func (m *MsgLiquidUnstakeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidUnstakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidUnstakeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidUnstakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidUnstakeResponse.Merge(m, src)
}
func (m *MsgLiquidUnstakeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidUnstakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidUnstakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidUnstakeResponse proto.InternalMessageInfo

func (m *MsgLiquidUnstakeResponse) GetCompletionTime() time.Time {
	if m != nil {
		return m.CompletionTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "sunrise.liquidstaking.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "sunrise.liquidstaking.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgLiquidStake)(nil), "sunrise.liquidstaking.MsgLiquidStake")
	proto.RegisterType((*MsgLiquidStakeResponse)(nil), "sunrise.liquidstaking.MsgLiquidStakeResponse")
	proto.RegisterType((*MsgLiquidUnstake)(nil), "sunrise.liquidstaking.MsgLiquidUnstake")
	proto.RegisterType((*MsgLiquidUnstakeResponse)(nil), "sunrise.liquidstaking.MsgLiquidUnstakeResponse")
}

func init() { proto.RegisterFile("sunrise/liquidstaking/tx.proto", fileDescriptor_7aae18fe0b821bcc) }

var fileDescriptor_7aae18fe0b821bcc = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0x0b, 0x54, 0xf4, 0x02, 0xfd, 0x61, 0x95, 0xd6, 0xb1, 0xc0, 0xa9, 0x2c, 0x01, 0x25,
	0x52, 0xee, 0x94, 0x20, 0x81, 0xc8, 0x04, 0xe9, 0x84, 0x44, 0x24, 0xe4, 0x50, 0x06, 0x96, 0xe8,
	0x92, 0x5c, 0xaf, 0xa7, 0xc6, 0x3e, 0xe3, 0x3b, 0x57, 0x0d, 0x13, 0x62, 0x64, 0xea, 0x9f, 0xc1,
	0x80, 0x50, 0x06, 0xfe, 0x04, 0x86, 0x8e, 0x15, 0x13, 0x53, 0x41, 0xc9, 0x90, 0x9d, 0xbf, 0x00,
	0xd9, 0xbe, 0x38, 0x4d, 0x9a, 0x8a, 0x88, 0x89, 0x25, 0xf1, 0x7b, 0xdf, 0xf7, 0xfc, 0xde, 0xf7,
	0xde, 0x67, 0x60, 0x89, 0xd0, 0x0b, 0x98, 0x20, 0xa8, 0xc3, 0xde, 0x86, 0xac, 0x2d, 0x24, 0x3e,
	0x60, 0x1e, 0x45, 0xf2, 0x08, 0xfa, 0x01, 0x97, 0x5c, 0xbf, 0xa5, 0x70, 0x38, 0x81, 0x9b, 0x6b,
	0xd8, 0x65, 0x1e, 0x47, 0xf1, 0x6f, 0xc2, 0x34, 0x37, 0x5b, 0x5c, 0xb8, 0x5c, 0x20, 0x57, 0x50,
	0x74, 0x58, 0x8a, 0xfe, 0x14, 0x90, 0x4b, 0x80, 0x46, 0x1c, 0xa1, 0x24, 0x50, 0xd0, 0x3a, 0xe5,
	0x94, 0x27, 0xf9, 0xe8, 0x49, 0x65, 0xed, 0xd9, 0x33, 0xf9, 0x38, 0xc0, 0xee, 0xa8, 0xd2, 0x52,
	0xdd, 0x9a, 0x58, 0x10, 0x74, 0x58, 0x6a, 0x12, 0x89, 0x4b, 0xa8, 0xc5, 0x99, 0xa7, 0xf0, 0x3c,
	0xe5, 0x9c, 0x76, 0x08, 0x8a, 0xa3, 0x66, 0xb8, 0x87, 0x24, 0x73, 0x89, 0x90, 0xd8, 0xf5, 0x15,
	0xe1, 0xc1, 0xec, 0x26, 0x13, 0x51, 0x42, 0xb5, 0x07, 0x1a, 0x58, 0xa9, 0x09, 0xba, 0xeb, 0xb7,
	0xb1, 0x24, 0x2f, 0xe3, 0x29, 0xf4, 0x47, 0x60, 0x09, 0x87, 0x72, 0x9f, 0x07, 0x4c, 0x76, 0x0d,
	0x6d, 0x4b, 0xdb, 0x5e, 0xaa, 0x1a, 0xdf, 0xbf, 0x16, 0xd7, 0x95, 0xbc, 0x67, 0xed, 0x76, 0x40,
	0x84, 0xa8, 0xcb, 0x80, 0x79, 0xd4, 0x19, 0x53, 0xf5, 0xa7, 0x60, 0x31, 0xd1, 0x61, 0x2c, 0x6c,
	0x69, 0xdb, 0xd9, 0xf2, 0x1d, 0x38, 0x73, 0xc1, 0x30, 0x69, 0x53, 0x5d, 0x3a, 0x39, 0xcb, 0x67,
	0x3e, 0x0d, 0x7b, 0x05, 0xcd, 0x51, 0x75, 0x95, 0xd7, 0x1f, 0x86, 0xbd, 0xc2, 0xf8, 0x8d, 0x1f,
	0x87, 0xbd, 0xc2, 0x0e, 0x65, 0x72, 0x3f, 0x6c, 0xc2, 0x16, 0x77, 0x91, 0x7a, 0x5d, 0xf1, 0x1d,
	0xf7, 0x48, 0x1a, 0x60, 0xdf, 0x47, 0x47, 0x53, 0x4a, 0xa7, 0x14, 0xd9, 0x39, 0xb0, 0x39, 0x95,
	0x72, 0x88, 0xf0, 0xb9, 0x27, 0x88, 0xfd, 0x59, 0x03, 0xcb, 0x35, 0x41, 0x5f, 0xc4, 0xf5, 0x75,
	0x89, 0x0f, 0x88, 0xfe, 0x1c, 0xac, 0xb5, 0x49, 0x87, 0x50, 0x2c, 0x79, 0xd0, 0xc0, 0x89, 0x5a,
	0xb5, 0x87, 0xdb, 0xbf, 0xcf, 0xf2, 0x46, 0x17, 0xbb, 0x9d, 0x8a, 0x7d, 0x81, 0x62, 0x3b, 0xab,
	0x69, 0x4e, 0xed, 0x48, 0x7f, 0x0c, 0x16, 0xb1, 0xcb, 0x43, 0x4f, 0xaa, 0x95, 0xe4, 0xa0, 0x5a,
	0x62, 0x74, 0x5b, 0xa8, 0x6e, 0x0b, 0x77, 0x38, 0xf3, 0xaa, 0x57, 0xa3, 0x75, 0x38, 0x8a, 0x5e,
	0xd9, 0x88, 0x36, 0x71, 0x71, 0x0c, 0xdb, 0x00, 0x1b, 0x93, 0xd3, 0xa6, 0x42, 0xbe, 0x68, 0x60,
	0x35, 0x85, 0x76, 0x3d, 0xf1, 0xdf, 0x4b, 0x61, 0xc0, 0x98, 0x9e, 0x77, 0x24, 0x46, 0xaf, 0x81,
	0x95, 0x16, 0x77, 0xfd, 0x0e, 0x91, 0x8c, 0x7b, 0x8d, 0xc8, 0xdf, 0xf1, 0xd4, 0xd9, 0xb2, 0x09,
	0x13, 0xf3, 0xc3, 0x91, 0xf9, 0xe1, 0xab, 0x91, 0xf9, 0xab, 0xd7, 0xa3, 0xb6, 0xc7, 0x3f, 0xf3,
	0x9a, 0xb3, 0x3c, 0x2e, 0x8e, 0xe0, 0xf2, 0xb7, 0x05, 0x70, 0xa5, 0x26, 0xa8, 0xbe, 0x07, 0x6e,
	0x4c, 0x38, 0xfd, 0xde, 0x25, 0x0e, 0x9d, 0x32, 0x8b, 0x09, 0xe7, 0xe3, 0xa5, 0xe3, 0xb7, 0x40,
	0xf6, 0xbc, 0xa1, 0xee, 0x5e, 0x5e, 0x7e, 0x8e, 0x66, 0x16, 0xe7, 0xa2, 0xa5, 0x4d, 0x18, 0xb8,
	0x39, 0x79, 0xec, 0xfb, 0x7f, 0xab, 0x57, 0x44, 0x13, 0xcd, 0x49, 0x1c, 0xb5, 0x32, 0xaf, 0xbd,
	0x8f, 0x3e, 0xd3, 0x6a, 0xfd, 0xa4, 0x6f, 0x69, 0xa7, 0x7d, 0x4b, 0xfb, 0xd5, 0xb7, 0xb4, 0xe3,
	0x81, 0x95, 0x39, 0x1d, 0x58, 0x99, 0x1f, 0x03, 0x2b, 0xf3, 0xe6, 0xc9, 0xbf, 0x7c, 0xa5, 0xb2,
	0xeb, 0x13, 0xd1, 0x5c, 0x8c, 0x2f, 0xf9, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xa1,
	0xac, 0x27, 0xae, 0x05, 0x00, 0x00,
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
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	LiquidStake(ctx context.Context, in *MsgLiquidStake, opts ...grpc.CallOption) (*MsgLiquidStakeResponse, error)
	LiquidUnstake(ctx context.Context, in *MsgLiquidUnstake, opts ...grpc.CallOption) (*MsgLiquidUnstakeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/sunrise.liquidstaking.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LiquidStake(ctx context.Context, in *MsgLiquidStake, opts ...grpc.CallOption) (*MsgLiquidStakeResponse, error) {
	out := new(MsgLiquidStakeResponse)
	err := c.cc.Invoke(ctx, "/sunrise.liquidstaking.Msg/LiquidStake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LiquidUnstake(ctx context.Context, in *MsgLiquidUnstake, opts ...grpc.CallOption) (*MsgLiquidUnstakeResponse, error) {
	out := new(MsgLiquidUnstakeResponse)
	err := c.cc.Invoke(ctx, "/sunrise.liquidstaking.Msg/LiquidUnstake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	LiquidStake(context.Context, *MsgLiquidStake) (*MsgLiquidStakeResponse, error)
	LiquidUnstake(context.Context, *MsgLiquidUnstake) (*MsgLiquidUnstakeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) LiquidStake(ctx context.Context, req *MsgLiquidStake) (*MsgLiquidStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidStake not implemented")
}
func (*UnimplementedMsgServer) LiquidUnstake(ctx context.Context, req *MsgLiquidUnstake) (*MsgLiquidUnstakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidUnstake not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sunrise.liquidstaking.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LiquidStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sunrise.liquidstaking.Msg/LiquidStake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidStake(ctx, req.(*MsgLiquidStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LiquidUnstake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidUnstake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidUnstake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sunrise.liquidstaking.Msg/LiquidUnstake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidUnstake(ctx, req.(*MsgLiquidUnstake))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sunrise.liquidstaking.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "LiquidStake",
			Handler:    _Msg_LiquidStake_Handler,
		},
		{
			MethodName: "LiquidUnstake",
			Handler:    _Msg_LiquidUnstake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sunrise/liquidstaking/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgLiquidStake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidStake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidStake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLiquidStakeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidStakeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidStakeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgLiquidUnstake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidUnstake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidUnstake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLiquidUnstakeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidUnstakeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidUnstakeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CompletionTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintTx(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0xa
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
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgLiquidStake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgLiquidStakeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgLiquidUnstake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgLiquidUnstakeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgLiquidStake) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgLiquidStake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidStake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgLiquidStakeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgLiquidStakeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidStakeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgLiquidUnstake) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgLiquidUnstake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidUnstake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgLiquidUnstakeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgLiquidUnstakeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidUnstakeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
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
