package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/data-market/internal/logs"
	"github.com/data-market/server"
	"github.com/spf13/cobra"
)

var logger = logs.Logger("db")

var (
	port string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a server",
	Run: func(cmd *cobra.Command, args []string) {
		// new http server with port
		srv := server.StartServer(port)

		go func() {
			logger.Debugf("start http server at port:%s", port)
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
	runCmd.Flags().StringVarP(&port, "port", "p", "8080", "listen port")

	DaemonCmd.AddCommand(runCmd, stopCmd)
}
