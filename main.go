package main

import (
	"context"
	"fmt"
	"go_blog/pkg/setting"
	"go_blog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "go_blog/docs"
)

//@title Swagger Example API
//@version 0.0.1
//@description This is a sample server blog server
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//ListenAndServer总是返回一个错误
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//开始关闭
	log.Printf("Shutdown Server... ")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//ctx会等待5秒让Shutdown执行，如果这5秒Shutdown执行完了，那么err就是server自己的错误。
	//如果5秒内没执行完，那么就会返回ctx的错误
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
