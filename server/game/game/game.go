package game

import (
	"google.golang.org/protobuf/proto"
	"server/game/core/log"
	"server/game/gate"
	"server/game/pb"
)

func NewGame() *Game {
	return &Game{
		users:   make(map[uint64]*User, 10),
		clients: make([]*gate.Client, 0),
		userID:  10000,
	}
}

type Game struct {
	users   map[uint64]*User //容器
	clients []*gate.Client
	userID  int64
}

func (g *Game) OnAccept(c *gate.Client) {
	log.Debug("game accpet client")
	c.RegisterMsgHandler(g)
	g.clients = append(g.clients, c)
}

// 游戏循环
func (g *Game) Loop() {
	for _, v := range g.clients {
		v.Update()
	}

	for _, u := range g.users {
		u.Update()
	}
}

func (g *Game) HandleMsg(c *gate.Client, id uint32, data []byte) {
	if id == uint32(pb.MsgID_LOGIN_REQ) {
		g.OnLogin(c, id, data)
	} else {
		if u, ok := g.users[c.GetFD()]; ok {
			u.OnMessage(id, data)
		} else {
			log.Error("not find fd %d", c.GetFD())
		}
	}
}

func (g *Game) OnLogin(c *gate.Client, id uint32, data []byte) {
	req := &pb.PBLoginReq{}
	err := proto.Unmarshal(data, req)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Debug("game on user login uid %s", req.GetUid())

	u := NewUser(c, g.userID, req.GetUid())
	g.users[c.GetFD()] = u
	//u.registerMsgHandler()
	g.userID++

	rsp := &pb.PBLoginRsp{}
	rsp.Uid = uint64(g.userID)
	rsp.Account = req.GetUid()
	c.Send(uint32(pb.MsgID_LOGIN_RSP), rsp)
}
