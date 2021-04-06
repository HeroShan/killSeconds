package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := new(Config)
	conf.GetConfig()
	fmt.Printf("%#v\n",conf)
}
