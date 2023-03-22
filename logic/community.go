package logic

import (
	"ezTikTok/dao/mysql"
	"ezTikTok/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库 查找community并返回

	return mysql.GetCommunityList()
}
