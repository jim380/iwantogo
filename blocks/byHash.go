package blocks

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type blockExecutorHash interface {
	getBlockByHash(key string, c *websocket.Conn)
	getBlockTransactionCountByHash(k string, c *websocket.Conn)
}

type blockHashParamsPreSign struct {
	BlockHash string `json:"blockHash"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type blockHashMessagePreSign struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  blockHashParamsPreSign `json:"params"`
	ID      int64                  `json:"id"`
}

type blockHashParamsPostSign struct {
	BlockHash string `json:"blockHash"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type blockHashMessagePostSign struct {
	JSONRPC string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  blockHashParamsPostSign `json:"params"`
	ID      int64                   `json:"id"`
}

// NewReqByHash instantiates a new RPC-JSON call
func NewReqByHash(hash string) *blockHashMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &blockHashMessagePreSign{
		JSONRPC: "2.0",
		Params: blockHashParamsPreSign{
			BlockHash: hash,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *blockHashMessagePreSign) getBlockByHash(key string, c *websocket.Conn) {
	m.Method = "getBlockByHash"
	sig := common.GenSig(m, key)
	msgSend := &blockHashMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBlockByHash",
		Params: blockHashParamsPostSign{
			BlockHash: m.Params.BlockHash,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetBlockByHash returns info of the block hash provided
func GetBlockByHash(be blockExecutorHash, k string, c *websocket.Conn) {
	be.getBlockByHash(k, c)
}

func (m *blockHashMessagePreSign) getBlockTransactionCountByHash(key string, c *websocket.Conn) {
	m.Method = "getBlockTransactionCount"
	sig := common.GenSig(m, key)
	msgSend := &blockHashMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBlockTransactionCount",
		Params: blockHashParamsPostSign{
			BlockHash: m.Params.BlockHash,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetBlockTransactionCountByHash returns the transaction count of the block hash provided
func GetBlockTransactionCountByHash(be blockExecutorHash, k string, c *websocket.Conn) {
	be.getBlockTransactionCountByHash(k, c)
}
