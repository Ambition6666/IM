package repository

import (
	"context"
	"fmt"
	"im/sql"

	"github.com/redis/go-redis/v9"
)

// -----------------------------------------------------
// 以下为群聊操作
// 创建群聊名称,本人名字
// Score:1为群主,2为管理员,3为普通成员

// 创建群聊
func CreateGroup(gname string, name string) (error, error) {
	rdb := sql.GetRedis()
	err := rdb.ZAdd(context.Background(), gname, redis.Z{Score: 1, Member: name}).Err()
	err2 := rdb.ZAdd(context.Background(), name, redis.Z{Score: 1, Member: gname}).Err()
	return err, err2
}

// 加入群聊
func JoinGroup(gname string, name string) (error, error) {
	rdb := sql.GetRedis()
	err := rdb.ZAdd(context.Background(), gname, redis.Z{Score: 3, Member: name}).Err()
	err2 := rdb.ZAdd(context.Background(), name, redis.Z{Score: 3, Member: gname}).Err()
	return err, err2
}

// 设置群管理员
func SetGroupManger(gname string, name string) (error, error) {
	rdb := sql.GetRedis()
	err := rdb.ZAdd(context.Background(), gname, redis.Z{Score: 2, Member: name}).Err()
	err2 := rdb.ZAdd(context.Background(), name, redis.Z{Score: 2, Member: gname}).Err()
	return err, err2
}

// 删除群成员或者成员退出该群聊
func DltMember(gname string, name string) (error, error) {
	rdb := sql.GetRedis()
	err := rdb.ZRem(context.Background(), gname, name).Err()
	err2 := rdb.ZRem(context.Background(), name, gname).Err()
	return err, err2
}

// 查询群成员或者查询某人所加群列表
func SearchMembers(gname string) ([]string, error) {
	rdb := sql.GetRedis()
	vals, err := rdb.ZRange(context.Background(), gname, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return vals, nil
}

// 查询该群人数
func SearchGroup(gname string) int64 {
	rdb := sql.GetRedis()
	size, err := rdb.ZCard(context.Background(), gname).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return size
}

//查询自己所加群的列表
