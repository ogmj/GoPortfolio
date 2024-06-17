package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	}
	return text
}

func main() {
	fmt.Println("test")
	if runtime.GOOS == "windows" {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			text := readString(reader)
			cmd := strings.Split(text, " ")
			if len(cmd) > 1 {
				if cmd[0] == "exit" {
					break
				}
			}
		}
	}
}
