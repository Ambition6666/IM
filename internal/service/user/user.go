package user

import (
	"encoding/json"
	"fmt"
	"im/internal/repository"
	"im/internal/service/models"
	websokcet "im/internal/service/websocket"
	tools "im/pkg"
	"im/vo"

	"github.com/gin-gonic/gin"
)

//个人操作

// 添加好友实现操作
func Add(v string, f *models.User) (int, string) {
	u := repository.SearchUser(v)
	fmt.Println(u)
	if u.ID == 0 {
		return 200, "没有此人"
	}
	repository.Create_friend_application(v, f)
	return 200, "申请成功"
}

// 同意好友申请
func Agree(n string, s string) (int, string) {
	repository.Create_friend(n, s)             //数据库插入好友记录
	repository.Create_friend(s, n)             //数据库插入好友记录
	repository.Delete_friend_application(n, s) //删除好友请求
	msg := vo.Message{
		Fr:  s,
		To:  n,
		Ctx: "我已添加你为好友",
		T:   1,
		TT:  1,
	} //添加一条同意好友申请的消息
	v, _ := json.Marshal(msg)
	//json封装
	websokcet.ParseDate(websokcet.Manager.Users[s], v)
	//处理发送
	return 200, "添加成功"
	//返回响应
}

// 获取好友申请记录
func GetPostList(n string) (int, []map[string]any) {
	v := repository.SearchPostList(n)
	a := make([]map[string]any, 0)
	for i := len(v) - 1; i >= 0; i-- {
		a = append(a, gin.H{
			"name": v[i].FromId,
		})
	}
	return 200, a
}

// 获取好友列表
func GetList(n string) (int, []map[string]any) {
	v := repository.SearchFriendList(n)
	a := make([]map[string]any, 0)
	for i := len(v) - 1; i >= 0; i-- {
		a = append(a, gin.H{
			"name": v[i].FellowName,
		})
	}
	return 200, a
}

// 获取聊天记录,第一个自己,第二个他人
func Get_messages_list(n string, s string) (int, []map[string]any) {
	v := repository.GetMessages(n, s)
	v = append(v, repository.GetMessages(s, n)...)
	tools.QuickSort(0, len(v)-1, &v)
	a := make([]map[string]any, 0)
	for i := len(v) - 1; i >= 0; i-- {
		a = append(a, gin.H{
			"id":  v[i].ID,
			"ctx": v[i].Context,
			"fr":  v[i].FromId,
		})
	}
	return 200, a
}

// 获取群消息
func Get_group_messages_list(g string) (int, []map[string]any) {
	v := repository.GetGroupMessages(g)
	a := make([]map[string]any, 0)
	for i := len(v) - 1; i >= 0; i-- {
		a = append(a, gin.H{
			"id":  v[i].ID,
			"ctx": v[i].Context,
			"fr":  v[i].FromId,
		})
	}
	return 200, a
}

// 群聊操作
func CreateGroup(n string, g string) (int, map[string]any) {
	num := repository.SearchGroup(g)
	if num != 0 {
		return 500, gin.H{
			"data": "已有该群",
		}
	}
	e1, e2 := repository.CreateGroup(g, n)
	if e1 != nil || e2 != nil {
		fmt.Println(e1, e2)
		return 500, gin.H{
			"data": "创建失败",
		}
	}
	return 200, nil
}

// 加入群聊
func JoinGroup(g string, n string) (int, map[string]any) {
	num := repository.SearchGroup(g)
	if num == 0 {
		return 500, gin.H{
			"data": "没有该群",
		}
	}
	e1, e2 := repository.JoinGroup(g, n)
	if e1 != nil || e2 != nil {
		fmt.Println(e1, e2)
		return 500, gin.H{
			"data": "创建失败",
		}
	}
	return 200, nil
}

// 查看自己所加群聊列表
func SearchGrouplist(n string) (int, []map[string]any) {
	vals, err := repository.SearchMembers(n)
	if err != nil {
		fmt.Println(err)
		return 500, nil
	}
	a := make([]map[string]any, 0)
	for _, v := range vals {
		a = append(a, gin.H{
			"gname": v,
		})
	}
	return 200, a
}
