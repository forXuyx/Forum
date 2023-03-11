package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"ezTikTok/models"
)

const secret = "hais"

// 把每一步数据库操作封装成函数
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行sql语句入库
	sqlStr := `insert into user(user_id, username, password, phone, email) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Phone, user.Email)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
