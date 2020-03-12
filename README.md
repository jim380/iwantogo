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
		}
	}()

	//***************************************//
	//                Test                   //
	//***************************************//
	msgAcct := account.NewReq(address)
	accounts.GetBalance(msgAcct, secretKey, c)
	msgPOS := pos.NewReq(address)
	pos.GetValidatorInfo(msgPOS, secretKey, c)

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

