package go_tokenbases

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// withdraw
func (cli *Client) Withdraw(req TokenBasesWithdrawReq) (*WithdrawRespDataBodyContent, error) {

	rawURL := cli.Params.BaseUrl + "/mch/withdraw"

	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	params["merchantId"] = cli.Params.MerchantId
	bd, _ := json.Marshal(params)

	//构造最终请求
	nonce, _ := utils.RandInt32()
	request := TokenBasesReq{
		Timestamp: time.Now().Unix(),
		Nonce:     nonce,
		Body:      string(bd),
	}

	//----------------------------

	var params3 map[string]interface{}
	mapstructure.Decode(request, &params3)

	//签名
	signStr := utils.Sign(params3, cli.Params.AccessKey)
	params3["sign"] = signStr

	//返回值会放到这里
	var result TokenBasesWithdrawResp

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params3).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil || resp.StatusCode() != 200 || result.Code != 200 {
		if err == nil {
			err = fmt.Errorf("statusCode:%d, code:%d", resp.StatusCode(), result.Code)
		}
		return nil, err
	}

	//验证签名
	var params2 map[string]interface{}
	mapstructure.Decode(result.Data, &params2)
	verifyResult := utils.VerifySign(params2, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return nil, errors.New("verify sign error!")
	}

	//解析body
	var bodyContent WithdrawRespDataBodyContent
	err = json.Unmarshal([]byte(result.Data.Body), &bodyContent)
	if err != nil {
		return nil, err
	}

	return &bodyContent, err
}
