package go_tokenbases

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCreateAddress(t *testing.T) {

	vlog := VLog{}
	//构造client
	cli := NewClient(vlog, &TokenBasesInitParams{MERCHANT_ID, ACCESS_KEY, BASE_URL})
	cli.SetDebugModel(true)
	//发请求
	resp, err := cli.CreateAddress(GenCreateAddressDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("final-resp:%+v\n", resp)
}

func GenCreateAddressDemo() TokenBasesCreateAddressReq {

	return TokenBasesCreateAddressReq{
		ChainName: "TRX", //商户uid
		Count:     5,
	}
}
