package metrics

import (
    "fmt"
    "net/http"
)

func StartServer() {
    addr := ":9091"
    fmt.Println("📈 Starting metrics server on", addr)
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "# HELP dockshield_status DockShield status\n# TYPE dockshield_status gauge\ndockshield_status 1")
    })
    _ = http.ListenAndServe(addr, nil)
}
