package mysql

import (
	"database/sql"
	"ezTikTok/models"
	"go.uber.org/zap"
)

// CreatePost 插入用户到数据库
func CreatePost(post *models.Post) (err error) {
	// 执行sql语句入库
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}

// GetPostDetailByID 根据ID获取帖子详情
func GetPostDetailByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from  post where post_id = ?`
	if err := db.Get(post, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return post, err
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (postList []*models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time 
	from  post 
	limit ?,?`
	postList = make([]*models.Post, 0, 4)
	err = db.Select(&postList, sqlStr, (page-1)*size, size)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no post in db")
		err = nil
		return
	}
	return
}
