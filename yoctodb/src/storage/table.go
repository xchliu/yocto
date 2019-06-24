package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Create_table(name string) bool {
	//TODO read from the config file
	name = strings.Replace(name, "\n", "", -1)
	name = strings.Replace(name, " ", "", -1)
	fmt.Println("Create new table :" + name)
	datadir := "/Users/xchliu/Documents/workspace/yoctodb/yoctodb/data/"
	tabledir := filepath.Join(datadir, name)
	fmt.Println(tabledir)
	err := os.Mkdir(tabledir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	os.Chmod(tabledir, os.ModePerm)
	return true
}
