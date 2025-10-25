package monitor

import (
    "fmt"
    "os/exec"
    "time"

    "github.com/spf13/viper"
)

func CheckDockerHealth() {
    fmt.Println("🔍 Checking Docker daemon health...")
    cmd := exec.Command("docker", "info")
    if err := cmd.Run(); err != nil {
        fmt.Println("❌ Docker is not running or not installed.")
        if viper.GetBool("monitor.auto_restart_docker") {
            fmt.Println("♻️ Attempting to restart Docker...")
            _ = exec.Command("systemctl", "restart", "docker").Run()
            time.Sleep(5 * time.Second)
        }
    } else {
        fmt.Println("✅ Docker daemon is running.")
        if viper.GetBool("optimize.auto_prune") {
            fmt.Println("🧹 Running docker system prune -af")
            _ = exec.Command("docker", "system", "prune", "-af").Run()
        }
    }
}

func StartContinuousMonitor() {
    interval := viper.GetInt("monitor.interval_seconds")
    if interval <= 0 {
        interval = 30
    }
    for {
        CheckDockerHealth()
        time.Sleep(time.Duration(interval) * time.Second)
    }
}
