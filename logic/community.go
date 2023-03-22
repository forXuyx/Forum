package logic

import (
	"ezTikTok/dao/mysql"
	"ezTikTok/models"
)

// GetCommunityList 获取社区列表
func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库 查找community并返回

	return mysql.GetCommunityList()
}

// CommunityDetail 获取社区详情
func CommunityDetail(id int64) (*models.CommunityDetail, error) {
	// 查询数据库 查找community并返回

	return mysql.GetCommunityDetailByID(id)
}
