package main

import "github.com/gorilla/websocket"

func (m *messagePreSign) getBalance(key string, c *websocket.Conn) {
	m.Method = "getBalance"
	sig := m.genSig(key)
	msgSend := &messagePostSign{
		JSONRPC: "2.0",
		Method:  "getBalance",
		Params: paramsPostSign{
			Address:   m.Params.Address,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	msgSend.sendMessage(c)
}

// GetBalance fetches the balance of a specific account
func GetBalance(p perform, k string, c *websocket.Conn) {
	p.getBalance(k, c)
}
