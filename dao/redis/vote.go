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

func PostVote(userID, postID string, value float64) error {
	postTime := rdb.ZScore(getRedisKey(KeyPostScoreZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInsecond {
		return ErrVoteTimeExpire
	}
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZsetPF+postID), userID).Val()
	var op float64
	if value > ov {
		op = -1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}
	if value == 0 {
		_, err = rdb.ZRem(getRedisKey(KeyPostVotedZsetPF+postID), userID).Result()
	} else {
		_, err = rdb.ZAdd(getRedisKey(KeyPostVotedZsetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		}).Result()
	}
	return err
}
