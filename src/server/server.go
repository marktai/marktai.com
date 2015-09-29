package server 

import (
   "fmt"
    "log"
    "net/http"
    "time"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "test") 
}

func Run(port uint16) {
    http.HandleFunc("/healthz", healthzHandler)
    for {
        log.Println("[healthz server] %s", http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
        time.Sleep(1 * time.Second)
    }
}
