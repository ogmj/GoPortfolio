package socket

import (
	"errors"
	"fmt"
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

// socketBuffer 관련 함수 /////////////////////////////////////////
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

// TCP 통신 관련 함수 /////////////////////////////////////////////
func (t *TCP) Connect(address string, port uint) bool {
	var err error
	host := address + ":" + fmt.Sprint(port)
	t.connection, err = net.Dial("tcp", host)
	if err != nil {
		return false
	}
	t.connected = true
	t.buffer.initSocketBuffer()
	return t.connected
}

func (t *TCP) Close() {
	_ = t.connection.Close()
	t.connected = false
}

func (t *TCP) Send(buf []byte) {
	_, _ = t.connection.Write(buf)
}

func (t *TCP) Peek(size int) ([]byte, error) {
	return t.buffer.peek(size)
}

func (t *TCP) Read(buffer []byte, size int) error {
	return t.buffer.read(buffer, size)
}
