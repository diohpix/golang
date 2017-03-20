package main

import (
	"net"
	//"os"
	"fmt"
	"os"

	"./lib"


)


func main() {
	//addr :=net.ParseIP("106.243.233.236:80");
	//fmt.Println(addr.String());
	fanuc.Init();
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "it.yhsbearing.com:8193")
	checkError(err);
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	fanuc.Connect(conn);
	fanuc.SysInfo(conn);



//	fmt.Println("count :", binary.BigEndian.Uint16(tmp[:2]))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
