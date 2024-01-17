package user

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/o-ga09/api/internal/usecase"
)

type UserHandler struct {
	usecase usecase.FindUserUsecase
}

// GetUsers godoc
// @Summary ユーザー一覧を取得
// @Tags GetUsers
// @Accept json
// @Produce json
// @Success 200 {object} []ResponseUser
// @Router /v1/users [get]
func (h *UserHandler) GetUsers(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")

	res, err := h.usecase.Run(ctx)
	if err != nil {
		slog.Error("debug", "error", err)
		ctx.JSON(500, gin.H{"status": "Error"})
		return
	}

	slog.Info("debug", "limit", limit, "offset", offset)
	slog.Info("debug", "response", res)
	ctx.JSON(200, gin.H{"status": "OK"})
}

// GetUserById godoc
// @Summary ユーザーの詳細情報を取得
// @Tags GetUserById
// @Accept json
// @Produce json
// @Param request path string ture "ユーザーID"
// @Success 200 {object} ResponseUser
// @Router /v1/users/:id [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {}

// EditUser godoc
// @Summary ユーザー情報を編集
// @Tags EditUser
// @Accept json
// @Produce json
// @Param request body RequestUserParam ture "ユーザー情報"
// @Success 200 {object} Response
// @Router /v1/users [post]
func (h *UserHandler) EditUser(ctx *gin.Context) {}

// DeleteUser godoc
// @Summary ユーザー情報を削除
// @Tags DeleteUser
// @Accept json
// @Produce json
// @Param request path string ture "ユーザーID"
// @Success 200 {object} Response
// @Router /v1/users [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {}

func NewUserHandler(usecase usecase.FindUserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
