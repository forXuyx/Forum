package controller

import (
	"ezTikTok/logic"
	"ezTikTok/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回json
		zap.L().Error("CreatePostHandler with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取社区详情
func GetPostDetailHandler(c *gin.Context) {
	// 获取帖子id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询该帖子
	data, err := logic.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	page, size := GetPageInfo(c)
	// 查询到所有帖子
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
