package go_buy365

import (
	"github.com/asaka1234/go-buy365/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *Buy365InitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *Buy365InitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debugMode bool) {
	cli.debugMode = debugMode
}
