package pkg

import "github.com/gorilla/websocket"

func (m *accountMessagePreSign) getValidatorInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorInfo"
	sig := m.genSig(key)
	msgSend := &accountMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorInfo",
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

// GetValidatorInfo returns the info of a specific validator account
func GetValidatorInfo(p perform, k string, c *websocket.Conn) {
	p.getValidatorInfo(k, c)
}
