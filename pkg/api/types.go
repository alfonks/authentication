package api

import (
	"time"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	EchoContext echo.Context `json:"-"`
	StartTime   time.Time    `json:"-"`
	Status      int64        `json:"status"`
	StatusText  string       `json:"status_text"`
	ProcessTime string       `json:"process_time"`
	Data        HTTPData     `json:"data"`
}

type HTTPData struct {
	Request interface{} `json:"request"`
	Reponse interface{} `json:"response"`
	Message string      `json:"message"`
}
