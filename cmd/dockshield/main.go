package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/abubakargithub/dockshield/internal/config"
	"github.com/abubakargithub/dockshield/internal/logging"
	"github.com/abubakargithub/dockshield/internal/monitor"
	"github.com/abubakargithub/dockshield/internal/metrics"
	"github.com/abubakargithub/dockshield/internal/trivy"
)

func main() {
	config.Load()
	logging.Init(false)
	defer logging.Sync()

	rootCmd := &cobra.Command{
		Use:   "dockshield",
		Short: "DockShield - Docker Guardian and Security Agent",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "check",
		Short: "Run Docker health diagnostics",
		Run: func(cmd *cobra.Command, args []string) {
			monitor.CheckDockerHealth()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "scan",
		Short: "Run container vulnerability scan using Trivy",
		Run: func(cmd *cobra.Command, args []string) {
			trivy.Scan()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "metrics",
		Short: "Start Prometheus metrics server",
		Run: func(cmd *cobra.Command, args []string) {
			metrics.StartServer()
		},
	})

	if err := rootCmd.Execute(); err != nil {
		zap.L().Fatal("command failed", zap.Error(err))
	}
}
