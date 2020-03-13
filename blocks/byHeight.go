package blocks

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type blockHeightPerform interface {
	getBlockByNumber(key string, c *websocket.Conn)
	getBlockTransactionCountByHeight(key string, c *websocket.Conn)
}

type blockHeightParamsPreSign struct {
	BlockNumber string `json:"blockNumber"`
	ChainType   string `json:"chainType"`
	Timestamp   string `json:"timestamp"`
}

type blockHeightMessagePreSign struct {
	JSONRPC string                   `json:"jsonrpc"`
	Method  string                   `json:"method"`
	Params  blockHeightParamsPreSign `json:"params"`
	ID      int64                    `json:"id"`
}

type blockHeightParamsPostSign struct {
	BlockNumber string `json:"blockNumber"`
	ChainType   string `json:"chainType"`
	Timestamp   string `json:"timestamp"`
	Signature   string `json:"signature"`
}

type blockHeightMessagePostSign struct {
	JSONRPC string                    `json:"jsonrpc"`
	Method  string                    `json:"method"`
	Params  blockHeightParamsPostSign `json:"params"`
	ID      int64                     `json:"id"`
}

// genSig generates a hmac sha256 signature
func (m *blockHeightMessagePreSign) genSig(k string) string {
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

func (m *blockHeightMessagePostSign) sendMessage(c *websocket.Conn) {
	// ---	check JSON	--- //
	// result, _ := json.Marshal(m)
	// stringJSON := string(result)
	// fmt.Println("\nJSON:", stringJSON)

	connectionErr := c.WriteJSON(m)
	if connectionErr != nil {
		log.Println("write:", connectionErr)
	}
}

// NewReqByHeight instantiates a new RPC-JSON call
func NewReqByHeight(height string) *blockHeightMessagePreSign {
	timeStamp := nowAsUnixMilli()
	msg := &blockHeightMessagePreSign{
		JSONRPC: "2.0",
		Params: blockHeightParamsPreSign{
			BlockNumber: height,
			ChainType:   "WAN",
			Timestamp:   timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *blockHeightMessagePreSign) getBlockByNumber(key string, c *websocket.Conn) {
	m.Method = "getBlockByNumber"
	sig := m.genSig(key)
	msgSend := &blockHeightMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBlockByNumber",
		Params: blockHeightParamsPostSign{
			BlockNumber: m.Params.BlockNumber,
			ChainType:   "WAN",
			Timestamp:   m.Params.Timestamp,
			Signature:   sig,
		},
		ID: 1,
	}

	msgSend.sendMessage(c)
}

// GetBlockByNumber returns info of the block number provided
func GetBlockByNumber(p blockHeightPerform, k string, c *websocket.Conn) {
	p.getBlockByNumber(k, c)
}

func (m *blockHeightMessagePreSign) getBlockTransactionCountByHeight(key string, c *websocket.Conn) {
	m.Method = "getBlockTransactionCount"
	sig := m.genSig(key)
	msgSend := &blockHeightMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getBlockTransactionCount",
		Params: blockHeightParamsPostSign{
			BlockNumber: m.Params.BlockNumber,
			ChainType:   "WAN",
			Timestamp:   m.Params.Timestamp,
			Signature:   sig,
		},
		ID: 1,
	}

	msgSend.sendMessage(c)
}

// GetBlockTransactionCountByHeight returns the info of a specific validator account
func GetBlockTransactionCountByHeight(p blockHeightPerform, k string, c *websocket.Conn) {
	p.getBlockTransactionCountByHeight(k, c)
}
