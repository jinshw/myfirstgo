package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

//var conn net.Conn
var conn, err = net.Dial("tcp", "localhost:8000")
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/template/").Handler(http.StripPrefix("/template/",http.FileServer(http.Dir("template/"))))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == "GET" { //如果请求方法为get显示login.html,并相应给前端
			t, _ := template.ParseFiles("template/html/chat.html")
			t.Execute(writer, nil)
		}
	})
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("login success!")
		connectClient(conn)
		fmt.Println(conn)
	})
	http.HandleFunc("/send", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		request.ParseForm()
		content := request.Form["content"][0]
		fmt.Println(content)
		connectClient(conn)
		ioCopy(conn, strings.NewReader(content))
		defer conn.Close()
	})

	port := "127.0.0.1:8080"
	fmt.Println("http service start success!port"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}

