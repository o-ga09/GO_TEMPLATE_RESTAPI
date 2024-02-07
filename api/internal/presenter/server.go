package presenter

import (
	"context"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/o-ga09/api/internal/controller/system"
	userHandler "github.com/o-ga09/api/internal/controller/user"
	adminDomain "github.com/o-ga09/api/internal/domain/administrator"
	userDomain "github.com/o-ga09/api/internal/domain/user"
	"github.com/o-ga09/api/internal/driver/mysql"
	"github.com/o-ga09/api/internal/driver/mysql/repository"
	"github.com/o-ga09/api/internal/middleware"
	"github.com/o-ga09/api/internal/usecase"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()

	// ロガーを設定
	logger := middleware.New()
	httpLogger := middleware.RequestLogger(logger)

	// CORS設定関数
	cors := middleware.CORS()

	// ginにCORSを設定する
	r.Use(cors)

	// ginを使用してリクエスト情報を取得する
	r.Use(httpLogger)

	// request idを付与する
	r.Use(requestid.New())

	v1 := r.Group(latest)
	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	// dependecy injection
	conn := mysql.New(ctx)
	UserDriver := repository.NewUserDriver(conn)
	AdminDriver := repository.NewAdminDriver(conn)
	UserDomainService := userDomain.NewUserDomainService(UserDriver)
	AdminDomainService := adminDomain.NewAdminDomainService(AdminDriver)
	finduser_usecase := usecase.NewFindUserUsecase(UserDomainService, AdminDomainService)
	findByIduser_usecase := usecase.NewFindUserByIdUsecase(UserDomainService, AdminDomainService)
	saveuser_usecase := usecase.NewSaveUserUsecase(UserDomainService, AdminDomainService)
	deleteuser_usecase := usecase.NewDeleteUserUsecase(UserDomainService, AdminDomainService)
	handler := userHandler.NewUserHandler(*finduser_usecase, *findByIduser_usecase, *saveuser_usecase, *deleteuser_usecase)

	// ユーザー管理機能
	users := v1.Group("/users")
	{
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUserById)
		users.POST("", handler.EditUser)
		users.DELETE("/:id", handler.DeleteUser)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
