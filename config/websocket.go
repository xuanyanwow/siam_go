package config

import (
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
	"main/core"
	"net/http"
)

func WebSocket(){
	// 获取用户id标识
	idGen := func(ctx iris.Context) string {
		if username := ctx.GetHeader("X-Username"); username != "" {
			return username
		}

		return websocket.DefaultIDGenerator(ctx)
	}

	websocketServer := neffos.New(
		gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool{return true}}),
		neffos.Events{
			websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
				// 消息处理
				log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

				nsConn.Conn.Server().Broadcast(nsConn, msg)
				return nil
			},
		},)

	websocketServer.OnConnect = func(c *websocket.Conn)  error {
		log.Println("on connect")
		return nil
	}
	websocketServer.OnUpgradeError = func(err error) {
		log.Println(err.Error())
	}

	core.S_Iris.Handle("GET",  "/websocket_endpoint", websocket.Handler(websocketServer, idGen))
}