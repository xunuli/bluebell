package model

import "time"

// 定义community数据表的模型
type Community struct {
	Id            int       `gorm:"column:id"`
	CommunityID   int       `gorm:"column:community_id"`
	CommunityName string    `gorm:"column:community_name"`
	Introduction  string    `gorm:"column:introduction"`
	CreateTime    time.Time `gorm:"autoCreateTime"`
	UpdateTime    time.Time `gorm:"autoUpdateTime"`
}

func (t *Community) TableName() string {
	return "community"
}

// 社区查询列表的响应
type CommunityRsp struct {
	CommunityID   int    `json:"community_id"`
	CommunityName string `json:"community_name"`
}

// 查询社区的响应信息
type CommunityDetailRsp struct {
	CommunityID   int       `json:"community_id"`
	CommunityName string    `json:"community_name"`
	Introduction  string    `json:"introduction"`
	CreateTime    time.Time `json:"create_time"`
}
