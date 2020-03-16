package pos

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/pkg/common"
)

type posHeightExecutor interface {
	getStakerInfo(key string, c *websocket.Conn)
}

type posHeightParamsPreSign struct {
	BlockNumber int    `json:"blockNumber"`
	ChainType   string `json:"chainType"`
	Timestamp   string `json:"timestamp"`
}

type posHeightMessagePreSign struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  posHeightParamsPreSign `json:"params"`
	ID      int64                  `json:"id"`
}

type posHeightParamsPostSign struct {
	BlockNumber int    `json:"blockNumber"`
	ChainType   string `json:"chainType"`
	Timestamp   string `json:"timestamp"`
	Signature   string `json:"signature"`
}

type posHeightMessagePostSign struct {
	JSONRPC string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  posHeightParamsPostSign `json:"params"`
	ID      int64                   `json:"id"`
}

// NewReqByHeight instantiates a new RPC-JSON call
func NewReqByHeight(height int) *posHeightMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &posHeightMessagePreSign{
		JSONRPC: "2.0",
		Params: posHeightParamsPreSign{
			BlockNumber: height,
			ChainType:   "WAN",
			Timestamp:   timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *posHeightMessagePreSign) getStakerInfo(key string, c *websocket.Conn) {
	m.Method = "getStakerInfo"
	sig := common.GenSig(m, key)
	msgSend := &posHeightMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getStakerInfo",
		Params: posHeightParamsPostSign{
			BlockNumber: m.Params.BlockNumber,
			ChainType:   "WAN",
			Timestamp:   m.Params.Timestamp,
			Signature:   sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetStakerInfo returns the info of a specific validator account
func GetStakerInfo(pe posHeightExecutor, k string, c *websocket.Conn) {
	pe.getStakerInfo(k, c)
}
