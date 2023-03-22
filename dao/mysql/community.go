package mysql

import (
	"database/sql"
	"ezTikTok/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db")
		err = nil
		return
	}
	return
}
