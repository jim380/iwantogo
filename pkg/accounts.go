package pkg

import "github.com/gorilla/websocket"

func (m *accountMessagePreSign) getBalance(key string, c *websocket.Conn) {
	m.Method = "getBalance"
	sig := m.genSig(key)
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

	msgSend.sendMessage(c)
}

// GetBalance returns the balance of a specific account
func GetBalance(p perform, k string, c *websocket.Conn) {
	p.getBalance(k, c)
}

func (m *accountMessagePreSign) getNonce(key string, c *websocket.Conn) {
	m.Method = "getNonce"
	sig := m.genSig(key)
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

	msgSend.sendMessage(c)
}

// GetNonce returns the nonce of an account
func GetNonce(p perform, k string, c *websocket.Conn) {
	p.getNonce(k, c)
}

func (m *accountMessagePreSign) getNonceIncludePending(key string, c *websocket.Conn) {
	m.Method = "getNonceIncludePending"
	sig := m.genSig(key)
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

	msgSend.sendMessage(c)
}

// GetNonceIncludePending returns the nonce of an account
func GetNonceIncludePending(p perform, k string, c *websocket.Conn) {
	p.getNonceIncludePending(k, c)
}

func (m *accountMessagePreSign) importAddress(key string, c *websocket.Conn) {
	m.Method = "importAddress"
	sig := m.genSig(key)
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

	msgSend.sendMessage(c)
}

// ImportAddress sends an import address to BTC.
func ImportAddress(p perform, k string, c *websocket.Conn) {
	p.importAddress(k, c)
}
