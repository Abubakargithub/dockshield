package monitor

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)

// CheckDockerHealth runs an initial Docker health check.
func CheckDockerHealth() {
	fmt.Println("🔍 Checking Docker daemon health...")

	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Docker is not running or not installed.")

		if viper.GetBool("monitor.auto_restart_docker") {
			fmt.Println("♻️ Attempting to restart Docker service...")
			restart := exec.Command("net", "start", "com.docker.service")
			if err := restart.Run(); err != nil {
				fmt.Println("⚠️  Failed to restart Docker:", err)
			} else {
				fmt.Println("✅ Docker restarted successfully.")
			}
		}
	} else {
		fmt.Println("✅ Docker daemon is running.")
		if viper.GetBool("optimize.auto_prune") {
			fmt.Println("🧹 Running Docker system prune...")
			exec.Command("docker", "system", "prune", "-af").Run()
		}
	}
}

// StartContinuousMonitor runs periodic health checks.
func StartContinuousMonitor() {
	interval := viper.GetInt("monitor.interval_seconds")
	if interval <= 0 {
		interval = 60
	}

	for {
		CheckDockerHealth()
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
