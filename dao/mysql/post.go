package mysql

import (
	"database/sql"
	"ezTikTok/models"
)

// CreatePost 插入用户到数据库
func CreatePost(post *models.Post) (err error) {
	// 执行sql语句入库
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}

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
