package go_tokenbases

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaka1234/go-tokenbases/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) CreateAddress(req TokenBasesCreateAddressReq) (*AddressContent, error) {

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
	var bodyContent AddressContent
	err = json.Unmarshal([]byte(result.Data.Body), &bodyContent)
	if err != nil {
		return nil, err
	}

	fmt.Printf("url:%s\n", rawURL)
	fmt.Printf("resp:%d, %+v\n", resp.StatusCode(), string(resp.Body()))

	return &bodyContent, nil
}
