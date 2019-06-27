package storage

var REDOBUFFER = make(chan string, 128)
var CHANGEBUFFER = make(chan string, 128)

//replace as the config value
var MEMTABLE = make(map[string]map[string]string)
