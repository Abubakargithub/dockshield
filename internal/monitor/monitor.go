package monitor

import (
	"fmt"
	"os/exec"
)

func CheckDockerHealth() {
	fmt.Println("🔍 Checking Docker daemon health...")

	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Docker is not running or not installed.")
	} else {
		fmt.Println("✅ Docker daemon is running.")
	}
}
