package main

import (
    "sugar-level-client/client"
    "sugar-level-client/reports"
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    var backend, err = client.NewFromUrlAddress("http://localhost:8080/results")

    if err == nil {
        reports.NewFromBackendClient(backend)
        var report, err = reports.ProduceSugarReport()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        byteArray, _ := json.MarshalIndent(&report, "", "  ")
        fmt.Println(string(byteArray))
    }
}
