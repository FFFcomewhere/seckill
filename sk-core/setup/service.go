package setup

import (
	"fmt"
	"github.com/FFFcomewhere/seckill/pkg/bootstrap"
	register "github.com/FFFcomewhere/seckill/pkg/discover"
	"github.com/FFFcomewhere/seckill/sk-core/service/srv_redis"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func RunService() {
	//启动处理线程
	srv_redis.RunProcess()
	errChan := make(chan error)
	//http server

	go func() {
		//启动前执行注册
		register.Register()
		//TODO 做测试 可以删除
		r := gin.Default()
		r.GET("/hello", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello core",
			})
		})

		r.Run(":" + bootstrap.HttpConfig.Port)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	register.Deregister()
	fmt.Println(error)
}
