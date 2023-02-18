package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// SetURL 请求配置地址
func SetURL() {

	var tmp = make(chan interface{}, 1)

	router := gin.Default()
	router.POST("/webhook/event", func(context *gin.Context) {

		resqMap := getMessage(context)
		//fmt.Println(resqMap["challenge"])
		context.JSON(200, gin.H{
			"challenge": resqMap["challenge"],
		})
		tmp <- true
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 用协程初始化
	// 将不会阻塞停止服务
	go func() {
		fmt.Println("1")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	<-tmp
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func getMessage(context *gin.Context) map[string]interface{} {
	var resqMap map[string]interface{}
	err := context.Bind(&resqMap)
	if err != nil {
		fmt.Println(err)
	}
	return resqMap
}
