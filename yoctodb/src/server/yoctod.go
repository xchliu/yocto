package main

import (
	"bytes"
	"fmt"
	//"lib"
	"log"
	"net"
	"os"
	"runtime"
	"storage"
	"strconv"
	"strings"
)

var SERVICE_ADDR = "0.0.0.0:6180"

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

// handle the service
func main() {
	listener, err := net.Listen("tcp", SERVICE_ADDR)
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		go cmd(conn) //thread handle?
	}
}

// deal the commands
func cmd(conn net.Conn) {
	//	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client, data: %v\n", str)
		var buffer bytes.Buffer
		buffer.WriteString(Goid())
		buffer.WriteString("\nRoger for ")
		buffer.Write(buf[:n])
		conn.Write(buffer.Bytes())
	}
}

// handle buffer io etc
func deamon() {
	return
}

func Goid() string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	fmt.Print(id)
	return idField
}
