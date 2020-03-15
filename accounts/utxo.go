package accounts

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type accountUTXOExecutor interface {
	getUTXO(key string, c *websocket.Conn)
}

type accountUTXOParamsPreSign struct {
	Address   []string `json:"address"`
	Minconf   int      `json:"minconf"`
	Maxconf   int      `json:"maxconf"`
	ChainType string   `json:"chainType"`
	Timestamp string   `json:"timestamp"`
}

type accountUTXOMessagePreSign struct {
	JSONRPC string                   `json:"jsonrpc"`
	Method  string                   `json:"method"`
	Params  accountUTXOParamsPreSign `json:"params"`
	ID      int                      `json:"id"`
}

type accountUTXOParamsPostSign struct {
	Address   []string `json:"address"`
	Minconf   int      `json:"minconf"`
	Maxconf   int      `json:"maxconf"`
	ChainType string   `json:"chainType"`
	Timestamp string   `json:"timestamp"`
	Signature string   `json:"signature"`
}

type accountUTXOMessagePostSign struct {
	JSONRPC string                    `json:"jsonrpc"`
	Method  string                    `json:"method"`
	Params  accountUTXOParamsPostSign `json:"params"`
	ID      int                       `json:"id"`
}

// NewReqUTXO instantiates a new RPC-JSON call
func NewReqUTXO(addr []string, minc int, maxc int) *accountUTXOMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &accountUTXOMessagePreSign{
		JSONRPC: "2.0",
		Params: accountUTXOParamsPreSign{
			Address:   addr,
			Minconf:   minc,
			Maxconf:   maxc,
			ChainType: "BTC",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *accountUTXOMessagePreSign) getUTXO(key string, c *websocket.Conn) {
	m.Method = "getUTXO"
	// sig := m.common.GenSig(key)
	sig := common.GenSig(m, key)
	msgSend := &accountUTXOMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getUTXO",
		Params: accountUTXOParamsPostSign{
			Address:   m.Params.Address,
			Minconf:   m.Params.Minconf,
			Maxconf:   m.Params.Maxconf,
			ChainType: "BTC",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetUTXO returns UTXO info for BTC addresses provided.
func GetUTXO(ae accountUTXOExecutor, k string, c *websocket.Conn) {
	ae.getUTXO(k, c)
}
