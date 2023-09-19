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
	Currency_NTD Currency = 0
	Currency_USD Currency = 1
	Currency_JPY Currency = 2
	Currency_CNY Currency = 3
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "NTD",
		1: "USD",
		2: "JPY",
		3: "CNY",
	}
	Currency_value = map[string]int32{
		"NTD": 0,
		"USD": 1,
		"JPY": 2,
		"CNY": 3,
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
	Unit_One   Unit = 0
	Unit_Cent  Unit = 1
	Unit_Milli Unit = 2
	Unit_Micro Unit = 3
	Unit_Nano  Unit = 4
	Unit_Pico  Unit = 5
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "One",
		1: "Cent",
		2: "Milli",
		3: "Micro",
		4: "Nano",
		5: "Pico",
	}
	Unit_value = map[string]int32{
		"One":   0,
		"Cent":  1,
		"Milli": 2,
		"Micro": 3,
		"Nano":  4,
		"Pico":  5,
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

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return Currency_NTD
}

func (x *Balance) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_proto_template7_wallet_v1_proto protoreflect.FileDescriptor

var file_proto_template7_wallet_v1_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x37, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x18, 0x0a, 0x06, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x2e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x07, 0x0a, 0x03, 0x4e, 0x54, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x53, 0x44,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x59, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x43,
	0x4e, 0x59, 0x10, 0x03, 0x2a, 0x43, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x61, 0x6e, 0x6f, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x69, 0x63, 0x6f, 0x10, 0x05, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x37, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x37, 0x2f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_proto_template7_wallet_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_template7_wallet_v1_proto_goTypes = []interface{}{
	(Currency)(0),   // 0: wallet.currency
	(Unit)(0),       // 1: wallet.unit
	(*Wallet)(nil),  // 2: wallet.wallet
	(*Balance)(nil), // 3: wallet.balance
}
var file_proto_template7_wallet_v1_proto_depIdxs = []int32{
	0, // 0: wallet.balance.currency:type_name -> wallet.currency
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_template7_wallet_v1_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
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