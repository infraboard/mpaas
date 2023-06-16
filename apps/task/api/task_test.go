package api_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
)

func TestWatchTaskLog(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		impl.WatchTaskLog(restful.NewRequest(r), restful.NewResponse(w))
	}))
	u, _ := url.Parse(srv.URL)
	u.Scheme = "ws"

	fmt.Println(u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("cannot make websocket connection: %v", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("xxxx"))
	if err != nil {
		log.Fatalf("cannot write message: %v", err)
	}

	// err = conn.WriteMessage(websocket.BinaryMessage, []byte("world"))
	// if err != nil {
	// 	log.Fatalf("cannot write message: %v", err)
	// }
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("cannot read message: %v", err)
	}

	fmt.Printf("success: received response: %q\n", p)
}
