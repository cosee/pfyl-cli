package external

import (
	"gitlab.cosee.biz/pfyl/pfyl-cli/analysis"
	"gitlab.cosee.biz/pfyl/pfyl-cli/configuration"
)

type Client struct {
	config configuration.Configuration
}

func NewClient(config configuration.Configuration) *Client {
	return &Client{config: config}
}

func (c *Client) ConsumeSymbolTable(symbolTable []analysis.SymbolTableEntry) error {
	return nil
}
