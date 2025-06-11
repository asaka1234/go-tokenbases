package go_tokenbases

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
)

// withdraw
func (cli *Client) Withdraw(req TokenBasesWithdrawReq) (*TokenBasesWithdrawResp, error) {

	rawURL := cli.Params.WithdrawUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	body := params["body"].(map[string]interface{})
	body["merchantId"] = cli.Params.MerchantId

	bd, _ := json.Marshal(body)
	params["body"] = string(bd)

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result TokenBasesWithdrawResp

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil || resp.StatusCode() != 200 {
		return nil, err
	}

	fmt.Printf("url:%s\n", rawURL)
	fmt.Printf("resp:%d, %+v\n", resp.StatusCode(), string(resp.Body()))

	return &result, err
}
