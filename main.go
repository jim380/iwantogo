package main

//                                                  jim380 <admin@cyphercore.io>
//  ============================================================================
//
//  Copyright (C) 2020 jim380
//
//  Permission is hereby granted, free of charge, to any person obtaining
//  a copy of this software and associated documentation files (the
//  "Software"), to deal in the Software without restriction, including
//  without limitation the rights to use, copy, modify, merge, publish,
//  distribute, sublicense, and/or sell copies of the Software, and to
//  permit persons to whom the Software is furnished to do so, subject to
//  the following conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
//  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
//  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
//  CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN setup OF CONTRACT,
//  TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
//  SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
//  ============================================================================

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/node_tooling/Celo/util"
)

var addr = flag.String("addr", "api.wanchain.org:8443", "http service address")

type params struct {
	Address   string `json:"address"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type message struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  params `json:"params"`
	ID      int64  `json:"id"`
}

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	address := os.Getenv("ADDRESS")

	util.SetEnv()
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{
		Scheme: "wss",
		Host:   *addr,
		Path:   "/ws/v3/" + apiKey,
	}
	log.Printf("connecting to %s", u.String())

	// establish connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			// receive message
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("received: %s", message)
		}
	}()

	timeStamp := nowAsUnixMilli()
	msg := &message{
		JSONRPC: "2.0",
		Method:  "getBalance",
		Params: params{
			Address:   address,
			ChainType: "WAN",
			Timestamp: timeStamp,
			// Signature: sig,
		},
		ID: 1,
	}
	sig := msg.getSig(secretKey)
	msg.Params.Signature = sig

	json, _ := json.Marshal(msg)
	stringJSON := string(json)
	fmt.Println("\nJSON:", stringJSON)
	// send message
	connectionErr := c.WriteJSON(msg)
	if connectionErr != nil {
		log.Println("write:", connectionErr)
	}

	for {
		select {
		case <-done:
			return
		// case t := <-ticker.C:
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")

			// close connection gracefully
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection closed"))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func nowAsUnixMilli() string {
	timeInt := time.Now().UnixNano() / 1e6
	return strconv.FormatInt(timeInt, 10)
}

// getSig generates a hmac sha256 signature
func (m *message) getSig(k string) string {
	key := []byte(k)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(m)

	message := []byte(reqBodyBytes.Bytes())

	// Create a new HMAC instance
	hash := hmac.New(sha256.New, key)

	// Write Data to it
	hash.Write(message)

	// Get signature and encode in base64
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// Get signature and encode in hex
	// signature := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("\nSignature: " + signature)
	return signature
}
