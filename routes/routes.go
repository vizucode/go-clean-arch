package routes

import (
	"echo-boilerplate/config"

	userhandler "echo-boilerplate/domains/users/handlers"
	userModel "echo-boilerplate/domains/users/models"
	userrepo "echo-boilerplate/domains/users/repositories"
	userusecase "echo-boilerplate/domains/users/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	userRepo := userrepo.New([]userModel.User{})
	userUsecase := userusecase.New(userRepo)
	userHandler := userhandler.New(userUsecase)

	e.GET("user/:id", userHandler.GetById)
	e.POST("users", userHandler.Store)
	e.PUT("user/:id", userHandler.Update)
	e.DELETE("user/:id", userHandler.Delete)
}
