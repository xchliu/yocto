package storage

import (
	"strings"
	"yocto/src/lib"
	"yocto/src/log"

	"github.com/Unknwon/goconfig"
)

type buffer []string

var REDOBUFFER = make(chan string, 256)
var CHANGEBUFFER = make(chan string, 128)

//replace as the config value
var MEMTABLE = make(map[string]map[string]string)

//buffers
var CONFIGBUFFER = make(map[string]map[string]string)

//buffer pool
type BUFFERPOOL struct {
	name  string
	buf   buffer
	size  int
	value interface{}
}

// sequense
var SEQUENCE = make(map[string]int)

func buffer_init() bool {
	//config buffer
	cfg, err := goconfig.LoadConfigFile("../../data/yocto.cnf")
	if err != nil {
		log.Error.Println(err)
		//TODO reset as the default values
	} else {
		for _, section := range cfg.GetSectionList() {
			CONFIGBUFFER[section] = make(map[string]string)
			CONFIGBUFFER[section], _ = cfg.GetSection(section)
		}
	}
	//table meta
	tablebuffer_int()
	//sequence table
	return true
}

func tablebuffer_int() bool {
	datadir := GetConf("datadir")
	dbs, _ := lib.DirScan(datadir)
	//TODO put the table meta into table cache
	for index := range dbs {
		files, _ := lib.FileScan(dbs[index])
		for index := range files {
			ok := strings.HasSuffix(files[index], ".def")
			if ok {
				continue
			}
		}
	}
	return true
}

func BufferRecycle() {

}

func GetConf(key string) string {
	//TODO return the values match the given key
	//deal with section
	value, ok := CONFIGBUFFER[key]
	if ok {
		return value[key]
	} else {
		return ""
	}
}
