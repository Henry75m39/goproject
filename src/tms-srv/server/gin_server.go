package server

import (
	"GRM/src/common/utils/log"
	"GRM/src/tms-srv/provider"
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func GinServer() {
	router := gin.Default()
	logger := log.Instance()
	//Please keep in mind to comment out this debug mode when release to prod env
	gin.SetMode("debug")
	// server management. init server
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// invoke server router group.
	provider.ServiceRouterGroup(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error", zap.Any("listen: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Info", zap.Any("Gin:", "Shutdown Server ..."))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Error", zap.Any("TMS Server Shutdown:", err))
	}
	logger.Info("Info", zap.Any("Gin:", "TMS Server exiting ..."))

}
