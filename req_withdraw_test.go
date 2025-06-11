package go_tokenbases

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {

	vlog := VLog{}
	//构造client
	cli := NewClient(vlog, &TokenBasesInitParams{MERCHANT_ID, ACCESS_KEY, BASE_URL})
	cli.SetDebugModel(true)

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() TokenBasesWithdrawReq {

	return TokenBasesWithdrawReq{
		Address:    "TYBFrHn2AV6V1Ce8kZsXbRfsvk8cpZNmXV",
		Amount:     "2.1",
		ChainName:  "TRX", //商户uid
		BusinessID: "123",
		Memo:       "",
		TokenName:  "TUSDT",
	}
}
