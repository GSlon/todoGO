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

    api := router.Group("/api")
    {
        list := api.Group("/list")
        {
            list.GET("/:id", h.getListById)    // id юзера

            item := list.Group("/item")
            {
                item.POST("/:id", h.createItem)   // id item'a
                item.GET("/:id", h.getItemById)
                item.DELETE("/:id", h.deleteItem)
                item.PUT("/:id", h.updateItem)
            }
        }
    }

    return router
}
