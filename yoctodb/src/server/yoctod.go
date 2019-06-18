package server


import (
	"os"
	"net"
	"fmt"
)

var SERVICE_ADDR="0.0.0.0:6180"

// handle the service
func main (){
	listener, err := net.Listen("tcp", SERVICE_ADDR)
    if err != nil {
        fmt.Printf("listen fail, err: %v\n", err)
        os.Exit(1)
    }	
	for (
		conn,err := listener.Accept()
		if err != nil {
            fmt.Printf("accept fail, err: %v\n", err)
            continue
        } 
        go cmd(conn)		//thread ?
	)
}
// deal the commands
func cmd(conn net.Conn){
	defer conn.Close()
	for {
		var buf [128]byte
 	    n, err := conn.Read(buf[:])
        if err != nil {
            fmt.Printf("read from connect failed, err: %v\n", err)
            break
        }
        str := string(buf[:n])
        fmt.Printf("receive from client, data: %v\n", str)	
	}
	conn.Write(str)
}
// handle buffer io etc
func deamon(){
	
}