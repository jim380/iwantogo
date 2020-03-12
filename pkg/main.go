package pkg

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type perform interface {
	// Accounts
	getBalance(key string, c *websocket.Conn)
	getNonce(key string, c *websocket.Conn)
	getNonceIncludePending(key string, c *websocket.Conn)
	importAddress(key string, c *websocket.Conn)
	// Blocks

	getValidatorInfo(key string, c *websocket.Conn)
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

type messageRecv struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int64  `json:"id"`
}

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

func nowAsUnixMilli() string {
	timeInt := time.Now().UnixNano() / 1e6
	return strconv.FormatInt(timeInt, 10)
}

// NewReq instantiates a new RPC-JSON call
func NewReq(addr string) *accountMessagePreSign {
	timeStamp := nowAsUnixMilli()
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
