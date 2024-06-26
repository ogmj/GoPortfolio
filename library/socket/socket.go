﻿package socket

import (
	"encoding/binary"
	"errors"
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"net"
)

var (
	fbHeader flatbuffers.UOffsetT = 4
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
	listener net.Listener
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

func (t *TCP) IsConnected() bool {
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

func (t *TCP) ConnectionHandler(f func(), d func()) {
	bufBytes := make([]byte, 65536)
	for {
		n, err := t.connection.Read(bufBytes)
		if err != nil {
			if n == 0 {
				t.connected = false
				d()
			}
			break
		}
		if n > 0 {
			t.buffer.write(bufBytes[:n])
			f()
		}
	}
}

// TODO : 플랫버퍼 헤더 세팅
func (t *TCP) CreateHeaderForFlatBuffer(msg uint16, size uint16) []byte {
	hp := make([]byte, fbHeader)
	binary.LittleEndian.PutUint16(hp[:2], msg)
	binary.LittleEndian.PutUint16(hp[2:4], size)
	return hp
}

// TODO : 바이트를 실제로 네트워크로 보낼때 사이즈(헤더포함)
func (t *TCP) NetworkSize(packet []byte) uint16 {
	size := len(packet) + int(fbHeader)
	return uint16(size)
}

// TCP Listener 함수 관련 /////////////////////////////////////////////////
func (l *Listener) Listen(port uint) error {
	ipNport := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", ipNport)
	if err != nil {
		return err
	}
	l.listener = listener
	return nil
}

func (l *Listener) AsyncAccept(acceptCallback func(*TCP)) {
	go func() {
		for {
			conn, _ := l.listener.Accept()
			connection := new(TCP)
			connection.connection = conn
			connection.connected = true
			connection.buffer.initSocketBuffer()

			acceptCallback(connection)
		}
	}()
}
