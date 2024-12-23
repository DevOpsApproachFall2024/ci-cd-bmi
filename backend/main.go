package main

import (
    "net/http"

    "github.com/DevOpsApproachFall2024/ci-cd-bmi/handlers"
)

func main() {
    http.HandleFunc("/calculate-bmi", handlers.CalculateBMI)
    http.ListenAndServe(":8080", nil)
}
