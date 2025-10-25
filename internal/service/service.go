package service

import (
    "errors"
    "os/exec"
    "runtime"
)

func IsRunning() bool {
    return false
}

func Stop() error {
    if runtime.GOOS == "windows" {
        _ = exec.Command("powershell", "-Command", "Stop-Service -Name 'dockshield' -ErrorAction SilentlyContinue").Run()
        return nil
    }
    err := exec.Command("systemctl", "stop", "dockshield").Run()
    if err != nil {
        return errors.New("failed to stop systemd service (maybe not installed)")
    }
    return nil
}

func RestartDocker() error {
    if runtime.GOOS == "windows" {
        _ = exec.Command("powershell", "-Command", "Restart-Service -Name 'com.docker.service'").Run()
        return nil
    }
    err := exec.Command("systemctl", "restart", "docker").Run()
    if err != nil {
        return err
    }
    return nil
}
