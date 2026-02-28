package todos

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ybuilds/todo-api/internal/models"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) GetTodos(ctx *gin.Context) {
	todos, err := h.svc.GetTodos()
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "no todos found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error retreiving todos from service",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"todos":   todos,
		"message": "retreived successfully",
	})
}

func (h *Handler) GetTodoById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse todo id from url",
		})
		return
	}

	todo, err := h.svc.GetTodoById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "no todo with requested id found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error retreiving todo from service",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"todo":    todo,
		"message": "retreived successfully",
	})
}

func (h *Handler) AddTodo(ctx *gin.Context) {
	var todo models.Todo

	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing request body",
		})
		return
	}

	id, err := h.svc.AddTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error inserting todo from service",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("todo with id %d created successfully", id),
	})
}

func (h *Handler) UpdateTodoById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse todo id from url",
		})
		return
	}

	var todo models.Todo

	err = ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing request body",
		})
		return
	}

	updId, err := h.svc.UpdateTodoById(id, todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error updating todo from service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("todo with id %d updated successfully", updId),
	})
}

func (h *Handler) DeleteTodoById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse todo id from url",
		})
		return
	}

	delId, err := h.svc.DeleteTodoById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "no todo with requested id found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting todo from service",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("todo with id %d deleted successfully", delId),
	})
}
