package main

import (
	"net"
	//"os"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"


)
var fnHeader = []byte{ 0xa0,  0xa0,  0xa0, 0xa0, 0x00, 0x01, 0x21, 0x01 };
var key_systeminfo = 		[]byte{ 0x00, 0x01, 0x00, 0x18 };
var key_systeminfo2 = 	[]byte{ 0x00, 0x02, 0x00, 0x18 };
var DEFAULT_REQ_LENGTH = 0x1c;
var DEFAULT_REQ_COUNT_LENGTH = 2;
var DEFAULT_TOTAL_SIZE_LENGTH = 2;
var DEFALUT_REQUEST_HEADER_LENGTH = 8;

func main() {
	//addr :=net.ParseIP("106.243.233.236:80");
	//fmt.Println(addr.String());
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "it.yhsbearing.com:8193")
	checkError(err);
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.Write([]byte{  0xa0,0xa0,0xa0,0xa0,0x00,0x01,0x01,0x01,0x00,0x02,0x00,0x01,0x02 });
	v :=makeRequestPacket(2);
	makeRequest(v, key_systeminfo, 0);
	makeRequest(v, key_systeminfo2, 0);
	wc,_:=conn.Write(v.Bytes());
	fmt.Println("write ",wc);
	result ,err := ioutil.ReadAll(conn);
	fmt.Println(string(result));
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func  makeRequestPacket( count int) *bytes.Buffer {
	var  totalLen = DEFAULT_REQ_COUNT_LENGTH + DEFAULT_REQ_LENGTH * count;
	var  len = DEFALUT_REQUEST_HEADER_LENGTH + DEFAULT_TOTAL_SIZE_LENGTH+ totalLen;
	var b bytes.Buffer;
	b.Grow(len);
	b.Reset();
	fmt.Print("cap ",b.Cap(),"\n");
	b.Write(fnHeader);
	fmt.Print("current#1 ",b.Len(),"\n");
	b.WriteByte(byte(totalLen))
	fmt.Print("current#2 ",b.Len(),"\n");
	b.WriteByte(byte(count));
	fmt.Print("current#3 ",b.Len(),"\n");
	return &b;
}
func makeRequest(buf *bytes.Buffer, key []byte , parameter ...int)  {
	var  p = buf.Len();

	buf.WriteByte(byte(DEFAULT_REQ_LENGTH)); // 요청길이 2byte
	fmt.Print("c#1 ",buf.Len(),"\n");
	buf.WriteByte(byte(1)); // 함수이름 시작 2byte
	fmt.Print("c#2 ",buf.Len(),"\n");
	buf.Write(key); // key.length
	fmt.Print("c#3 ",buf.Len(),"\n");
	for  _, num:= range parameter{
		buf.WriteByte(byte(num));
	}
	buf.WriteByte( byte(p + DEFAULT_REQ_LENGTH));
	fmt.Print("last ",buf.Len(),"\n");
}
