package status

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/pkg/common"
)

type statusExecutor interface {
	getGasPrice(key string, c *websocket.Conn)
}

type statusParamsPreSign struct {
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type statusMessagePreSign struct {
	JSONRPC string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  statusParamsPreSign `json:"params"`
	ID      int64               `json:"id"`
}

type statusParamsPostSign struct {
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type statusMessagePostSign struct {
	JSONRPC string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  statusParamsPostSign `json:"params"`
	ID      int64                `json:"id"`
}

// NewReq instantiates a new RPC-JSON call
func NewReq() *statusMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &statusMessagePreSign{
		JSONRPC: "2.0",
		Params: statusParamsPreSign{
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *statusMessagePreSign) getGasPrice(key string, c *websocket.Conn) {
	m.Method = "getGasPrice"
	sig := common.GenSig(m, key)
	msgSend := &statusMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getGasPrice",
		Params: statusParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetGasPrice returns the current gas price in wei as bigNumber type
func GetGasPrice(pe statusExecutor, k string, c *websocket.Conn) {
	pe.getGasPrice(k, c)
}
