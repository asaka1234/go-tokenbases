package go_tokenbases

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) CreateAddress(req TokenBasesCreateAddressReq) (*TokenBasesCreateAddressResp, error) {

	rawURL := cli.Params.CreateAddressUrl

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
	var result TokenBasesCreateAddressResp

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil || resp.StatusCode() != 200 {
		return nil, err
	}

	fmt.Printf("url:%s\n", rawURL)
	fmt.Printf("resp:%d, %+v\n", resp.StatusCode(), string(resp.Body()))

	return &result, nil
}
