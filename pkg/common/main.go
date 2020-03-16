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

type resultMap struct {
	Address          string  `json:"address"`
	StakeIn          float64 `json:"stakeIn"`
	StakeInTimestamp float64 `json:"stakeInTimestamp"`
}

type messageRecvMap struct {
	JSONRPC string `json:"jsonrpc"`
	resultMap
	ID float64 `json:"id"`
}

var messageRecvWrapper map[string]interface{}

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

// ParseRes dynamically unmarshals JSON messages
func ParseRes(res []byte) {
	var m messageRecv
	err := json.Unmarshal(res, &messageRecvWrapper)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range messageRecvWrapper {
		switch key {
		case "jsonrpc":
			// fmt.Println("jsonrpc:", value)
			m.JSONRPC = value.(string)
		case "id":
			// fmt.Println("id:", value)
			m.ID = value.(float64)
		case "result":
			v := reflect.TypeOf(value).Kind()
			switch v {
			case reflect.Float64:
				m.Result = value.(float64)
				fmt.Println("Result:", m.Result)
			case reflect.Slice:
				var mrm messageRecvMap
				val := reflect.ValueOf(value)
				if val.Len() == 0 {
					log.Println("Result is empty!")
					break
				}
				for i := 0; i < val.Len(); i++ {
					elm := val.Index(i).Interface().(map[string]interface{})
					mrm.JSONRPC = m.JSONRPC
					mrm.ID = m.ID
					mrm.Address = elm["address"].(string)
					mrm.StakeIn = elm["stakeIn"].(float64)
					mrm.StakeInTimestamp = elm["stakeInTimestamp"].(float64)
					// fmt.Println(mrm.Address, mrm.StakeIn, mrm.StakeInTimestamp)
					fmt.Println("Result:", mrm.resultMap)
				}
			default:
				fmt.Printf("value: %v\n", reflect.ValueOf(value))
				fmt.Printf("type: %v\n", reflect.ValueOf(value).Kind())
				fmt.Printf("error: unrecognized type (%v)\n", v.String())
			}
		}
	}
}
