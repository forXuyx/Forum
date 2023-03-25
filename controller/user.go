package controller

import (
	"errors"
	"ezTikTok/dao/mysql"
	"ezTikTok/logic"
	"ezTikTok/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignupHandler 注册请求接口
// @Summary 注册请求接口
// @Description 用于注册的接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.ParamSignUp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _SignUp
// @Router /signup [post]
func SignupHandler(c *gin.Context) {
	// 1.参数校验和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回json
		zap.L().Error("Signup with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录请求接口
// @Summary 登录请求接口
// @Description 用于登录的接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.ParamLogin false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _Login
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1.参数校验和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回json
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}
	// 2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
