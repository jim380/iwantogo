package pos

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type posChainTypeExecutor interface {
	getCurrentStakerInfo(key string, c *websocket.Conn)
	getCurrentEpochInfo(key string, c *websocket.Conn)
	getMaxStableBlkNumber(key string, c *websocket.Conn)
	getEpochID(key string, c *websocket.Conn)
	getSlotID(key string, c *websocket.Conn)
	getSlotTime(key string, c *websocket.Conn)
}

type posChainTypeParamsPreSign struct {
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type posChainTypeMessagePreSign struct {
	JSONRPC string                    `json:"jsonrpc"`
	Method  string                    `json:"method"`
	Params  posChainTypeParamsPreSign `json:"params"`
	ID      int64                     `json:"id"`
}

type posChainTypeParamsPostSign struct {
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type posChainTypeMessagePostSign struct {
	JSONRPC string                     `json:"jsonrpc"`
	Method  string                     `json:"method"`
	Params  posChainTypeParamsPostSign `json:"params"`
	ID      int64                      `json:"id"`
}

// NewReq instantiates a new RPC-JSON call
func NewReq() *posChainTypeMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &posChainTypeMessagePreSign{
		JSONRPC: "2.0",
		Params: posChainTypeParamsPreSign{
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *posChainTypeMessagePreSign) getCurrentStakerInfo(key string, c *websocket.Conn) {
	m.Method = "getCurrentStakerInfo"
	sig := common.GenSig(m, key)
	msgSend := &posChainTypeMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getCurrentStakerInfo",
		Params: posChainTypeParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetCurrentStakerInfo returns an array with info on each
// of the current validators
func GetCurrentStakerInfo(pe posChainTypeExecutor, k string, c *websocket.Conn) {
	pe.getCurrentStakerInfo(k, c)
}

func (m *posChainTypeMessagePreSign) getEpochID(key string, c *websocket.Conn) {
	m.Method = "getEpochID"
	sig := common.GenSig(m, key)
	msgSend := &posChainTypeMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getEpochID",
		Params: posChainTypeParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetEpochID returns the current Epoch ID
func GetEpochID(pe posChainTypeExecutor, k string, c *websocket.Conn) {
	pe.getEpochID(k, c)
}

func (m *posChainTypeMessagePreSign) getCurrentEpochInfo(key string, c *websocket.Conn) {
	m.Method = "getCurrentEpochInfo"
	sig := common.GenSig(m, key)
	msgSend := &posChainTypeMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getCurrentEpochInfo",
		Params: posChainTypeParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetCurrentEpochInfo returns the current Epoch ID
func GetCurrentEpochInfo(pe posChainTypeExecutor, k string, c *websocket.Conn) {
	pe.getCurrentEpochInfo(k, c)
}

func (m *posChainTypeMessagePreSign) getMaxStableBlkNumber(key string, c *websocket.Conn) {
	m.Method = "getMaxStableBlkNumber"
	sig := common.GenSig(m, key)
	msgSend := &posChainTypeMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getMaxStableBlkNumber",
		Params: posChainTypeParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetMaxStableBlkNumber returns  current highest stable block number
func GetMaxStableBlkNumber(pe posChainTypeExecutor, k string, c *websocket.Conn) {
	pe.getMaxStableBlkNumber(k, c)
}

func (m *posChainTypeMessagePreSign) getSlotID(key string, c *websocket.Conn) {
	m.Method = "getSlotID"
	sig := common.GenSig(m, key)
	msgSend := &posChainTypeMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getSlotID",
		Params: posChainTypeParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetSlotID returns the current epoch slot ID
func GetSlotID(pe posChainTypeExecutor, k string, c *websocket.Conn) {
	pe.getSlotID(k, c)
}

func (m *posChainTypeMessagePreSign) getSlotTime(key string, c *websocket.Conn) {
	m.Method = "getSlotTime"
	sig := common.GenSig(m, key)
	msgSend := &posChainTypeMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getSlotTime",
		Params: posChainTypeParamsPostSign{
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetSlotTime returns the duration of each slot in seconds
func GetSlotTime(pe posChainTypeExecutor, k string, c *websocket.Conn) {
	pe.getSlotTime(k, c)
}
