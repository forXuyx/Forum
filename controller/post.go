package controller

import (
	"ezTikTok/logic"
	"ezTikTok/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子接口
// @Summary 创建帖子接口
// @Description 用于创建帖子的接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.Post false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _CreatePost
// @Router /post [post]
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

// GetPostDetailHandler 获取帖子详情接口
// @Summary 获取帖子详情接口
// @Description 用于获取帖子详情的接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ApiPostDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _GetPostDetail
// @Router /post/:id [get]
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

// GetPostListHandler 获取帖子列表接口
// @Summary 获取帖子列表接口
// @Description 普通获取帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts [get]
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

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询到所有帖子
	data, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
