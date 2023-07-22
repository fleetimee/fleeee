package routes

import (
	"github.com/fleetimee/flee/handlers"
	"github.com/fleetimee/flee/repositories"
	"github.com/fleetimee/flee/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouteContact(db *gorm.DB, router *gin.Engine) {
	repository := repositories.NewRepositoryContact(db)
	service := services.NewServiceContact(repository)
	handler := handlers.NewHandlerContact(service)

	route := router.Group("/api/v1/contact")

	router.GET("/api/v1/contact/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)

}
