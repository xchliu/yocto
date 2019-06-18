package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var WELCOME string = "Welcome to the Yoctodb!"
var PROMOTE string = "yoo> "
var CMDEOF byte = '\n'
var SERVER = "127.0.0.1:6180"

func main() {
	fmt.Println(WELCOME)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(PROMOTE)
		cmd, err := reader.ReadString(CMDEOF)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Print(cmd)
		if cmd == "exit" {
			os.Exit(0)
		}
	}
}
