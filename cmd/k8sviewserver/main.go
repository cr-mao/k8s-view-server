package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/bootstrap"
	"github.com/cr-mao/k8s-view-server/infra/logger"
)

var (
	BuildTime string
	Version   string
)

func init() {
	// 注入编译时间和git版本  后面用
	log.Printf("Version: %s", Version)
	log.Printf("BuildTime: %s", BuildTime)
}

func main() {
	env := flag.String("env", "local", "--env")
	port := flag.String("port", "8088", "--port")
	flag.Parse()
	bootstrap.Bootstrap(*env)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//http start
	srv := k8sadmin.NewServe(*port)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.ErrorString("http", "serve", err.Error())
		}
	}()
	<-quit
	// shutdown http server
	newCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(newCtx); err != nil {
		logger.WarnString("http", "shut_down", err.Error())
	}
}
