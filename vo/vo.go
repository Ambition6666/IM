package vo

//该包是对前端发来的数据进行脱敏

//接收用户信息
type UserInfo struct {
	N string `json:"n"`
	P string `json:"p"`
}

//接受用户注册信息
type UserR struct {
	N string `json:"n"`
	P string `json:"p"`
	O string `json:"o"`
	E string `json:"e"`
}

//消息
type Message struct {
	Fr  string `json:"fr"`
	To  string `json:"to"`
	Ctx string `json:"ctx"`
	T   uint   `json:"t"`
	TT  uint   `json:"tt"`
}

//
type GroupName struct {
	Gname string
}
