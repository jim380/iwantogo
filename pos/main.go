package pkg

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type posExecutor interface {
	getValidatorInfo(key string, c *websocket.Conn)
}

type posParamsPreSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type posMessagePreSign struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  posParamsPreSign `json:"params"`
	ID      int64            `json:"id"`
}

type posParamsPostSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type posMessagePostSign struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  posParamsPostSign `json:"params"`
	ID      int64             `json:"id"`
}

// NewReq instantiates a new RPC-JSON call
func NewReq(addr string) *posMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &posMessagePreSign{
		JSONRPC: "2.0",
		Params: posParamsPreSign{
			Address:   addr,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *posMessagePreSign) getValidatorInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorInfo"
	sig := common.GenSig(m, key)
	msgSend := &posMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorInfo",
		Params: posParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetValidatorInfo returns the info of a specific validator account
func GetValidatorInfo(pe posExecutor, k string, c *websocket.Conn) {
	pe.getValidatorInfo(k, c)
}
