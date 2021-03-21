package controller

// HealthCheckController :
const (
	HealthCheckController = `
	package controllers

import (
	"encoding/json"
	"net/http"
)

// HealthCheckController :
type HealthCheckController struct {
}

// NewHealthCheckController : creates a controller instance
// add parameters to inject controller dependencies e.g
// casandra client, kafka client, etc
func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// HealthCheckerHandler :
func (ctrl *HealthCheckController) HealthCheckerHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(NewAPIResponse(r.Context(), StatusOK, "alive"))
	return
}`
)
