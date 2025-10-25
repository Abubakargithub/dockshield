package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/spf13/cobra"
    "go.uber.org/zap"

    "github.com/abubakargithub/dockshield/internal/config"
    "github.com/abubakargithub/dockshield/internal/logging"
    "github.com/abubakargithub/dockshield/internal/monitor"
    "github.com/abubakargithub/dockshield/internal/metrics"
    "github.com/abubakargithub/dockshield/internal/trivy"
    "github.com/abubakargithub/dockshield/internal/optimizer"
    "github.com/abubakargithub/dockshield/internal/service"
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
        Use:   "start",
        Short: "Start monitor in foreground",
        Run: func(cmd *cobra.Command, args []string) {
            c := make(chan os.Signal, 1)
            signal.Notify(c, os.Interrupt, syscall.SIGTERM)
            go monitor.StartContinuousMonitor()
            fmt.Println("Monitoring started. Ctrl+C to stop.")
            <-c
            fmt.Println("Stopping monitor...")
        },
    })

    rootCmd.AddCommand(&cobra.Command{
        Use:   "metrics",
        Short: "Start metrics server",
        Run: func(cmd *cobra.Command, args []string) {
            metrics.StartServer()
        },
    })

    rootCmd.AddCommand(&cobra.Command{
        Use:   "scan",
        Short: "Scan images with Trivy",
        Run: func(cmd *cobra.Command, args []string) {
            trivy.ScanAllLocalImages()
        },
    })

    rootCmd.AddCommand(&cobra.Command{
        Use:   "optimize",
        Short: "Run prune optimizer",
        Run: func(cmd *cobra.Command, args []string) {
            optimizer.RunPrune()
        },
    })

    rootCmd.AddCommand(&cobra.Command{
        Use:   "version",
        Short: "Print version",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("DockShield v1.2.0 - abubakargithub/dockshield")
        },
    })

    if err := rootCmd.Execute(); err != nil {
        zap.L().Fatal("command failed", zap.Error(err))
    }
}
