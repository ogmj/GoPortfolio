package main

import (
	"bufio"
	"fmt"
	"library/socket"
	"os"
	"runtime"
	"strings"
)

var (
	jhcGOClient *socket.TCP
)

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	}
	return text
}

func main() {
	fmt.Println("JhcGoClient Start")

	initConnect()

	if runtime.GOOS == "windows" {
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
}

func initConnect() {
	//TODO : 세션을 하나 가져온다.
	//TODO : 해당 세션을 이용하여 서버와 연결을 시도한다.
}
