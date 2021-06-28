package main

import (
	"fmt"
	"golang-demo/socket_demo/proto"
	"log"
	"net"
	"time"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("connect to server error", err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	var version uint16 = 1
	var op uint32 = 2
	var seq uint32 = 0

	for {
		seq = seq + 1
		content := fmt.Sprintf("say hello, cur_seq: %d", seq)
		message := proto.Message{
			Ver:  version,
			Op:   op,
			Seq:  seq,
			Body: []byte(content),
		}
		messageBytes := proto.Encode(&message)
		_, err := conn.Write(messageBytes)
		if err != nil {
			log.Fatalf("write bytes error:%v\n", err)
		}
		log.Printf("send message: %s\n", message.String())
		log.Printf("message bytes: %v\n", messageBytes)
		time.Sleep(1 * time.Second)
	}

}
