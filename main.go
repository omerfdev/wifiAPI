package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strings"
)

func getWifiNetworks() ([]string, error) {
    cmd := exec.Command("netsh", "wlan", "show", "network")
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    fmt.Println("Output:", string(output)) // Çıktıyı yazdır

    networks := make([]string, 0)
    lines := strings.Split(string(output), "\n")
    for _, line := range lines {
        if strings.Contains(line, "SSID") {
            parts := strings.Split(line, ":")
            if len(parts) == 2 {
                networks = append(networks, strings.TrimSpace(parts[1]))
            }
        }
    }

    return networks, nil
}


func getWifiHandler(w http.ResponseWriter, r *http.Request) {
    networks, err := getWifiNetworks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response, err := json.Marshal(networks)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func main() {
    http.HandleFunc("/getWifi", getWifiHandler)
    fmt.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
