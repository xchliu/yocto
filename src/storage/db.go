package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type database struct {
	name    string
	charset string
}

func DDL_db(name string, action string) bool {
	var db database
	db.name = name
	db.charset = "utf8"
	if action == "create" {
		return create_db(db.name)
	}
	return false
}

func create_db(name string) bool {
	//TODO read from the config file
	name = strings.Replace(name, "\n", "", -1)
	name = strings.Replace(name, " ", "", -1)
	fmt.Println("Create new database :" + name)
	datadir := "/Users/xchliu/Documents/workspace/yoctodb/yoctodb/data/"
	dbdir := filepath.Join(datadir, name)
	fmt.Println(dbdir)
	err := os.Mkdir(dbdir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	os.Chmod(dbdir, os.ModePerm)
	return true
}
