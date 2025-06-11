package go_tokenbases

type TokenBasesInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`

	CreateAddressUrl string `json:"createAddressUrl" mapstructure:"createAddressUrl" config:"createAddressUrl"  yaml:"createAddressUrl"`
	WithdrawUrl      string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
}

// ----------create address-------------------------

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
	Code    int               `json:"code" mapstructure:"code"`
	Message string            `json:"message" mapstructure:"message"`
	Data    CreateAddressData `json:"data" mapstructure:"data"`
}

type CreateAddressData struct {
	Body      string `json:"body" mapstructure:"body"`
	Nonce     int    `json:"nonce" mapstructure:"nonce"`
	Sign      string `json:"sign" mapstructure:"sign"`
	Timestamp string `json:"timestamp" mapstructure:"timestamp"`
}

// CreateAddressData的Body是如下struct的json字符串
type AddressContent struct {
	Address   []string `json:"address" mapstructure:"address"`
	ChainName string   `json:"chainName" mapstructure:"chainName"`
}

//===========withdraw===================================

type TokenBasesWithdrawReq struct {
	Timestamp int64               `json:"timestamp" mapstructure:"timestamp"`
	Nonce     int32               `json:"nonce" mapstructure:"nonce"`
	Body      WithdrawBodyContent `json:"body" mapstructure:"body"`
	//Sign      string              `json:"sign" mapstructure:"sign"`
}

// 对应上边的body
type WithdrawBodyContent struct {
	Address    string `json:"address" mapstructure:"address"`       //接收地址，必须为有效区块链地址
	Amount     string `json:"amount" mapstructure:"amount"`         //提现金额
	ChainName  string `json:"chainName" mapstructure:"chainName"`   //主链类型，例如 `ETH` 或 `BTC`
	BusinessID string `json:"businessId" mapstructure:"businessId"` //merchant订单号
	Memo       string `json:"memo" mapstructure:"memo"`             //特定链需要
	TokenName  string `json:"tokenName" mapstructure:"tokenName"`   //提现的 Token 名称，例如 `ETH` 或 `USDT`
	//MerchantID string `json:"merchantId" mapstructure:"merchantId"`
}

// 返回response
type TokenBasesWithdrawResp struct {
	Code    int              `json:"code" mapstructure:"code"`
	Message string           `json:"message" mapstructure:"message"`
	Data    WithdrawRespData `json:"data" mapstructure:"data"`
}

type WithdrawRespData struct {
	Body      string `json:"body" mapstructure:"body"` //再解析
	Nonce     int    `json:"nonce" mapstructure:"nonce"`
	Sign      string `json:"sign" mapstructure:"sign"`
	Timestamp string `json:"timestamp" mapstructure:"timestamp"`
}

// WithdrawRespData的body是如下的Json字符串
type WithdrawRespDataBodyContent struct {
	Success bool  `json:"success" mapstructure:"success"` //是否成功
	TransID int64 `json:"transId" mapstructure:"transId"` //提现id
}

// ----------充值 回调-------------------------

type TokenBasesDepositCallbackReq struct {
	Body      string `json:"body" mapstructure:"body"`
	Nonce     int    `json:"nonce" mapstructure:"nonce"`
	Sign      string `json:"sign" mapstructure:"sign"`
	Timestamp int64  `json:"timestamp" mapstructure:"timestamp"`
}

// body的内容
type DepositCallbackBodyContent struct {
	AddressFrom string `json:"addressFrom" mapstructure:"addressFrom"`
	AddressTo   string `json:"addressTo" mapstructure:"addressTo"`
	Amount      string `json:"amount" mapstructure:"amount"`
	ChainName   string `json:"chainName" mapstructure:"chainName"`
	Confirm     int    `json:"confirm" mapstructure:"confirm"` //确认数
	ConfirmTime int64  `json:"confirmTime" mapstructure:"confirmTime"`
	Fee         string `json:"fee" mapstructure:"fee"`
	MerchantID  string `json:"merchantId" mapstructure:"merchantId"`
	Symbol      string `json:"symbol" mapstructure:"symbol"`
	TxID        string `json:"txId" mapstructure:"txId"`
	Type        int    `json:"type" mapstructure:"type"` //交易类型 1：充值，2：提现，3：归集
}

// 这个是返回
type TokenBasesDepositCallbackResp struct {
	Errno  string                `json:"errno" mapstructure:"errno"` //“000“:成功
	Errmsg string                `json:"errmsg" mapstructure:"errmsg"`
	Data   []DepositCallbackData `json:"data" mapstructure:"data"` //响应数据内容(用于验证签名),正常/异常响应都必须保证字段完整性，包括签名信息----适用于所有回调接口
}

// 上述data实际是 []DepositConfirmData
type DepositCallbackData struct {
	Body      string `json:"body" mapstructure:"body"`
	Nonce     int32  `json:"nonce" mapstructure:"nonce"`
	Sign      string `json:"sign" mapstructure:"sign"`
	Timestamp int64  `json:"timestamp" mapstructure:"timestamp"`
}

//========================================================

type TokenBasesWithdrawCallbackReq struct {
	Body      string `json:"body" mapstructure:"body"`
	Nonce     int    `json:"nonce" mapstructure:"nonce"`
	Timestamp int64  `json:"timestamp" mapstructure:"timestamp"`
	Sign      string `json:"sign" mapstructure:"sign"`
}

// 对上上边的Body
type WithdrawCallbackBodyContent struct {
	AddressFrom string `json:"addressFrom" mapstructure:"addressFrom"`
	AddressTo   string `json:"addressTo" mapstructure:"addressTo"`
	TxID        string `json:"txId" mapstructure:"txId"`
	Amount      string `json:"amount" mapstructure:"amount"`
	Confirm     int    `json:"confirm" mapstructure:"confirm"`
	ConfirmTime int64  `json:"confirmTime" mapstructure:"confirmTime"`
	ChainName   string `json:"chainName" mapstructure:"chainName"`
	MerchantID  string `json:"merchantId" mapstructure:"merchantId"`
	Fee         string `json:"fee" mapstructure:"fee"`
	Symbol      string `json:"symbol" mapstructure:"symbol"`
	Type        int    `json:"type" mapstructure:"type"`
	BusinessID  string `json:"businessId" mapstructure:"businessId"` //业务id
	TransID     int64  `json:"transId" mapstructure:"transId"`
}

// response
type TokenBasesWithdrawCallbackResp struct {
	Errno  string                               `json:"errno" mapstructure:"errno"`
	Errmsg string                               `json:"errmsg" mapstructure:"errmsg"`
	Data   []TokenBasesWithdrawCallbackDataItem `json:"data" mapstructure:"data"`
}

type TokenBasesWithdrawCallbackDataItem struct {
	Body      string `json:"body" mapstructure:"body"`
	Nonce     int32  `json:"nonce" mapstructure:"nonce"`
	Sign      string `json:"sign" mapstructure:"sign"`
	Timestamp int64  `json:"timestamp" mapstructure:"timestamp"`
}
