package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"ush/pkg/config"
	"ush/internal/shortener/controller/rsrv"
	"ush/internal/shortener/repository/sqlite"
	"ush/internal/shortener/service"

	"github.com/spf13/cobra"
)

// shortenerCmd represents the shortener command
var shortenerCmd = &cobra.Command{
	Use:   "shortener",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(configPath)
    if err := startShortener(&cfg.ShortenerConfig); err != nil {
      panic(err)
    }
	},
}

func init() {
	rootCmd.AddCommand(shortenerCmd)
}

func startShortener(cfg *config.ShortenerConfig) error {
	if cfg == nil {
		panic("config struct is nil")
	}

	repo, err := sqlite.New(cfg)
	if err != nil {
		return err
	}

	service := service.New(repo)

	ctrl, err := rsrv.New(cfg, service)
	if err != nil {
		return err
	}

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-sig

		if err := ctrl.Stop(); err != nil {
			log.Fatal(err)
		}

	}()

	if err = ctrl.Start(fmt.Sprintf("0.0.0.0:%d", cfg.ListenPort)); err != nil {
		return err
	}

	return nil
}
