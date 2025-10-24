# 🛡️ DockShield — Docker Guardian

DockShield is a Go-based open-source daemon that monitors, secures, and optimizes your Docker runtime.

## ✨ Features
- Docker daemon health monitoring
- Auto-restart on crash
- Vulnerability scanning (Trivy)
- Optimization & pruning
- Prometheus metrics
- Systemd integration

## 🚀 Installation
### Go
```
go install github.com/Abubakargithub/dockshield@latest
```
### Binary
```
curl -L https://github.com/Abubakargithub/dockshield/releases/latest/download/dockshield-linux-amd64.tar.gz | tar xz
sudo mv dockshield /usr/local/bin/
```
### Docker
```
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock yourorg/dockshield check
```
