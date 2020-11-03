package server

import (
	"net"
	"SecondsKill/core"
)

//TcpServer
func TcpServer() {
	var over bool
	var writeMsg string
	//初始化秒杀链表
	kill := core.InitList(2)
	go kill.Monitor()
	listener,_ := net.Listen("tcp","127.0.0.1:9999")
	
	for{
		conn ,_ := listener.Accept()
		//接收用户的请求
		buf := make([]byte, 512) //1024大小的缓冲区
		n, _ := conn.Read(buf)
	 
		if string(buf[:n]) == "kill"{
			//秒杀动作
			over = kill.GetKillOption()
			if over == true{
				writeMsg = "秒杀结束"
			}else{
				writeMsg = "秒杀成功"
			}
			conn.Write([]byte(writeMsg))
		}
 
	}

}


