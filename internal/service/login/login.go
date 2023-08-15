package login

import (
	"fmt"
	"im/internal/repository"
	"im/vo"
)

func Login(n string, m string) (code int, resp string) {

	u := repository.SearchUser(n)
	//ID为0表示没有该用户
	if u.ID == 0 {
		resp = "没有该用户"
		code = 401
		return
	}
	if u.Pwd != m {
		resp = "密码错误"
		code = 401
		return
	}
	resp, err := GetToken(Msk, n)
	code = 200
	if err != nil {
		fmt.Println(err)
	}
	return
}
func Register(u *vo.UserR) int {
	n := repository.SearchUser(u.N)
	if n.ID != 0 {
		return 500
	}
	repository.CreateUser(u.N, u.P, u.E, u.P)
	return 200
}
