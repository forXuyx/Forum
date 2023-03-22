package mysql

import "ezTikTok/models"

// CreatePost 插入用户到数据库
func CreatePost(post *models.Post) (err error) {
	// 执行sql语句入库
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}
