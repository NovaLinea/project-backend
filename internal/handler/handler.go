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
		auth := api.Group("/auth")
		{
			auth.POST("/signup", h.Register)
			auth.POST("/signin", h.Login)
			auth.GET("/refresh", h.Refresh)
			auth.GET("/logout", h.Logout)
		}

		user := api.Group("/user")
		{
			user.GET("/:userID/fetch-data-profile", h.GetDataProfile)
			user.GET("/:userID/fetch-data-settings", h.GetDataSettings)
			user.POST("/:userID/save-data", h.SaveDataUser)
			user.POST("/:userID/change-password", h.ChangePassword)
			user.GET("/:userID/delete-account", h.DeleteAccount)
			user.GET("/:userID/get-likes-favorites", h.GetLikesFavorites)
			user.GET("/:userID/get-data-params", h.GetDataParams)
			user.GET("/:userID/subscribe/:accountID", h.SubscribeUser)
			user.GET("/:userID/unsubscribe/:accountID", h.UnSubscribeUser)
		}

		project := api.Group("/project")
		{
			project.POST("/:ID/create", h.CreateProject)
			project.GET("/:ID/get-projects-popular", h.GetProjectsPopular)
			project.GET("/:ID/get-projects-user", h.GetProjectsUser)
			project.GET("/:ID/get-favorites-projects", h.GetFavoritesProjects)
			project.GET("/:ID/like/:userID", h.LikeProject)
			project.GET("/:ID/dislike/:userID", h.DislikeProject)
			project.GET("/:ID/favorite/:userID", h.FavoriteProject)
			project.GET("/:ID/remove-favorite/:userID", h.RemoveFavoriteProject)
		}
	}

	return router
}
