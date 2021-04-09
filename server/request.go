package server

import (
	"SecondsKill/register"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
)

// StartTime 缺少一个初始化秒杀端口
type StartTime struct {
	Now string
}

func (timeString *StartTime) Start() bool {
	t, _ := time.Parse("2006-01-02 15:04:05", timeString.Now)
	nowTime := time.Now()
	return t.Before(nowTime)
}

func ks(c *gin.Context) {
	var t StartTime
	t.Now = "2021-03-30 19:48:00"
	log.Printf("%v\n\n", t.Start())
	if t.Start() == true {
		result := TcpClient()
		c.String(200, result)
	} else {
		c.String(200, "秒杀开始时间是："+t.Now)
	}

}

func GinHttp() {
	router := gin.Default()
	//服务注册
	var service register.ServiceSource
	localIps, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln("register ip to consul failed")
	}
	var ipv4 string
	//获取网卡地址ipv4
	for _, addr := range localIps {
		if ipNet, isIpNet := addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String()
			}
		}
	}
	service.Name = "ginHTTP"
	service.Port = 1997
	service.Tags = append(service.Tags, "service kill tcp")
	service.Addr = ipv4
	serviceRestul := register.RegisterService(&service)
	if !serviceRestul {
		panic("gin server is shutdown")
	}
	root := router.Group("/")
	{
		root.GET("/ks", ks)
	}
	router.Run("0.0.0.0:1997")
}
