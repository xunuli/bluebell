package snowflake

import (
	"bluebell/config"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

//雪花算法生成全局唯一id(趋势递增)

// 定义一个id生成器
var node *sf.Node

// 初始化node
func Init(cfg *config.SfConf) {
	startTime := cfg.StartTime
	machineID := cfg.MachineID
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// 生成id
func GenID() int64 {
	return node.Generate().Int64()
}
