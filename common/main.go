package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type messageRecv struct {
	JSONRPC string  `json:"jsonrpc"`
	Result  float64 `json:"result"`
	ID      float64 `json:"id"`
}

var arbitraryJSON map[string]interface{}

// GetTimeStamp returns a timestamp in milliseconds
func GetTimeStamp() string {
	timeInt := time.Now().UnixNano() / 1e6
	return strconv.FormatInt(timeInt, 10)
}

// GenSig generates a hmac sha256 signature
func GenSig(m interface{}, k string) string {
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

func SendMessage(m interface{}, c *websocket.Conn) {
	// ---	check JSON	--- //
	// result, _ := json.Marshal(m)
	// stringJSON := string(result)
	// fmt.Println("\nJSON:", stringJSON)

	connectionErr := c.WriteJSON(m)
	if connectionErr != nil {
		log.Println("write:", connectionErr)
	}
}

// ParseRes unmarshals the JSON result
func ParseRes(res []byte) {
	var m messageRecv
	err := json.Unmarshal(res, &arbitraryJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range arbitraryJSON {
		switch key {
		case "jsonrpc":
			// fmt.Println("jsonrpc:", value)
			m.JSONRPC = value.(string)
		case "id":
			// fmt.Println("id:", value)
			m.ID = value.(float64)
		case "result":
			// fmt.Println("result:", value)
			switch v := value.(type) {
			case float64:
				// fmt.Printf("result: result(%v) is of type(%v)\n", v, reflect.TypeOf(v).Name())
				m.Result = value.(float64)
			default:
				fmt.Printf("result: unrecognized type (%v)\n", reflect.TypeOf(v).Name())
				fmt.Println(v)
				m.Result = 0.0
			}

		}
	}
	fmt.Println("Parsed Result:", m.Result)
}
