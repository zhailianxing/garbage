package mysql

import (
	"fmt"
	"garbage/database"
	"strconv"
)

type GarbageIndex struct {
	Id       int64
	Name     string
	Category string
	Count    int64
}

func GetCategoryByName(name string) (error, GarbageIndex) {
	rows := make([]GarbageIndex, 0)
	err := database.Engine.
		Where("name = ?", name).
		Find(&rows)
	if err != nil {
		fmt.Print("GetCategoryByName error,:", err)
		return err, GarbageIndex{}
	}
	if len(rows) <= 0 {
		fmt.Printf("GetCategoryByName , name:%s can not find\n", name)
		return err, GarbageIndex{}
	}
	fmt.Println("mysql:")
	fmt.Println(rows[0])
	return nil, rows[0]
}

func GetCategorysByLikeName(name string, limit int) (err error, ret []GarbageIndex) {
	rows := make([]GarbageIndex, 0)
	// 1.error,模糊查询不起效果啊
	//err := database.Engine.
	//	Where("name like ?", "%"+name+"%").Limit(0, limit).
	//	Find(&rows)
	// 2. success
	err = database.Engine.SQL("select * from garbage_index where name like " + "'" + "%" + name + "%" + "'" + " limit 0," + strconv.Itoa(limit)).Find(&rows)
	// 3. error:
	// 原生语句要加 单引号： '%艾草%' ，所以也要加 单引号，官方文档坑啊, 但是还是不行啊
	//err := database.Engine.
	//	Where("name like ?", "'"+"%"+name+"%"+"'").Limit(0, limit).
	//	Find(&rows)
	if err != nil {
		fmt.Print("GetCategorysByLikeName error,:", err)
		return
	}
	if len(rows) <= 0 {
		fmt.Printf("GetCategorysByLikeName , name like:%s can not find\n", name)
		return nil, rows
	}
	return nil, rows
}
