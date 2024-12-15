package cmd

import (
	"fmt"
	"log"
	"ush/internal/config"

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

var configPath *string
var cfg = config.Config{}

func init() {
	rootCmd.AddCommand(gatewayCmd)

	configPath = gatewayCmd.Flags().StringP("config", "c", "config.json", "configuration file pass")
	if err := config.Parse(&cfg, *configPath); err != nil {
		log.Fatal(err)
	}
}

func startGateway(cfg *config.GatewayConfig) error {
	// new gateway service
	// new controller with service inside it
	// start controller to serve http
	// graceful shutdown
	return nil
}
