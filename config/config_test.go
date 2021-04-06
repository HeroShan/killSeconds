package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	c := GetConfig()
	fmt.Printf("%#v\n",c)
}
