package main

import (
	"net"
	//"os"
	"fmt"
	"os"

	"./lib"


)


func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8193")
	checkError(err);
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	fanuc.Connect(conn);
	fanuc.SysInfo(conn);
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
