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
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/conf"
)

func TestWatchTaskLog(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		impl.JobTaskLog(restful.NewRequest(r), restful.NewResponse(w))
	}))
	u, _ := url.Parse(srv.URL)
	u.Scheme = "ws"

	fmt.Println(u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("cannot make websocket connection: %v", err)
	}

	req := task.NewWatchJobTaskLogRequest(conf.C.MCENTER_BUILD_TASK_ID)
	err = conn.WriteMessage(websocket.TextMessage, []byte(req.ToJSON()))
	if err != nil {
		log.Fatalf("cannot write message: %v", err)
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Print(string(p))
	}
}
