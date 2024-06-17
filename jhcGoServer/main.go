package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return text
}

func main() {
	fmt.Println("JhcGoServer")
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
