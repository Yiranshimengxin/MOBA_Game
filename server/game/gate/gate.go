package gate

import (
	"fmt"
	"net"
	"server/game/core/log"
)

type Server struct {
	host       string //服务器监听的IP地址
	port       int    //服务器监听的端口号
	AcceptChan chan *Client
	close      chan struct{} //信号量，查看服务器是否关闭
	listener   net.Listener
}

func NewServer(host string, port int) *Server {
	return &Server{
		host:       host,
		port:       port,
		AcceptChan: make(chan *Client, 100),
		close:      make(chan struct{}),
	}
}

func (s *Server) Run() {
	//1. 监听端口
	//2. 接受客户端连接
	//3. 创建一个session
	//4. 将session加入到sessionMgr中
	//5. 启动session
	//6. 处理session的信息
	//7. 将session从sessionMgr中删除
	//8. 关闭session
	//9. 关闭连接

	//在返回之前会先执行defer
	defer func() {
		log.Info("gate stop run ...")
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	s.listener = listener

	//主循环
	for {
		select {
		//退出条件
		case <-s.close:
			log.Info("gate server close")
			return
		default:
		}
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		client := NewClient(conn, &MsgProtocol{})
		s.AcceptChan <- client
	}
}

func (s *Server) ReadMQ() <-chan *Client {
	return s.AcceptChan
}

func (s *Server) Stop() {
	log.Info("gate server start stop ...")
	close(s.close)
	err := s.listener.Close()
	if err != nil {
		log.Error(err.Error())
		return
	}
}
