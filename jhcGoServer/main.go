package main

import (
	"bufio"
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

	jhcGOTcpServer.AsyncAccept(func(connection *socket.TCP) {
		//TODO : 통신 버퍼를 셋하자.
		buffer := make([]byte, 32768)
		//TODO : 셋한 통신 버퍼에 TCP 소켓에 있는 값을 리시브한다.
		err := connection.ConnectionHandler(buffer)
		//TODO : 에러인 경우 처리하는 함수를 추가한다.
		if err != nil {
			connection.Close()
		}
		//TODO : 위의 작업을 처리하는 함수는 TCP스트럭트에서 처리하도록 한다.
	})
}

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return text
}

func main() {
	fmt.Println("JhcGoServer")

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
