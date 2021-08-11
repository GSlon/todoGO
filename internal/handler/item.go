package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/GSlon/todoGO/internal/entity"
    _ "github.com/sirupsen/logrus"
    "strconv"
    "net/http"
)

func parseContextParam(param string, c *gin.Context) (int, error) {
    id, err := strconv.Atoi(c.Param(param))
    if err != nil {
        errorResponse(c, http.StatusBadRequest, "invalid item id")
        return 0, err
    }

    return id, nil
}

func parseContextJSON(c *gin.Context) (entity.TodoItem, error) {
    var input entity.TodoItem
    if err := c.BindJSON(&input); err != nil {
        errorResponse(c, http.StatusBadRequest, err.Error())
        return input, err
    }

    return input, nil
}

func parseInputContextJSON(c *gin.Context) (entity.UpdateTodoItem, error) {
    var input entity.UpdateTodoItem
    if err := c.BindJSON(&input); err != nil {
        errorResponse(c, http.StatusBadRequest, err.Error())
        return input, err
    }

    return input, nil
}

func (h *Handler) createItem(c *gin.Context) {
    userId, err := parseContextParam("id", c)
    if err != nil {
        return
    }

    input, err :=  parseContextJSON(c)
    if err != nil {
        return
    }

    id, err := h.services.TodoItem.Create(userId, input)
    if err != nil {
        errorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{
        "id": id,
    })
}

func (h *Handler) deleteItem(c *gin.Context) {
    itemId, err := parseContextParam("itemid", c)
    if err != nil {
        return
    }

	err = h.services.TodoItem.Delete(itemId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) updateItem(c *gin.Context) {
    itemId, err := parseContextParam("itemid", c)
    if err != nil {
        return
    }

    input, err :=  parseInputContextJSON(c)
    if err != nil {
        return
    }

	err = h.services.TodoItem.Update(itemId, input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getItemById(c *gin.Context) {
    itemId, err := parseContextParam("itemid", c)
    if err != nil {
        return
    }

    item, err := h.services.TodoItem.GetItemById(itemId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) getAllItems(c *gin.Context) {
    userId, err := parseContextParam("id", c)
    if err != nil {
        return
    }

    items, err := h.services.TodoItem.GetAllItems(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
