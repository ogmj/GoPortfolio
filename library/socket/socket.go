package socket

import (
	"errors"
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

func (b *socketBuffer) initSocketBuffer() {
	b.data = make([]byte, 65536)
}

func (b *socketBuffer) write(packet []byte) {
	packetLength := len(packet)
	if dataLength := copy(b.data[b.offset:], packet); dataLength < packetLength {
		b.data = append(b.data, packet[dataLength:]...)
	}
	b.offset = b.offset + len(packet)
}

func (b *socketBuffer) peek(size int) ([]byte, error) {
	if size > b.offset {
		return nil, errors.New("peek func Packet Overflow")
	}
	return b.data[:size], nil
}

func (b *socketBuffer) read(buffer []byte, size int) error {
	if size > b.offset {
		return errors.New("read func Packet Overflow")
	}
	if len(buffer) < size {
		panic("read func panic")
	}
	b.offset = b.offset - size
	copy(buffer, b.data[:size])
	copy(buffer, b.data[size:])

	return nil
}
