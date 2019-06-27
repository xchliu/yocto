package storage

import (
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
	log.Info.Printf("Storage init start...")
	//TODO init upto the config size with zero
	//redo
	go redo_loop()
	//memtable
	go memtable_loop()
}

func redo_loop() {
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

// scan the memory tables that big enough ,and push it into the merge flow
//
func memtable_loop() {
	for {
		// if int(len(MEMTABLE)) > 0 {

		// } else {
		time.Sleep(time.Duration(1) * time.Second)
		// }
		log.Error.Printf("Table cache usage: %d\n", len(MEMTABLE))
	}
}

type IORequest struct {
	iotype   int    //
	metadata string //db.table
	key      string
	data     string
}

//IO handle for threads
//step 1 write the redo log
//step 2 write the memory table
func (ior IORequest) save() {
	//TODO trans to redo formate
	redo_log := ior.data
	REDOBUFFER <- redo_log
	//memtable
	if val, ok := MEMTABLE[ior.metadata]; ok {
		val[ior.key] = ior.data
	} else {
		//TODO
		if len(MEMTABLE) == 32 {
			log.Trace.Println("Table cache is null,wait for perge!")
		} else {
			MEMTABLE[ior.metadata][ior.key] = ior.data
		}
	}
}

func (ior IORequest) get() {

}

func file_read() {

}

func file_write() {

}
