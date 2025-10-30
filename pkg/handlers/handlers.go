// HTTP Handlers for api_server
package handlers

import (
    "encoding/json"
    "net/http"
    "time"
)

type Response struct {
    Success   bool        `json:"success"`
    Message   string      `json:"message"`
    Data      interface{} `json:"data,omitempty"`
    Timestamp time.Time   `json:"timestamp"`
}

// StatusHandler returns system status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Success:   true,
        Message:   "System operational",
        Data:      map[string]interface{}{"uptime": "100s"},
        Timestamp: time.Now(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// DataHandler handles data operations
func DataHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Get data
        response := Response{
            Success:   true,
            Data:      []string{"item1", "item2"},
            Timestamp: time.Now(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
        
    case http.MethodPost:
        // Create data
        var body map[string]interface{}
        json.NewDecoder(r.Body).Decode(&body)
        
        response := Response{
            Success:   true,
            Message:   "Data created",
            Data:      body,
            Timestamp: time.Now(),
        }
        w.WriteHeader(http.StatusCreated)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
