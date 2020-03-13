package blocks

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type blockExecutorHeight interface {
	getBlockByNumber(key string, c *websocket.Conn)
	getBlockTransactionCountByHeight(key string, c *websocket.Conn)
}

type blockHeightParamsPreSign struct {
	BlockNumber string `json:"blockNumber"`
	ChainType   string `json:"chainType"`
	Timestamp   string `json:"timestamp"`
}

type blockHeightMessagePreSign struct {
	JSONRPC string                   `json:"jsonrpc"`
	Method  string                   `json:"method"`
	Params  blockHeightParamsPreSign `json:"params"`
	ID      int64                    `json:"id"`
}

type blockHeightParamsPostSign struct {
	BlockNumber string `json:"blockNumber"`
	ChainType   string `json:"chainType"`
	Timestamp   string `json:"timestamp"`
	Signature   string `json:"signature"`
}

type blockHeightMessagePostSign struct {
	JSONRPC string                    `json:"jsonrpc"`
	Method  string                    `json:"method"`
	Params  blockHeightParamsPostSign `json:"params"`
	ID      int64                     `json:"id"`
}

// NewReqByHeight instantiates a new RPC-JSON call
func NewReqByHeight(height string) *blockHeightMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &blockHeightMessagePreSign{
		JSONRPC: "2.0",
		Params: blockHeightParamsPreSign{
			BlockNumber: height,
			ChainType:   "WAN",
			Timestamp:   timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *blockHeightMessagePreSign) getBlockByNumber(key string, c *websocket.Conn) {
	m.Method = "getBlockByNumber"
	sig := common.GenSig(m, key)
	msgSend := &blockHeightMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBlockByNumber",
		Params: blockHeightParamsPostSign{
			BlockNumber: m.Params.BlockNumber,
			ChainType:   "WAN",
			Timestamp:   m.Params.Timestamp,
			Signature:   sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetBlockByNumber returns info of the block number provided
func GetBlockByNumber(be blockExecutorHeight, k string, c *websocket.Conn) {
	be.getBlockByNumber(k, c)
}

func (m *blockHeightMessagePreSign) getBlockTransactionCountByHeight(key string, c *websocket.Conn) {
	m.Method = "getBlockTransactionCount"
	sig := common.GenSig(m, key)
	msgSend := &blockHeightMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBlockTransactionCount",
		Params: blockHeightParamsPostSign{
			BlockNumber: m.Params.BlockNumber,
			ChainType:   "WAN",
			Timestamp:   m.Params.Timestamp,
			Signature:   sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetBlockTransactionCountByHeight returns the info of a specific validator account
func GetBlockTransactionCountByHeight(be blockExecutorHeight, k string, c *websocket.Conn) {
	be.getBlockTransactionCountByHeight(k, c)
}
