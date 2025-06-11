package go_tokenbases

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallBack(req TokenBasesDepositCallbackReq, processor func(DepositCallbackBodyContent) error) error {
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

	//开始处理body内容
	return processor(bodyContent)
}

// 返回
func (cli *Client) DepositCallBackResp(code string, msg string) TokenBasesDepositCallbackResp {
	//构造data
	nonce, _ := utils.RandInt32()
	data := DepositCallbackData{
		Body:      "",
		Nonce:     nonce,
		Timestamp: time.Now().Unix(),
	}

	//赋值sign
	var params map[string]interface{}
	mapstructure.Decode(data, &params)
	signStr := utils.Sign(params, cli.Params.AccessKey)
	data.Sign = signStr

	return TokenBasesDepositCallbackResp{
		Errno:  code,
		Errmsg: msg,
		Data:   []DepositCallbackData{data},
	}
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallBack(req TokenBasesWithdrawCallbackReq, processor func(WithdrawCallbackBodyContent) error) error {
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

	//开始处理
	return processor(bodyContent)
}

// code='000'是成功
func (cli *Client) WithdrawCallBackResp(code string, msg string) TokenBasesWithdrawCallbackResp {
	//构造data
	nonce, _ := utils.RandInt32()
	data := TokenBasesWithdrawCallbackDataItem{
		Body:      "",
		Nonce:     nonce,
		Timestamp: time.Now().Unix(),
	}

	//赋值sign
	var params map[string]interface{}
	mapstructure.Decode(data, &params)
	signStr := utils.Sign(params, cli.Params.AccessKey)
	data.Sign = signStr

	return TokenBasesWithdrawCallbackResp{
		Errno:  code,
		Errmsg: msg,
		Data:   []TokenBasesWithdrawCallbackDataItem{data},
	}
}
