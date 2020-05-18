package external

import (
	"gitlab.cosee.biz/pfyl/pfyl-cli/analysis"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) ConsumeSymbolTable(symbolTable []analysis.SymbolTableEntry) error {
	return nil
}


