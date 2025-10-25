package optimizer

import (
    "fmt"
    "os/exec"
)

func RunPrune() {
    fmt.Println("🧹 Running prune: docker system prune -af")
    cmd := exec.Command("docker", "system", "prune", "-af")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Prune failed:", err)
    }
    fmt.Println(string(out))
}

func RunPruneSummary() {
    fmt.Println("🧹 Prune summary (docker system df)")
    cmd := exec.Command("docker", "system", "df")
    out, _ := cmd.CombinedOutput()
    fmt.Println(string(out))
}
