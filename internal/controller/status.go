package controller

import (
	"context"

	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/pkg/api"
	"github.com/labstack/echo/v4"
)

type HealthCheck struct {
}

func NewHealthCheckController() internal.HealthCheckController {
	return &HealthCheck{}
}

// StatusHealthCheck godoc
// @Summary health check for system
// @Description check system status, should return status ok,
// @Tags Health Check
// @Accept json
// @Produce json
//
// @Success 200 {object} api.HTTP{data=api.HTTPData}
// @Failure 500 {object} api.HTTP{data=api.HTTPData}
//
// @Router /status [get]
func (h *HealthCheck) StatusHealthCheck(c echo.Context) error {
	_, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	res := api.DefaultResponse(c)
	return res.ReponseOK(nil, constant.ServerStatusOK)
}
