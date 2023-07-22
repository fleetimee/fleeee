package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fleetimee/flee/handlers"
	"github.com/fleetimee/flee/middlewares"
	"github.com/fleetimee/flee/repositories"
	"github.com/fleetimee/flee/services"
)

func NewRouteMerchant(db *gorm.DB, router *gin.Engine) {
	repository := repositories.NewRepositoryMerchant(db)
	service := services.NewServiceMerchant(repository)
	handler := handlers.NewHandlerMerchant(service)

	route := router.Group("/api/v1/merchant")

	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "merchant": true}))

	router.GET("/api/v1/merchant/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
