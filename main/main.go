package main

import (
	//db "aze.org/database"
	"flag"
	"fmt"
	"garbage/conf"
	"garbage/database"
)

func main() {
	//defer db.SqlDB.Close()
	//env := flag.String("env", "local", "-env local | qa | product")
	//platform := flag.String("platform", "qutoutiao", "-platform qutoutiao | midu")
	//flag.Parse()
	//confFile := fmt.Sprintf("../conf/%s/%sConf.json", *platform,*env)

	env := flag.String("env", "local", "-env local | qa | product")
	platform := flag.String("platform", "garbage", "-platform qutoutiao | midu")
	flag.Parse()
	confFile := fmt.Sprintf("../conf/%s/%sConf.json", *platform, *env)
	conf.InitConfig(confFile)

	database.InitDatabase() //TODO: 释放数据库资源
	router := initRouter()
	router.Run(":8000")

}
