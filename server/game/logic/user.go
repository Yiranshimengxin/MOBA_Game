package logic

import "time"

func NewUser() *User {
	//u:=new(User)
	//return u

	u := &User{
		Uid:        10000,
		Account:    "test",
		Name:       "Ganyu",
		CreateTime: time.Now().Unix(),
	}
	return u
}

type User struct {
	Uid             UserID // 用户ID
	Account         string // 账号
	Name            string // 名字
	CreateTime      int64  // 创建时间
	LastOffLineTime int64  // 最后下线时间
	loginTIme       int64  // 登录时间
	Money           int64  // 金币
}
