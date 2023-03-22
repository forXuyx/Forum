package logic

import (
	"ezTikTok/dao/mysql"
	"ezTikTok/models"
	"ezTikTok/pkg/snowflake"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	// 生成UID
	postID := snowflake.GetID()
	// 构造一个user实例
	post := &models.Post{
		ID:          postID,
		AuthorID:    p.AuthorID,
		CommunityID: p.CommunityID,
		Status:      p.Status,
		Title:       p.Title,
		Content:     p.Content,
		CreateTime:  p.CreateTime,
	}
	// 保存到数据库
	return mysql.CreatePost(post)
}
