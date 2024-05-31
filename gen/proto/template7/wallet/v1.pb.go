// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
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

type Direction int32

const (
	Direction_unknown     Direction = 0
	Direction_deposit     Direction = 1
	Direction_withdraw    Direction = 2
	Direction_transferIn  Direction = 3
	Direction_transferOut Direction = 4
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "unknown",
		1: "deposit",
		2: "withdraw",
		3: "transferIn",
		4: "transferOut",
	}
	Direction_value = map[string]int32{
		"unknown":     0,
		"deposit":     1,
		"withdraw":    2,
		"transferIn":  3,
		"transferOut": 4,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_template7_wallet_v1_proto_enumTypes[2].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_proto_template7_wallet_v1_proto_enumTypes[2]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{2}
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

type CurrencyBalanceHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Direction     Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=wallet.Direction" json:"direction,omitempty"`
	Amount        string    `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	BalanceBefore string    `protobuf:"bytes,4,opt,name=balanceBefore,proto3" json:"balanceBefore,omitempty"`
	BalanceAfter  string    `protobuf:"bytes,5,opt,name=balanceAfter,proto3" json:"balanceAfter,omitempty"`
	Timestamp     uint64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Note          string    `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CurrencyBalanceHistory) Reset() {
	*x = CurrencyBalanceHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyBalanceHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyBalanceHistory) ProtoMessage() {}

func (x *CurrencyBalanceHistory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyBalanceHistory.ProtoReflect.Descriptor instead.
func (*CurrencyBalanceHistory) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CurrencyBalanceHistory) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CurrencyBalanceHistory) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_unknown
}

func (x *CurrencyBalanceHistory) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CurrencyBalanceHistory) GetBalanceBefore() string {
	if x != nil {
		return x.BalanceBefore
	}
	return ""
}

func (x *CurrencyBalanceHistory) GetBalanceAfter() string {
	if x != nil {
		return x.BalanceAfter
	}
	return ""
}

func (x *CurrencyBalanceHistory) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CurrencyBalanceHistory) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type BalanceHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Direction     Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=wallet.Direction" json:"direction,omitempty"`
	Currency      Currency  `protobuf:"varint,3,opt,name=currency,proto3,enum=wallet.Currency" json:"currency,omitempty"`
	Amount        string    `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	BalanceBefore string    `protobuf:"bytes,5,opt,name=balanceBefore,proto3" json:"balanceBefore,omitempty"`
	BalanceAfter  string    `protobuf:"bytes,6,opt,name=balanceAfter,proto3" json:"balanceAfter,omitempty"`
	Timestamp     int64     `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Note          string    `protobuf:"bytes,8,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *BalanceHistory) Reset() {
	*x = BalanceHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template7_wallet_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceHistory) ProtoMessage() {}

func (x *BalanceHistory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template7_wallet_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceHistory.ProtoReflect.Descriptor instead.
func (*BalanceHistory) Descriptor() ([]byte, []int) {
	return file_proto_template7_wallet_v1_proto_rawDescGZIP(), []int{6}
}

func (x *BalanceHistory) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BalanceHistory) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_unknown
}

func (x *BalanceHistory) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_ntd
}

func (x *BalanceHistory) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *BalanceHistory) GetBalanceBefore() string {
	if x != nil {
		return x.BalanceBefore
	}
	return ""
}

func (x *BalanceHistory) GetBalanceAfter() string {
	if x != nil {
		return x.BalanceAfter
	}
	return ""
}

func (x *BalanceHistory) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BalanceHistory) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
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
	0x22, 0xed, 0x01, 0x0a, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x22, 0x93, 0x02, 0x0a, 0x0e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x2a, 0x2e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x74, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x75,
	0x73, 0x64, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x6a, 0x70, 0x79, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x63, 0x6e, 0x79, 0x10, 0x03, 0x2a, 0x43, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x07,
	0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x63, 0x65, 0x6e, 0x74, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x61, 0x6e, 0x6f, 0x10,
	0x04, 0x12, 0x08, 0x0a, 0x04, 0x70, 0x69, 0x63, 0x6f, 0x10, 0x05, 0x2a, 0x54, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x10, 0x03,
	0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x10,
	0x04, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x37, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x37, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_template7_wallet_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_template7_wallet_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_template7_wallet_v1_proto_goTypes = []interface{}{
	(Currency)(0),                  // 0: wallet.currency
	(Unit)(0),                      // 1: wallet.unit
	(Direction)(0),                 // 2: wallet.direction
	(*Wallet)(nil),                 // 3: wallet.wallet
	(*Balance)(nil),                // 4: wallet.balance
	(*DepositRequest)(nil),         // 5: wallet.depositRequest
	(*WithdrawRequest)(nil),        // 6: wallet.withdrawRequest
	(*TransferRequest)(nil),        // 7: wallet.transferRequest
	(*CurrencyBalanceHistory)(nil), // 8: wallet.currencyBalanceHistory
	(*BalanceHistory)(nil),         // 9: wallet.balanceHistory
}
var file_proto_template7_wallet_v1_proto_depIdxs = []int32{
	4, // 0: wallet.wallet.balances:type_name -> wallet.balance
	0, // 1: wallet.balance.currency:type_name -> wallet.currency
	0, // 2: wallet.depositRequest.currency:type_name -> wallet.currency
	0, // 3: wallet.withdrawRequest.currency:type_name -> wallet.currency
	0, // 4: wallet.transferRequest.currency:type_name -> wallet.currency
	2, // 5: wallet.currencyBalanceHistory.direction:type_name -> wallet.direction
	2, // 6: wallet.balanceHistory.direction:type_name -> wallet.direction
	0, // 7: wallet.balanceHistory.currency:type_name -> wallet.currency
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
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
		file_proto_template7_wallet_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyBalanceHistory); i {
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
		file_proto_template7_wallet_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceHistory); i {
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
			NumEnums:      3,
			NumMessages:   7,
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
