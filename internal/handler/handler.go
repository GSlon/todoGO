package handler

import (
        "github.com/gin-gonic/gin"
        "github.com/GSlon/todoGO/internal/service"
)

type Handler struct {
    services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
    router := gin.New()

    auth := router.Group("/auth")
    {
        auth.POST("/signin", h.signin)
        auth.POST("/signup", h.signup)
    }

    list := router.Group("/list")
    {
        list.GET("/:id", h.getAllItems)       // id юзера
        list.POST("/:id", h.createItem)
    }

    item := router.Group("/item")
    {
        item.GET("/:itemid", h.getItemById)
        item.DELETE("/:itemid", h.deleteItem)
        item.PUT("/:itemid", h.updateItem)
    }


    return router
}
