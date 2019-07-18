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

func StorageInit() {
	log.Info.Printf("Storage init start...")
	//buffer pools
	go buffer_init()
	//TODO init upto the config size with zero
	//redo
	go redo_loop()
	//memtable
	go memtable_loop()

}

func redo_loop() {
	redofile := filepath.Join(GetConf("datadir"), "ydblog")
	f, err := os.OpenFile(redofile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Error.Println("Open REDO file failed: ", err)
	}
	fmt.Println(redofile)
	for {
		if int(len(REDOBUFFER)) > 0 {
			redo := <-REDOBUFFER
			_, err := f.WriteString(redo)
			if err != nil {
				log.Error.Println("Flush redo log failed:", err)
				fmt.Println("Flush redo log failed:", err)
			}
		} else {
			time.Sleep(time.Duration(1) * time.Second)
		}
		log.Trace.Printf("Redo buffer usage: %d / %d\n", len(REDOBUFFER), cap(REDOBUFFER))
	}
	f.Close()
}

// scan the memory tables that big enough ,and push it into the merge flow
//
func memtable_loop() {
	for {
		// if int(len(MEMTABLE)) > 0 {

		// } else {
		time.Sleep(time.Duration(1) * time.Second)
		// }
		log.Trace.Printf("Table cache usage: %d\n", len(MEMTABLE))
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
func (ior IORequest) Save() bool {
	//TODO trans to redo format
	redo_log := ior.data
	REDOBUFFER <- redo_log
	//memtable
	if val, ok := MEMTABLE[ior.metadata]; ok {
		val[ior.key] = ior.data
	} else {
		//TODO block for free space
		if len(MEMTABLE) == 32 {
			log.Trace.Println("Table cache is null,wait for perge!")
		} else {
			if MEMTABLE[ior.metadata] == nil {
				MEMTABLE[ior.metadata] = make(map[string]string)
			} else {
				MEMTABLE[ior.metadata][ior.key] = ior.data
			}
		}
	}
	return true
}

func (ior IORequest) get() {

}

func file_read() {

}

func file_write() {

}
