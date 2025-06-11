package go_tokenbases

import (
	"fmt"
	"github.com/asaka1234/go-tokenbases/utils"
	"testing"
	"time"
)

func TestCreateAddress(t *testing.T) {

	//构造client
	cli := NewClient(nil, &TokenBasesInitParams{MERCHANT_ID, ACCESS_KEY, CreateAddressUrl, WITHDRAW_URL})
	cli.SetDebugModel(true)
	//发请求
	resp, err := cli.CreateAddress(GenCreateAddressDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenCreateAddressDemo() TokenBasesCreateAddressReq {

	nonce, _ := utils.RandInt32()

	return TokenBasesCreateAddressReq{
		Timestamp: time.Now().Unix(),
		Nonce:     nonce,
		Body: CreateAddressBody{
			ChainName: "TRX", //商户uid
			Count:     5,
		},
	}
}
