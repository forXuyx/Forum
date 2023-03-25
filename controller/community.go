package controller

import (
	"ezTikTok/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 获取社区列表接口
// @Summary 获取社区列表接口
// @Description 用于查询所有社区的接口
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.Community false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityList
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	// 查询到所有社区（community_id, community_name）以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 获取社区详情
// @Summary 获取社区详情
// @Description 用于获取社区详情的接口
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.CommunityDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityDetail
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get community detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询帖子详情
	data, err := logic.CommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.CommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
