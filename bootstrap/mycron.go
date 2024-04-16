package bootstrap

import (
	"machinesearch/Cron"
)

var (
	Count = 10
)

//var once sync.Once

func InitializeCron() {
	Cron.Corns()
}
