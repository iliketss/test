package utils

import (
	"fmt"
	"net"
	"os"
)

func GetLocalIp() net.IP {
	// 获取本机主机名
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get hostname:", err)

	}
	// 解析主机名的IP地址
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("Failed to lookup IP address:", err)

	}
	return addrs[3]
}
