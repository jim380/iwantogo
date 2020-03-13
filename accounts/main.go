package accounts

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"

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

// genSig generates a hmac sha256 signature
func (m *accountMessagePreSign) genSig(k string) string {
	key := []byte(k)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(m)
	message, _ := json.Marshal(m)

	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	// ---	get signature and encode in base64	--- //
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// ---	get signature and encode in hex	--- //
	// signature := hex.EncodeToString(hash.Sum(nil))

	// ---	check signature	--- //
	// fmt.Println("\nSignature: " + signature)
	return signature
}

func (m *accountMessagePostSign) sendMessage(c *websocket.Conn) {
	// ---	check JSON	--- //
	// result, _ := json.Marshal(m)
	// stringJSON := string(result)
	// fmt.Println("\nJSON:", stringJSON)

	connectionErr := c.WriteJSON(m)
	if connectionErr != nil {
		log.Println("write:", connectionErr)
	}
}

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
func GetBalance(ae accountExecutor, k string, c *websocket.Conn) {
	ae.getBalance(k, c)
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
func GetNonce(ae accountExecutor, k string, c *websocket.Conn) {
	ae.getNonce(k, c)
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
func GetNonceIncludePending(ae accountExecutor, k string, c *websocket.Conn) {
	ae.getNonceIncludePending(k, c)
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
func ImportAddress(ae accountExecutor, k string, c *websocket.Conn) {
	ae.importAddress(k, c)
}
