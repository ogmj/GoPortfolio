package socket

import (
	"net"
)

// socketBuffer 리시브용 소켓 버퍼
type socketBuffer struct {
	data   []byte
	offset int
}

// TCP
type TCP struct {
	connection net.Conn
	connected  bool
	buffer     socketBuffer
}

// Listener
type Listener struct {
	ln       net.Listener
	flagStop bool
}
