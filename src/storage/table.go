package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type table struct {
	name    string
	db      string
	charset string
	define  string
}

func DDL_Table(db string, name string, action string, define string) bool {
	var tb table
	tb.name = name
	tb.db = db
	tb.charset = "utf8"
	tb.define = define
	tb.name = strings.Replace(tb.name, "\n", "", -1)
	tb.name = strings.Replace(tb.name, " ", "", -1)

	if action == "create" {
		return Create_table(tb)
	}
	return false
}

func DML_Table(db string, name string, data string) bool {
	var ior IORequest
	ior.data=data
	ior.metadata=strings.Join([db,name],'.')
	ior.iotype=1
	ior.key=get_next_id(ior.metadata)
	return ior.save()
}

func Create_table(tb table) bool {
	//TODO read from the config file
	fmt.Println("Create new table :" + tb.name)
	datadir := "/Users/xchliu/Documents/workspace/yoctodb/yoctodb/data/"
	tabledir := filepath.Join(datadir, tb.db, tb.name)
	fmt.Println("Init table in :" + tabledir)
	obj_def, err := os.OpenFile(tabledir+".def", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	obj_def.WriteString(tb.define)
	obj_data, err := os.OpenFile(tabledir+".ydb", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	obj_data.WriteString("")
	return true
}

func Get_Table_Define(db string,name string){
	cfg
}