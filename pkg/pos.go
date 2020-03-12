package pkg

import "github.com/gorilla/websocket"

func (m *messagePreSign) getValidatorInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorInfo"
	sig := m.genSig(key)
	msgSend := &messagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorInfo",
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

// GetValidatorInfo fetches the info of a specific validator account
func GetValidatorInfo(p perform, k string, c *websocket.Conn) {
	p.getValidatorInfo(k, c)
}
