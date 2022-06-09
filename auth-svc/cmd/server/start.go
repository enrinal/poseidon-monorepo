package main

import (
	"strings"

	v1 "github.com/enrinal/poseidon-monorepo/auth-svc/internal/api/v1"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/config"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/entities"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/errors"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/log"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/repository"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/services"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	configPath string
	startCmd   = &cobra.Command{
		Use:   "start",
		Short: "start server",
		Long:  `start server, default port is 5000`,
		Run:   startServer,
	}
	enablePprof bool
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is $PWD/config/default.yaml)")
	startCmd.PersistentFlags().Int("port", 5000, "Port to run Application server on")
	startCmd.PersistentFlags().BoolVarP(&enablePprof, "pprof", "p", false, "enable pprof mode (default: false)")
	config.Viper().BindPFlag("port", startCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	if len(configPath) != 0 {
		config.Viper().SetConfigFile(configPath)
	} else {
		config.Viper().AddConfigPath("./config")
		config.Viper().SetConfigName("default")
	}
	config.Viper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.Viper().AutomaticEnv()
	if err := config.Viper().ReadInConfig(); err != nil {
		log.Fatalf("Load config from file [%s]: %v", config.Viper().ConfigFileUsed(), err)
	}
	config.Parse()
}

func startServer(cmd *cobra.Command, agrs []string) {
	log.Info("Start http-server")
	db, err := gorm.Open(sqlite.Open(config.Default.Database.URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&entities.User{})

	// repositories
	authRepo := repository.NewAuthRepository(db)

	// services
	authService := services.NewAuthService(authRepo)

	router := gin.New()
	router.Use(errors.GinError())
	router.Use(gin.Recovery())
	if enablePprof {
		pprof.Register(router, "monitor/pprof")
	}
	apiV1Router := router.Group("/api/v1")
	v1.RegisterRouterAPIV1(apiV1Router, authService)

	router.Run(config.Default.Server.Port)
}
