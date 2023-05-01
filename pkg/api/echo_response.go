package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func DefaultResponse(echoContext echo.Context) HTTP {
	return HTTP{
		Status:      http.StatusInternalServerError,
		StatusText:  http.StatusText(http.StatusInternalServerError),
		EchoContext: echoContext,
		StartTime:   time.Now(),
		Data: HTTPData{
			Request: echoContext.Request().URL.Query(),
		},
	}
}

func (h *HTTP) ReponseOK(response interface{}, msg string) error {
	h.Status = http.StatusOK
	h.StatusText = http.StatusText(http.StatusOK)
	h.ProcessTime = fmt.Sprintf("%v ms", time.Since(h.StartTime).Milliseconds())
	h.Data.Reponse = response
	h.Data.Message = msg
	return h.EchoContext.JSON(http.StatusOK, h)
}

func (h *HTTP) ResponseInternalServerError(response interface{}, msg string) error {
	h.ProcessTime = fmt.Sprintf("%v ms", time.Since(h.StartTime).Milliseconds())
	h.Data.Message = msg
	h.Data.Reponse = response
	return h.EchoContext.JSON(http.StatusInternalServerError, h)
}

func (h *HTTP) ResponseBadRequestError(msg string) error {
	h.Status = http.StatusBadRequest
	h.StatusText = http.StatusText(http.StatusBadRequest)
	h.ProcessTime = fmt.Sprintf("%v ms", time.Since(h.StartTime).Milliseconds())
	h.Data.Message = msg
	return h.EchoContext.JSON(http.StatusBadRequest, h)
}

func (h *HTTP) ResponseUnauthorized(msg string) error {
	h.Status = http.StatusUnauthorized
	h.StatusText = http.StatusText(http.StatusUnauthorized)
	h.ProcessTime = fmt.Sprintf("%v ms", time.Since(h.StartTime).Milliseconds())
	h.Data.Message = msg
	return h.EchoContext.JSON(http.StatusUnauthorized, h)
}
