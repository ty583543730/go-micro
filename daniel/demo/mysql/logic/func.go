package logic

import (
	"../common"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewMysql() (*sql.DB, error) {
	dbConfig := common.DBConfig{
		Host:     "rdsqafzen2y2i7f278.mysql.rds.aliyuncs.com",
		Port:     3306,
		Database: "xproject",
		Username: "xuser",
		Password: "xpass611",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	DB, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return DB, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func GetTime() string {
	const format = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(format)
	return str
}
