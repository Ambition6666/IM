package middleware

import (
	"im/internal/repository"
	"im/internal/service/login"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 验证用户是否登录的中间件
func Mid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.GetHeader("Authorization") //得到字串开头
		if t == "" || !strings.HasPrefix(t, "Bearer ") {
			ctx.JSON(401, "解析失败")
			ctx.Abort()
			return
		}

		t = t[7:]                       //扔掉头部
		tk, c, e := login.ParseToken(t) //c为claim结构体的实例
		if e != nil || !tk.Valid {
			ctx.JSON(401, "解析失败")
			ctx.Abort() //中间件不通过
			return
		}
		//查找用户
		u := repository.SearchUser(c.N)
		//存储用户信息
		ctx.Set("user", *u)
		ctx.Next()
	}

}

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
