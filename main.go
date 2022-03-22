package main

import (
	"context"
	"errors"
	"fmt"
	"ginson/pkg/conf"
	"ginson/pkg/log"
	"ginson/pkg/middleware"
	"ginson/pkg/router"
	"ginson/platform/cache"
	"ginson/platform/database"
	"ginson/platform/kafka"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// 初始化相关依赖组件
	err := initialize(conf.AppConf)
	if err != nil {
		panic(fmt.Sprintf("initialize failed: %s", err))
	}

	// 启动Web服务
	err = startServer(conf.AppConf)
	if err != nil {
		panic(fmt.Sprintf("Server started failed: %s", err))
	}
}

// initialize 初始化组件依赖
func initialize(cfg *conf.AppConfig) error {
	// 初始化日志
	log.Init(cfg)

	// 初始化数据库
	err := database.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("init db failed, error: %s", err)
	}

	err = database.InitMongo(cfg)
	if err != nil {
		return fmt.Errorf("init db failed, error: %s", err)
	}

	// 初始化Redis
	err = cache.InitRedis(cfg)
	if err != nil {
		return fmt.Errorf("init redis failed, error: %s", err)
	}

	// 初始化Redis Cluster
	err = cache.InitRedisCluster(cfg)
	if err != nil {
		return fmt.Errorf("init redis cluster failed, error: %s", err)
	}

	// 初始化kafka
	err = kafka.InitProducer(cfg)
	if err != nil {
		return fmt.Errorf("init kafka producer failed, error: %s", err)
	}

	// 初始化kafka
	err = kafka.InitConsumer(cfg)
	if err != nil {
		return fmt.Errorf("init kafka consumer failed, error: %s", err)
	}

	return nil
}

// startServer 启动服务
func startServer(cfg *conf.AppConfig) error {
	server := &http.Server{
		Addr:    cfg.HttpAddr + ":" + strconv.Itoa(cfg.HttpPort),
		Handler: initEngine(cfg),
	}
	ctx, cancel := context.WithCancel(context.Background())
	go listenToSystemSignals(cancel)

	go func() {
		<-ctx.Done()
		shutdown(server)
	}()

	log.Info(context.Background(), "Server started success")
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Info(context.Background(), "Server was shutdown gracefully")
		return nil
	}

	return err
}

// 初始化gin路由
func initEngine(cfg *conf.AppConfig) *gin.Engine {
	gin.SetMode(func() string {
		if cfg.IsDevEnv() {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())

	engine := gin.New()

	engine.Use(middleware.Trace())
	engine.Use(middleware.Logger())
	engine.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		log.Error(c, "panic recovery: %v", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "服务器内部错误，请稍后再试！",
		})
	}))

	router.RegisterRoutes(engine)

	return engine
}

// 监听系统命令
func listenToSystemSignals(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			cancel()
			return
		}
	}
}

// 关闭端口
func shutdown(server *http.Server) {
	time.Sleep(5 * time.Second)
	// 最后释放log
	defer func() {
		if err := log.Logger.Sync(); err != nil {
			fmt.Printf("FailedWithCode to close log: %s\n", err)
		}
	}()

	// 资源释放
	cache.Close()
	kafka.Close()

	// 关闭server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error(context.Background(), "FailedWithCode to shutdown server: %v", err)
	}
}
