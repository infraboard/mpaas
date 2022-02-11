package http

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mpaas/provider/k8s"
)

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: 60 * time.Second,
		ReadBufferSize:   8192,
		WriteBufferSize:  8192,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var (
	defaultCmd = `TERM=xterm-256color; export TERM; [ -x /bin/bash ] && ([ -x /usr/bin/script ] && /usr/bin/script -q -c "/bin/bash" /dev/null || exec /bin/bash) || exec /bin/sh`
)

func (h *handler) websocketAuth(payload string) error {
	return nil
}

// Login Container Websocket
func (h *handler) LoginContainer(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	// websocket handshake
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			h.log.Errorf("websocket handshake error, %s", err)
		}
		return
	}
	defer ws.Close()

	term := k8s.NewWebsocketTerminal(ws)

	// websocket 认证
	term.Auth(h.websocketAuth)

	// 获取参数
	req := k8s.NewLoginContainerRequest([]string{"sh", "-c", defaultCmd}, term)
	term.ParseParame(req)

	err = client.LoginContainer(req)
	if err != nil {
		_, err := term.Write([]byte(err.Error()))
		if err != nil {
			h.log.Errorf("term write error, %s", err)
		}
		return
	}
}

// Watch Container Log Websocket
func (h *handler) WatchConainterLog(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	client, err := h.GetClient(r.Context(), ctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	// websocket handshake
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			h.log.Errorf("websocket handshake error, %s", err)
		}
		return
	}
	defer ws.Close()

	term := k8s.NewWebsocketTerminal(ws)

	// websocket 认证
	term.Auth(h.websocketAuth)

	// 获取参数
	req := k8s.NewWatchConainterLogRequest()
	term.ParseParame(req)

	reader, err := client.WatchConainterLog(r.Context(), req)
	if err != nil {
		_, err := term.Write([]byte(err.Error()))
		if err != nil {
			h.log.Errorf("term write error, %s", err)
		}
		return
	}

	// 读取出来的数据流 copy到term
	_, err = io.Copy(term, reader)
	if err != nil {
		h.log.Errorf("copy log to weboscket error, %s", err)
	}
}
