package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func Goid() int {
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
	return id
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func DirScan(dirPth string) (dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			//			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		}
	}
	return dirs, nil
}

func FileScan(dirPth string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if !fi.IsDir() {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
