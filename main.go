package main

import(
	"SecondsKill/server"
)
func main(){
	go server.TcpServer()
	server.GinHttp()
	
}
