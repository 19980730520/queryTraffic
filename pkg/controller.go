package pkg

import (
	"fmt"

	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

// XrayController 是一个用于控制 Xray 服务的结构体
type XrayController struct {
	SsClient statsService.StatsServiceClient
	CmdConn  *grpc.ClientConn // gRPC 连接
}

// BaseConfig 是 XrayController 的配置结构体
type BaseConfig struct {
	APIAddress string
	APIPort    int
}

// Init 初始化 XrayController 并建立连接
func (xrayCtl *XrayController) Init(cfg *BaseConfig) (err error) {
	// 先取得 ClientConn, 用完记得 close
	xrayCtl.CmdConn, err = grpc.Dial(fmt.Sprintf("%s:%d", cfg.APIAddress, cfg.APIPort), grpc.WithInsecure())
	if err != nil {
		return err
	}

	// 依次获取 API Client, 可根据需求删减
	xrayCtl.SsClient = statsService.NewStatsServiceClient(xrayCtl.CmdConn)

	return
}
