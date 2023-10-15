package gate

import (
	"encoding/binary"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"server/game/core/log"
)

const (
	DataLen      = 4                  //消息长度
	MinPacketLen = DataLen            //最小消息长度
	MaxPacketLen = (2 << 8) * DataLen //最大消息长度
)

// 数据包结构（长度+数据）
type Packet struct {
	len  uint32
	data []byte
}

// 返回Packet里的数据流
func (p *Packet) GetData() []byte {
	return p.data
}

// 序列化
func (p *Packet) Serialize() []byte {
	dataLen := uint32(len(p.data))
	buff := make([]byte, MinPacketLen+dataLen)
	binary.BigEndian.PutUint32(buff, dataLen)
	if dataLen > 0 {
		copy(buff[MinPacketLen:], p.data)
	}
	return buff
}

// 解码
func (p *Packet) UnmarshalPB(msg proto.Message) error {
	return proto.Unmarshal(p.data, msg)
}

func NewPacket(msg proto.Message) *Packet {
	p := &Packet{}
	if data, err := proto.Marshal(msg); err == nil {
		p.data = data
		p.len = uint32(len(p.data))
	} else {
		log.Error(fmt.Sprintf("[NewPacket] proto marshal error:%v", err))
		return nil
	}
	return p
}

// 从客户端读一个协议
type MsgProtocol struct {
}

func (p *MsgProtocol) ReadPacket(r io.Reader) (*Packet, error) {
	buff := make([]byte, MinPacketLen)

	//data length
	if size, err := io.ReadFull(r, buff); err != nil {
		log.Debug(fmt.Sprintf("MsgProtocol ReadFull size %d", size))
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(buff)

	if dataLen > MaxPacketLen {
		return nil, errors.New("data max")
	}

	//id
	msg := &Packet{
		len: dataLen,
	}

	//data
	if msg.len > 0 {
		msg.data = make([]byte, msg.len)
		if _, err := io.ReadFull(r, msg.data); err != nil {
			return nil, err
		}
	}

	return msg, nil
}
