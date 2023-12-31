// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/template7/wallet/v1.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Currency int32

const (
	Currency_ntd Currency = 0
	Currency_usd Currency = 1
	Currency_jpy Currency = 2
	Currency_cny Currency = 3
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "ntd",
		1: "usd",
		2: "jpy",
		3: "cny",
	}
	Currency_value = map[string]int32{
		"ntd": 0,
		"usd": 1,
		"jpy": 2,
		"cny": 3,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_template7_wallet_v1_proto_enumTypes[0].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_proto_template7_wallet_v1_proto_enumTypes[0]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{0}
}

type Unit int32

const (
	Unit_one   Unit = 0
	Unit_cent  Unit = 1
	Unit_milli Unit = 2
	Unit_micro Unit = 3
	Unit_nano  Unit = 4
	Unit_pico  Unit = 5
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "one",
		1: "cent",
		2: "milli",
		3: "micro",
		4: "nano",
		5: "pico",
	}
	Unit_value = map[string]int32{
		"one":   0,
		"cent":  1,
		"milli": 2,
		"micro": 3,
		"nano":  4,
		"pico":  5,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_template7_wallet_v1_proto_enumTypes[1].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_proto_template7_wallet_v1_proto_enumTypes[1]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{1}
}

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Balances []*Balance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wallet) GetBalances() []*Balance {
	if x != nil {
		return x.Balances
	}
	return nil
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency Currency `protobuf:"varint,1,opt,name=currency,proto3,enum=wallet.Currency" json:"currency,omitempty"`
	Amount   string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{1}
}

func (x *Balance) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_ntd
}

func (x *Balance) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type DepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency Currency `protobuf:"varint,1,opt,name=currency,proto3,enum=wallet.Currency" json:"currency,omitempty"`
	Amount   uint32   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DepositRequest) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_ntd
}

func (x *DepositRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency Currency `protobuf:"varint,1,opt,name=currency,proto3,enum=wallet.Currency" json:"currency,omitempty"`
	Amount   uint32   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{3}
}

func (x *WithdrawRequest) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_ntd
}

func (x *WithdrawRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromWalletId string   `protobuf:"bytes,1,opt,name=fromWalletId,proto3" json:"fromWalletId,omitempty"`
	ToWalletId   string   `protobuf:"bytes,2,opt,name=toWalletId,proto3" json:"toWalletId,omitempty"`
	Currency     Currency `protobuf:"varint,3,opt,name=currency,proto3,enum=wallet.Currency" json:"currency,omitempty"`
	Amount       uint32   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{4}
}

func (x *TransferRequest) GetFromWalletId() string {
	if x != nil {
		return x.FromWalletId
	}
	return ""
}

func (x *TransferRequest) GetToWalletId() string {
	if x != nil {
		return x.ToWalletId
	}
	return ""
}

func (x *TransferRequest) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_ntd
}

func (x *TransferRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_proto_template7_wallet_v1_proto protoreflect.FileDescriptor

var file_proto_template7_wallet_v1_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x37, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x45, 0x0a, 0x06, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x22, 0x4f, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x56, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x6f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x2a, 0x2e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x07, 0x0a, 0x03,
	0x6e, 0x74, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x75, 0x73, 0x64, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x6a, 0x70, 0x79, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x63, 0x6e, 0x79, 0x10, 0x03,
	0x2a, 0x43, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x63, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6d,
	0x69, 0x6c, 0x6c, 0x69, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x70,
	0x69, 0x63, 0x6f, 0x10, 0x05, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x37, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x37, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_template7_wallet_v1_proto_rawDescOnce sync.Once
	file_proto_template7_wallet_v1_proto_rawDescData = file_proto_template7_wallet_v1_proto_rawDesc
)

func file_proto_template7_wallet_v1_proto_rawDescGZIP() []byte {
	file_proto_template7_wallet_v1_proto_rawDescOnce.Do(func() {
		file_proto_template7_wallet_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_template7_wallet_v1_proto_rawDescData)
	})
	return file_proto_template7_wallet_v1_proto_rawDescData
}

var file_proto_template7_wallet_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_template7_wallet_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_template7_wallet_v1_proto_goTypes = []interface{}{
	(Currency)(0),           // 0: wallet.currency
	(Unit)(0),               // 1: wallet.unit
	(*Wallet)(nil),          // 2: wallet.wallet
	(*Balance)(nil),         // 3: wallet.balance
	(*DepositRequest)(nil),  // 4: wallet.depositRequest
	(*WithdrawRequest)(nil), // 5: wallet.withdrawRequest
	(*TransferRequest)(nil), // 6: wallet.transferRequest
}
var file_proto_template7_wallet_v1_proto_depIdxs = []int32{
	3, // 0: wallet.wallet.balances:type_name -> wallet.balance
	0, // 1: wallet.balance.currency:type_name -> wallet.currency
	0, // 2: wallet.depositRequest.currency:type_name -> wallet.currency
	0, // 3: wallet.withdrawRequest.currency:type_name -> wallet.currency
	0, // 4: wallet.transferRequest.currency:type_name -> wallet.currency
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_template7_wallet_v1_proto_init() }
func file_proto_template7_wallet_v1_proto_init() {
	if File_proto_template7_wallet_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_template7_wallet_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template7_wallet_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template7_wallet_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template7_wallet_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template7_wallet_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_template7_wallet_v1_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_template7_wallet_v1_proto_goTypes,
		DependencyIndexes: file_proto_template7_wallet_v1_proto_depIdxs,
		EnumInfos:         file_proto_template7_wallet_v1_proto_enumTypes,
		MessageInfos:      file_proto_template7_wallet_v1_proto_msgTypes,
	}.Build()
	File_proto_template7_wallet_v1_proto = out.File
	file_proto_template7_wallet_v1_proto_rawDesc = nil
	file_proto_template7_wallet_v1_proto_goTypes = nil
	file_proto_template7_wallet_v1_proto_depIdxs = nil
}
