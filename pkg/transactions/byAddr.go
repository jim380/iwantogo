package transactions

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/pkg/common"
)

type txAddrExecutor interface {
	getTransByAddress(key string, c *websocket.Conn)
}

type txAddrBlkExecutor interface {
	getTransByAddressBetweenBlocks(key string, c *websocket.Conn)
}

type txAddrParamsPreSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type txAddrMessagePreSign struct {
	JSONRPC string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  txAddrParamsPreSign `json:"params"`
	ID      int64               `json:"id"`
}

type txAddrParamsPostSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type txAddrMessagePostSign struct {
	JSONRPC string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  txAddrParamsPostSign `json:"params"`
	ID      int64                `json:"id"`
}

// Transactions between blocks
type txAddrBlkParamsPreSign struct {
	Address      string `json:"address"`
	StartBlockNo int64  `json:"startBlockNo"`
	EndBlockNo   int64  `json:"endBlockNo"`
	ChainType    string `json:"chainType"`
	Timestamp    string `json:"timestamp"`
}

type txAddrBlkMessagePreSign struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  txAddrBlkParamsPreSign `json:"params"`
	ID      int64                  `json:"id"`
}

type txAddrBlkParamsPostSign struct {
	Address      string `json:"address"`
	StartBlockNo int64  `json:"startBlockNo"`
	EndBlockNo   int64  `json:"endBlockNo"`
	ChainType    string `json:"chainType"`
	Timestamp    string `json:"timestamp"`
	Signature    string `json:"signature"`
}

type txAddrBlkMessagePostSign struct {
	JSONRPC string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  txAddrBlkParamsPostSign `json:"params"`
	ID      int64                   `json:"id"`
}

// NewReqByAddr instantiates a new RPC-JSON call
func NewReqByAddr(addr string) *txAddrMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &txAddrMessagePreSign{
		JSONRPC: "2.0",
		Params: txAddrParamsPreSign{
			Address:   addr,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

// NewReqByAddrBlk instantiates a new RPC-JSON call
func NewReqByAddrBlk(addr string, startBlk int64, endBlk int64) *txAddrBlkMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &txAddrBlkMessagePreSign{
		JSONRPC: "2.0",
		Params: txAddrBlkParamsPreSign{
			Address:      addr,
			StartBlockNo: startBlk,
			EndBlockNo:   endBlk,
			ChainType:    "WAN",
			Timestamp:    timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *txAddrMessagePreSign) getTransByAddress(key string, c *websocket.Conn) {
	m.Method = "getTransByAddress"
	sig := common.GenSig(m, key)
	msgSend := &txAddrMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getTransByAddress",
		Params: txAddrParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetTransByAddress returns transaction info of the specific address provided
func GetTransByAddress(e txAddrExecutor, k string, c *websocket.Conn) {
	e.getTransByAddress(k, c)
}

func (m *txAddrBlkMessagePreSign) getTransByAddressBetweenBlocks(key string, c *websocket.Conn) {
	m.Method = "getTransByAddressBetweenBlocks"
	sig := common.GenSig(m, key)
	msgSend := &txAddrBlkMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getTransByAddressBetweenBlocks",
		Params: txAddrBlkParamsPostSign{
			Address:      m.Params.Address,
			StartBlockNo: m.Params.StartBlockNo,
			EndBlockNo:   m.Params.EndBlockNo,
			ChainType:    "WAN",
			Timestamp:    m.Params.Timestamp,
			Signature:    sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetTransByAddressBetweenBlocks returns transaction info of the specific address provided
func GetTransByAddressBetweenBlocks(e txAddrBlkExecutor, k string, c *websocket.Conn) {
	e.getTransByAddressBetweenBlocks(k, c)
}
