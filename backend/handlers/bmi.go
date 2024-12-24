package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
)

type BMIRequest struct {
    Height float64 `json:"height"`
    Weight float64 `json:"weight"`
}

type BMIResponse struct {
    BMI float64 `json:"bmi"`
    Status string  `json:"status"`
}

func CalculateBMI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")


    if r.Method != http.MethodGet {
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract height and weight from query parameters
    heightStr := r.URL.Query().Get("height")
    weightStr := r.URL.Query().Get("weight")

    // Check if both height and weight are provided
    if heightStr == "" || weightStr == "" {
        http.Error(w, "Height and Weight parameters are required", http.StatusBadRequest)
        return
    }

    // Convert height and weight to float64
    height, err := strconv.ParseFloat(heightStr, 64)
    if err != nil || height <= 0 {
        http.Error(w, "Invalid or non-positive height", http.StatusBadRequest)
        return
    }

    weight, err := strconv.ParseFloat(weightStr, 64)
    if err != nil || weight <= 0 {
        http.Error(w, "Invalid or non-positive weight", http.StatusBadRequest)
        return
    }

    // Calculate BMI
    bmi := weight / (height * height)
    status := ""
    switch {
    case bmi < 18.5:
        status = "Underweight"
    case bmi < 24.9:
        status = "Normal weight"
    case bmi < 29.9:
        status = "Overweight"
    default:
        status = "Obese"
    }

    // Respond with the BMI and status
    res := BMIResponse{
        BMI:    bmi,
        Status: status,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
