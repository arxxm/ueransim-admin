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

		auth := api.Group("/emulator-objects")
		{
			auth.GET("/", h.getListEmulatorObjects)
			//auth.POST("/sign-up", h.signUp)
			//auth.POST("/sign-in", h.signIn)
		}

	}

	return router
}
