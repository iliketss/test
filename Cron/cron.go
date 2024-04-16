package Cron

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/robfig/cron/v3"
	"machinesearch/app/models"
	"machinesearch/global"
	"machinesearch/utils"
	"time"
)

// Corns
/*
@yearly (or @annually)	每年1月1日凌晨执行一次	0 0 0 1 1 *
@monthly	每个月第一天的凌晨执行一次	0 0 0 1 * *
@weekly	每周周六的凌晨执行一次	0 0 0 * * 0
@daily (or @midnight)	每天凌晨0点执行一次	0 0 0 * * *
@hourly	每小时执行一次	0 0 * * * *
*/
var c *cron.Cron
var i int

func Corns() {
	//获取除本机外的所有已注册IP
	var machines []models.MachineMsg

	global.App.DB.Where("local_ip != ? ", utils.GetLocalIp().String()).Find(&machines)
	fmt.Println("cron:", len(machines))
	c = cron.New(cron.WithSeconds())
	for i = 0; i < len(machines); i++ {
		// 添加定时探测任务 测试每五秒来一次
		res := models.Result{}

		res.DetectingSourceIP = utils.GetLocalIp().String()
		res.DetectionDestinationIP = machines[i].LocalIP
		_, err := c.AddFunc("*/5 * * * * *", func() {

			//需要在任务中重置ID
			res.ID = 0
			res.CreatedAt = time.Now()
			res.UpdatedAt = time.Now()
			p := utils.PingByIp(res.DetectionDestinationIP, 10)

			p.OnFinish = func(stats *ping.Statistics) {
				res.PacketLossRate = float32(stats.PacketLoss) / 100
				res.Delayed = time.Duration(stats.AvgRtt.Seconds() * 1000)
			}

			// 开始Ping
			err := p.Run()
			if err != nil {
				fmt.Printf("Error running pinger: %v\n", err)
			}
			global.App.DB.Create(&res)
			fmt.Println("当前执行ping ", res.DetectionDestinationIP, "操作")
		})
		if err != nil {
			fmt.Println("添加定时任务失败：", err)
			return
		} else {
			fmt.Println("添加定时任务成功")
		}
	}
	// 启动调度器
	c.Start()

	// 阻塞主线程
	select {}
}
