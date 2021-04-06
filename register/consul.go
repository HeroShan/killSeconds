package register

import (
	"SecondsKill/config"
	"fmt"
	consular "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"log"
)

type ServiceSource struct {
	Name string
	Port int
	Tags []string
	Addr string
}


func RegisterService(ser *ServiceSource) bool {
	conf := new(config.Config)
	conf.GetConfig()
	config := consular.DefaultConfig()
	config.Address = conf.Consul.Addr+":"+conf.Consul.Port
	client, err := consular.NewClient(config);if err != nil{
		panic(err)
	}
	//服务注册元数据
	registration := new(consular.AgentServiceRegistration)
	registration.ID = uuid.Must(uuid.NewV4(),nil).String()
	registration.Name = ser.Name
	registration.Port = ser.Port
	registration.Tags = ser.Tags
	registration.Address = ser.Addr
	log.Printf("ser %#v\n",registration)
	// 增加consul健康检查回调函数
	check := new(consular.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil{
		log.Fatalf("service:%#v\n register failed\n error is:%v",ser,err)
		return false
	}
	return true
}

func CheckService(){

}