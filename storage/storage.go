package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"yocto/src/log"
)

//main loop
//init the memory tables
//init log files

var datadir string = "/Users/xchliu/Documents/workspace/yoctodb/yoctodb/data/"

func StorageInit() {
	fmt.Printf("storage init")
	redofile := filepath.Join(datadir, "ydblog")
	f, _ := os.OpenFile(redofile, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	for {
		if int(len(REDOBUFFER)) > 0 {
			redo := <-REDOBUFFER
			f.WriteString(redo)
		} else {
			time.Sleep(time.Duration(1) * time.Second)
		}
		log.Error.Printf("Redo buffer usage: %d / %d\n", len(REDOBUFFER), cap(REDOBUFFER))
	}
}

//IO handle for threads
type IORequest struct {
	iotype   int    //
	metadata string //db.table
	key      string
	data     string
}

func (ior IORequest) save() {
	//TODO trans to redo formate
	redo_log := ior.data
	REDOBUFFER <- redo_log
}

func (ior IORequest) get() {

}

func file_read() {

}

func file_write() {

}
