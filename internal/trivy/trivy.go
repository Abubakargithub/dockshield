package trivy

import (
    "fmt"
    "os/exec"
    "strings"
)

func ScanAllLocalImages() {
    fmt.Println("🛡️ Scanning local images with Trivy...")
    out, err := exec.Command("docker", "images", "--format", "{{.Repository}}:{{.Tag}}").CombinedOutput()
    if err != nil {
        fmt.Println("Failed to list images:", err)
        return
    }
    images := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, img := range images {
        if img == "" || strings.HasPrefix(img, "<none>") {
            continue
        }
        fmt.Println("Scanning:", img)
        cmd := exec.Command("trivy", "image", "--severity", "CRITICAL,HIGH", img)
        data, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Println("Trivy scan error for", img, ":", err)
        }
        fmt.Println(string(data))
    }
}
