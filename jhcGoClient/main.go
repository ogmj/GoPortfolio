package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

func readString(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	} else {
		text = strings.TrimRight(text, "\n")
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
	} else {
		sigs := make(chan os.Signal, 1)
		pipeline := make(chan bool, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			pipeline <- true
		}()
		<-pipeline
	}
}
