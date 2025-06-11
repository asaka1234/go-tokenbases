package go_tokenbases

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"testing"
)

func TestDeposit(t *testing.T) {
	req := TokenBasesCreateAddressReq{
		Timestamp: 111,
		Nonce:     222,
		Body: CreateAddressBody{
			ChainName: "aaa",
			Count:     20,
		},
	}

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["body"].(map[string]interface{})["cc"] = 1

	fmt.Printf("==>%+v\n", params)
}
