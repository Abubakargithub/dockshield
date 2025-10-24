package metrics

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fmt.Println("📈 Starting metrics server on :9091")
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "# HELP dockshield_status DockShield status\n# TYPE dockshield_status gauge\ndockshield_status 1")
	})
	_ = http.ListenAndServe(":9091", nil)
}
