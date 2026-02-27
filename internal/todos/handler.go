package todos

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	todos := h.svc.GetTodos()
	if todos == nil {
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

}

func (h *Handler) UpdateTodoById(ctx *gin.Context) {

}

func (h *Handler) DeleteTodoById(ctx *gin.Context) {

}
