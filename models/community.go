package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`     // 社区ID
	Name string `json:"name" db:"community_name"` // 社区名
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`                     // 社区ID
	Name         string    `json:"name" db:"community_name"`                 // 社区名
	Introduction string    `json:"introduction,omitempty" db:"introduction"` // 社区介绍
	CreateTime   time.Time `json:"create_time" db:"create_time"`             // 社区创建时间
}
