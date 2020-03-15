package accounts

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type accountMultiExecutor interface {
	getMultiBalances(key string, c *websocket.Conn)
}

type accountMultiParamsPreSign struct {
	Address   []string `json:"address"`
	ChainType string   `json:"chainType"`
	Timestamp string   `json:"timestamp"`
}

type accountMultiMessagePreSign struct {
	JSONRPC string                    `json:"jsonrpc"`
	Method  string                    `json:"method"`
	Params  accountMultiParamsPreSign `json:"params"`
	ID      int64                     `json:"id"`
}

type accountMultiParamsPostSign struct {
	Address   []string `json:"address"`
	ChainType string   `json:"chainType"`
	Timestamp string   `json:"timestamp"`
	Signature string   `json:"signature"`
}

type accountMultiMessagePostSign struct {
	JSONRPC string                     `json:"jsonrpc"`
	Method  string                     `json:"method"`
	Params  accountMultiParamsPostSign `json:"params"`
	ID      int64                      `json:"id"`
}

// NewReqMulti instantiates a new RPC-JSON call
func NewReqMulti(addr []string) *accountMultiMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &accountMultiMessagePreSign{
		JSONRPC: "2.0",
		// Method:  "getBalance",
		Params: accountMultiParamsPreSign{
			Address:   addr,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *accountMultiMessagePreSign) getMultiBalances(key string, c *websocket.Conn) {
	m.Method = "getMultiBalances"
	sig := common.GenSig(m, key)
	msgSend := &accountMultiMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getMultiBalances",
		Params: accountMultiParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetMultiBalances returns balances of multiple addresses provided
func GetMultiBalances(ae accountMultiExecutor, k string, c *websocket.Conn) {
	ae.getMultiBalances(k, c)
}
