package main

import (
	"SecondsKill/core"
	"SecondsKill/register"
	"log"
	"net"
	"time"
)

//TcpServer
func main() {
	var over bool
	var writeMsg string
	//初始化秒杀链表
	kill := core.InitList(500)
	go kill.Monitor()
	listener, _ := net.Listen("tcp", "127.0.0.1:9999")
	//服务注册
	var service register.ServiceSource
	service.Name = "killTCP"
	service.Port = 9999
	service.Tags = append(service.Tags, "service kill tcp")
	service.Addr = "127.0.0.1"
	register.RegisterService(&service)
	for {
		conn, _ := listener.Accept()
		//接收用户的请求
		buf := make([]byte, 512) //1024大小的缓冲区
		n, _ := conn.Read(buf)

		if string(buf[:n]) == "kill" {
			//秒杀动作
			over = kill.GetKillOption()
			if over == true {
				writeMsg = "秒杀结束"
			} else {
				writeMsg = "秒杀成功"
			}
			log.Println(writeMsg, "timeis:", time.Now().UnixNano())
			conn.Write([]byte(writeMsg))
		}
	}

}
