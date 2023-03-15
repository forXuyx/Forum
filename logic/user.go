package logic

import (
	"ezTikTok/dao/mysql"
	"ezTikTok/models"
	"ezTikTok/pkg/jwt"
	"ezTikTok/pkg/snowflake"
)

// SignUp 处理注册业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GetID()
	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		Phone:    p.Phone,
		Email:    p.Email,
	}
	// 保存到数据库
	return mysql.InsertUser(user)
}

// Login 处理登录业务逻辑
func Login(p *models.ParamLogin) (token string, err error) {
	// 直接登录
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
