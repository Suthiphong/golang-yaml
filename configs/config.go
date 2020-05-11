package configs

import "fmt"

type Config struct {
	Yaml YamlConfig
}

func (c *Config) ShowConfig() {
	fmt.Println("Show config")
}

func GetConfig() Config {
	return Config{
		Yaml: GetYamlInfo(),
	}
}
