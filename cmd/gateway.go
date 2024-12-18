package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"ush/internal/gateway/controller/http"
	"ush/internal/gateway/service"
	"ush/internal/pkg/config"

	"github.com/spf13/cobra"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Runs the gateway service",
	Long: `
	Gateway runs the gateway service for this microservice app,
	Reading configuration files from 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gateway called")
		if err := startGateway(&cfg.GatewayConfig); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

func startGateway(cfg *config.GatewayConfig) error {
  if cfg== nil {
    panic("config struct is nil")
  }
	// new gateway service
	srv, err := service.NewService(cfg)
	if err != nil {
		return err
	}

	// new controller with service inside it
	ctrl := http.New(srv)

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-sig

		if err := ctrl.Stop(); err != nil {
			log.Fatal(err)
		}
	}()

	// start controller to serve http
	if err = ctrl.Start(fmt.Sprintf("0.0.0.0:%d", cfg.ListenPort)); err != nil {
		return err
	}

	return nil
}
