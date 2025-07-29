package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samber/do/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/smartmemos/memos/internal/api"
	"github.com/smartmemos/memos/internal/config"
	"github.com/smartmemos/memos/internal/memos/service"
	"github.com/smartmemos/memos/internal/server"
)

var (
	cfgFile   string
	container do.Injector
	Version   string = "unknown"
	BuildTime string = "unknown"
	GitCommit string = "unknown"

	rootCmd = &cobra.Command{
		Use:   "memos",
		Short: "",
		Long:  "",
	}

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			serverCfg := config.GetConfig().Server
			ctx, cancel := context.WithCancel(context.Background())
			sv, err := server.NewServer(&server.Profile{
				Addr:      fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.Port),
				Container: container,
			})
			if err != nil {
				log.Fatalln(err)
			}

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			if err = sv.Start(ctx); err != nil {
				log.Fatalln(err)
			}
			go func() {
				<-quit
				sv.Shutdown(ctx)
				cancel()
			}()
			<-ctx.Done()
		},
	}
)

func main() {
	config.Version = Version
	config.BuildTime = BuildTime
	config.GitCommit = GitCommit
	container = do.New()

	log.SetFormatter(&log.JSONFormatter{})

	// 设置固定时区, Docker默认时区是UTC
	if l, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		time.Local = l
	}

	cobra.OnInitialize(func() {
		config.Init(cfgFile)
		service.Init(container)
		api.Init(container)
	})

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.memos/config.toml)")
	rootCmd.AddCommand(serveCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Exit(0)
}
