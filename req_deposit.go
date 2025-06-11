package go_tokenbases

/*

// pre-order
func (cli *Client) Deposit(req TokenBasesDepositReq) (*TokenBasesDepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["sys_no"] = cli.Params.MerchantId
	params["order_time"] = time.Now().Format("2006-01-02 15:04:05")

	//签名
	signStr := utils.SignDeposit(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result TokenBasesDepositCommonResponse

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
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

	//------------------------------------------------------
	if result.Code == 111 && result.Status == "success" {
		//说明成功

		//step-1
		var data map[string]interface{}
		if err := json.Unmarshal(resp2.Body(), &data); err != nil {
			return nil, err
		}

		//step-2
		var resp3 TokenBasesDepositResponse
		if err := mapstructure.Decode(data, &resp3); err != nil {
			return nil, err
		}

		return &resp3, nil
	}

	return &TokenBasesDepositResponse{
		Code:   result.Code,
		Status: result.Status,
		Msg:    result.Msg,
	}, fmt.Errorf("result is failed")
}

*/
