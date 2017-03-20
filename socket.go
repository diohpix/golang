package main

import (
	"net"
	//"os"
	"fmt"
	"os"

	"bytes"
	 "encoding/binary"


)
var fnHeader = []byte{ 0xa0,  0xa0,  0xa0, 0xa0, 0x00, 0x01, 0x21, 0x01 };
var key_systeminfo = []int8{ 0x00, 0x01, 0x00, 0x18 };
var key_systeminfo2 = 	[]int8{ 0x00, 0x02, 0x00, 0x18 };
var DEFAULT_REQ_LENGTH int16 = 0x1c;
var DEFAULT_REQ_COUNT_LENGTH int16= 2;
var DEFAULT_TOTAL_SIZE_LENGTH int16= 2;
var DEFALUT_REQUEST_HEADER_LENGTH int16= 8;


func main() {
	//addr :=net.ParseIP("106.243.233.236:80");
	//fmt.Println(addr.String());



	header := make([]byte, 10)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "it.yhsbearing.com:8193")
	checkError(err);
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.Write([]byte{  0xa0, 0xa0,  0xa0,  0xa0, 0x00, 0x01, 0x01, 0x01, 0x00, 0x02, 0x00,0x02 });
	 conn.Read(header);
	readBody := binary.BigEndian.Uint16(header[8:]);
	fmt.Println("header ",binary.BigEndian.Uint16(header[8:]));
	tmp := make([]byte, readBody)     // using small tmo buffer for demonstrating
	conn.Read(tmp);
	//

	v :=makeRequestPacket(2);
	makeRequest(v, key_systeminfo, 0);
	makeRequest(v, key_systeminfo2, 0);
	for i := v.Len() ; i < v.Cap() ; i++ {
		binary.Write(v,binary.BigEndian,byte(0))
	}
	conn.Write(v.Bytes());


	 conn.Read(header);
	readBody = binary.BigEndian.Uint16(header[8:]);
	tmp = make([]byte, readBody)     // using small tmo buffer for demonstrating
	n,_:=conn.Read(tmp);
	fmt.Println("read", n,tmp[:n])



	fmt.Println("count :", binary.BigEndian.Uint16(tmp[:2]))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
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

func cnc_sysinfo( buf []byte) {
	buffer := bytes.NewBuffer(buf);
	resultcount,_ := binary.ReadVarint(buffer);
	fmt.Println(resultcount)
	/*for  i :=0 , i < int32(resultCount) ; i++ {

		int endPos = parseFunc(buf);
		short dataLen = buf.readShort();
		if(dataLen ==0) break;
		ODBSYS odbsys = new ODBSYS();
		odbsys.addinfo = buf.readShort();
		odbsys.max_axis =buf.readShort();

		byte [] cnc_type_buf=new byte[2];
		buf.readBytes(cnc_type_buf);
		odbsys.cnc_type = new String(cnc_type_buf);

		byte [] mt_type_buf=new byte[2];
		buf.readBytes(mt_type_buf);
		odbsys.mt_type = new String(mt_type_buf);

		byte [] serise_buf=new byte[4];
		buf.readBytes(serise_buf);
		odbsys.serise = new String(serise_buf);

		byte [] version_buf=new byte[4];
		buf.readBytes(version_buf);
		odbsys.version = new String(version_buf);

		byte [] axe_buf=new byte[2];
		buf.readBytes(axe_buf);
		odbsys.axe = new String(axe_buf);
		l.add(odbsys);
		buf.readerIndex(endPos);
	}*/

}