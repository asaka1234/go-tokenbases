package go_tokenbases

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallBack(req TokenBasesCallbackReq, processor func(DepositCallbackBodyContent) error) error {
	//1. 验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//2. 解析body
	var bodyContent DepositCallbackBodyContent
	err := json.Unmarshal([]byte(req.Body), &bodyContent)
	if err != nil {
		return err
	}

	if bodyContent.MerchantID != cli.Params.MerchantId {
		return errors.New("merchantId is unmatch!")
	}

	// 交易类型 1：充值，2：提现，3：归集
	if bodyContent.Type != 1 {
		return errors.New("type is unmatch!")
	}

	//开始处理body内容
	return processor(bodyContent)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallBack(req TokenBasesCallbackReq, processor func(WithdrawCallbackBodyContent) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//2. 解析body
	var bodyContent WithdrawCallbackBodyContent
	err := json.Unmarshal([]byte(req.Body), &bodyContent)
	if err != nil {
		return err
	}

	if bodyContent.MerchantID != cli.Params.MerchantId {
		return errors.New("merchantId is unmatch!")
	}

	// 交易类型 1：充值，2：提现，3：归集，4：提现失败
	if bodyContent.Type != 2 && bodyContent.Type != 4 {
		return errors.New("type is unmatch!")
	}

	//开始处理
	return processor(bodyContent)
}

//-----------callback的resp========

// code='000'是成功
func (cli *Client) CallBackResp(code string, msg string) TokenBasesCallbackResp {
	//构造data
	nonce, _ := utils.RandInt32()
	data := CallbackData{
		Body:      "",
		Nonce:     nonce,
		Timestamp: time.Now().Unix(),
	}

	//赋值sign
	var params map[string]interface{}
	mapstructure.Decode(data, &params)
	signStr := utils.Sign(params, cli.Params.AccessKey)
	data.Sign = signStr

	return TokenBasesCallbackResp{
		Errno:  code,
		Errmsg: msg,
		Data:   []CallbackData{data},
	}
}
