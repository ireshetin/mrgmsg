package Helper

import (
	"io/ioutil"
	"log"
)

type tConfig struct {
	App tConfigApp `json:"app"`
}

type tConfigApp struct {
	Common tConfigCommon `json:"common"`
	Db     tConfigDb     `json:"db"`
}

type tConfigCommon struct {
	Parama string `json:"param_a"`
	Paramb string `json:"param_b"`
}

type tConfigDb struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Net    string `json:"net"`
	Addr   string `json:"addr"`
	Dbname string `json:"dbname"`
}

func GetConfig() *tConfig {
	var config = &tConfig{}
	data, err := ioutil.ReadFile("app.json")
	if err != nil {
		log.Fatal(err)
	}
	err = ParseJsonIntoStruct(data, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
