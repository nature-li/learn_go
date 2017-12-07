package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	ServerPort int `json:"server_port"`
	LocalPort int `json:"local_port"`
	PassWord string `json:"password"`
	Timeout int `json:"timeout"`
}

func ParseConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	config := &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(config.LocalPort)
}

func main() {
	ParseConfig("./config.json")
}

/*
{
  "server_port": 8388,
  "local_port": 1081,
  "password": "123456!",
  "timeout": 60
}
*/