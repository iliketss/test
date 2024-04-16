package models

import (
	"gorm.io/gorm"
)

// MachineMsg
// 机器信息表
// 注册时间、机器IP地址
type MachineMsg struct {
	gorm.Model

	LocalIP string `gorm:"unique"`
}
