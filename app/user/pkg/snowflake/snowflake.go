package snowflake

import (
	"dzug/conf"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

//调用方法：
//在主程序先snowflake.Init(conf.Config.StartTime, conf.Config.MachineID)
//然后在需要生成id时，直接	id := snowflake.GenID() 即可获取

var node *sf.Node

func Init() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", conf.Config.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(conf.Config.MachineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
