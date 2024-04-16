package router

import (
	"github.com/gin-gonic/gin"
	"machinesearch/service"
)

func SetApiGroupRoutes(r *gin.RouterGroup) {

	r.GET("/ping", service.Ping)

	r.GET("ips", service.GetIpList)

	//获取IP之间的探测关系
	r.GET("/rp", service.GetIpsRelationship)
}
