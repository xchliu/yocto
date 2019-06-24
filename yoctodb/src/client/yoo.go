package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var WELCOME string = "Welcome to the Yoctodb!"
var PROMOTE string = "yoo> "
var CMDEOF byte = ';'
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
		//fmt.Print(cmd)
		client(cmd)
		if cmd == "exit" {
			os.Exit(0)
		}
	}
}

func client(cmd string) {
	conn, err := net.Dial("tcp", SERVER)
	//	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	_, err = conn.Write([]byte(cmd))
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
