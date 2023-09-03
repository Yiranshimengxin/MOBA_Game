package logic

import (
	"fmt"
	"time"
)

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

type Writer interface {
	Write([]byte) (int, error)
	Send([]byte, int) (int, error)
}

//type UserInterface interface {
//	Login(name string,password string)
//	Logout()
//}

func (u *User) Write([]byte) (int, error) {
	return 0, nil
}

// User实现了Send接口
func (u *User) Send([]byte, int) (int, error) {
	return 0, nil
}

func (u *User) Update() {
	fmt.Printf("玩家%d更新\n", u.Uid)
}
