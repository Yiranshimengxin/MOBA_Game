package gate

import (
	"google.golang.org/protobuf/proto"
	"net"
	"server/game/core/log"
	"server/game/pb"
	"time"
)

type MsgHandle interface {
	HandleMsg(c *Client, id uint32, data []byte)
}

func NewClient(conn net.Conn, protocol *MsgProtocol) *Client {
	c := &Client{
		conn:     conn,
		protocol: protocol,
	}
	if s, ok := conn.(*net.TCPConn); !ok {
		log.Error("not tcp conn")
	} else {
		f, err := s.File()
		if err != nil {
			log.Error("new connection %s", err.Error())
			return c
		}
		c.fd = uint64(f.Fd())
	}
	return c
}

type Client struct {
	conn     net.Conn //给每一个链接分配一个类
	protocol *MsgProtocol
	handler  MsgHandle //接口
	fd       uint64
}

// 获取文件的句柄
func (c *Client) GetFD() uint64 {
	return c.fd
}

// 注册回调函数
func (c *Client) RegisterMsgHandler(handler MsgHandle) {
	c.handler = handler
}

func (c *Client) Update() {
	c.Read()
}

func (c *Client) Read() {
	//设置超时时间
	err := c.conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return
	}

	p, err := c.protocol.ReadPacket(c.conn)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return
		}

		log.Error(err.Error())
		return
	}

	c.OnMessage(p)
}

func (c *Client) Send(id uint32, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	head := &pb.Head{
		Id:   int32(id),
		Data: data,
	}

	p := NewPacket(head)
	_, err = c.conn.Write(p.Serialize())
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func (c *Client) OnMessage(p *Packet) {
	head := &pb.Head{}
	err := proto.Unmarshal(p.GetData(), head)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Info("receive msg id:", head.GetId())

	if c.handler != nil {
		c.handler.HandleMsg(c, uint32(head.GetId()), head.GetData())
		return
	} else {
		log.Error("msg.handler is nil")
		return
	}
}
