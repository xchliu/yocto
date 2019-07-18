package test

import (
	"fmt"
	"net"
	"testing"
)

var SERVER = "127.0.0.1:6180"

func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	b.StartTimer() //重新开始时间
	cmd := "INSERT INTO A.T VALUES (1, 'ABC',2,3,'bbb');"
	conn, _ := net.Dial("tcp", SERVER)
	for i := 0; i < b.N; i++ {
		conn.Write([]byte(cmd))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed , err : %v\n", err)
		}
		res := string(buf[:n])
		fmt.Printf(res)
	}
}
