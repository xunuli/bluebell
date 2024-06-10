package mysql

import (
	"bluebell/internal/model"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

var ErrorInvalidID = errors.New("无效的ID")

// 根据字段获取社区列表
func GetCommunityList() ([]*model.CommunityRsp, error) {
	var community []*model.CommunityRsp
	//根据字段查询community表
	fmt.Println("正在查询3。。。")
	mdb := GetDB()
	//result := mdb.Model(model.Community{}).Select("community_name", "community_id").Find(&community)
	result := mdb.Model(model.Community{}).Scan(&community)
	fmt.Println(*community[0])
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("there is no community in db")
	}
	return community, nil
}

// 根据社区id获取社区的详细信息
func GetCommunityDetailById(id int64) (*model.CommunityDetailRsp, error) {
	var community *model.CommunityDetailRsp
	mdb := GetDB()
	result := mdb.Model(model.Community{}).Where("community_id = ?", id).Scan(&community)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrorInvalidID
	}
	return community, nil
}
