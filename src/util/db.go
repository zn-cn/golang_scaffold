package util

import (
	"config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetDBSession 获取数据库 sess
func GetDBSession() (*sql.DB, error) {
	conf := config.Conf
	username := conf.MySQL.UserName
	userpw := conf.MySQL.UserPW
	if conf.MySQL.UserName == "" || conf.MySQL.UserPW == "" {
		username = "root"
		userpw = conf.MySQL.RootPW
	}
	db, err := sql.Open(conf.MySQL.DriverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, userpw, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.DBName))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
