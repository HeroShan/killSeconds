package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// StartTime 缺少一个初始化秒杀端口
type StartTime struct{
	Now string
}
func (timeString *StartTime)Start()(bool){
	t, _ := time.Parse("2006-01-02 15:04:05", timeString.Now)
	nowTime := time.Now()
	return t.Before(nowTime)
}

func ks(c *gin.Context){	
	var t StartTime
	t.Now = "2021-03-30 19:48:00"
	log.Printf("%v\n\n",t.Start())
	if t.Start() == true {
		result := TcpClient()
		c.String(200,result)
	}else{
		c.String(200,"秒杀开始时间是："+t.Now)
	}
	

}

func GinHttp(){
	router := gin.Default()
	root := router.Group("/")
	{
		root.GET("/ks",ks)
	}
	router.Run()
}