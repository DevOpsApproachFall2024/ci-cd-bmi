package handlers

import (
    "encoding/json"
    "net/http"
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
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is  allowed", http.StatusMethodNotAllowed)
        return
    }

    var req BMIRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON body", http.StatusBadRequest)
        return
    }

    if req.Height <= 0 || req.Weight <= 0 {
        http.Error(w, "Height and Weight must be positive numbers", http.StatusBadRequest)
        return
    }

    bmi := req.Weight / (req.Height * req.Height)
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

    res := BMIResponse{
        BMI:    bmi,
        Status: status,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
