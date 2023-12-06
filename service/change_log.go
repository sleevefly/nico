package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Changelog struct {
	Level string `json:"level"`
}

func ChangeLog(ctx *gin.Context) {
	var req Changelog
	ctx.ShouldBindQuery(&req)
	change(req.Level)
}

func change(level string) {
	// 注册日志级别变更处理器
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return
	}
	logrus.SetLevel(parseLevel)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	// 记录日志消息
	log.Println("This is a log message.")
	fmt.Fprintf(w, "Hello, world!")
}
