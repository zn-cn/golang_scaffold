package util

import (
	"config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化MySQL数据库
func init() {
	conf := config.Conf
	username := conf.MySQL.UserName
	userpw := conf.MySQL.UserPW
	db, err := sql.Open(conf.MySQL.DriverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", username, userpw, conf.MySQL.Host, conf.MySQL.Port))
	if err != nil {
		return
	}
	defer db.Close()
	fmt.Println("init db")
	// 创建数据库
	db.Exec("CREATE DATABASE IF NOT EXISTS member")
	db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INT UNSIGNED AUTO_INCREMENT,
			name VARCHAR(32) NOT NULL,
			nickname VARCHAR(32) NOT NULL,
			email VARCHAR(32) NOT NULL,
			phone_number VARCHAR(32) NOT NULL,
			bcrypt_pw VARCHAR(64) NOT NULL,
			group VARCHAR(32) NOT NULL,
			status INT NOT NULL,
			create_date DATE NOT NULL,
			PRIMARY KEY(id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8
		`)
}

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
