package configs

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Configure struct {
		Serial   string `yaml:"serial"`
		Baudrate int    `yaml:"baudrate"`
	}
}

func (y *YamlConfig) ReadFile() *YamlConfig {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Println("Cannot found file ./config.yaml")
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Println("Cannot convert")
	}
	return y
}

func (y *YamlConfig) YamlDetail() string {
	return fmt.Sprintf("serial=%s baudrate=%d", y.Configure.Serial, y.Configure.Baudrate)
}

func GetYamlInfo() YamlConfig {
	return YamlConfig{}
}
