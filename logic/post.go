package logic

import (
	"ezTikTok/dao/mysql"
	"ezTikTok/models"
	"ezTikTok/pkg/snowflake"
	"go.uber.org/zap"
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

// GetPostDetailByID 获取帖子详情
func GetPostDetailByID(id int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合接口数据
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID failed", zap.Error(err))
		return
	}
	// 根据id查询作者
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
		return
	}
	// 根据社区id查询社区详情
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		// 根据id查询作者
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
			continue
		}
		// 根据社区id查询社区详情
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
