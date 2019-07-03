package conf

import (
	"fmt"
	"github.com/akkuman/parseConfig"
)

type GlobalConfig struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	Charset    string

	KafkaServer string
	KafkaTopic  string
	KafkaCmd    string
	KafkaGroup  string

	RedisHost           string
	RedisMaxIdle        int
	RedisIdleTimeoutSec int
	RedisDbNum          string
	RedisPassword       string
}

var (
	GlobalConf GlobalConfig
	//GlobalChan chan []byte
)

func InitConfig(confFile string) {
	fmt.Println("init config")
	fmt.Println(confFile)
	// config.json路径是在package main下; config.json配置文件如果有注释（// ooxx）就会提示无法工作
	var config = parseConfig.New(confFile)

	var dbUser = config.Get("db > user").(string)
	var dbPassword = config.Get("db > password").(string)
	var dbHost = config.Get("db > host").(string)
	var dbName = config.Get("db > db_name").(string)
	var charset = config.Get("db > charset").(string)

	var kafkaServer = config.Get("kafka > kafkaServer").(string)
	var kafkaTopic = config.Get("kafka > kafkaTopic").(string)
	var kafkaCmd = config.Get("kafka > kafkaCmd").(string)
	var kafkaGroup = config.Get("kafka > kafkaGroup").(string)

	var redisHost = config.Get("redis > host").(string) // 配置文件中配置的 整数，这里获得的都是float
	var redisMaxIdle = int(config.Get("redis > redisMaxIdle").(float64))
	var redisIdleTimeoutSec = int(config.Get("redis > redisIdleTimeoutSec").(float64))
	var redisDbNum = config.Get("redis > dbNum").(string)            // 配置文件中配置的 整数，这里获得的都是float
	var redisPassword = config.Get("redis > redisPassword").(string) // 配置文件中配置的 整数，这里获得的都是float

	GlobalConf.DbUser = dbUser
	GlobalConf.DbPassword = dbPassword
	GlobalConf.DbHost = dbHost
	GlobalConf.DbName = dbName
	GlobalConf.Charset = charset

	GlobalConf.KafkaServer = kafkaServer
	GlobalConf.KafkaTopic = kafkaTopic
	GlobalConf.KafkaCmd = kafkaCmd
	GlobalConf.KafkaGroup = kafkaGroup

	GlobalConf.RedisHost = redisHost
	GlobalConf.RedisMaxIdle = redisMaxIdle
	GlobalConf.RedisIdleTimeoutSec = redisIdleTimeoutSec
	GlobalConf.RedisDbNum = redisDbNum
	GlobalConf.RedisPassword = redisPassword

	//GlobalChan = make(chan []byte, 2000000)

	//a := []byte{123}
	//GlobalChan  <- a
	//b :=<- GlobalChan
	//fmt.Print(b)

}
