// Client command
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    resp, err := http.Get("http://localhost:8080/api/status")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Status: %v\n", result)
}
