package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) getMobileTerminalsList(c *gin.Context) {

	objectsString, err := h.service.MobileTerminal.GetList(c)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	mobileTerminals := strings.Split(objectsString, "\n")

	var bs []string
	for _, o := range mobileTerminals {
		if o != "" {
			bs = append(bs, o)
		}
	}

	type response struct {
		MobileTerminals []string `json:"mobile_terminals"`
	}

	var res = response{
		MobileTerminals: mobileTerminals,
	}

	response200(c, res)
}

func (h *Handler) getMobileTerminalStatus(c *gin.Context) {

	var mobileTerminal string
	if c.Query("base_station") != "" {
		mobileTerminal = c.Query("base_station")
	}

	status, err := h.service.MobileTerminal.GetStatus(c, mobileTerminal)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	type response struct {
		Status string `json:"status"`
	}

	var res = response{
		Status: status,
	}

	response200(c, res)
}

func (h *Handler) getMobileTerminalConnectionStatus(c *gin.Context) {

	var mobileTerminal string
	if c.Query("base_station") != "" {
		mobileTerminal = c.Query("base_station")
	}

	status, err := h.service.MobileTerminal.GetNetworkConnectionStatus(c, mobileTerminal)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	type response struct {
		ConnectionStatus string `json:"connection_status"`
	}

	var res = response{
		ConnectionStatus: status,
	}

	response200(c, res)
}
