package database

import (
	"fmt"
	"garbage/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

// 初始化数据库池
func InitDatabase() {
	fmt.Println("init dtabase")
	user := conf.GlobalConf.DbUser
	password := conf.GlobalConf.DbPassword
	host := conf.GlobalConf.DbHost
	dbName := conf.GlobalConf.DbName
	charset := conf.GlobalConf.Charset

	if host != "" {
		host = fmt.Sprintf("tcp(%s)", host)
	}

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=%s", user, password, host, dbName, charset)
	fmt.Println(dsn)
	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		msg := "create db engine fail"
		//log.AppLogger.Critical(msg)
		fmt.Println("init db mysql error")
		panic(msg)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	db.ShowSQL(true) // 显示打印出sql语句
	Engine = db

	//log.AppLogger.Info("init db engine was successful")
}
