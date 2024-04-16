package utils

import (
	"fmt"
	"github.com/go-ping/ping"
	"time"
)

// PingByIp 模拟ping进行简单探测
func PingByIp(ipAddr string, count int) *ping.Pinger {
	// 创建Ping
	pinger, err := ping.NewPinger(ipAddr)
	if err != nil {
		fmt.Printf("Error creating pinger: %v\n", err)

	}
	// 设置Ping参数
	pinger.Count = count                     // 发送count个Ping请求
	pinger.Timeout = 3 * time.Second         // 设置超时时间
	pinger.Interval = 100 * time.Millisecond // 设置Ping请求间隔
	err = pinger.Run()
	if err != nil {
		fmt.Printf("Error running pinger: %v\n", err)
	}
	return pinger

}
