package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) getBaseStationsList(c *gin.Context) {

	objectsString, err := h.service.BaseStation.GetList(c)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	baseStations := strings.Split(objectsString, "\n")

	var bs []string
	for _, o := range baseStations {
		if o != "" {
			bs = append(bs, o)
		}
	}

	type response struct {
		BaseStations []string `json:"objects"`
	}

	var res = response{
		BaseStations: bs,
	}

	response200(c, res)
}

func (h *Handler) getBaseStationsStatus(c *gin.Context) {

	var baseStation string
	if c.Query("base_station") != "" {
		baseStation = c.Query("base_station")
	}

	status, err := h.service.BaseStation.GetStatus(c, baseStation)
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

func (h *Handler) getBaeStationsCountConnections(c *gin.Context) {

	var baseStation string
	if c.Query("base_station") != "" {
		baseStation = c.Query("base_station")
	}

	count, err := h.service.BaseStation.GetCountConnections(c, baseStation)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	type response struct {
		CountConnections string `json:"count_connections"`
	}

	var res = response{
		CountConnections: count,
	}

	response200(c, res)
}

func (h *Handler) getBaseStationInfo(c *gin.Context) {

	var baseStation string
	if c.Query("base_station") != "" {
		baseStation = c.Query("base_station")
	}

	info, err := h.service.BaseStation.GetInfo(c, baseStation)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	type response struct {
		Info string `json:"info"`
	}

	var res = response{
		Info: info,
	}

	response200(c, res)
}
