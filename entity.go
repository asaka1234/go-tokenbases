package go_tokenbases

type TokenBasesInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`

	CreateAddressUrl string `json:"createAddressUrl" mapstructure:"createAddressUrl" config:"createAddressUrl"  yaml:"createAddressUrl"`
	WithdrawUrl      string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
}

// ----------pre order-------------------------

type TokenBasesCreateAddressReq struct {
	Timestamp int64             `json:"timestamp" mapstructure:"timestamp"`
	Nonce     int32             `json:"nonce" mapstructure:"nonce"`
	Body      CreateAddressBody `json:"body" mapstructure:"body"`
	//sdk赋值
	//Sign string `json:"sign" mapstructure:"sign"`
}

type CreateAddressBody struct {
	ChainName string `json:"chainName" mapstructure:"chainName"`
	Count     int    `json:"count" mapstructure:"count"` //TODO 实际没用,因为每次只返回唯一一个.
	//sdk赋值
	//MerchantID string `json:"merchantId" mapstructure:"merchantId"`
}

type TokenBasesCreateAddressResp struct {
	Code    int    `json:"code" mapstructure:"code"`
	Message string `json:"message" mapstructure:"message"`
	Data    Data   `json:"data" mapstructure:"data"`
}

type Data struct {
	Body      string `json:"body" mapstructure:"body"`
	Nonce     int    `json:"nonce" mapstructure:"nonce"`
	Sign      string `json:"sign" mapstructure:"sign"`
	Timestamp int64  `json:"timestamp" mapstructure:"timestamp"`
}

// If you need to parse the body content separately
type AddressContent struct {
	Address   []string `json:"address" mapstructure:"address"`
	ChainName string   `json:"chainName" mapstructure:"chainName"`
}

//=========

type TokenBasesDepositReq struct {
	OrderId     string `json:"order_id" mapstructure:"order_id"`           // 订单ID
	OrderAmount string `json:"order_amount" mapstructure:"order_amount"`   // 订单金额
	UserId      string `json:"user_id" mapstructure:"user_id"`             // 用户ID
	OrderIp     string `json:"order_ip" mapstructure:"order_ip"`           // 订单IP
	PayUserName string `json:"pay_user_name" mapstructure:"pay_user_name"` // 付款人姓名
	//这个是TokenBases给分配的商户id
	//SysNo string `json:"sys_no" mapstructure:"sys_no"` // 系统编号
	//这个让sdk来赋值
	//OrderTime string `json:"order_time" mapstructure:"order_time"` // 订单时间 格式:yyyy-MM-dd HH:mm:ss
}

// 不管是正确/失败的通用字段返回
type TokenBasesDepositCommonResponse struct {
	Code   int    `json:"code"`   // 111 是正确
	Status string `json:"status"` //success 是正确
	Msg    string `json:"msg"`
}

type TokenBasesDepositResponse struct {
	Code   int                            `json:"code" mapstructure:"code"`     // 111 是正确
	Status string                         `json:"status" mapstructure:"status"` //success 是正确
	Msg    string                         `json:"msg" mapstructure:"msg"`
	Data   *TokenBasesDepositResponseData `json:"data" mapstructure:"data"`
}

type TokenBasesDepositResponseData struct {
	OrderNo string `json:"order_no" mapstructure:"order_no"` // 订单编号
	SendUrl string `json:"send_url" mapstructure:"send_url"` // 发送URL
	UserId  string `json:"user_id" mapstructure:"user_id"`   // 用户ID
}

// ------------------------------------------------------------
type TokenBasesDepositCancelBackReq struct {
	BillNo     string `json:"bill_no" mapstructure:"bill_no"`         // 唯一订单号，商户下单时传过来的order_id
	BillStatus int    `json:"bill_status" mapstructure:"bill_status"` // 订单状态：1=订单已取消；2=订单已激活
	SysNo      string `json:"sys_no" mapstructure:"sys_no"`           // 商户编号
	Sign       string `json:"sign" mapstructure:"sign"`               // 签名，参照验签规范
}

type TokenBasesDepositSucceedBackReq struct {
	BillNo string `json:"bill_no" mapstructure:"bill_no"` // 必须包含订单号
	Amount string `json:"amount" mapstructure:"amount"`   // 必须是数字字符串
	SysNo  string `json:"sys_no" mapstructure:"sys_no"`   // 必须包含商户号
	Sign   string `json:"sign" mapstructure:"sign"`       // 必须包含签名
}

//===========withdraw===================================

type TokenBasesWithdrawReq struct {
	Data []TokenBasesWithdrawData `json:"data" mapstructure:"data"` // 申请sys_no唯一标识 610001
	//这个是TokenBases给分配的商户id ,sdk来赋值
	//SysNo string `json:"sys_no" mapstructure:"sys_no"` // 申请sys_no唯一标识 610001号
}

type TokenBasesWithdrawData struct {
	UserName    string `json:"user_name" mapstructure:"user_name"`       // 真实姓名
	BankCardNo  string `json:"bankcard_no" mapstructure:"bankcard_no"`   // 卡号
	SerialNo    string `json:"serial_no" mapstructure:"serial_no"`       // 订单号
	BankAddress string `json:"bank_address" mapstructure:"bank_address"` // 支行地址
	Amount      string `json:"amount" mapstructure:"amount"`             // 金额
}

type TokenBasesWithdrawResponse struct {
	Code int    `json:"code"` //200是成功
	Msg  string `json:"msg"`
}

type TokenBasesWithdrawCancelBackReq struct {
	BillNo     string `json:"bill_no" mapstructure:"bill_no"`         // 唯一订单号，商户下单时传过来的order_id
	BillStatus int    `json:"bill_status" mapstructure:"bill_status"` // 订单状态：1=订单已取消；2=订单已激活
	SysNo      string `json:"sys_no" mapstructure:"sys_no"`           // 商户编号
	Sign       string `json:"sign" mapstructure:"sign"`               // 签名，参照验签规范
}

type TokenBasesWithdrawSucceedBackReq struct {
	BillNo string `json:"bill_no" mapstructure:"bill_no"` // 唯一订单号，商户下单时传过来的order_id
	Amount string `json:"amount" mapstructure:"amount"`   //订单金额
	SysNo  string `json:"sys_no" mapstructure:"sys_no"`   //商户编号
	Sign   string `json:"sign" mapstructure:"sign"`       //签名，参照验签规范
}

// ----------withdraw confirm-------------------------

// callback以后,还要单独发个请求再来查询下.
type TokenBasesWithdrawConfirmReq struct {
	Ids string `json:"ids" mapstructure:"ids"` //确认收款订单列表接口中获取的id，用英文逗号“,”拼接起来
	//这个是TokenBases给分配的商户id ,sdk来赋值
	//SysNo string `json:"sys_no" mapstructure:"sys_no"` // 申请sys_no唯一标识 610001号
}

type TokenBasesWithdrawConfirmResponse struct {
	Code string `json:"code"` //
	Msg  string `json:"msg"`
}

// =================单独请求===============================

type TokenBasesOrderListRsp struct {
	Code   string                    `json:"code"` //
	Msg    string                    `json:"msg"`
	Result TokenBasesOrderPageResult `json:"result"`
}

type TokenBasesOrderPageResult struct {
	TotalCount string                 `json:"totalCount"` // 总记录数
	TotalPage  int64                  `json:"totalPage"`  // 总页数
	Page       int64                  `json:"page"`       // 当前页码
	Data       []*TokenBasesOrderData `json:"data"`       // 订单数据列表
}

type TokenBasesOrderData struct {
	ID                       string `json:"id"`
	SysSerialNo              string `json:"sysSerialNo"`
	Amount                   string `json:"amount"`
	PayType                  string `json:"payType"`
	UserName                 string `json:"userName"`
	BankCardNo               string `json:"bankCardNo"`
	BankAddress              string `json:"bankAddress"`
	ChangeRate               string `json:"changeRate"`
	HandlingFee              string `json:"handlingFee"`
	MerchantSettleUSDTNumber string `json:"merchantSettleUSDTNumber"`
	SerialNo                 string `json:"serialNo"`
	CreateTime               string `json:"createTime"`
	Remark                   string `json:"remark"`
	NumRow                   string `json:"numRow"`
	StatusName               string `json:"statusName"`
}
