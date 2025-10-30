// Server command
package main

import (
    "log"
    "net/http"
    "your-project/pkg/handlers"
    "your-project/pkg/middleware"
)

func main() {
    mux := http.NewServeMux()
    
    // Register routes
    mux.HandleFunc("/api/status", handlers.StatusHandler)
    mux.HandleFunc("/api/data", handlers.DataHandler)
    
    // Apply middleware
    handler := middleware.Logger(middleware.CORS(middleware.Recovery(mux)))
    
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}
