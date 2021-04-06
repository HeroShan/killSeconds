package register

import (
	"fmt"
	"testing"
)

func TestRegisterService(t *testing.T) {
	var service ServiceSource
	service.Name = "kill"
	service.Port = 8080
	service.Tags = append(service.Tags,"test service")
	service.Addr = "127.0.0.1"
	b := RegisterService(&service)
	fmt.Printf("%t",b)
}
