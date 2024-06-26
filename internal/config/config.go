package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Img2Siyuan struct {
	Token   string `json:"token"`
	Host    string `json:"host"`
	BlockId string `json:"block_id"`
}

type Module struct {
	Img2Siyuan Img2Siyuan `json:"img2siyuan"`
}

type Model struct {
}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Feishu struct {
	AppID        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	EncryptKey   string `json:"encrypt_key"`
	Verification string `json:"verification"`
}

type Config struct {
	Feishu Feishu `json:"feishu"`
	Server Server `json:"server"`
	Module Module `json:"module"`
	Model  Model  `json:"model"`
}

var (
	GlobalConfig *Config
)

func InitConfig() error {
	configFile := "config.json"
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		data, err = ioutil.ReadFile("./config/" + configFile)
		if err != nil {
			log.Println("Read config error!")
			return err
		}
	}

	config := &Config{}

	err = json.Unmarshal(data, config)

	if err != nil {
		log.Println("Unmarshal config error!")
		log.Panic(err)
		return err
	}

	GlobalConfig = config
	log.Println("Config " + configFile + " loaded.")
	return nil
}

func InitWithPath(configPath string) error {
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		return err
	}

	config := &Config{}

	err = json.Unmarshal(data, config)

	if err != nil {
		log.Println("Unmarshal config error!")
		return err
	}

	GlobalConfig = config
	log.Println("Config " + configPath + " loaded.")
	return nil
}
