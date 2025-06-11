package go_tokenbases

import (
	"fmt"
	"github.com/asaka1234/go-tokenbases/utils"
	"testing"
	"time"
)

func TestWithdraw(t *testing.T) {

	//构造client
	cli := NewClient(nil, &TokenBasesInitParams{MERCHANT_ID, ACCESS_KEY, CreateAddressUrl, WITHDRAW_URL})
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

	nonce, _ := utils.RandInt32()

	return TokenBasesWithdrawReq{
		Timestamp: time.Now().Unix(),
		Nonce:     nonce,
		Body: WithdrawBodyContent{
			Address:    "TYBFrHn2AV6V1Ce8kZsXbRfsvk8cpZNmXV",
			Amount:     "2.1",
			ChainName:  "TRX", //商户uid
			BusinessID: "123",
			Memo:       "",
			TokenName:  "TUSDT",
		},
	}
}
