package models

import (
	"gorm.io/gorm"
	"time"
)

// Result
// 探测结果表
// 数据插入时间，探测源IP，探测目的IP，延时，丢包率
type Result struct {
	gorm.Model
	DetectingSourceIP      string
	DetectionDestinationIP string
	Delayed                time.Duration
	PacketLossRate         float32
}
