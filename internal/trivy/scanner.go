package trivy

import (
	"fmt"
	"os/exec"
)

func Scan() {
	fmt.Println("🛡️ Starting Trivy vulnerability scan...")
	cmd := exec.Command("trivy", "image", "--severity", "CRITICAL,HIGH", "alpine:latest")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("⚠️ Trivy scan failed:", err)
	}
	fmt.Println(string(out))
}
