package game

import (
	"google.golang.org/protobuf/proto"
	"server/game/gate"
	"time"
)

func NewUser(c *gate.Client, uid int64, acc string) *User {
	u := &User{
		Client:     c,
		Uid:        uid,
		Account:    acc,
		Name:       "Ganyu",
		CreateTime: time.Now().Unix(),
	}
	return u
}

type User struct {
	Uid             int64  // 用户ID
	Account         string // 账号
	Name            string // 名字
	CreateTime      int64  // 创建时间
	LastOffLineTime int64  // 最后下线时间
	LoginTIme       int64  // 登录时间
	Money           int64  // 金币
	Client          *gate.Client
}

func (u *User) Send(id uint32, msg proto.Message) {
	u.Client.Send(id, msg)
}

func (u *User) Update() {

}

func (u *User) OnMessage(id uint32, data []byte) {

}
