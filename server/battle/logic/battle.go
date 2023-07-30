package logic

import (
	"fmt"
	"time"
)

// 服务器初始化
func Init() {
	fmt.Print("服务器初始化")
}

// 服务器关闭
func Close() {
	fmt.Print("服务器初始化")
}

// 游戏循环
func Loop() {
	timer := time.Tick(1 * time.Second)
	for {
		select {
		case <-timer:
			println("一秒钟过去了")
		}
	}
}
