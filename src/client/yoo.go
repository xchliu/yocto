package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var WELCOME string = "Welcome to the Yoctodb!"
var PROMOTE string = "yoo> "
var CMDEOF byte = ';'
var SERVER = "127.0.0.1:6180"

func main() {
	fmt.Println(WELCOME)
	reader := bufio.NewReader(os.Stdin)
	conn, err := net.Dial("tcp", SERVER)
	if err != nil {
		fmt.Printf("Connect failed, err : %v\n", err.Error())
		return
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(8 * time.Hour))
	for {
		fmt.Print(PROMOTE)
		cmd, err := reader.ReadString(CMDEOF)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if cmd == "exit" {
			os.Exit(0)
		}
		// fmt.Print(cmd)
		client(conn, cmd)
	}
}

func client(conn net.Conn, cmd string) {
	_, err := conn.Write([]byte(cmd))
	if err != nil {
		fmt.Printf("write failed , err : %v\n", err)
	} else {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed , err : %v\n", err)
		}
		res := string(buf[:n])
		fmt.Printf(res)
	}
}
