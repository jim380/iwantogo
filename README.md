# iwantoGo

## Go Library for iWan RPC Server

### Instructions

1. Create a `config.env` file under root directory and fill in parameters as shown in `config.env.example`
2. `$ go get -u github.com/jim380/iwantogo`

### Example

```go
import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	account "github.com/iwantogo/pkg/accounts"
	block "github.com/iwantogo/pkg/blocks"
	common "github.com/iwantogo/pkg/common"
	pos "github.com/iwantogo/pkg/pos"
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
	addresses := []string{"0x7212b9e259792879d85ca3227384f1005437e5f5", "0xfc2730f75330bb75cb28fcff12f0aea5b6e433e1", "0x4ee67553ab5fa994bc6a9cefecc93ff134083343"}
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
	msgAcct := account.NewReqMulti(addresses)
	account.GetMultiBalances(msgAcct, secretKey, c)
	msgAcct := account.NewReqUTXO(btcAddresses, 0, 6)
	account.GetUTXO(msgAcct, secretKey, c)
	msgPOSByAddr := pos.NewReqByAddr(address)
	pos.GetValidatorSupStakeInfo(msgPOSByAddr, secretKey, c)
	msgPOSByEpochID := pos.NewReqByEpochID(epochID)
	pos.GetTimeByEpochID(msgPOSByEpochID, secretKey, c)
	msgBlkByHash := block.NewReqByHash(hash)
	block.GetBlockTransactionCountByHash(msgBlkByHash, secretKey, c)
	msgBlkByHeight := block.NewReqByHeight(height)
	block.GetBlockTransactionCountByHeight(msgBlkByHeight, secretKey, c)

	for {
		select {
		case <-done:
			return
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
```