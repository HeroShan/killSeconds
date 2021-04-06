package config
import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Consul struct {
		Port string	`yaml:"port"`
		Addr string `yaml:"addr"`
	}
}

func (conf *Config)GetConfig() {
	yamlFile, err := ioutil.ReadFile("D:\\WorkSpace\\GoWork\\src\\killSeconds\\config\\config.yaml");if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return
}
