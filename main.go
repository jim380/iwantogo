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
	iwantogo "github.com/iwantogo/pkg"
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
	msg := iwantogo.NewReq(address)
	iwantogo.GetBalance(msg, secretKey, c)
	msg = iwantogo.NewReq(address)
	iwantogo.GetValidatorInfo(msg, secretKey, c)

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
