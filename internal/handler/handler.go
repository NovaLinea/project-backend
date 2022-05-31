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
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
			auth.GET("/get-me", h.userIdentity)

			userAuth := auth.Group("/user")
			{
				userAuth.GET("/:userID/get-settings", h.GetDataSettings)
				userAuth.POST("/:userID/save", h.SaveDataUser)
				userAuth.POST("/:userID/change-password", h.ChangePassword)
				userAuth.GET("/:userID/delete", h.DeleteAccount)
				userAuth.GET("/:userID/get-likes-favorites", h.GetLikesFavorites)
				userAuth.GET("/:userID/get-followings", h.GetListFollowings)
				userAuth.GET("/:userID/check-subscribe/:accountID", h.CheckSubscribe)
				userAuth.GET("/:userID/subscribe/:accountID", h.SubscribeUser)
				userAuth.GET("/:userID/unsubscribe/:accountID", h.UnSubscribeUser)
			}

			projectAuth := auth.Group("/project")
			{
				projectAuth.POST("/:ID/create", h.Create)
				projectAuth.GET("/:ID/get-projects-user", h.GetProjectsUser)
				projectAuth.GET("/:ID/get-home", h.GetHome)
				projectAuth.GET("/:ID/get-favorites", h.GetFavorites)
				projectAuth.GET("/:ID/like/:userID", h.Like)
				projectAuth.GET("/:ID/dislike/:userID", h.Dislike)
				projectAuth.GET("/:ID/favorite/:userID", h.Favorite)
				projectAuth.GET("/:ID/remove-favorite/:userID", h.RemoveFavorite)
				projectAuth.GET("/:ID/delete", h.DeleteProject)
				projectAuth.POST("/:ID/save", h.Update)
			}
		}

		user := api.Group("/user")
		{
			user.GET("/:userID/get-data", h.GetDataProfile)
			user.GET("/:userID/get-params", h.GetDataParams)
			user.GET("/:userID/get-follows", h.GetFollows)
			user.GET("/:userID/get-followings", h.GetFollowings)
		}

		project := api.Group("/project")
		{
			project.GET("/:ID/get-popular", h.GetPopular)
			project.GET("/:ID/get-data", h.GetDataProject)
		}
	}

	return router
}
