package redis

const (
	KeyPrefix          = "blog:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZsetPF = "post:voted:"
)
