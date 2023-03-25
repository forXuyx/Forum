package logic

import (
	"ezTikTok/dao/redis"
	"ezTikTok/models"
	"go.uber.org/zap"
	"strconv"
)

// PostVote 投票功能
func PostVote(userID int64, p *models.ParamVote) error {
	zap.L().Debug("PostVote",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.PostVote(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
