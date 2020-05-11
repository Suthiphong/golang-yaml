package app

import (
	"fmt"
	"localhost/configs"
)

//Serial for global variable
var (
	Serial   string //Serial global
	Baudrate int    //Baudrate Global
	Yaml     = configs.GetConfig().Yaml
)

/*
==== INITIAL CONFIG YAML ====
*/
func init() {
	Yaml.ReadFile()
	Serial = Yaml.Configure.Serial
	Baudrate = Yaml.Configure.Baudrate
}

//Run start program.
func Run() {
	/*
	  === Setup config && Setup config file *.yaml
	*/
	fmt.Printf("%+v", Yaml)
	fmt.Println("Program already.")

}
