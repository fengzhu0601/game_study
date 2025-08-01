// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: pb_cs_error.proto

package pb_cs_error

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 错误码,前1000给服务器内部用
type Code int32

const (
	Code_ServerNotExit                Code = 0    // 服务器不存在
	Code_ServerNotConn                Code = 1    // 服务器未连接
	Code_RpcTimeout                   Code = 2    // rpc访问超时
	Code_RpcServerError               Code = 3    // 目标服发生错误
	Code_RpcLocalNotRegist            Code = 4    // 本地Rpc消息没注册
	Code_RpcNotRegist                 Code = 5    // 远程Rpc消息没注册
	Code_RpcAccExit                   Code = 6    // 角色已存在，无法创建角色
	Code_ServerWorkerNotExit          Code = 7    // 进程不存在
	Code_ServerCantFind               Code = 8    // 没有可用服务器
	Code_ServerAccNotExit             Code = 9    // 角色不存在
	Code_ServerError                  Code = 1001 // 服务器内部错误:%1s
	Code_MsgError                     Code = 1002 // 未定义cmd:%1i
	Code_NameEmptyError               Code = 1003 // 昵称为空
	Code_NameInvalidLenError          Code = 1004 // 无效的昵称长度
	Code_NameInvalidError             Code = 1005 // 昵称存在非法字符
	Code_NameUsedError                Code = 1006 // 该昵称已被使用
	Code_ItemNotEnough                Code = 1007 // %1s不足
	Code_ItemUseIlleganotEnoughlError Code = 1008 // 道具非法使用
	Code_ConfigError                  Code = 1009 // 配置错误
	Code_RewardUnmetError             Code = 1010 // 未满足领取条件
	Code_RewardReceivedError          Code = 1011 // 该奖励已领取
	Code_MailNotExit                  Code = 1141 // 邮件不存在
	Code_ChargeActTypeClose           Code = 1151 // 该活动类型未开放充值
	Code_ChargeProductNull            Code = 1152 // 商品不存在
	Code_ChargeClose                  Code = 1153 // 充值入口未开放
	Code_ChargeGenOrderErr            Code = 1154 // 生成订单号失败
	Code_ChargeGetOrderErr            Code = 1155 // 获取订单号失败
	Code_ChargeNotExit                Code = 1156 // 订单不存在
	Code_ChargeHadDeliver             Code = 1157 // 订单已发货
	Code_ChargeInfoNotMatch           Code = 1158 // 订单核对失败
	Code_ChargeProductNotExit         Code = 1159 // 商品已过期
	Code_ChargeGearError              Code = 1160 // 档位配置不存在
	Code_GiftCodeError                Code = 1171 // 兑换码错误
	Code_FirstChargeStatusErr         Code = 1181 // 没有可领取的首充奖励
	Code_ShopNull                     Code = 1300 // 商店未开启
	Code_ShopTimeFreeLimit            Code = 1301 // 免费次数不足
	Code_ShopTimeVideoLimit           Code = 1302 // 广告次数不足
	Code_ShopTimeBuyLimit             Code = 1303 // 购买次数不足
	Code_ShopTimeGoodNull             Code = 1304 // 商品不存在
	Code_Plant                        Code = 1305 // 商品不存在
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:    "ServerNotExit",
		1:    "ServerNotConn",
		2:    "RpcTimeout",
		3:    "RpcServerError",
		4:    "RpcLocalNotRegist",
		5:    "RpcNotRegist",
		6:    "RpcAccExit",
		7:    "ServerWorkerNotExit",
		8:    "ServerCantFind",
		9:    "ServerAccNotExit",
		1001: "ServerError",
		1002: "MsgError",
		1003: "NameEmptyError",
		1004: "NameInvalidLenError",
		1005: "NameInvalidError",
		1006: "NameUsedError",
		1007: "ItemNotEnough",
		1008: "ItemUseIlleganotEnoughlError",
		1009: "ConfigError",
		1010: "RewardUnmetError",
		1011: "RewardReceivedError",
		1141: "MailNotExit",
		1151: "ChargeActTypeClose",
		1152: "ChargeProductNull",
		1153: "ChargeClose",
		1154: "ChargeGenOrderErr",
		1155: "ChargeGetOrderErr",
		1156: "ChargeNotExit",
		1157: "ChargeHadDeliver",
		1158: "ChargeInfoNotMatch",
		1159: "ChargeProductNotExit",
		1160: "ChargeGearError",
		1171: "GiftCodeError",
		1181: "FirstChargeStatusErr",
		1300: "ShopNull",
		1301: "ShopTimeFreeLimit",
		1302: "ShopTimeVideoLimit",
		1303: "ShopTimeBuyLimit",
		1304: "ShopTimeGoodNull",
		1305: "Plant",
	}
	Code_value = map[string]int32{
		"ServerNotExit":                0,
		"ServerNotConn":                1,
		"RpcTimeout":                   2,
		"RpcServerError":               3,
		"RpcLocalNotRegist":            4,
		"RpcNotRegist":                 5,
		"RpcAccExit":                   6,
		"ServerWorkerNotExit":          7,
		"ServerCantFind":               8,
		"ServerAccNotExit":             9,
		"ServerError":                  1001,
		"MsgError":                     1002,
		"NameEmptyError":               1003,
		"NameInvalidLenError":          1004,
		"NameInvalidError":             1005,
		"NameUsedError":                1006,
		"ItemNotEnough":                1007,
		"ItemUseIlleganotEnoughlError": 1008,
		"ConfigError":                  1009,
		"RewardUnmetError":             1010,
		"RewardReceivedError":          1011,
		"MailNotExit":                  1141,
		"ChargeActTypeClose":           1151,
		"ChargeProductNull":            1152,
		"ChargeClose":                  1153,
		"ChargeGenOrderErr":            1154,
		"ChargeGetOrderErr":            1155,
		"ChargeNotExit":                1156,
		"ChargeHadDeliver":             1157,
		"ChargeInfoNotMatch":           1158,
		"ChargeProductNotExit":         1159,
		"ChargeGearError":              1160,
		"GiftCodeError":                1171,
		"FirstChargeStatusErr":         1181,
		"ShopNull":                     1300,
		"ShopTimeFreeLimit":            1301,
		"ShopTimeVideoLimit":           1302,
		"ShopTimeBuyLimit":             1303,
		"ShopTimeGoodNull":             1304,
		"Plant":                        1305,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_cs_error_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_pb_cs_error_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_pb_cs_error_proto_rawDescGZIP(), []int{0}
}

// 错误码提示
type RepError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          Code                   `protobuf:"varint,1,opt,name=Code,proto3,enum=pb_cs_error.Code" json:"Code,omitempty"` // 错误码
	ArgsIntList   []uint32               `protobuf:"varint,2,rep,packed,name=ArgsIntList,proto3" json:"ArgsIntList,omitempty"`  // int参数列表
	ArgsStrList   []string               `protobuf:"bytes,3,rep,name=ArgsStrList,proto3" json:"ArgsStrList,omitempty"`          // string参数列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RepError) Reset() {
	*x = RepError{}
	mi := &file_pb_cs_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RepError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepError) ProtoMessage() {}

func (x *RepError) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepError.ProtoReflect.Descriptor instead.
func (*RepError) Descriptor() ([]byte, []int) {
	return file_pb_cs_error_proto_rawDescGZIP(), []int{0}
}

func (x *RepError) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_ServerNotExit
}

func (x *RepError) GetArgsIntList() []uint32 {
	if x != nil {
		return x.ArgsIntList
	}
	return nil
}

func (x *RepError) GetArgsStrList() []string {
	if x != nil {
		return x.ArgsStrList
	}
	return nil
}

var File_pb_cs_error_proto protoreflect.FileDescriptor

const file_pb_cs_error_proto_rawDesc = "" +
	"\n" +
	"\x11pb_cs_error.proto\x12\vpb_cs_error\"u\n" +
	"\bRepError\x12%\n" +
	"\x04Code\x18\x01 \x01(\x0e2\x11.pb_cs_error.CodeR\x04Code\x12 \n" +
	"\vArgsIntList\x18\x02 \x03(\rR\vArgsIntList\x12 \n" +
	"\vArgsStrList\x18\x03 \x03(\tR\vArgsStrList*\xe4\x06\n" +
	"\x04Code\x12\x11\n" +
	"\rServerNotExit\x10\x00\x12\x11\n" +
	"\rServerNotConn\x10\x01\x12\x0e\n" +
	"\n" +
	"RpcTimeout\x10\x02\x12\x12\n" +
	"\x0eRpcServerError\x10\x03\x12\x15\n" +
	"\x11RpcLocalNotRegist\x10\x04\x12\x10\n" +
	"\fRpcNotRegist\x10\x05\x12\x0e\n" +
	"\n" +
	"RpcAccExit\x10\x06\x12\x17\n" +
	"\x13ServerWorkerNotExit\x10\a\x12\x12\n" +
	"\x0eServerCantFind\x10\b\x12\x14\n" +
	"\x10ServerAccNotExit\x10\t\x12\x10\n" +
	"\vServerError\x10\xe9\a\x12\r\n" +
	"\bMsgError\x10\xea\a\x12\x13\n" +
	"\x0eNameEmptyError\x10\xeb\a\x12\x18\n" +
	"\x13NameInvalidLenError\x10\xec\a\x12\x15\n" +
	"\x10NameInvalidError\x10\xed\a\x12\x12\n" +
	"\rNameUsedError\x10\xee\a\x12\x12\n" +
	"\rItemNotEnough\x10\xef\a\x12!\n" +
	"\x1cItemUseIlleganotEnoughlError\x10\xf0\a\x12\x10\n" +
	"\vConfigError\x10\xf1\a\x12\x15\n" +
	"\x10RewardUnmetError\x10\xf2\a\x12\x18\n" +
	"\x13RewardReceivedError\x10\xf3\a\x12\x10\n" +
	"\vMailNotExit\x10\xf5\b\x12\x17\n" +
	"\x12ChargeActTypeClose\x10\xff\b\x12\x16\n" +
	"\x11ChargeProductNull\x10\x80\t\x12\x10\n" +
	"\vChargeClose\x10\x81\t\x12\x16\n" +
	"\x11ChargeGenOrderErr\x10\x82\t\x12\x16\n" +
	"\x11ChargeGetOrderErr\x10\x83\t\x12\x12\n" +
	"\rChargeNotExit\x10\x84\t\x12\x15\n" +
	"\x10ChargeHadDeliver\x10\x85\t\x12\x17\n" +
	"\x12ChargeInfoNotMatch\x10\x86\t\x12\x19\n" +
	"\x14ChargeProductNotExit\x10\x87\t\x12\x14\n" +
	"\x0fChargeGearError\x10\x88\t\x12\x12\n" +
	"\rGiftCodeError\x10\x93\t\x12\x19\n" +
	"\x14FirstChargeStatusErr\x10\x9d\t\x12\r\n" +
	"\bShopNull\x10\x94\n" +
	"\x12\x16\n" +
	"\x11ShopTimeFreeLimit\x10\x95\n" +
	"\x12\x17\n" +
	"\x12ShopTimeVideoLimit\x10\x96\n" +
	"\x12\x15\n" +
	"\x10ShopTimeBuyLimit\x10\x97\n" +
	"\x12\x15\n" +
	"\x10ShopTimeGoodNull\x10\x98\n" +
	"\x12\n" +
	"\n" +
	"\x05Plant\x10\x99\n" +
	"B\x18Z\x16game/pb/cs/pb_cs_errorb\x06proto3"

var (
	file_pb_cs_error_proto_rawDescOnce sync.Once
	file_pb_cs_error_proto_rawDescData []byte
)

func file_pb_cs_error_proto_rawDescGZIP() []byte {
	file_pb_cs_error_proto_rawDescOnce.Do(func() {
		file_pb_cs_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pb_cs_error_proto_rawDesc), len(file_pb_cs_error_proto_rawDesc)))
	})
	return file_pb_cs_error_proto_rawDescData
}

var file_pb_cs_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_cs_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_cs_error_proto_goTypes = []any{
	(Code)(0),        // 0: pb_cs_error.Code
	(*RepError)(nil), // 1: pb_cs_error.RepError
}
var file_pb_cs_error_proto_depIdxs = []int32{
	0, // 0: pb_cs_error.RepError.Code:type_name -> pb_cs_error.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_cs_error_proto_init() }
func file_pb_cs_error_proto_init() {
	if File_pb_cs_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pb_cs_error_proto_rawDesc), len(file_pb_cs_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_cs_error_proto_goTypes,
		DependencyIndexes: file_pb_cs_error_proto_depIdxs,
		EnumInfos:         file_pb_cs_error_proto_enumTypes,
		MessageInfos:      file_pb_cs_error_proto_msgTypes,
	}.Build()
	File_pb_cs_error_proto = out.File
	file_pb_cs_error_proto_goTypes = nil
	file_pb_cs_error_proto_depIdxs = nil
}
