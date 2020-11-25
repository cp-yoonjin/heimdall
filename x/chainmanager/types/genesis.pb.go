// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chainmanager/v1beta/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_maticnetwork_heimdall_types_common "github.com/maticnetwork/heimdall/types/common"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the chainmanager module's genesis state.
type GenesisState struct {
	// heimdall.types.chainmanager.x.params
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f2fccb4e0c3c0d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisState.Unmarshal(m, b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return xxx_messageInfo_GenesisState.Size(m)
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type ChainParams struct {
	BorChainID            string                                                        `protobuf:"bytes,1,opt,name=bor_chain_id,json=borChainId,proto3" json:"bor_chain_id,omitempty"`
	MaticTokenAddress     github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,2,opt,name=matic_token_address,json=maticTokenAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"matic_token_address,omitempty"`
	StakingManagerAddress github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,3,opt,name=staking_manager_address,json=stakingManagerAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"staking_manager_address,omitempty"`
	SlashManagerAddress   github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,4,opt,name=slash_manager_address,json=slashManagerAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"slash_manager_address,omitempty"`
	RootChainAddress      github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,5,opt,name=root_chain_address,json=rootChainAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"root_chain_address,omitempty"`
	StakingInfoAddress    github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,6,opt,name=staking_info_address,json=stakingInfoAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"staking_info_address,omitempty"`
	StateSenderAddress    github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,7,opt,name=state_sender_address,json=stateSenderAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"state_sender_address,omitempty"`
	StateReceiverAddress  github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,8,opt,name=state_receiver_address,json=stateReceiverAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"state_receiver_address,omitempty"`
	ValidatorSetAddress   github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,9,opt,name=validator_set_address,json=validatorSetAddress,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"validator_set_address,omitempty"`
}

func (m *ChainParams) Reset()      { *m = ChainParams{} }
func (*ChainParams) ProtoMessage() {}
func (*ChainParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f2fccb4e0c3c0d, []int{1}
}
func (m *ChainParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainParams.Unmarshal(m, b)
}
func (m *ChainParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainParams.Marshal(b, m, deterministic)
}
func (m *ChainParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainParams.Merge(m, src)
}
func (m *ChainParams) XXX_Size() int {
	return xxx_messageInfo_ChainParams.Size(m)
}
func (m *ChainParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainParams.DiscardUnknown(m)
}

var xxx_messageInfo_ChainParams proto.InternalMessageInfo

type Params struct {
	MainchainTxConfirmations  uint64       `protobuf:"varint,1,opt,name=mainchain_tx_confirmations,json=mainchainTxConfirmations,proto3" json:"mainchain_tx_confirmations,omitempty"`
	MaticchainTxConfirmations uint64       `protobuf:"varint,2,opt,name=maticchain_tx_confirmations,json=maticchainTxConfirmations,proto3" json:"maticchain_tx_confirmations,omitempty"`
	ChainParams               *ChainParams `protobuf:"bytes,3,opt,name=chain_params,json=chainParams,proto3" json:"chain_params,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f2fccb4e0c3c0d, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "heimdall.chainmanager.v1beta1.GenesisState")
	proto.RegisterType((*ChainParams)(nil), "heimdall.chainmanager.v1beta1.ChainParams")
	proto.RegisterType((*Params)(nil), "heimdall.chainmanager.v1beta1.Params")
}

func init() { proto.RegisterFile("chainmanager/v1beta/genesis.proto", fileDescriptor_f2f2fccb4e0c3c0d) }

var fileDescriptor_f2f2fccb4e0c3c0d = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xd3, 0x6d, 0xff, 0x6e, 0x73, 0xab, 0xbf, 0x20, 0xdb, 0xa0, 0x0c, 0x91, 0x8e, 0x4a,
	0x48, 0x13, 0x87, 0x84, 0x81, 0xb8, 0x20, 0x86, 0xb4, 0x0e, 0x09, 0x76, 0xa8, 0x40, 0xe9, 0x4e,
	0x5c, 0x22, 0x37, 0x71, 0x53, 0xab, 0x8d, 0xdf, 0x62, 0x7b, 0x5d, 0xf7, 0x0d, 0x38, 0x72, 0xe4,
	0xb8, 0x8f, 0xc3, 0x71, 0x47, 0x4e, 0x03, 0xb5, 0x9f, 0x02, 0x4e, 0x28, 0xb6, 0x93, 0x66, 0x43,
	0xb0, 0x4b, 0x6e, 0xee, 0xdb, 0xe7, 0x79, 0x7e, 0xf6, 0x6b, 0xe7, 0x45, 0x0f, 0xc3, 0x01, 0xa6,
	0x2c, 0xc1, 0x0c, 0xc7, 0x84, 0x7b, 0x93, 0xbd, 0x1e, 0x91, 0xd8, 0x8b, 0x09, 0x23, 0x82, 0x0a,
	0x77, 0xcc, 0x41, 0x82, 0xfd, 0x60, 0x40, 0x68, 0x12, 0xe1, 0xd1, 0xc8, 0x2d, 0x6a, 0x5d, 0xad,
	0xdd, 0xdb, 0xde, 0x8c, 0x21, 0x06, 0xa5, 0xf4, 0xd2, 0x95, 0x36, 0xb5, 0x3a, 0xa8, 0xfe, 0x46,
	0xa7, 0x74, 0x25, 0x96, 0xc4, 0xde, 0x47, 0xd5, 0x31, 0xe6, 0x38, 0x11, 0x8d, 0xca, 0x4e, 0x65,
	0xb7, 0xf6, 0xf4, 0x91, 0xfb, 0xcf, 0x54, 0xf7, 0xbd, 0x12, 0xfb, 0xc6, 0xd4, 0xfa, 0xb9, 0x8a,
	0x6a, 0x87, 0xa9, 0x4e, 0xd7, 0xed, 0x27, 0xa8, 0xde, 0x03, 0x1e, 0x28, 0x6b, 0x40, 0x23, 0x15,
	0xba, 0xde, 0xfe, 0x7f, 0x76, 0xd9, 0x44, 0x6d, 0xe0, 0x4a, 0x79, 0xf4, 0xda, 0x47, 0xbd, 0x6c,
	0x1d, 0xd9, 0x1f, 0xd1, 0x46, 0x82, 0x25, 0x0d, 0x03, 0x09, 0x43, 0xc2, 0x02, 0x1c, 0x45, 0x9c,
	0x08, 0xd1, 0x58, 0x52, 0xc6, 0x83, 0x5f, 0x97, 0xcd, 0xfd, 0x98, 0xca, 0xc1, 0x49, 0xcf, 0x0d,
	0x21, 0xf1, 0x94, 0x92, 0x11, 0x79, 0x0a, 0x7c, 0xe8, 0x65, 0x1b, 0xf5, 0xe4, 0xd9, 0x98, 0x08,
	0x2f, 0x84, 0x24, 0x01, 0xe6, 0xbe, 0x35, 0xd5, 0x03, 0x1d, 0xe4, 0xdf, 0x56, 0x9e, 0xe3, 0x34,
	0xdc, 0x94, 0xec, 0x33, 0x74, 0x57, 0x48, 0x3c, 0xa4, 0x2c, 0x0e, 0xcc, 0xf1, 0x72, 0xec, 0x72,
	0x59, 0xd8, 0x2d, 0x43, 0xe8, 0x68, 0x40, 0x86, 0x3e, 0x41, 0x5b, 0x62, 0x84, 0xc5, 0xe0, 0x0f,
	0xf0, 0x4a, 0x59, 0xe0, 0x0d, 0x95, 0x7f, 0x0d, 0x0b, 0xc8, 0xe6, 0x00, 0xd2, 0xdc, 0x4b, 0xc6,
	0xfc, 0xaf, 0x2c, 0xe6, 0xad, 0x34, 0x5c, 0x5d, 0x68, 0x06, 0x14, 0x68, 0x33, 0x6b, 0x31, 0x65,
	0x7d, 0xc8, 0x91, 0xd5, 0xb2, 0x90, 0xb6, 0x89, 0x3f, 0x62, 0x7d, 0xb8, 0x0a, 0x95, 0x24, 0x10,
	0x84, 0x45, 0x85, 0xde, 0xae, 0x96, 0x09, 0x95, 0xa4, 0xab, 0xd2, 0x33, 0xe8, 0x29, 0xba, 0xa3,
	0xa1, 0x9c, 0x84, 0x84, 0x4e, 0x0a, 0xd8, 0xb5, 0xb2, 0xb0, 0xfa, 0x54, 0xbe, 0xc9, 0x2f, 0x3c,
	0xa5, 0x09, 0x1e, 0xd1, 0x08, 0x4b, 0xe0, 0x81, 0x20, 0x32, 0xe7, 0xae, 0xef, 0x54, 0x76, 0xeb,
	0xa5, 0x3c, 0xa5, 0x3c, 0xbf, 0x4b, 0xa4, 0x29, 0xbe, 0x58, 0xfb, 0x74, 0xde, 0xb4, 0xbe, 0x9c,
	0x37, 0xad, 0xd6, 0xf7, 0x0a, 0xaa, 0x9a, 0xcf, 0xfe, 0x25, 0xda, 0x4e, 0x30, 0x65, 0xfa, 0x79,
	0xc9, 0x69, 0x10, 0x02, 0xeb, 0x53, 0x9e, 0x22, 0x81, 0xe9, 0xc9, 0xb2, 0xe2, 0x37, 0x72, 0xc5,
	0xf1, 0xf4, 0xb0, 0xf8, 0xbf, 0xfd, 0x0a, 0xdd, 0x57, 0xbb, 0xfb, 0x8b, 0x7d, 0x49, 0xd9, 0xef,
	0x2d, 0x24, 0xd7, 0xfd, 0x1d, 0x54, 0xd7, 0x56, 0x33, 0xc9, 0x96, 0xd5, 0x24, 0x7b, 0x7c, 0xc3,
	0x24, 0x2b, 0x8c, 0x2d, 0xbf, 0x16, 0x2e, 0x7e, 0x2c, 0x4e, 0xd8, 0x7e, 0xf7, 0x75, 0xe6, 0x58,
	0x17, 0x33, 0xc7, 0xfa, 0x31, 0x73, 0xac, 0xcf, 0x73, 0xc7, 0xba, 0x98, 0x3b, 0xd6, 0xb7, 0xb9,
	0x63, 0x7d, 0x78, 0x7e, 0x63, 0x67, 0xa7, 0xde, 0x95, 0x09, 0xae, 0x1a, 0xdd, 0xab, 0xaa, 0x21,
	0xfc, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0xbf, 0x43, 0xc8, 0xde, 0x05, 0x00, 0x00,
}