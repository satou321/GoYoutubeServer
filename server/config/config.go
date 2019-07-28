package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strconv"
)

type configList struct {
	Port      int
	DbName    string
	SQLDriver string
	APIKey    string
}

//global
var Config configList

func init() {
	//load the config file
	cfg, err := ini.Load("server/config.ini")
	//cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	//port setting
	portStr := os.Getenv("PORT")
	port, _ := strconv.Atoi(portStr)
	if port == 0 {
		port = cfg.Section("web").Key("port").MustInt()
	}

	//generate Config
	Config = configList{
		Port:      port,
		DbName:    cfg.Section("db").Key("name").MustString("tmp.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		APIKey:    cfg.Section("youtube").Key("api_key").String(),
	}
	//fmt.Println(Config)
}
