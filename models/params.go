package models

// 定义请求参数结构体

// ParamSignUp 注册结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required"`
}

// ParamLogin 登录结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
