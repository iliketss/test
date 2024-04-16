package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
	"machinesearch/app/models"
	"machinesearch/global"
	"machinesearch/utils"
	"net/http"
	"time"
)

func GetIpList(c *gin.Context) {

	//获取除本机外的所有已注册IP
	var machines []models.MachineMsg
	res := models.Result{}
	res.DetectingSourceIP = utils.GetLocalIp().String()

	global.App.DB.Where("local_ip != ? ", utils.GetLocalIp().String()).Find(&machines)

	//实时获取ping结果
	for i := 0; i < len(machines); i++ {
		res.DetectionDestinationIP = machines[i].LocalIP
		fmt.Println(machines[i].LocalIP)
		p := utils.PingByIp(machines[i].LocalIP, 10)

		p.OnFinish = func(stats *ping.Statistics) {
			res.PacketLossRate = float32(stats.PacketLoss) / 100
			res.Delayed = time.Duration(stats.AvgRtt.Seconds() * 1000)
			fmt.Printf("%d packets transmitted, %d packets received, %.1f%% packet loss\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Printf("round-trip min/avg/max/stddev = %.3f/%.3f/%.3f/%.3f ms\n",
				stats.MinRtt.Seconds()*1000, stats.AvgRtt.Seconds()*1000, stats.MaxRtt.Seconds()*1000, stats.StdDevRtt.Seconds()*1000)
		}

		// 开始Ping
		err := p.Run()
		if err != nil {
			fmt.Printf("Error running pinger: %v\n", err)

		}

	}

	global.App.DB.Create(&res)

	c.JSON(http.StatusOK, res)
}

func GetIpsRelationship(c *gin.Context) {
	var res []models.Result
	global.App.DB.Find(&res)
	c.JSON(http.StatusOK, res)
}
