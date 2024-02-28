package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/ryoook/endless"
	"net/http"
	"os"
	"strconv"
)

var (
	port int
)

func run(ctx context.Context, routerHandler func(router *gin.Engine)) {
	// 初始化环境变量
	err := initVariable()
	if err != nil {
		color.Error.Printf("encounter error when starting server with %s\n", err.Error())
		return
	}

	// 初始化路由
	router := gin.Default()
	routerHandler(router)

	// 启动
	errChan := make(chan error)
	server := endless.NewServer(fmt.Sprintf(":%d", port), router)
	go func() {
		errChan <- server.ListenAndServe()
	}()
	color.Info.Printf("server has been started, port:%v\n", port)
	select {
	case err = <-errChan:
		if err != http.ErrServerClosed {
			color.Error.Printf("encounter error when starting server with %s", err.Error())
		}
	}
}

func initVariable() error {
	p, _ := os.LookupEnv("PORT")
	port, _ = strconv.Atoi(p)

	if port == 0 {
		return fmt.Errorf("invalid port: %s", p)
	}
	return nil
}
