package controller

import (
	"ezTikTok/logic"
	"ezTikTok/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 帖子投票功能接口
// @Summary 帖子投票功能接口
// @Description 用于给帖子投票的接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamVote false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _PostVote
// @Router /vote [post]
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVote)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回json
		zap.L().Error("PostVote with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		return
	}
	// 获取当前用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票逻辑
	if err := logic.PostVote(userID, p); err != nil {
		zap.L().Error("logic.PostVote failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
