package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"server/internal/config"
	"server/internal/router"
	"server/internal/shared/database"
	"server/internal/shared/jwt"
	"server/internal/shared/logger"
	"server/internal/shared/utils"
	"server/internal/shared/validator"
	"server/internal/system/loginlog"
	"server/internal/system/operlog"
	"server/internal/system/user"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yaml", "配置文件路径")
}

func main() {
	flag.Parse()

	// 初始化配置
	if err := config.Init(configPath); err != nil {
		panic(fmt.Sprintf("初始化配置失败: %v", err))
	}

	// 初始化日志
	logger.Init(
		config.Conf.Log.Filename,
		config.Conf.Log.Level,
		config.Conf.Log.MaxSize,
		config.Conf.Log.MaxBackups,
		config.Conf.Log.MaxAge,
		config.Conf.Log.Compress,
	)

	// 初始化验证器
	validator.Init()

	// 初始化数据库
	dbConfig := database.Config{
		Driver: config.Conf.Database.Driver,
		MySQL: database.MySQL{
			Host:            config.Conf.Database.Host,
			Port:            config.Conf.Database.Port,
			Username:        config.Conf.Database.Username,
			Password:        config.Conf.Database.Password,
			Database:        config.Conf.Database.Database,
			Charset:         config.Conf.Database.Charset,
			MaxIdleConns:    config.Conf.Database.MaxIdleConns,
			MaxOpenConns:    config.Conf.Database.MaxOpenConns,
			ConnMaxLifetime: config.Conf.Database.ConnMaxLifetime,
		},
		PostgreSQL: database.PostgreSQL{
			Host:            config.Conf.Database.Host,
			Port:            config.Conf.Database.Port,
			Username:        config.Conf.Database.Username,
			Password:        config.Conf.Database.Password,
			Database:        config.Conf.Database.Database,
			MaxIdleConns:    config.Conf.Database.MaxIdleConns,
			MaxOpenConns:    config.Conf.Database.MaxOpenConns,
			ConnMaxLifetime: config.Conf.Database.ConnMaxLifetime,
		},
		SQLite: database.SQLite{
			Path: config.Conf.Database.Path,
		},
	}
	if err := database.Init(dbConfig); err != nil {
		logger.Error("数据库初始化失败")
		panic(fmt.Sprintf("数据库初始化失败: %v", err))
	}

	// 自动建表
	if err := database.DB.AutoMigrate(
		&user.User{},
		&operlog.OperLog{},
		&loginlog.LoginLog{},
	); err != nil {
		logger.Error("自动建表失败")
		panic(fmt.Sprintf("自动建表失败: %v", err))
	}

	// 初始化验证码存储
	if config.Conf.Captcha.Enabled {
		utils.NewCaptchaStore(config.Conf.Captcha.Expire)
	}

	// 初始化JWT
	jwtManager := jwt.New(config.Conf.Jwt.Secret, config.Conf.Jwt.Expire)

	// 设置Gin模式
	gin.SetMode(config.Conf.Server.Mode)

	// 创建引擎
	r := gin.New()

	// 注册路由
	router.Register(r, jwtManager)

	// 启动服务
	addr := fmt.Sprintf(":%d", config.Conf.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Duration(config.Conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Conf.Server.WriteTimeout) * time.Second,
	}

	logger.Info(fmt.Sprintf("服务启动成功，监听端口: %s", addr))

	// 优雅关闭
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("服务启动失败: %v", err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("服务正在关闭...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("服务关闭异常: %v", err))
	}

	logger.Info("服务已关闭")
}