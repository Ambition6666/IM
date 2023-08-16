package websokcet

import (
	"encoding/json"
	"fmt"
	"im/internal/repository"
	"im/internal/service/models"
	"im/vo"
	"runtime/debug"

	web "github.com/gorilla/websocket"
)

type Client struct {
	Name string
	Addr string
	Send chan []byte
	Conn *web.Conn
}

func NewClient(n string, addr string, cc *web.Conn) *Client {
	return &Client{
		Name: n,
		Addr: addr,
		Send: make(chan []byte, 100),
		Conn: cc,
	}
}

// 发送给客户端
func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)

		}
	}()

	defer func() {
		Manager.Unregister <- c
		c.Conn.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for v := range c.Send {
		err := c.Conn.WriteMessage(web.TextMessage, v)
		if err != nil {
			Manager.Unregister <- c
			c.Conn.Close()
			fmt.Println("Client发送数据 defer", c)
		}
	}
}

// 写数据给客户端
func (c *Client) SendMessage(msg []byte) {
	if c == nil {

		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

// 接收客户端发来的数据
func (c *Client) Read() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("write stop", string(debug.Stack()), e)
		}
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		ParseDate(c, msg)
	}
}
func ParseDate(cl *Client, msg []byte) {
	m := new(vo.Message)
	err := json.Unmarshal(msg, m)
	if err != nil {
		fmt.Println(err)
		return
	}
	//cl.SendMessage(msg)
	if m.TT == 1 {
		acl := Manager.Users[m.To]
		//如果该目标用户在线,则直接发给目标用户,并且将消息存储到数据库中
		//否则则存到消息队列中由后台对其进行处理
		if Manager.Clients[acl] {
			acl.SendMessage(msg)
			ms := &models.Message{
				FromId:     m.Fr,
				TargetId:   m.To,
				Context:    m.Ctx,
				Type:       m.T,
				TargetType: m.TT,
			}
			repository.InsertMessage(ms)
		} else {
			repository.InsertCache(string(msg))
		}
	} else if m.TT == 2 {
		vals, err := repository.SearchMembers(m.To)
		if err != nil {
			fmt.Println(err)
			return
		}
		//先把消息存进数据库里
		ms := &models.Message{
			FromId:     m.Fr,
			TargetId:   m.To,
			Context:    m.Ctx,
			Type:       m.T,
			TargetType: m.TT,
		}
		repository.InsertMessage(ms)
		for _, v := range vals {
			acl := Manager.Users[v]
			//如果该目标用户在线,则直接发给目标用户,并且将消息存储到数据库中
			//否则则存到消息队列中由后台对其进行处理
			if Manager.Clients[acl] {
				acl.SendMessage(msg)
			}
		}
	} else if m.TT == 3 {
		Manager.Broadcast <- msg
		ms := &models.Message{
			FromId:     m.Fr,
			TargetId:   m.To,
			Context:    m.Ctx,
			Type:       m.T,
			TargetType: m.TT,
		}
		repository.InsertMessage(ms)
	}

}
