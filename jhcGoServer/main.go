package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"library/socket"
	"os"
	"strings"
)

var (
	jhcGOTcpServer *socket.Listener
)

func initNetwork() {
	initTcpSocketNetwork()
}

func initTcpSocketNetwork() {
	jhcGOTcpServer = new(socket.Listener)
	err := jhcGOTcpServer.Listen(9999)

	if err != nil {
		panic("Failed to initTcpSocketNetword")
	}

	//TODO : 비동기 accept함수 구현
	jhcGOTcpServer.AsyncAccept(func(connection *socket.TCP) {
		go func() {
			//TODO : 통신 버퍼를 셋하자.
			buffer := make([]byte, 32768)
			//TODO : 셋한 통신 버퍼에 TCP 소켓에 있는 값을 리시브한다.
			connection.ConnectionHandler(func() {
				//TODO : TCP통신 특성상 나뉘어 오는 패킷 처리
				for extractPacket(connection, buffer) != nil {
					//TODO : 받은 패킷 처리
					packetReceiver(connection, buffer)
				}
			}, func() {
				//TODO : 에러인 경우 접속을 해제한다.
				closeTcpSocket(connection)
			})
		}()
	})
}

func extractPacket(connection *socket.TCP, buffer []byte) []byte {
	const headerSize int = 4

	rawHeader, err := connection.Peek(headerSize)
	if err != nil {
		return nil
	}
	rawSize := rawHeader[2:4]
	size := binary.LittleEndian.Uint16(rawSize)

	err = connection.Read(buffer, int(size))
	return nil
}

func closeTcpSocket(connection *socket.TCP) {
	connection.Close()
}

func packetReceiver(connection *socket.TCP, packet []byte) {
	const headerSize int = 4
	rawMsg := packet[0:2]
	msg := binary.LittleEndian.Uint16(rawMsg)

	switch msg {
	case 10001:
		fmt.Println("test")
	}

}

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return text
}

func main() {
	fmt.Println("JhcGoServer Start")

	initNetwork()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text := readString(reader)
		cmd := strings.Split(text, " ")
		if len(cmd) > 0 {
			if cmd[0] == "exit" {
				break
			}
		}
	}
}
