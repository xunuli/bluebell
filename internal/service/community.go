package service

import (
	"bluebell/internal/dao/mysql"
	"bluebell/internal/model"
	"fmt"
)

// 查询数据库community表中的信息（community_id, community_name）并以列表的形式返回
func GetCommunityList() ([]*model.CommunityRsp, error) {
	fmt.Println("正在查询2。。。")
	return mysql.GetCommunityList()
}

// 根据社区id获取某个社区的详细信息
func GetCommunityDetail(id int64) (*model.CommunityDetailRsp, error) {
	return mysql.GetCommunityDetailById(id)
}
