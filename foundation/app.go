package foundation

import (
	"context"
	"errors"
	"fmt"
	"ginson/foundation/cache"
	"ginson/foundation/cfg"
	"ginson/foundation/database"
	"ginson/foundation/kafka"
	"ginson/foundation/log"
	middleware2 "ginson/foundation/middleware"
	"ginson/foundation/mongo"
	"ginson/pkg/code"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// StartServer 应用入口点
func StartServer(registerRoutes func(*gin.Engine)) {
	// 初始化相关依赖组件
	err := initialize(cfg.AppConf)
	if err != nil {
		panic(fmt.Sprintf("initialize failed: %s", err))
	}

	// 启动Web服务
	err = startServer(cfg.AppConf, registerRoutes)
	if err != nil {
		panic(fmt.Sprintf("Server started failed: %s", err))
	}
}

// initialize 初始化组件依赖
func initialize(cfg *cfg.AppConfig) error {
	// 初始化日志
	log.Init(cfg)

	// 初始化数据库
	err := database.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("init db failed, error: %s", err)
	}

	err = mongo.InitMongo(cfg)
	if err != nil {
		return fmt.Errorf("init db failed, error: %s", err)
	}

	// 初始化Redis
	err = cache.InitRedis(cfg)
	if err != nil {
		return fmt.Errorf("init foundation failed, error: %s", err)
	}

	// 初始化Redis Cluster
	err = cache.InitRedisCluster(cfg)
	if err != nil {
		return fmt.Errorf("init foundation cluster failed, error: %s", err)
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
func startServer(cfg *cfg.AppConfig, registerRoutes func(*gin.Engine)) error {
	server := &http.Server{
		Addr:    cfg.HttpAddr + ":" + strconv.Itoa(cfg.HttpPort),
		Handler: initEngine(cfg, registerRoutes),
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
func initEngine(cfg *cfg.AppConfig, registerRoutes func(*gin.Engine)) *gin.Engine {
	gin.SetMode(func() string {
		if cfg.IsDevEnv() {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())

	engine := gin.New()

	// 性能监控中间件
	// to look at the heap profile: go tool ip:port/dev/pprof/heap
	// to look at a 30-second CPU profile: go tool ip:port/dev/pprof/profile
	// to look at the goroutine blocking profile: go tool ip:port/dev/pprof/block
	// to collect a 5-second execution trace: wget ip:port/debug/pprof/trace?seconds=5
	pprof.Register(engine, "dev/pprof")

	engine.Use(middleware2.Trace())
	engine.Use(middleware2.Logger())
	engine.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		log.Error(c, "panic recovery: %v", err)
		c.AbortWithStatusJSON(http.StatusOK, code.ServerError)
	}))

	registerRoutes(engine)

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
			fmt.Printf("Failed to close logger: %s\n", err)
		}
		if err := log.AccessLogger.Sync(); err != nil {
			fmt.Printf("Failed to close access logger: %s\n", err)
		}
	}()

	// 资源释放
	cache.Close()
	kafka.Close()

	// 关闭server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error(context.Background(), "Shutdown server failed, error: %v", err)
	}
}
