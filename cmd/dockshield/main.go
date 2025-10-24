package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "github.com/yourorg/dockshield/internal/config"
    "github.com/yourorg/dockshield/internal/logging"
)

func main() {
    config.Load()
    logging.Init(false)
    defer logging.Sync()

    rootCmd := &cobra.Command{Use: "dockshield", Short: "DockShield - Docker Guardian"}
    rootCmd.AddCommand(&cobra.Command{
        Use: "check",
        Short: "Run diagnostic checks",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Running DockShield diagnostics...")
        },
    })
    if err := rootCmd.Execute(); err != nil {
        zap.L().Fatal("Command failed", zap.Error(err))
    }
}
