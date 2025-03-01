package id_generator

import (
	"github.com/DopamineNone/bubblePro/src/config"
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init() {
	startTime, machineID := config.GetConf().StartTime, config.GetConf().MachineID
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1e6 // ms time precision
	node, err = snowflake.NewNode(machineID)
}

func GetID() int64 {
	return node.Generate().Int64()
}
