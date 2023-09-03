package logic

import (
	"fmt"
	"time"
)

// 创建玩家的切片
var users []*User

// 服务器初始化
func Init() {
	fmt.Println("服务器初始化")

	//var userWriter Writer
	//userWriter = NewUser()
	//userWriter.Send(nil, 10)

	users = make([]*User, 0)
	for i := 0; i < 100; i++ {
		u := NewUser()
		users = append(users, u)
	}
}

// 服务器关闭
func Close() {
	fmt.Println("服务器初始化")
}

// 游戏循环
func Loop() {
	timer := time.Tick(1 * time.Millisecond)
	for {
		select {
		case <-timer:
			println("一秒钟过去了")
			UpdateUser()
		}
	}
}

func UpdateUser() {
	for _, u := range users {
		go u.Update()
	}
}
