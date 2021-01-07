
package server
 
import (
    "net"
)
 
func TcpClient()(string){
    //主动连接服务器
    conn, _ := net.Dial("tcp", "127.0.0.1:9999")
 
    defer conn.Close()
 
    //发送数据
    conn.Write([]byte("kill"))
    buf := make([]byte, 512) //1024大小的缓冲区
		n, _ := conn.Read(buf)
	 
	return string(buf[:n])
    
 
}

