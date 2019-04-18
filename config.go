/*

configuration management

[2019-04-18] pretty printing values to log status

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

//Configuration struct with all variables imported from config.json
type Configuration struct {
	Dummy string `json:"dummy"`

	ProjectRoots []string `json:"project_roots"`
}

var conf Configuration

//ReadConfig reads 'config.json' and fills Configuration struct
func ReadConfig() {
	configFile := "config.json"
	log.Printf("ReadConfig: reading 'Configuration' from '%s'", configFile)
	var err error
	conf, err = loadConfig(configFile)
	if err == nil {
		//log.Printf("%+v\n", conf)
		log.Printf("Configuration: \n%v", prettyPrintConf())

	} else {
		log.Println(err)
	}
}

func saveConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func loadConfig(filename string) (Configuration, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}

//prettyPrintStructLocal prints a struct using reflection, simple version
//la stessa funzione di pretty print è riportata nelle goutils in due forme:
//PrettyPrintStruct() printing to stdout and PrettyFormatStruct() returning a string
func prettyPrintConf() string {
	s := ""
	typ := reflect.TypeOf(conf)
	val := reflect.ValueOf(conf)
	for i := 0; i < typ.NumField(); i++ {
		if val.Field(i).CanInterface() {
			fieldValue := val.Field(i).Interface()
			//fmt.Println(fieldValue)
			s += fmt.Sprintf("%d: %s %s = %v\n", i,
				typ.Field(i).Name, val.Field(i).Type(), fieldValue)
		} else {
			s += fmt.Sprintf("%d: %s (private value)\n", i, typ.Field(i).Name)
		}
	}
	return s
}

/*
func main() {
	configuration, err := loadConfig("config.json")
	if err == nil {
		fmt.Printf("%+v\n", configuration)
	} else {
		fmt.Println(err)
	}
}
*/
