package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"yocto/src/log"
	"yocto/src/storage"
	"yocto/src/yoctoparser"
	"yocto/src/yoctoparser/grammer/parser"

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
		fmt.Println("Next loop")
	}
}

// deal the commands
func cmd(conn net.Conn) {
	defer conn.Close()
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
		// buffer.WriteString("Result for ")
		buffer.Write(buf[:n])
		buffer.WriteString(strconv.FormatBool(cmd_parse(str)) + "\n")
		buffer.WriteString(conn.RemoteAddr().String())
		conn.Write(buffer.Bytes())
		log.Trace.Println("Waiting for command .")
	}
}

//TODO add the yoctoparser,and the cmd should be a json or parsetree
func cmd_parse(cmd string) bool {
	cmd = strings.Replace(cmd, "\n", "", -1)
	cmd = strings.Replace(cmd, ";", "", -1)

	return cmd_run(cmd)
}

//TODO to be rebuild for yoctoparser
func cmd_run(cmd string) bool {
	fmt.Println(cmd)
	sqlObject := yoctoparser.YoctoPaser(cmd, "")
	fmt.Println(sqlObject.SQLType)
	switch sqlObject.SQLType {

	case parser.MySqlParserRULE_ddlStatement:
		{
			return cmd_ddl(*sqlObject)
		}

	case parser.MySqlParserRULE_dmlStatement:
		{
			return cmd_dml(*sqlObject)
		}

	case parser.MySqlParserRULE_transactionStatement:
		return cmd_tx(*sqlObject)

	default:
		fmt.Printf("statement type doesn't support yet")
		return false
	}

	//if strings.HasPrefix(cmd, "create") {
	//	return cmd_ddl(cmd)
	//}
}

func cmd_dml(obj yoctoparser.SQLObject) bool {
	// meta := obj.TableList
	fmt.Println(obj.SQLCommand)
	switch obj.SQLCommand {
	case parser.MySqlParserRULE_insertStatement:
		{
			io_insert := new(storage.IORequest)
			io_insert.Save()
		}
	case parser.MySqlParserRULE_updateStatement:
		{
			return true
		}
	}
	return true
}

func cmd_tx(obj yoctoparser.SQLObject) bool {
	return true
}

func cmd_ddl(obj yoctoparser.SQLObject) bool {
	switch obj.SQLCommand {
	case parser.MySqlParserRULE_createTable:
		{
			return storage.DDL_ColumnCreateTable(obj)
		}
	case parser.MySqlParserRULE_alterTable:
		{

		}
	default:
		fmt.Println("sqlCommand doesn't support yet ")
	}

	return false
}

func cmd_dml(obj yoctoparser.SQLObject) bool {
	switch obj.SQLCommand {
	case parser.MySqlParserRULE_insertStatement:
		{
			return storage.DML_InsertStatement(obj)
		}
	case parser.MySqlParserRULE_deleteStatement:
		{

		}
	case parser.MySqlParserRULE_updateStatement:
		{

		}

	case parser.MySqlParserRULE_selectStatement:
		{

		}

	default:
		fmt.Println("sqlCommand doesn't support yet ")
	}

	return false
}

func cmd_tx(obj yoctoparser.SQLObject) bool {
	return true
}

// handle buffer io etc
func deamon() {
	return
}
