package main

import (
	"fmt"
	_ "lottery/model"
	"lottery/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	//config.AllowOrigins = []string{"http://zlican.com"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"}
	config.AllowHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Authorization",
		"Accept",
		"Origin",
		"Referer",
		"User-Agent",
		"X-Requested-With",
	}
	config.MaxAge = 86400

	r.Use(cors.New(config))

	router.SetupRouter(r)

	fmt.Println("服务器正在启动，监听端口 8000...")
	if err := r.Run(":8000"); err != nil {
		fmt.Printf("启动服务器失败: %v\n", err)
		panic(err)
	}
}
