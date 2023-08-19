package websokcet

import (
	"sync"
	"time"
)

type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // appId+uuid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	IsLogin     chan *Client       // 用户登录处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

var Manager ClientManager

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		IsLogin:    make(chan *Client, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
	return
}
func (m *ClientManager) Rest(c *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()
	m.Clients[c] = true
}
func (m *ClientManager) UnRest(c *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()
	m.Clients[c] = false
}
func (m *ClientManager) CreateUser(c *Client) {
	m.UserLock.Lock()
	defer m.UserLock.Unlock()
	m.Users[c.Name] = c
}
func (m *ClientManager) BroadCast(msg []byte) {
	for k, v := range m.Clients {
		if v {
			k.SendMessage(msg)
		}
	}
}
func (m *ClientManager) Heartbeat(c *Client) {
	msg := []byte("pong")
	c.SendMessage(msg)
	c.Login_time = time.Now()
}
func (m *ClientManager) Start() {
	for {
		select {
		case conn := <-m.Register:
			m.Rest(conn)
			m.CreateUser(conn)
		case conn := <-m.Unregister:
			m.UnRest(conn)
		case conn := <-m.IsLogin:
			m.Heartbeat(conn)
		case msg := <-m.Broadcast:
			m.BroadCast(msg)
		}
	}
}
