package pkg

// func (m *accountMessagePreSign) getBlockByHash(key string, c *websocket.Conn) {
// 	m.Method = "getBlockByHash"
// 	sig := m.genSig(key)
// 	msgSend := &accountMessagePostSign{
// 		JSONRPC: "2.0",
// 		Method:  "getBlockByHash",
// 		Params: accountParamsPostSign{
// 			Address:   m.Params.Address,
// 			ChainType: "WAN",
// 			Timestamp: m.Params.Timestamp,
// 			Signature: sig,
// 		},
// 		ID: 1,
// 	}

// 	msgSend.sendMessage(c)
// }

// // GetBlockByHash returns info of the block being queried
// func GetBlockByHash(p perform, k string, c *websocket.Conn) {
// 	p.getBlockByHash(k, c)
// }
