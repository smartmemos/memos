package config

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v8"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"

	"github.com/smartmemos/memos/internal/pkg/db"
)

// 版本信息，在编译时自动生成
var (
	Version   = "unknown"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

type Config struct {
	Server   serverConfig `envPrefix:"SERVER_"`
	Logger   loggerConfig `envPrefix:"LOGGER_"`
	Database db.Config    `envPrefix:"DATABASE_"`
	JWT      JWT          `envPrefix:"JWT_"`
}

type loggerConfig struct {
	Format string
	Level  string
	Output string
}

type JWT struct {
	Key string
}

type serverConfig struct {
	Host  string
	Port  int
	Debug bool
}

var appPath string
var config *Config

// GetConfig 获取配置
func GetConfig() *Config {
	return config
}

// Init 初始化配置
func Init(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config err: %v", err)
	}
	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	// 将env配置加载到环境变量
	configPath := filepath.Dir(viper.ConfigFileUsed())
	if err := gotenv.Load(filepath.Join(configPath, ".env")); err != nil {
		log.Debug(err)
	}
	// 读取环境变量
	if err := env.ParseWithOptions(config, env.Options{Prefix: "MEMOS_", UseFieldNameByDefault: true}); err != nil {
		log.Fatalf("bind environment err: %v", err)
	}

	log.Infof("config: %v", config)
	db.Init(config.Database)
	if err = initLogger(config.Logger); err != nil {
		log.Fatalf("init logger err: %v", err)
	}
}

func initLogger(cfg loggerConfig) error {
	if cfg.Format == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	lvl, err := log.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	if cfg.Output != "" {
		logFile, err := os.OpenFile(cfg.Output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		log.SetOutput(logFile)
		log.RegisterExitHandler(func() {
			log.Info("退出logger......")
			logFile.Close()
		})
	}
	return nil
}
