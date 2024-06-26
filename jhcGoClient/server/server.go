package server

import "library/socket"

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
	return nil
}

// TODO : 패킷을 보내는 부분
func Send() {
}

// TODO : TCP통신의 특성을 감안하여 여러 개의 스트림에 한 개의 패킷이 올 것을 대비하여 처리
func extractPacket() {
}

// TODO : 패킷을 받는 부분
func receiver() {
}
