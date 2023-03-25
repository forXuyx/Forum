package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekInsecond = 7 * 24 * 3600
	scorePerVote    = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

// CreatePost 创建帖子
func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

// PostVote 投票功能
func PostVote(userID, postID string, value float64) error {
	// 1.判断发帖时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInsecond {
		return ErrVoteTimeExpire
	}
	// 2.更新帖子分数
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZsetPF+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := rdb.Pipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3.记录用户投票数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZsetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZsetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
