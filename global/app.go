package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"machinesearch/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
}

var App = new(Application)
