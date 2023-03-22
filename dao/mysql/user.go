package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"ezTikTok/models"
)

const secret = "hais"

// 注册业务
// CheckUserExist 检查用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 插入用户到数据库
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行sql语句入库
	sqlStr := `insert into user(user_id, username, password, phone, email) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Phone, user.Email)
	return
}

// encryptPassword md5加密密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login 登录业务
func Login(user *models.User) (err error) {
	// 查询用户对应密码
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	if err = db.Get(user, sqlStr, user.Username); err != nil {
		return err
	}
	// 查询密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
