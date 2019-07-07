package garbage

import (
	"garbage/model/mysql"
)

type ReturnData struct {
	Name     string
	Category string
}

//用户输入查询词，到数据库中查找
func GetCategoryByName(name string) []*ReturnData {
	// 精准匹配
	ret := make([]*ReturnData, 0)
	err, matchCategory := mysql.GetCategoryByName(name)

	if err == nil && len(matchCategory.Category) > 0 {
		ret = append(ret, &ReturnData{
			Name:     matchCategory.Name,
			Category: matchCategory.Category,
		})
	}
	// 再模糊匹配 10条
	err, likeCategory := mysql.GetCategorysByLikeName(name, 10)
	if err == nil {
		for _, category := range likeCategory {
			if category.Name != matchCategory.Name {
				ret = append(ret, &ReturnData{
					Name:     category.Name,
					Category: category.Category,
				})
			}
		}
	}
	return ret
}
