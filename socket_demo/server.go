package main

import (
	"bufio"
	"golang-demo/socket_demo/proto"
	"log"
	"net"
	"net/rpc"
)

type MessageHandler interface {
	Handler(channel *Channel, message *proto.Message)
}

type Channel struct {
	conn        net.Conn
	rd          bufio.Reader
	wr          bufio.Writer
	messageChan chan proto.Message
}

func (c *Channel) SendMsg(msg string) {

	if _, err := c.wr.WriteString(msg); err != nil {
		log.Printf("write buf error:%v\n", err)
	}
	if err := c.wr.Flush(); err != nil {
		log.Printf("flush write buf error:%v\n", err)
	}
}

func (c *Channel) Close() {
	if err := c.wr.Flush(); err != nil {
		log.Printf("flush write buf error:%v\n", err)
	}
	c.conn.Close()
}

func handleRead(rd *bufio.Reader) {

	//  获取消息包长度 ，4个字节，32位,
	packageLenBytes := make([]byte, 4)
	n, err := rd.Read(packageLenBytes)
	if err != nil {
		log.Panicf("read message error：%v\n", err)
	}
	if n < 4 {
		log.Panicf("read package length fail, readed bytes size:%d, expect: 4\n")
	}
	packageLen := proto.BytesToUint32(packageLenBytes)
	// 创建字节数组存储本条消息
	messageBytes := make([]byte, packageLen-4)

	n, err = rd.Read(messageBytes)
	if err != nil {
		log.Panicf("read message error：%v\n")
	}
	if n < int(packageLen)-4 {
		log.Panicf("read message package fail, readed bytes size:%d, expect: %d\n", n, packageLen-4)
	}

	message := proto.Decode(&messageBytes, packageLen)
	log.Printf("receive message: %s\n", message.String())
	// 组装message
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	for {
		handleRead(rd)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		go handleConn(conn)
	}

	rpc.NewServer()

}
