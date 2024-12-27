package survey_router

import (
	mysql_configs "github.com/celpung/gocleanarch/configs/database/mysql"
	middlewares "github.com/celpung/gocleanarch/configs/middlewares/gin"
	"github.com/celpung/gocleanarch/configs/role"
	survey_delivery_implementation "github.com/celpung/gocleanarch/domain/survey/delivery/implementation"
	survey_repository_implementation "github.com/celpung/gocleanarch/domain/survey/repository/implementation"
	survey_usecase_implementation "github.com/celpung/gocleanarch/domain/survey/usecase/implementation"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	repository := survey_repository_implementation.NewSurveyRepositry(mysql_configs.DB)
	usecase := survey_usecase_implementation.NewSurveyUsecase(repository)
	delivery := survey_delivery_implementation.NewSurveyDelivery(usecase)

	routes := r.Group("/surveys")
	{
		routes.POST("", middlewares.UserMiddleware(role.User), delivery.Create)
		routes.GET("", middlewares.UserMiddleware(role.User), delivery.Read)
		routes.GET("/:id", middlewares.UserMiddleware(role.User), delivery.ReadByID)
		routes.PUT("", middlewares.UserMiddleware(role.User), delivery.Update)
		routes.DELETE("/:id", middlewares.UserMiddleware(role.User), delivery.Delete)
	}
}
