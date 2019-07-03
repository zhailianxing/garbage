package garbage

import (
	"fmt"
	"garbage/model/mysql"
)

type ReturnData struct {
	Name     string
	Category string
}

func GetCategoryByName(name string) (ret []ReturnData) {
	// 精准匹配
	err, matchCategory := mysql.GetCategoryByName(name)

	if err == nil && len(matchCategory.Category) > 0 {
		ret = append(ret, ReturnData{
			Name:     matchCategory.Name,
			Category: matchCategory.Category,
		})
	}
	// 再模糊匹配 10条
	err, likeCategory := mysql.GetCategorysByLikeName(name, 10)
	fmt.Println("mohu match")
	fmt.Println(err)
	fmt.Println(likeCategory)
	if err == nil {
		for _, category := range likeCategory {
			if category.Name != matchCategory.Name {
				ret = append(ret, ReturnData{
					Name:     category.Name,
					Category: category.Category,
				})
			}
		}
	}
	fmt.Println(ret)
	return ret
}
