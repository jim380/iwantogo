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

type posPerform interface {
	getValidatorInfo(key string, c *websocket.Conn)
}

type posParamsPreSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type posMessagePreSign struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  posParamsPreSign `json:"params"`
	ID      int64            `json:"id"`
}

type posParamsPostSign struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type posMessagePostSign struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  posParamsPostSign `json:"params"`
	ID      int64             `json:"id"`
}

// genSig generates a hmac sha256 signature
func (m *posMessagePreSign) genSig(k string) string {
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

func (m *posMessagePostSign) sendMessage(c *websocket.Conn) {
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
func NewReq(addr string) *posMessagePreSign {
	timeStamp := nowAsUnixMilli()
	msg := &posMessagePreSign{
		JSONRPC: "2.0",
		Params: posParamsPreSign{
			Address:   addr,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *posMessagePreSign) getValidatorInfo(key string, c *websocket.Conn) {
	m.Method = "getValidatorInfo"
	sig := m.genSig(key)
	msgSend := &posMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorInfo",
		Params: posParamsPostSign{
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
func GetValidatorInfo(p posPerform, k string, c *websocket.Conn) {
	p.getValidatorInfo(k, c)
}
