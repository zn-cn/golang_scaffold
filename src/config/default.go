package config

import (
	"log"
	"os"
)

// Conf 配置
var Conf *Config

// Config 配置
type Config struct {
	Env  string
	Addr string

	MySQL     *mysqlConf
	SecretKey string
}

type mysqlConf struct {
	Host       string
	Port       string
	DriverName string
	UserName   string
	DBName     string
	RootPW     string
	UserPW     string
}

func init() {

	log.Println("init config")
	Conf := &Config{
		Env:  "local",
		Addr: ":1323",

		MySQL: &mysqlConf{
			Host:       "localhost",
			Port:       "3306",
			DriverName: "mysql",
		},
		SecretKey: "secret-key",
	}
	if v, ok := os.LookupEnv("ENV"); ok {
		if v == "dev" || v == "prod" {
			Conf.MySQL.Host = "mysql"
		}
		Conf.Env = v
	}

	if v, ok := os.LookupEnv("SecretKey"); ok {
		Conf.SecretKey = v
	}
	if v, ok := os.LookupEnv("APP_ADDR"); ok {
		Conf.Addr = v
	}

	initDBConf()

}

func initDBConf() {
	if v, ok := os.LookupEnv("MySQL_HOST"); ok {
		Conf.MySQL.Host = v
	}
	if v, ok := os.LookupEnv("MySQL_PORT"); ok {
		Conf.MySQL.Port = v
	}
	if v, ok := os.LookupEnv("MySQL_USER"); ok {
		Conf.MySQL.UserName = v
	}
	if v, ok := os.LookupEnv("MySQL_USERPW"); ok {
		Conf.MySQL.UserPW = v
	}
	if v, ok := os.LookupEnv("MySQL_ROOTPW"); ok {
		Conf.MySQL.RootPW = v
	}
	if v, ok := os.LookupEnv("MySQL_DBNAME"); ok {
		Conf.MySQL.DBName = v
	}
}
