package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Message struct {
	Ver  uint16 // 版本号
	Op   uint32 // 消息类型，如Ping，Pong, Text
	Seq  uint32 // 序列号
	Body []byte // 消息体 等于 PushMsg.Msg
}

func (m Message) String() string {
	return fmt.Sprintf("Ver:%d, Op:%d, Seq: %d, body: %s", m.Ver, m.Op, m.Seq, string(m.Body))
}

func Uint32ToBytes(n uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, n)
	return buf
}

func Uint16ToBytes(n uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, n)
	return buf
}

func BytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func Decode(messageBytes *[]byte, len uint32) Message {
	headerLen := BytesToUint16((*messageBytes)[0:2])
	return Message{
		Ver:  BytesToUint16((*messageBytes)[2:4]),
		Op:   BytesToUint32((*messageBytes)[4:8]),
		Seq:  BytesToUint32((*messageBytes)[8:12]),
		Body: (*messageBytes)[headerLen-4:],
	}
}

/**
 * Package Length     4 bytes
 * Header Length      2 bytes
 * Protocol Version   2 bytes
 * Operation          4 bytes
 * Sequence Id        4 bytes
 * Body        Package Length - Header Length
 */
func Encode(message *Message) []byte {

	headerLen := 16
	packageLen := 16 + len(message.Body)

	messageBytes := make([]byte, 0)

	bytes := Uint32ToBytes(uint32(packageLen))
	messageBytes = append(messageBytes, bytes...)
	messageBytes = append(messageBytes, Uint16ToBytes(uint16(headerLen))...)
	messageBytes = append(messageBytes, Uint16ToBytes(message.Ver)...)
	messageBytes = append(messageBytes, Uint32ToBytes(message.Op)...)
	messageBytes = append(messageBytes, Uint32ToBytes(message.Seq)...)
	messageBytes = append(messageBytes, message.Body...)

	return messageBytes

}
