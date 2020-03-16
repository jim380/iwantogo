package pos

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/pkg/common"
)

type posAddrExecutor interface {
	getValidatorInfo(key string, c *websocket.Conn)
	getDelegatorStakeInfo(key string, c *websocket.Conn)
	getDelegatorSupStakeInfo(key string, c *websocket.Conn)
	getValidatorStakeInfo(key string, c *websocket.Conn)
	getValidatorSupStakeInfo(key string, c *websocket.Conn)
}

type posAddrParamsPreSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type posAddrMessagePreSign struct {
	JSONRPC string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  posAddrParamsPreSign `json:"params"`
	ID      int64                `json:"id"`
}

type posAddrParamsPostSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type posAddrMessagePostSign struct {
	JSONRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  posAddrParamsPostSign `json:"params"`
	ID      int64                 `json:"id"`
}

// NewReqByAddr instantiates a new RPC-JSON call
func NewReqByAddr(addr string) *posAddrMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &posAddrMessagePreSign{
		JSONRPC: "2.0",
		Params: posAddrParamsPreSign{
			Address:   addr,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *posAddrMessagePreSign) getValidatorInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorInfo"
	sig := common.GenSig(m, key)
	msgSend := &posAddrMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorInfo",
		Params: posAddrParamsPostSign{
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
func GetValidatorInfo(pe posAddrExecutor, k string, c *websocket.Conn) {
	pe.getValidatorInfo(k, c)
}

func (m *posAddrMessagePreSign) getDelegatorStakeInfo(key string, c *websocket.Conn) {
	m.Method = "getDelegatorStakeInfo"
	sig := common.GenSig(m, key)
	msgSend := &posAddrMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getDelegatorStakeInfo",
		Params: posAddrParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetDelegatorStakeInfo returns staking info of a specific delegator account
func GetDelegatorStakeInfo(pe posAddrExecutor, k string, c *websocket.Conn) {
	pe.getDelegatorStakeInfo(k, c)
}

func (m *posAddrMessagePreSign) getDelegatorSupStakeInfo(key string, c *websocket.Conn) {
	m.Method = "getDelegatorSupStakeInfo"
	sig := common.GenSig(m, key)
	msgSend := &posAddrMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getDelegatorSupStakeInfo",
		Params: posAddrParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetDelegatorSupStakeInfo returns supplementary info of a specific delegator account
func GetDelegatorSupStakeInfo(pe posAddrExecutor, k string, c *websocket.Conn) {
	pe.getDelegatorSupStakeInfo(k, c)
}

func (m *posAddrMessagePreSign) getValidatorStakeInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorStakeInfo"
	sig := common.GenSig(m, key)
	msgSend := &posAddrMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorStakeInfo",
		Params: posAddrParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetValidatorStakeInfo returns staking info of a specific validator account
func GetValidatorStakeInfo(pe posAddrExecutor, k string, c *websocket.Conn) {
	pe.getValidatorStakeInfo(k, c)
}

func (m *posAddrMessagePreSign) getValidatorSupStakeInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorSupStakeInfo"
	sig := common.GenSig(m, key)
	msgSend := &posAddrMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorSupStakeInfo",
		Params: posAddrParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetValidatorSupStakeInfo returns supplementary info of a specific validator account
func GetValidatorSupStakeInfo(pe posAddrExecutor, k string, c *websocket.Conn) {
	pe.getValidatorSupStakeInfo(k, c)
}
