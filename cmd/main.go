package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"nico/conf"
	"nico/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	gin.SetMode("release")
	err := conf.Init("./logs/")
	if err != nil {
		fmt.Printf("log init err %s\n", err.Error())
		return
	}
	conf.Init_logrus("./logs/")

	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, "pong")
	})
	r.GET("/ws", service.WebSocketHandler)
	r.GET("/changelog", service.ChangeLog)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	logrus.Debug("init project")
	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Debug("Shutting down server...")
	logrus.Warn("Shutting down server..")
	logrus.Info("Shutting down server..")
	logrus.Error("Shutting down server..")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown: ", err)
	}

	logrus.Debug("Server exiting")
}
