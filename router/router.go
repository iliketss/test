package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"machinesearch/global"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	gin.SetMode("release")
	router := gin.New()
	// 注册 api 分组路由
	apiGroup := router.Group("")
	SetApiGroupRoutes(apiGroup)
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	//启动定时任务

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
