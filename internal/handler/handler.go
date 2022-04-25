package handler

import (
	"os"

	"github.com/ProjectUnion/project-backend.git/internal/service"
	"github.com/ProjectUnion/project-backend.git/pkg/logging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Handler struct {
	services *service.Service
	logger   logging.Logger
}

func NewHandler(services *service.Service) *Handler {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers,authorization,my-custom-header"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api := router.Group("/api")
	{
		api.POST("/signup", h.Register)
		api.POST("/signin", h.Login)
		api.GET("/refresh", h.Refresh)
		api.GET("/logout", h.Logout)
	}

	return router
}
