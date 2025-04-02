package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/data-market/internal/database"
	"github.com/data-market/internal/dumper"
	"github.com/data-market/internal/logs"
	"github.com/data-market/server"
	"github.com/spf13/cobra"
)

var logger = logs.Logger("db")

var (
	port string
	env  string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a server",
	Run: func(cmd *cobra.Command, args []string) {

		logger.Debug("env:", env)

		_ = database.G_DB

		d, err := dumper.NewDumper(env)
		if err != nil {
			log.Fatalf("new dumper failed: %s", err)
		}

		// start subscribe the block
		logger.Debug("start subscribe blocks")
		go d.Subscribe(context.Background())

		// new http server with port
		srv := server.StartServer(port)
		// run http server
		go func() {
			logger.Infof("start http server at port:%s", port)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("Shutting down server...")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "daemon commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

	},
}

func init() {
	// add flag info for run cmd
	runCmd.Flags().StringVarP(&port,
		"port",
		"p",
		"8080",
		"listen port")

	runCmd.Flags().StringVarP(
		&env,
		"env",
		"e",
		"test", // 默认值
		"设置运行环境\n可选值: test（测试环境）, dev（开发环境）, product（生产环境）\n默认使用 test 环境", // 详细说明
	)

	// 如果希望参数必填（默认值设为空字符串后）：
	runCmd.MarkFlagRequired("env")

	DaemonCmd.AddCommand(runCmd, stopCmd)
}
