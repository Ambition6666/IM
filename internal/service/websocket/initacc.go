package websokcet

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	web "github.com/gorilla/websocket"
)

func Up(c *gin.Context) *web.Conn {
	conn, err := (&web.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)

		return nil
	}
	return conn
}
