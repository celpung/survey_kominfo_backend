package survey_question_router

import (
	mysql_configs "github.com/celpung/gocleanarch/configs/database/mysql"
	middlewares "github.com/celpung/gocleanarch/configs/middlewares/gin"
	"github.com/celpung/gocleanarch/configs/role"
	survey_question_delivery_implementation "github.com/celpung/gocleanarch/domain/question/delivery/implementation"
	survey_question_repository_implementation "github.com/celpung/gocleanarch/domain/question/repository/implementation"
	survey_question_usecase_implementation "github.com/celpung/gocleanarch/domain/question/usecase/implementation"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	repository := survey_question_repository_implementation.NewSurveyQuestionRepository(mysql_configs.DB)
	usecase := survey_question_usecase_implementation.NewSurveyQuestionUsecase(repository)
	delivery := survey_question_delivery_implementation.NewSurveyQuestionDelivery(usecase)

	routes := r.Group("/questions")
	{
		routes.POST("", middlewares.UserMiddleware(role.User), delivery.Create)
		routes.GET("", middlewares.UserMiddleware(role.User), delivery.Read)
		routes.GET("/:id", middlewares.UserMiddleware(role.User), delivery.ReadById)
		routes.GET("/survey/:survey_id", middlewares.UserMiddleware(role.User), delivery.ReadBySurveyId)
		routes.PUT("", middlewares.UserMiddleware(role.User), delivery.Update)
		routes.DELETE("/:id", middlewares.UserMiddleware(role.User), delivery.Delete)
	}
}
