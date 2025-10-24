package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("⚠️  Using default configuration, config.yaml not found.")
		return
	}
	fmt.Println("✅ Configuration loaded:", viper.ConfigFileUsed())
}
