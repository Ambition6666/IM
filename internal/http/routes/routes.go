package routes

import (
	"im/api"
	"im/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/login", api.Login)       //用户登录
	r.POST("/register", api.Register) //用户注册
	r.GET("/acc", api.Chat)           //用户申请聊天
	identity := r.Group("/identity")  //已经过token验证
	identity.Use(middleware.Mid())
	{

		identity.POST("/addfriend", api.Addfriend)                            //申请添加好友
		identity.GET("/agree", api.AgreeFriendPost)                           //同意添加好友请求
		identity.GET("/getfriendpostlist", api.GetPostList)                   //获取添加好友申请列表
		identity.GET("/getfriendlist", api.GetList)                           //获取好友列表
		identity.GET("/get_messages_list", api.Get_messages_list)             //获取消息记录(私人聊天)
		identity.GET("/get_group_messages_list", api.Get_group_messages_list) //获取群聊记录
		identity.POST("/creategroup", api.CreateGroup)                        //创建群聊
		identity.POST("/joingroup", api.JoinGroup)                            //参加群聊
		identity.POST("/getgrouplist", api.SearchGrouplist)                   //获取已加群列表
	}
	return r
}
