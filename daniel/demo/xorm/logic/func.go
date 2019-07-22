package logic

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

type Model struct {
	engine *xorm.Engine
}

func CreateModel() (Model, error) {
	model := Model{}
	engine, err := xorm.NewEngine("mysql", "xuser:xpass611@tcp(rdsqafzen2y2i7f278.mysql.rds.aliyuncs.com:3306)/xproject?charset=utf8")
	CheckErr(err)

	fp, err := os.Open("../log/log.txt")
	checkErr(err)

	engine.ShowSQL(true)
	logger := xorm.NewSimpleLogger(fp)
	engine.SetLogger(logger)

	model.engine = engine
	return model, nil
}

func CheckErr(err error) {
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

