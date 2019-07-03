package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"yocto/src/log"
	"yocto/src/parser/grammer/parser"
	"yocto/src/storage"

	"github.com/Unknwon/goconfig"
)

var SERVICE_ADDR = "0.0.0.0:6180"
var CFG goconfig.ConfigFile

func init() {
	var logfile string
	CFG, err := goconfig.LoadConfigFile("../../data/yocto.cnf")
	if err != nil {
		fmt.Println(err)
	} else {
		logfile, _ = CFG.GetValue("global", "logfile")
	}
	log.LogInit(logfile)
	log.Info.Printf("init system")
	go storage.StorageInit()
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
		//thread handle? use sync.pool
		go cmd(conn)
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
	cmd = strings.Replace(cmd, "\n", "", -1)
	cmd = strings.Replace(cmd, ";", "", -1)
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
	obj_action := cmd_arrary[0]
	obj_type := cmd_arrary[1]
	obj_name := cmd_arrary[2]
	//TODO Get the session info
	//db=session.db
	db := "test"
	//TODO tbd
	//	obj_extra := cmd_arrary[3:]
	fmt.Println(obj_type)
	switch obj_type {
	case "database":
		return storage.DDL_db(obj_name, obj_action)
	case "table":
		return storage.DDL_Table()(db, obj_name, obj_action, strings.Join(cmd_arrary[3:], ""))
	case "insert":
		return storage.DML_Table(db, obj_name, strings.Join(cmd_arrary[3:], ""))
	default:
		fmt.Println("Unknown type %s", obj_type)
	}
	return false
}

// handle buffer io etc
func deamon() {
	return
}
