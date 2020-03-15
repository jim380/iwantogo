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
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	account "github.com/iwantogo/accounts"
	block "github.com/iwantogo/blocks"
	pos "github.com/iwantogo/pos"
	"github.com/joho/godotenv"
)

var addr = flag.String("addr", "api.wanchain.org:8443", "http service address")

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	address := os.Getenv("ADDRESS")
	hash := "0xa3c8e3e61c6f33af4125cbddb4792b284b980918fcd71db1e91a847a785a7ddd"
	// addresses := []string{"0x7212b9e259792879d85ca3227384f1005437e5f5", "0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1", "0x4ee67553ab5fa994bc6a9cefecc93ff134083343"}
	btcAddresses := []string{"19JEuWZbssQpXLMutL2cJeW2arqamnLCSP"}
	height := "500"
	epochID := 300

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

	//***************************************//
	//      Establish WS Connection          //
	//***************************************//
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			// ---	receive message	--- //
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Println("")
			log.Printf("received: %s", message)
			// ---	parse results	--- //
			// var msgRecv iwantogo.MessageRecv
			// json.Unmarshal([]byte(message), &msgRecv)
			// fmt.Println(msgRecv.Result)
		}
	}()

	//***************************************//
	//                Calls                  //
	//***************************************//
	// msgAcct := account.NewReq(address)
	// account.GetBalance(msgAcct, secretKey, c)
	// msgAcct := account.NewReqMulti(addresses)
	// account.GetMultiBalances(msgAcct, secretKey, c)
	msgAcct := account.NewReqUTXO(btcAddresses, 0, 6)
	account.GetUTXO(msgAcct, secretKey, c)
	msgPOSByAddr := pos.NewReqByAddr(address)
	pos.GetValidatorSupStakeInfo(msgPOSByAddr, secretKey, c)
	msgPOSByEpochID := pos.NewReqByEpochID(epochID)
	pos.GetTimeByEpochID(msgPOSByEpochID, secretKey, c)
	// msgPOS := pos.NewReq()
	// pos.GetCurrentStakerInfo(msgPOS, secretKey, c)
	// msgBlkByHash := block.NewReqByHash(hash)
	// block.GetBlockByHash(msgBlkByHash, secretKey, c)
	// msgBlkByHeight := block.NewReqByHeight(height)
	// block.GetBlockByNumber(msgBlkByHeight, secretKey, c)
	msgBlkByHash := block.NewReqByHash(hash)
	block.GetBlockTransactionCountByHash(msgBlkByHash, secretKey, c)
	msgBlkByHeight := block.NewReqByHeight(height)
	block.GetBlockTransactionCountByHeight(msgBlkByHeight, secretKey, c)

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

			//***************************************//
			//          close ws connection          //
			//***************************************//
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
