package user_router

import (
	mysql_configs "github.com/celpung/gocleanarch/configs/database/mysql"
	middlewares "github.com/celpung/gocleanarch/configs/middlewares/gin"
	"github.com/celpung/gocleanarch/configs/role"
	user_delivery_implementation "github.com/celpung/gocleanarch/domain/user/delivery/gin/implementation"
	user_repository_implementation "github.com/celpung/gocleanarch/domain/user/repository/implementation"
	user_usecase_implementation "github.com/celpung/gocleanarch/domain/user/usecase/implementation"
	jwt_services "github.com/celpung/gocleanarch/services/jwt"
	password_services "github.com/celpung/gocleanarch/services/password"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	passwordService := password_services.NewPasswordService()
	jwtService := jwt_services.NewJwtService()

	repository := user_repository_implementation.NewUserRepositry(mysql_configs.DB)
	usecase := user_usecase_implementation.NewUserUsecase(repository, passwordService, jwtService)
	delivery := user_delivery_implementation.NewUserDelivery(usecase)

	routes := r.Group("/users")
	{
		routes.POST("/register", delivery.Register)
		routes.POST("/login", delivery.Login)
		routes.GET("", middlewares.UserMiddleware(role.Admin), delivery.GetAllUserData)
		routes.GET("/:id", middlewares.UserMiddleware(role.User), delivery.GetUserById)
		routes.PATCH("", middlewares.UserMiddleware(role.User), delivery.UpdateUser)
	}
}
