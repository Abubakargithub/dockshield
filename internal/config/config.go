package config

import (
    "fmt"

    "github.com/spf13/viper"
)

func Load() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./configs")
    viper.AddConfigPath(".")

    viper.SetDefault("monitor.interval_seconds", 30)
    viper.SetDefault("monitor.auto_restart_docker", true)
    viper.SetDefault("optimize.auto_prune", true)
    viper.SetDefault("metrics.enabled", true)
    viper.SetDefault("metrics.addr", ":9091")

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("⚠️  Using default configuration, config.yaml not found.")
        return
    }
    fmt.Println("✅ Configuration loaded:", viper.ConfigFileUsed())
}

func Print() {
    fmt.Println(viper.AllSettings())
}
