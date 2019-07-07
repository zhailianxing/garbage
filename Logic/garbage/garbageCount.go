package garbage

import (
	"fmt"
	"garbage/model/mysql"
)

//用户输入查询词，到数据库中查找
func RecordSearchWords(garbageName string, isFound bool) (err error) {
	// 2. GarbageIndex中找到 查询词
	// 2.1 查询garbageSearchLog记录是否存在
	// 2.2 存在则更新counter和isExisted字段； 不存在则创建记录[isExisted为1, counter=1]
	fmt.Println("enter RecordSearchWords")
	if isFound {
		err, data := mysql.GetSearchLog(garbageName)
		if err == nil {
			if data != nil { //搜索log存在
				data.IsExisted = 1
				data.Counter = data.Counter + 1
				err, _ := mysql.UpdateSearchLog(garbageName, data)
				if err != nil {
					fmt.Println("Found RecordSearchWords UpdateSearchLog error", err)
				}
			} else { // 搜索log不存在
				newData := new(mysql.GarbageSearchLog)
				newData.IsExisted = 1
				newData.Counter = 1
				err, _ := mysql.InsertSearchLog(garbageName, newData)
				if err != nil {
					fmt.Println("Found RecordSearchWords InsertSearchLog error", err)
				}
			}
		}
	} else {
		// 1. GarbageIndex中没有找到 查询词
		// 1.1 查询garbageSearchLog记录是否存在
		// 1.2 存在则更新count字段； 不存在则创建记录[isExisted为0，counter=1]
		fmt.Println("notfound")
		err, data := mysql.GetSearchLog(garbageName)
		fmt.Println("GetSearchLog return , err:", err)
		fmt.Println("GetSearchLog return , data:", data)
		if err == nil {
			if data != nil { //搜索log存在
				data.IsExisted = 0 // 跟上面的区别
				data.Counter = data.Counter + 1
				err, _ := mysql.UpdateSearchLog(garbageName, data)
				fmt.Println("GetSearchLog UpdateSearchLog, err:", err)
				if err != nil {
					fmt.Println("NotFound RecordSearchWords UpdateSearchLog error", err)
				}
			} else { // 搜索log不存在
				newData := new(mysql.GarbageSearchLog)
				newData.IsExisted = 0 // 跟上面的区别
				newData.Counter = 1
				err, _ := mysql.InsertSearchLog(garbageName, newData)
				fmt.Println("GetSearchLog InsertSearchLog, err:", err)
				if err != nil {
					fmt.Println("NotFound RecordSearchWords InsertSearchLog error", err)
				}
			}
		}
	}
	fmt.Println("over RecordSearchWords")
	return
}
