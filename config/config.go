package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//UserStrut config struct
//use as map[string]UserStrut
type UserStrut struct {
	Pass  string
	Group []string
}

//TLSStrut config struct
type TLSStrut struct {
	Cert string
	Key  string
}

//ConfigStrut decide struct of config, only use for json Unmarshal
type ConfigStrut struct {
	ProjectPath string
	TLS         TLSStrut
}

var configStrut ConfigStrut

//value that will export
var (
	ProjectPath string
	TLS         TLSStrut
	Users       map[string]UserStrut
)

//Load Config file
func Load(file string) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	log.Println("Read Config \r\n", string(data))
	err = json.Unmarshal(data, &configStrut)
	if err != nil {
		panic(err)
	}

	//export config
	ProjectPath = configStrut.ProjectPath
	TLS = configStrut.TLS
	Users = make(map[string]UserStrut)
	Users["maple"] = UserStrut{
		Pass:  "maple",
		Group: []string{"admin"},
	}
}
