package config

import (
	"database/sql"
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	SQLDriver string
	DbName    string
	LogFile   string
}

var Db *sql.DB
var err error
var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
