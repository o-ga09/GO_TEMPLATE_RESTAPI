package user

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	userDomain "github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/internal/usecase"
	"github.com/o-ga09/api/pkg"
)

type UserHandler struct {
	FindUserUsecase usecase.FindUserUsecase
	FindByIdUsecase usecase.FindUserByIdUsecase
	SaveUserUsecase usecase.SaveUserUsecase
	DeleteUsecase   usecase.DeleteUserUsecase
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
	userId := ctx.Query("userid")
	request_id := ctx.GetHeader("X-Request-ID")

	reqCtx := context.WithValue(ctx, "ctxInfo", pkg.CtxInfo{PageLimit: limit, PageOffset: offset, UserId: userId, RequestId: request_id})

	res, err := h.FindUserUsecase.Run(reqCtx)
	if err != nil {
		slog.Error("can not process FindUser Usecase", "error msg", err, "request id", ctx.GetHeader("X-Request-ID"))
		ctx.JSON(500, gin.H{"status": "Error"})
		return
	}

	response := []ResponseUser{}
	for _, r := range res.User {
		user := &ResponseUser{User: ResponseUserModel{
			ID:               r.ID,
			Email:            r.Email,
			Password:         r.Password,
			User_ID:          r.User_ID,
			FirstName:        r.FirstName,
			LastName:         r.LastName,
			Gender:           r.Gender,
			BirthDay:         r.BirthDay,
			PhoneNumber:      r.PhoneNumber,
			PostOfficeNumber: r.PostOfficeNumber,
			Address:          r.Pref + r.City + r.Extra,
		}}
		response = append(response, *user)
	}

	slog.Error("process done FindUser Usecase", "request id", ctx.GetHeader("X-Request-ID"))
	ctx.JSON(200, response)
	ctx.Next()
}

// GetUserById godoc
// @Summary ユーザーの詳細情報を取得
// @Tags GetUserById
// @Accept json
// @Produce json
// @Param request path string ture "ユーザーID"
// @Success 200 {object} ResponseUser
// @Router /v1/users/:id [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	request_id := ctx.GetHeader("X-Request-ID")

	var reqCtx context.Context
	reqCtx = context.WithValue(ctx, "ctxInfo", pkg.CtxInfo{RequestId: request_id})
	res, err := h.FindByIdUsecase.Run(reqCtx, userId)
	if err != nil {
		slog.Error("can not process FindByID Usecase", "error msg", err, "request id", ctx.GetHeader("X-Request-ID"))
		ctx.JSON(http.StatusInternalServerError, Response{Status: "Internal Server Error"})
		return
	}

	user := &ResponseUser{User: ResponseUserModel{
		ID:               res.GetUUID(),
		Email:            res.GetEmail(),
		Password:         res.GetPassWord(),
		User_ID:          res.GetID(),
		FirstName:        res.GetFirstName(),
		LastName:         res.GetLastName(),
		Gender:           res.GetGender(),
		BirthDay:         res.GetBirthDay(),
		PhoneNumber:      res.GetPhoneNumber(),
		PostOfficeNumber: res.GetPostOfficeNumber(),
		Address:          res.GetPref() + res.GetCity() + res.GetExtra(),
	}}
	slog.Info("process done FindByID Usecase", "request id", ctx.GetHeader("X-Request-ID"))
	ctx.JSON(http.StatusOK, user)
}

// EditUser godoc
// @Summary ユーザー情報を編集
// @Tags EditUser
// @Accept json
// @Produce json
// @Param request body RequestUserParam ture "ユーザー情報"
// @Success 200 {object} Response
// @Router /v1/users [post]
func (h *UserHandler) EditUser(ctx *gin.Context) {
	request_id := ctx.GetHeader("X-Request-ID")
	param := &RequestUserParam{}

	var reqCtx context.Context
	reqCtx = context.WithValue(ctx, "ctxInfo", pkg.CtxInfo{RequestId: request_id})
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		slog.Error("can not process SaveUser Usecase", "error msg", err, "request id", ctx.GetHeader("X-Request-ID"))
		ctx.JSON(http.StatusBadRequest, Response{Status: "Bad Request"})
		return
	}

	user := userDomain.NewUser(param.ID, param.Email, param.Password, param.User_ID, param.FirstName, param.LastName, param.Gender, param.BirthDay, param.PhoneNumber, param.PostOfficeNumber, param.Pref, param.City, param.Extra)
	err = h.SaveUserUsecase.Run(reqCtx, user)
	if err != nil {
		slog.Error("can not process SaveUser Usecase", "error msg", err, "request id", ctx.GetHeader("X-Request-ID"))
		ctx.JSON(http.StatusBadRequest, Response{Status: "Internal Server Error"})
		return
	}

	slog.Info("process done SaveUser Usecase", "request id", ctx.GetHeader("X-Request-ID"))
	ctx.JSON(http.StatusOK, Response{Status: "process done"})
	return
}

// DeleteUser godoc
// @Summary ユーザー情報を削除
// @Tags DeleteUser
// @Accept json
// @Produce json
// @Param request path string ture "ユーザーID"
// @Success 200 {object} Response
// @Router /v1/users [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	request_id := ctx.GetHeader("X-Request-ID")
	userId := ctx.Param("id")
	var reqCtx context.Context
	reqCtx = context.WithValue(ctx, "ctxInfo", pkg.CtxInfo{RequestId: request_id})
	err := h.DeleteUsecase.Run(reqCtx, userId)
	if err != nil {
		slog.Error("can not process DeleteUser Usecase", "error msg", err, "request id", ctx.GetHeader("X-Request-ID"))
		ctx.JSON(http.StatusInternalServerError, Response{Status: "Internal Server Error"})
		return
	}

	slog.Info("process done DeleteUser Usecase", "request id", ctx.GetHeader("X-Request-ID"))
	ctx.JSON(http.StatusOK, Response{Status: "process complete"})
	return
}

func NewUserHandler(
	findUserusecase usecase.FindUserUsecase,
	findByIdUsecase usecase.FindUserByIdUsecase,
	saveUserusecase usecase.SaveUserUsecase,
	deleteUserusecase usecase.DeleteUserUsecase,
) *UserHandler {
	return &UserHandler{
		FindUserUsecase: findUserusecase,
		FindByIdUsecase: findByIdUsecase,
		SaveUserUsecase: saveUserusecase,
		DeleteUsecase:   deleteUserusecase,
	}
}
