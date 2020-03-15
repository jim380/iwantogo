package accounts

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type accountExecutor interface {
	// Accounts
	getBalance(key string, c *websocket.Conn)
	getNonce(key string, c *websocket.Conn)
	getNonceIncludePending(key string, c *websocket.Conn)
	importAddress(key string, c *websocket.Conn)
}

type accountParamsPreSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type accountMessagePreSign struct {
	JSONRPC string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  accountParamsPreSign `json:"params"`
	ID      int64                `json:"id"`
}

type accountParamsPostSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type accountMessagePostSign struct {
	JSONRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  accountParamsPostSign `json:"params"`
	ID      int64                 `json:"id"`
}

// type messageRecv struct {
// 	JSONRPC string `json:"jsonrpc"`
// 	Result  string `json:"result"`
// 	ID      int64  `json:"id"`
// }

// NewReq instantiates a new RPC-JSON call
func NewReq(addr string) *accountMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &accountMessagePreSign{
		JSONRPC: "2.0",
		// Method:  "getBalance",
		Params: accountParamsPreSign{
			Address:   addr,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *accountMessagePreSign) getBalance(key string, c *websocket.Conn) {
	m.Method = "getBalance"
	// sig := m.common.GenSig(key)
	sig := common.GenSig(m, key)
	msgSend := &accountMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBalance",
		Params: accountParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetBalance returns the balance of a specific account
func GetBalance(ae accountExecutor, k string, c *websocket.Conn) {
	ae.getBalance(k, c)
}

func (m *accountMessagePreSign) getNonce(key string, c *websocket.Conn) {
	m.Method = "getNonce"
	sig := common.GenSig(m, key)
	msgSend := &accountMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getNonce",
		Params: accountParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetNonce returns the nonce of an account
func GetNonce(ae accountExecutor, k string, c *websocket.Conn) {
	ae.getNonce(k, c)
}

func (m *accountMessagePreSign) getNonceIncludePending(key string, c *websocket.Conn) {
	m.Method = "getNonceIncludePending"
	sig := common.GenSig(m, key)
	msgSend := &accountMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getNonceIncludePending",
		Params: accountParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetNonceIncludePending returns the nonce of an account
func GetNonceIncludePending(ae accountExecutor, k string, c *websocket.Conn) {
	ae.getNonceIncludePending(k, c)
}

func (m *accountMessagePreSign) importAddress(key string, c *websocket.Conn) {
	m.Method = "importAddress"
	sig := common.GenSig(m, key)
	msgSend := &accountMessagePostSign{
		JSONRPC: "2.0",
		Method:  "importAddress",
		Params: accountParamsPostSign{
			Address:   m.Params.Address,
			ChainType: "BTC",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// ImportAddress sends an import address to BTC.
func ImportAddress(ae accountExecutor, k string, c *websocket.Conn) {
	ae.importAddress(k, c)
}
