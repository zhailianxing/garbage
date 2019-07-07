package mysql

import (
	"garbage/database"
	"time"
)

type GarbageSearchLog struct {
	Id         int64  `xorm:"pk autoincr"`
	Name       string `xorm:"varchar(25) notnull unique"` // TODO: 了解
	IsExisted  int32
	Counter    int32     `xorm:"default 0"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`

	// TODO: 了解: 下面的json用法
	// CreatTime  time.Time `xorm:"creat_time created" json:"creat_time" description:"创建时间"`
}

func GetSearchLog(name string) (error, *GarbageSearchLog) {
	rows := make([]*GarbageSearchLog, 0)
	err := database.Engine.
		Where("name = ?", name).
		Find(&rows)
	if err != nil {
		return err, nil
	}
	if len(rows) <= 0 {
		return nil, nil
	}
	return nil, rows[0]
}

func UpdateSearchLog(name string, row *GarbageSearchLog) (error, int64) {
	data := new(GarbageSearchLog)
	id := row.Id
	data.Name = name
	data.Counter = row.Counter
	data.IsExisted = row.IsExisted
	affect, err := database.Engine.Id(id).Update(data)
	if err != nil {
		return err, 0
	}
	return nil, affect
}

func InsertSearchLog(name string, row *GarbageSearchLog) (error, int64) {
	data := new(GarbageSearchLog)
	data.Name = name
	data.Counter = row.Counter
	data.IsExisted = row.IsExisted
	affect, err := database.Engine.Insert(data)
	if err != nil {
		return err, 0
	}
	return nil, affect
}
