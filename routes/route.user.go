package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fleetimee/flee/handlers"
	"github.com/fleetimee/flee/repositories"
	"github.com/fleetimee/flee/services"
)

func NewRouteUser(db *gorm.DB, router *gin.Engine) {
	repositoryUser := repositories.NewRepositoryUser(db)
	serviceUser := services.NewServiceUser(repositoryUser)
	handlerUser := handlers.NewHandlerUser(serviceUser)

	route := router.Group("/api/v1/auth")

	route.GET("/ping", handlerUser.HandlerPing)
	route.POST("/register", handlerUser.HandlerRegister)
	route.POST("/login", handlerUser.HandlerLogin)
}
