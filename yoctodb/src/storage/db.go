package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Create_db(name string) bool {
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
