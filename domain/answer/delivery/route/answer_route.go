package survey_answer_router

import (
	mysql_configs "github.com/celpung/gocleanarch/configs/database/mysql"
	middlewares "github.com/celpung/gocleanarch/configs/middlewares/gin"
	"github.com/celpung/gocleanarch/configs/role"
	survey_answer_delivery_impl "github.com/celpung/gocleanarch/domain/answer/delivery/implementation"
	survey_answer_repository_impl "github.com/celpung/gocleanarch/domain/answer/repository/implementation"
	survey_answer_usecase_impl "github.com/celpung/gocleanarch/domain/answer/usecase/implementation"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	// Initialize the dependencies
	repository := survey_answer_repository_impl.NewSurveyAnswerRepository(mysql_configs.DB)
	usecase := survey_answer_usecase_impl.NewSurveyAnswerUsecase(repository)
	delivery := survey_answer_delivery_impl.NewSurveyAnswerDelivery(usecase)

	// Define the routes under the /answers path
	routes := r.Group("/answers")
	{
		routes.POST("", middlewares.UserMiddleware(role.User), delivery.Create)
		routes.GET("", middlewares.UserMiddleware(role.User), delivery.Read)
		routes.GET("/:id", middlewares.UserMiddleware(role.User), delivery.ReadById)
		routes.GET("/question/:question_id", middlewares.UserMiddleware(role.User), delivery.ReadByQuestionId)
		routes.PUT("", middlewares.UserMiddleware(role.User), delivery.Update)
		routes.DELETE("/:id", middlewares.UserMiddleware(role.User), delivery.Delete)
	}
}
