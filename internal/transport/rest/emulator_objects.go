package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) getListEmulatorObjects(c *gin.Context) {

	objectString, err := h.service.EmulatorObjects.GetList(c)
	if err != nil {
		responseError(c, http.StatusInternalServerError, err)
		return
	}

	objects := strings.Split(objectString, "\n")

	var obj []string
	for _, o := range objects {
		if o != "" {
			obj = append(obj, o)
		}
	}

	type response struct {
		Objects []string `json:"objects"`
	}

	var res = response{
		Objects: obj,
	}

	response200(c, res)
}
