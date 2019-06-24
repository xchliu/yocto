package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
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
		buffer.WriteString("Result for ")
		buffer.Write(buf[:n])
		buffer.WriteString(strconv.FormatBool(cmd_parse(str)) + "\n")
		conn.Write(buffer.Bytes())
	}
}

//TODO add the parser,and the cmd should be a json or parsetree
func cmd_parse(cmd string) bool {
	return cmd_run(cmd)
}

//TODO to be rebuild for parser
func cmd_run(cmd string) bool {
	fmt.Println(cmd)
	if strings.HasPrefix(cmd, "create") {
		return cmd_ddl(cmd)
	}
	return true
}

//TODO to be rebuild for parser
func cmd_ddl(cmd string) bool {
	fmt.Println(cmd)
	cmd_arrary := strings.Split(cmd, " ")
	obj_type := cmd_arrary[1]
	obj_name := cmd_arrary[2]
	switch obj_type {
	case "database":
		return storage.Create_db(obj_name)
	case "table":
		return storage.Create_table(obj_name)
	default:
		fmt.Println("Unknown type %s", obj_type)
	}
	return false
}

// handle buffer io etc
func deamon() {
	return
}
