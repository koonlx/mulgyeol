package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Gin 라우터 생성
    router := gin.Default()

    // 핸들러 등록
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    // 서버 실행
    router.Run(":8080") // 기본적으로 8080 포트에서 실행
}
