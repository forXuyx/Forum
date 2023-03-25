package redis

const (
	KeyPrefix          = "blog:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZsetPF = "post:voted:"
)

// 给key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
