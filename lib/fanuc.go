package fanuc

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"net"

)

var cnchandle = []byte{  0xa0, 0xa0,  0xa0,  0xa0, 0x00, 0x01, 0x01, 0x01, 0x00, 0x02, 0x00,0x02 };
var fnHeader = []byte{ 0xa0,  0xa0,  0xa0, 0xa0, 0x00, 0x01, 0x21, 0x01 };
var key_systeminfo = []int8{ 0x00, 0x01, 0x00, 0x18 };
var key_systeminfo2 = 	[]int8{ 0x00, 0x02, 0x00, 0x18 };
var key_cnc_rdexeprog = 	[]int8{ 0x00, 0x01, 0x00, 0x20 };
var DEFAULT_REQ_LENGTH int16 = 0x1c;
var DEFAULT_REQ_COUNT_LENGTH int16= 2;
var DEFAULT_TOTAL_SIZE_LENGTH int16= 2;
var DEFALUT_REQUEST_HEADER_LENGTH int16= 8;

func Init(){
	fmt.Print("OK")
}
func Connect(conn *net.TCPConn ){
	conn.Write(cnchandle);
	readBodyLen :=readBodyLenth(conn);
	tmp := make([]byte, readBodyLen)
	conn.Read(tmp);
}
func SysInfo(conn * net.TCPConn){
	v :=makeRequestPacket(2);
	makeRequest(v, key_systeminfo, 0);
	makeRequest(v, key_systeminfo2, 0);
	for i := v.Len() ; i < v.Cap() ; i++ {
		binary.Write(v,binary.BigEndian,byte(0))
	}
	conn.Write(v.Bytes());
	readBodyLen := readBodyLenth(conn);
	tmp := make([]byte, readBodyLen)     // using small tmo buffer for demonstrating
	n,_:=conn.Read(tmp);
	fmt.Println("read", n,tmp[:n])
}
func Gcode(conn * net.TCPConn){

	v :=makeRequestPacket(1);
	makeRequest(v, key_cnc_rdexeprog, 1024);
	for i := v.Len() ; i < v.Cap() ; i++ {
		binary.Write(v,binary.BigEndian,byte(0))
	}
	conn.Write(v.Bytes());
	readBodyLen := readBodyLenth(conn);
	tmp := make([]byte, readBodyLen)     // using small tmo buffer for demonstrating
	n,_:=conn.Read(tmp);
	parseProgram(n,tmp);

}
func parseProgram(len int,buf [] byte){
	fmt.Println("read", len,buf[:len]);
	binary.BigEndian.Uint16(buf[0:]); // resultCount
	//parseFunc : resultLen 2 funcstrt 2 , functype 4, 2, dummy 2, dummy 2
	//dataLen 2 , 4
	s := string(buf[20:]);
	fmt.Println("count  resultLen ", s);
}

func readBodyLenth(conn * net.TCPConn) uint16{
	header := make([]byte, 10)
	conn.Read(header);
	return binary.BigEndian.Uint16(header[8:]);

}
func  makeRequestPacket( count int32) *bytes.Buffer {
	var  totalLen int16 = DEFAULT_REQ_COUNT_LENGTH + DEFAULT_REQ_LENGTH * int16(count);
	var  len int16 = int16(DEFALUT_REQUEST_HEADER_LENGTH) + int16(DEFAULT_TOTAL_SIZE_LENGTH)+ totalLen;
	b :=new(bytes.Buffer);
	b.Grow(int(len));
	binary.Write(b,binary.BigEndian,fnHeader);
	binary.Write(b,binary.BigEndian,totalLen);
	binary.Write(b,binary.BigEndian,int16(count));
	return b;
}
func makeRequest(buf *bytes.Buffer, key []int8 , parameter ...int32)  {
	var  p = buf.Len();
	binary.Write(buf,binary.BigEndian,DEFAULT_REQ_LENGTH);
	binary.Write(buf,binary.BigEndian,int16(1));
	binary.Write(buf,binary.BigEndian,key);
	for  _, num:= range parameter{
		binary.Write(buf,binary.BigEndian,num);
	}
	len :=(p + int(DEFAULT_REQ_LENGTH) - buf.Len() );
	for  i := 0 ; i < len;i++{
		binary.Write(buf,binary.BigEndian,byte(0))
	}
}
