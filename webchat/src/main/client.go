package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

/*func main() {



	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//获取服务端消息
	go ioCopy(os.Stdout, conn)
	//将用户输入的文本消息发送到到服务端
	ioCopy(conn, os.Stdin)
}*/

func connectClient(conn  net.Conn){

	//conn, err := net.Dial("tcp", "localhost:8000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer conn.Close()
	//获取服务端消息
	go ioCopy(os.Stdout, conn)

	var ioread io.Reader = strings.NewReader(string("hello, world"))
	io.Copy(conn,ioread)
	//将用户输入的文本消息发送到到服务端
	//ioCopy(conn, os.Stdin)

}

func ioCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}