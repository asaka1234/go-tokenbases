package go_tokenbases

import (
	"errors"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallBack(req TokenBasesDepositCallbackReq, processor func(TokenBasesDepositCallbackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallBack(req TokenBasesWithdrawCallbackReq, processor func(TokenBasesWithdrawCallbackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}
