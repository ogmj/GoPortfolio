package server

import (
	"encoding/binary"
	"errors"
	"fmt"
	"library/socket"
)

type Session struct {
	connection *socket.TCP
}

var (
	session Session
	buffer  []byte
)

func Init() {
	buffer = make([]byte, 65536)
}

// TODO : 한개만 전역적으로 사용하므로 싱글톤을 생성하자.
func GetSession() *Session {
	if session.connection == nil {
		session.connection = new(socket.TCP)
		Init()
	}
	return &session
}

// TODO : 연결 함수
func (s *Session) Connect(ip string, port uint) error {
	if session.connection.Connect(ip, port) {
		s.connection.ConnectionHandler(func() {
			for s.extractPacket(s.connection, buffer) != nil {
				s.receiver(buffer)
			}
		}, s.closer)
		return nil
	}
	return errors.New("Connection Refused")
}

// TODO : 패킷을 보내는 부분
func Send() {
}

// TODO : TCP통신의 특성을 감안하여 여러 개의 스트림에 한 개의 패킷이 올 것을 대비하여 처리
func (s *Session) extractPacket(connection *socket.TCP, buffer []byte) []byte {
	const headerSize int = 4
	rawHeader, err := connection.Peek(headerSize)
	if err != nil {
		return nil
	}

	rawSize := rawHeader[2:4]
	size := binary.LittleEndian.Uint16(rawSize)

	err = connection.Read(buffer, int(size))
	if err != nil {
		return nil
	}

	return buffer
}

// TODO : 패킷을 받는 부분
func (s *Session) receiver(packet []byte) {
	m := packet[0:2]
	msg := binary.LittleEndian.Uint16(m)
	switch msg {
	case 10000:
		fmt.Println("test")
	}
}

// TODO : 세션을 닫을때 처리하는 부분
func (s *Session) closer() {
	s.connection = nil
}
