package login

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Msk []byte = []byte("ztyyyds666")

type Claim struct {
	N string
	jwt.RegisteredClaims
} //创建用户登录标签
func GetToken(b []byte, c string) (string, error) { //得到token,c为用户名
	a := Claim{
		c,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), //token有效时间
			Issuer:    "zty",                                                   //签发人
		},
	} //获取claim实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a) //获取token
	return token.SignedString(b)                          //返回加密串
}
func ParseToken(a string) (*jwt.Token, *Claim, error) { //解析token
	claim := &Claim{}
	t, err := jwt.ParseWithClaims(a, claim, func(t *jwt.Token) (interface{}, error) {
		return Msk, nil
	}) //接收前端发来加密字段
	return t, claim, err
}
