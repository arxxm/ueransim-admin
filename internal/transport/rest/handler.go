package rest

import (
	"github.com/gin-gonic/gin"
	"ueransim-api/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.ErrorLogger())
	gin.SetMode(gin.DebugMode)

	api := router.Group("/api/v1")

	{
		api.GET("/health", func(context *gin.Context) {
			response200(context, nil)
		})

		baseStations := api.Group("/base-stations")
		{
			baseStations.GET("/", h.getBaseStationsList)
			baseStations.GET("/status", h.getBaseStationsStatus)
			baseStations.GET("/info", h.getBaseStationInfo)
			baseStations.GET("/cc", h.getBaeStationsCountConnections)
		}

		mobileTerminals := api.Group("/mobile-terminals")
		{
			mobileTerminals.GET("/", h.getMobileTerminalsList)
			mobileTerminals.GET("/status", h.getMobileTerminalStatus)
			mobileTerminals.GET("/cs", h.getMobileTerminalConnectionStatus)
		}

	}

	return router
}
