package storage

var REDOBUFFER = make(chan string, 128)
var CHANGEBUFFER = make(chan string, 128)
