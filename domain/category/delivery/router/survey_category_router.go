package survey_category_router

import (
	mysql_configs "github.com/celpung/gocleanarch/configs/database/mysql"
	middlewares "github.com/celpung/gocleanarch/configs/middlewares/gin"
	"github.com/celpung/gocleanarch/configs/role"
	survey_category_delivery_implementation "github.com/celpung/gocleanarch/domain/category/delivery/implementation"
	survey_category_repository_implementation "github.com/celpung/gocleanarch/domain/category/repository/implementation"
	survey_category_usecase_implementation "github.com/celpung/gocleanarch/domain/category/usecase/implementation"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	repository := survey_category_repository_implementation.NewSurveyCategoryRepository(mysql_configs.DB)
	usecase := survey_category_usecase_implementation.NewSurveyCategoryUsecase(repository)
	delivery := survey_category_delivery_implementation.NewSurveyCategoryDelivery(usecase)

	routes := r.Group("/categories")
	{
		routes.POST("", middlewares.UserMiddleware(role.User), delivery.Create)
		routes.GET("", middlewares.UserMiddleware(role.User), delivery.Read)
		routes.GET("/:id", middlewares.UserMiddleware(role.User), delivery.ReadById)
		routes.PUT("", middlewares.UserMiddleware(role.User), delivery.Update)
		routes.DELETE("/:id", middlewares.UserMiddleware(role.User), delivery.Delete)
	}
}
