package storage

import (
	"sync"
	"github.com/Unknwon/goconfig"
	"yocto/src/log"
)
)
type buffer []string

var REDOBUFFER = make(chan string, 128)
var CHANGEBUFFER = make(chan string, 128)

//replace as the config value
var MEMTABLE = make(map[string]map[string]string)


//buffers
var CONFIG = make(map[string]map[string]string)

//buffer pool
type BUFFERPOOL struct {
	name  string
	buf   buffer
	size  int
	value interface{}
}

var bp = sync.Pool{
		New: func() interface{} {
			return new(BUFFERPOOL)
		}
	}

func buffer_init() bool{
	//config buffer
	cfg, err := goconfig.LoadConfigFile("../../data/yocto.cnf")
	if err != nil {
		log.Error.Println(err)
		//TODO reset as the default values
	} else {
		for _,section := range cfg.GetSectionList(){
			CONFIG[section]=make(map[string][string])
			for key,value := range cfg.GetSection(section){
				CONFIG[section][key]=value
			}
		}
	}
	//table meta
	datadir = GetConf('datadir')
	
	var table_meta=bp.Get().(*BUFFERPOOL)
	table_meta.buf=CFG
	table_meta.value=	
}

func BufferRecycle(){
	
}

func GetConf(key string) string {
	value,ok := CONFIG[key]
	if (ok){
		return value
	}else{
		return nil
	}
}