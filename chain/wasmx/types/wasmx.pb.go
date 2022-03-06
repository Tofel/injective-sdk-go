// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/wasmx/v1/wasmx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Params struct {
	// Set the status to active to indicate that the contract is to be executed in begin blocker.
	IsExecutionEnabled bool `protobuf:"varint,1,opt,name=is_execution_enabled,json=isExecutionEnabled,proto3" json:"is_execution_enabled,omitempty"`
	// registry_contract_address is the address of the contract that will be used to register the contract executed in begin blocker.
	RegistryContractAddress string `protobuf:"bytes,2,opt,name=registry_contract_address,json=registryContractAddress,proto3" json:"registry_contract_address,omitempty"`
	// gas_price_begin_block defines the gas price in INJ used to execute contracts.
	GasPriceBeginBlock types.Coin `protobuf:"bytes,3,opt,name=gas_price_begin_block,json=gasPriceBeginBlock,proto3" json:"gas_price_begin_block"`
	// Maximum gas to be used for the contract execution in begin blocker.
	// Should always be less than max block gas
	MaxGasBeginBlock uint64 `protobuf:"varint,4,opt,name=max_gas_begin_block,json=maxGasBeginBlock,proto3" json:"max_gas_begin_block,omitempty"`
	// The gas allocated so far for all the registered active contracts.
	TotalGasActiveContracts uint64 `protobuf:"varint,5,opt,name=total_gas_active_contracts,json=totalGasActiveContracts,proto3" json:"total_gas_active_contracts,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6818ff331f2cddc4, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetIsExecutionEnabled() bool {
	if m != nil {
		return m.IsExecutionEnabled
	}
	return false
}

func (m *Params) GetRegistryContractAddress() string {
	if m != nil {
		return m.RegistryContractAddress
	}
	return ""
}

func (m *Params) GetGasPriceBeginBlock() types.Coin {
	if m != nil {
		return m.GasPriceBeginBlock
	}
	return types.Coin{}
}

func (m *Params) GetMaxGasBeginBlock() uint64 {
	if m != nil {
		return m.MaxGasBeginBlock
	}
	return 0
}

func (m *Params) GetTotalGasActiveContracts() uint64 {
	if m != nil {
		return m.TotalGasActiveContracts
	}
	return 0
}

type ContractRegistrationRequestProposal struct {
	Title                       string                      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description                 string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ContractRegistrationRequest ContractRegistrationRequest `protobuf:"bytes,3,opt,name=contract_registration_request,json=contractRegistrationRequest,proto3" json:"contract_registration_request"`
}

func (m *ContractRegistrationRequestProposal) Reset()         { *m = ContractRegistrationRequestProposal{} }
func (m *ContractRegistrationRequestProposal) String() string { return proto.CompactTextString(m) }
func (*ContractRegistrationRequestProposal) ProtoMessage()    {}
func (*ContractRegistrationRequestProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6818ff331f2cddc4, []int{1}
}
func (m *ContractRegistrationRequestProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractRegistrationRequestProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractRegistrationRequestProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractRegistrationRequestProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRegistrationRequestProposal.Merge(m, src)
}
func (m *ContractRegistrationRequestProposal) XXX_Size() int {
	return m.Size()
}
func (m *ContractRegistrationRequestProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRegistrationRequestProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRegistrationRequestProposal proto.InternalMessageInfo

type ContractRegistrationRequest struct {
	// CodeID is the unique identifier of the smart contract to be registered.
	CodeId int64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// Unique Identifier for contract instance to be registered.
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// Maximum gas to be used for the smart contract execution.
	GasLimit uint64 `protobuf:"varint,3,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Set the status to active to indicate that the smart contract is to be executed in begin blocker.
	Active bool `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *ContractRegistrationRequest) Reset()         { *m = ContractRegistrationRequest{} }
func (m *ContractRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*ContractRegistrationRequest) ProtoMessage()    {}
func (*ContractRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6818ff331f2cddc4, []int{2}
}
func (m *ContractRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractRegistrationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRegistrationRequest.Merge(m, src)
}
func (m *ContractRegistrationRequest) XXX_Size() int {
	return m.Size()
}
func (m *ContractRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRegistrationRequest proto.InternalMessageInfo

func (m *ContractRegistrationRequest) GetCodeId() int64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *ContractRegistrationRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractRegistrationRequest) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ContractRegistrationRequest) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "injective.wasmx.v1.Params")
	proto.RegisterType((*ContractRegistrationRequestProposal)(nil), "injective.wasmx.v1.ContractRegistrationRequestProposal")
	proto.RegisterType((*ContractRegistrationRequest)(nil), "injective.wasmx.v1.ContractRegistrationRequest")
}

func init() { proto.RegisterFile("injective/wasmx/v1/wasmx.proto", fileDescriptor_6818ff331f2cddc4) }

var fileDescriptor_6818ff331f2cddc4 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x95, 0x12, 0xc7, 0x75, 0x98, 0xa1, 0x01, 0xeb, 0xd6, 0x8e, 0x8d, 0xca, 0x86, 0xbb, 0xb8,
	0x43, 0xa4, 0xba, 0xdd, 0xdc, 0x29, 0x0e, 0x82, 0x20, 0x68, 0x06, 0x83, 0x63, 0x17, 0x81, 0xa2,
	0x08, 0x85, 0xad, 0x24, 0xba, 0x3c, 0xda, 0xb5, 0xff, 0xa0, 0x63, 0xa7, 0xce, 0xfe, 0x90, 0x7e,
	0x40, 0xc6, 0x8c, 0x9d, 0x8a, 0xc2, 0x5e, 0xf2, 0x19, 0x05, 0x45, 0xd9, 0x30, 0x5a, 0x24, 0x9b,
	0xee, 0xde, 0x7b, 0xa7, 0xbb, 0x77, 0x47, 0xe4, 0x89, 0xfc, 0x13, 0x67, 0x5a, 0xcc, 0x78, 0xf0,
	0x95, 0x42, 0x36, 0x0f, 0x66, 0x03, 0xfb, 0xe1, 0x4f, 0x94, 0xd4, 0x12, 0xe3, 0x2d, 0xee, 0xdb,
	0xf4, 0x6c, 0xd0, 0xaa, 0x27, 0x32, 0x91, 0x05, 0x1c, 0x98, 0x2f, 0xcb, 0x6c, 0x79, 0x4c, 0x42,
	0x26, 0x21, 0x88, 0x28, 0xf0, 0x60, 0x36, 0x88, 0xb8, 0xa6, 0x83, 0x80, 0x49, 0x91, 0x5b, 0xbc,
	0xf7, 0x73, 0x0f, 0x55, 0xc7, 0x54, 0xd1, 0x0c, 0xf0, 0x1b, 0x54, 0x17, 0x10, 0xf2, 0x39, 0x67,
	0x53, 0x2d, 0x64, 0x1e, 0xf2, 0x9c, 0x46, 0x29, 0x8f, 0x9b, 0x6e, 0xd7, 0xed, 0xd7, 0x08, 0x16,
	0x70, 0xb1, 0x81, 0x2e, 0x2c, 0x82, 0x87, 0xe8, 0x44, 0xf1, 0x44, 0x80, 0x56, 0x8b, 0x90, 0xc9,
	0x5c, 0x2b, 0xca, 0x74, 0x48, 0xe3, 0x58, 0x71, 0x80, 0xe6, 0x5e, 0xd7, 0xed, 0x1f, 0x92, 0xc6,
	0x86, 0x70, 0x5e, 0xe2, 0x67, 0x16, 0xc6, 0x04, 0x3d, 0x4f, 0x28, 0x84, 0x13, 0x25, 0x18, 0x0f,
	0x23, 0x9e, 0x88, 0x3c, 0x8c, 0x52, 0xc9, 0x3e, 0x37, 0xf7, 0xbb, 0x6e, 0xff, 0xe8, 0xed, 0x89,
	0x6f, 0x1b, 0xf7, 0x4d, 0xe3, 0x7e, 0xd9, 0xb8, 0x7f, 0x2e, 0x45, 0x3e, 0xaa, 0xdc, 0xfe, 0xee,
	0x38, 0x04, 0x27, 0x14, 0xc6, 0x46, 0x3c, 0x32, 0xda, 0x91, 0x91, 0xe2, 0x53, 0xf4, 0x2c, 0xa3,
	0xf3, 0xd0, 0xd4, 0xdd, 0xad, 0x58, 0xe9, 0xba, 0xfd, 0x0a, 0x39, 0xce, 0xe8, 0xfc, 0x92, 0xc2,
	0x0e, 0xfd, 0x3d, 0x6a, 0x69, 0xa9, 0x69, 0x5a, 0x08, 0x68, 0x61, 0xe7, 0x76, 0x0c, 0x68, 0x1e,
	0x14, 0xaa, 0x46, 0xc1, 0xb8, 0xa4, 0x70, 0x56, 0xe0, 0x9b, 0x29, 0x60, 0x58, 0xb9, 0x5f, 0x76,
	0xdc, 0xde, 0xca, 0x45, 0xaf, 0x36, 0x39, 0x62, 0x27, 0xa5, 0xc6, 0x21, 0xc2, 0xbf, 0x4c, 0x39,
	0xe8, 0xb1, 0x92, 0x13, 0x09, 0x34, 0xc5, 0x75, 0x74, 0xa0, 0x85, 0x4e, 0x79, 0x61, 0xe6, 0x21,
	0xb1, 0x01, 0xee, 0xa2, 0xa3, 0x98, 0x03, 0x53, 0x62, 0x62, 0x34, 0xa5, 0x63, 0xbb, 0x29, 0xbc,
	0x40, 0x2f, 0xb7, 0xc6, 0xaa, 0x9d, 0xfa, 0xa1, 0xb2, 0x3f, 0x28, 0xdd, 0x0a, 0xfc, 0xff, 0x0f,
	0xc2, 0x7f, 0xa4, 0xaf, 0xd2, 0xc3, 0x36, 0x7b, 0x98, 0x32, 0xac, 0x7d, 0x5b, 0x76, 0x9c, 0xfb,
	0x65, 0xc7, 0xe9, 0xfd, 0x70, 0x51, 0xfb, 0x91, 0x62, 0xb8, 0x81, 0x9e, 0x30, 0x19, 0xf3, 0x50,
	0xd8, 0x5b, 0xd9, 0x27, 0x55, 0x13, 0x5e, 0xc5, 0xf8, 0x35, 0x3a, 0x7e, 0xe0, 0x2c, 0x9e, 0xb2,
	0x7f, 0xce, 0xa1, 0x8d, 0x0e, 0xcd, 0x16, 0x52, 0x91, 0x09, 0x3b, 0x54, 0x85, 0xd4, 0x12, 0x0a,
	0xd7, 0x26, 0xc6, 0x2f, 0x50, 0xd5, 0xae, 0xa7, 0x58, 0x65, 0x8d, 0x94, 0xd1, 0x88, 0xdf, 0xae,
	0x3c, 0xf7, 0x6e, 0xe5, 0xb9, 0x7f, 0x56, 0x9e, 0xfb, 0x7d, 0xed, 0x39, 0x77, 0x6b, 0xcf, 0xf9,
	0xb5, 0xf6, 0x9c, 0x8f, 0x1f, 0x12, 0xa1, 0x6f, 0xa6, 0x91, 0xcf, 0x64, 0x16, 0x5c, 0x6d, 0xac,
	0xb9, 0xa6, 0x11, 0x04, 0x5b, 0xa3, 0x4e, 0x99, 0x54, 0x7c, 0x37, 0xbc, 0xa1, 0x22, 0x0f, 0x32,
	0x19, 0x4f, 0x53, 0x0e, 0xe5, 0xb3, 0xd3, 0x8b, 0x09, 0x87, 0xa8, 0x5a, 0x3c, 0x95, 0x77, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x11, 0xb9, 0xf6, 0x96, 0x03, 0x00, 0x00,
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
	if this.IsExecutionEnabled != that1.IsExecutionEnabled {
		return false
	}
	if this.RegistryContractAddress != that1.RegistryContractAddress {
		return false
	}
	if !this.GasPriceBeginBlock.Equal(&that1.GasPriceBeginBlock) {
		return false
	}
	if this.MaxGasBeginBlock != that1.MaxGasBeginBlock {
		return false
	}
	if this.TotalGasActiveContracts != that1.TotalGasActiveContracts {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalGasActiveContracts != 0 {
		i = encodeVarintWasmx(dAtA, i, uint64(m.TotalGasActiveContracts))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxGasBeginBlock != 0 {
		i = encodeVarintWasmx(dAtA, i, uint64(m.MaxGasBeginBlock))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.GasPriceBeginBlock.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWasmx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.RegistryContractAddress) > 0 {
		i -= len(m.RegistryContractAddress)
		copy(dAtA[i:], m.RegistryContractAddress)
		i = encodeVarintWasmx(dAtA, i, uint64(len(m.RegistryContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.IsExecutionEnabled {
		i--
		if m.IsExecutionEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContractRegistrationRequestProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractRegistrationRequestProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractRegistrationRequestProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ContractRegistrationRequest.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWasmx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintWasmx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintWasmx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractRegistrationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractRegistrationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractRegistrationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.GasLimit != 0 {
		i = encodeVarintWasmx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintWasmx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintWasmx(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWasmx(dAtA []byte, offset int, v uint64) int {
	offset -= sovWasmx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsExecutionEnabled {
		n += 2
	}
	l = len(m.RegistryContractAddress)
	if l > 0 {
		n += 1 + l + sovWasmx(uint64(l))
	}
	l = m.GasPriceBeginBlock.Size()
	n += 1 + l + sovWasmx(uint64(l))
	if m.MaxGasBeginBlock != 0 {
		n += 1 + sovWasmx(uint64(m.MaxGasBeginBlock))
	}
	if m.TotalGasActiveContracts != 0 {
		n += 1 + sovWasmx(uint64(m.TotalGasActiveContracts))
	}
	return n
}

func (m *ContractRegistrationRequestProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovWasmx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovWasmx(uint64(l))
	}
	l = m.ContractRegistrationRequest.Size()
	n += 1 + l + sovWasmx(uint64(l))
	return n
}

func (m *ContractRegistrationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovWasmx(uint64(m.CodeId))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovWasmx(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovWasmx(uint64(m.GasLimit))
	}
	if m.Active {
		n += 2
	}
	return n
}

func sovWasmx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWasmx(x uint64) (n int) {
	return sovWasmx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasmx
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsExecutionEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsExecutionEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
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
				return ErrInvalidLengthWasmx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasmx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegistryContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceBeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
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
				return ErrInvalidLengthWasmx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasmx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GasPriceBeginBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGasBeginBlock", wireType)
			}
			m.MaxGasBeginBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGasBeginBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalGasActiveContracts", wireType)
			}
			m.TotalGasActiveContracts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalGasActiveContracts |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWasmx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasmx
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
func (m *ContractRegistrationRequestProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasmx
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
			return fmt.Errorf("proto: ContractRegistrationRequestProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractRegistrationRequestProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
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
				return ErrInvalidLengthWasmx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasmx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
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
				return ErrInvalidLengthWasmx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasmx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractRegistrationRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
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
				return ErrInvalidLengthWasmx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasmx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractRegistrationRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasmx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasmx
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
func (m *ContractRegistrationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasmx
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
			return fmt.Errorf("proto: ContractRegistrationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractRegistrationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
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
				return ErrInvalidLengthWasmx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasmx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWasmx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasmx
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
func skipWasmx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWasmx
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
					return 0, ErrIntOverflowWasmx
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
					return 0, ErrIntOverflowWasmx
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
				return 0, ErrInvalidLengthWasmx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWasmx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWasmx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWasmx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWasmx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWasmx = fmt.Errorf("proto: unexpected end of group")
)