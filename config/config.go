package config
import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Consul struct {
		Port int	`yaml:"port"`
		Addr string `yaml:"addr"`
	}
}

func GetConfig() *Config {
	conf := new(Config)
	yamlFile, err := ioutil.ReadFile("config.yaml"); if err != nil{
		panic(err)
	}
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return conf
}
