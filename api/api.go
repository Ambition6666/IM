package api

import (
	"errors"
	"fmt"
	"im/internal/service/login"
	"im/internal/service/models"
	"im/internal/service/user"
	websokcet "im/internal/service/websocket"
	"im/vo"

	"github.com/gin-gonic/gin"
)

// 200 OK: 请求成功，服务器已成功处理请求。
// 201 Created: 请求成功并创建了新的资源。
// 204 No Content: 请求成功，但没有返回任何内容。
// 400 Bad Request: 服务器无法理解请求。
// 401 Unauthorized: 请求要求身份验证。
// 403 Forbidden: 服务器拒绝请求。
// 404 Not Found: 请求的资源不存在。
// 500 Internal Server Error: 服务器遇到错误，无法完成请求。
// 502 Bad Gateway: 服务器作为网关或代理，从上游服务器接收到无效的响应。
// 503 Service Unavailable: 服务器当前不可用
// 创建客户端与服务端的长链接,进行通信
func Chat(c *gin.Context) {
	//u, _ := GetUser(c)
	cc := websokcet.Up(c)
	cl := websokcet.NewClient(c.Query("name"), cc.RemoteAddr().String(), cc)
	go cl.Read()
	go cl.Write()
	go cl.TimeOutClose()
	websokcet.Manager.Register <- cl
}

// 用户登录
func Login(c *gin.Context) {
	u := new(vo.UserInfo)
	c.Bind(u)
	code, data := login.Login(u.N, u.P)
	switch code {
	case 401:
		c.JSON(code, data)
	case 200:
		c.JSON(code, gin.H{
			"msg":  "登录成功",
			"data": data,
		})
	}
}

// 用户注册
func Register(c *gin.Context) {
	u := new(vo.UserR)
	c.Bind(u)
	code := login.Register(u)
	if code == 200 {
		c.JSON(code, "注册成功")
	} else {
		c.JSON(code, "注册失败")
	}
}

// 申请好友接口
func Addfriend(c *gin.Context) {
	v := c.PostForm("name")
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := user.Add(v, u)
	c.JSON(code, data)
}

// 同意好友申请
func AgreeFriendPost(c *gin.Context) {
	v := c.Query("name")
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := user.Agree(v, u.Name)
	c.JSON(code, data)
}

// 拉取好友请求列表
func GetPostList(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := user.GetPostList(u.Name)
	c.JSON(code, data)
}

// 拉取好友列表
func GetList(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := user.GetList(u.Name)
	c.JSON(code, data)
}
func GetUser(c *gin.Context) (*models.User, error) {
	y, err := c.Get("user") //得到登录用户的信息
	fmt.Println(y)
	if !err {

		e := errors.New("提取用户失败")
		return nil, e
	}
	user1 := y.(models.User) //将any格式转化为USER格式
	return &user1, nil
}

// 获取消息记录
func Get_messages_list(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := user.Get_messages_list(u.Name, c.Query("name"))
	c.JSON(code, data)
}

// 创建群聊
func CreateGroup(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	a := new(vo.GroupName)
	c.Bind(a)
	//fmt.Println(a.Gname)
	code, data := user.CreateGroup(u.Name, a.Gname)
	if data != nil {
		c.JSON(code, data)
	} else {
		c.JSON(code, "创建成功")
	}
}

// 加入群聊
func JoinGroup(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	a := new(vo.GroupName)
	c.Bind(a)
	//fmt.Println(a.Gname)
	code, data := user.JoinGroup(a.Gname, u.Name)
	if data != nil {
		c.JSON(code, data)
	} else {
		c.JSON(code, "加群成功")
	}
}

// 查看自己所加群聊列表
func SearchGrouplist(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	code, data := user.SearchGrouplist(u.Name)
	c.JSON(code, data)
}

func Get_group_messages_list(c *gin.Context) {
	code, data := user.Get_group_messages_list(c.Query("gname"))
	c.JSON(code, data)
}
