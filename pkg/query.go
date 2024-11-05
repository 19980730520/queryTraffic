package pkg

import (
	"context"

	statsService "github.com/xtls/xray-core/app/stats/command"
)

// queryTraffic 查询 Xray 流量统计服务中的流量数据
func QueryTraffic(c statsService.StatsServiceClient, ptn string, reset bool) (traffic int64, err error) {
	// 初始化返回值
	traffic = -1

	// 调用查询方法
	resp, err := c.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		Pattern: ptn,
		Reset_:  reset,
	})
	if err != nil {
		return
	}

	// 获取流量数据
	stat := resp.GetStat()

	// 判断返回结果
	if len(stat) != 0 {
		traffic = stat[0].Value
	}

	return
}
