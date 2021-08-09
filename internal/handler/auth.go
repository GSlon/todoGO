package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/GSlon/todoGO/internal/entity"
    "net/http"
    "fmt"
)


func (h *Handler) signin(c *gin.Context) {
    var input entity.SignInUser

    if err := c.BindJSON(&input); err != nil {
        logrus.Error("invalid data")
        errorResponse(c, http.StatusBadRequest, err.Error())    // invalid request data
        return
    }

    id, err := h.services.Authorization.GetUser(input)
    if err != nil {
        errorResponse(c, http.StatusBadRequest, err.Error()) // user не найден
        return
    }

    logrus.Info(fmt.Sprintf("Redirect to %d", id))

    c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/api/lists/%d", id))
}

// валидация и передача в service(бизнес-логику)
func (h *Handler) signup(c *gin.Context) {
    var user entity.User

    if err := c.BindJSON(&user); err != nil {
        logrus.Error("invalid data")
        errorResponse(c, http.StatusBadRequest, err.Error())    // invalid request data
        return
    }

    id, err := h.services.Authorization.CreateUser(user)
    if err != nil {
        errorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{
        "id": id,
    })
}
