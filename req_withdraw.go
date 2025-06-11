package go_tokenbases

/*
// withdraw
func (cli *Client) Withdraw(req TokenBasesWithdrawReq) (*TokenBasesWithdrawResponse, error) {

	rawURL := cli.Params.WithdrawUrl

	jsonData, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["data"] = string(jsonData)
	params["sys_no"] = cli.Params.MerchantId

	//签名
	signStr := utils.SignWithdraw(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result TokenBasesWithdrawResponse

	_, err = cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}

*/
