package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pkg "github.com/19980730520/queryTraffic/pkg"
)

func main() {
	// 定义命令行标志
	var (
		serverAddr = flag.String("s", "", "Server address in the format 'host:port'")
		pattern    = flag.String("ptn", "", "Traffic pattern to query")
	)
	flag.Parse()

	// 检查必要参数是否提供
	if *serverAddr == "" || *pattern == "" {
		flag.Usage()
		os.Exit(1)
	}

	// 解析服务器地址
	addrParts := strings.Split(*serverAddr, ":")
	if len(addrParts) != 2 {
		log.Fatalf("Invalid server address format. Use 'host:port'.")
	}
	apiAddress := addrParts[0]
	apiPort, err := strconv.Atoi(addrParts[1])
	if err != nil {
		log.Fatalf("Invalid port number: %s", err)
	}

	// 创建 XrayController 实例
	xrayCtl := new(pkg.XrayController)

	// 配置信息
	cfg := &pkg.BaseConfig{
		APIAddress: apiAddress,
		APIPort:    apiPort,
	}

	// 初始化 XrayController 并捕获错误
	if err := xrayCtl.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize XrayController: %s", err)
	}
	defer xrayCtl.CmdConn.Close() // 确保程序退出时关闭连接

	// 调用 queryTraffic 函数查询流量数据
	trafficData, err := pkg.QueryTraffic(xrayCtl.SsClient, *pattern, false)
	if err != nil {
		log.Fatalf("Failed to query traffic: %s", err)
	}

	// 打印查询到的流量数据
	fmt.Printf("流量: %d bytes\n", trafficData)

	// 程序正常退出
	os.Exit(0)
}
