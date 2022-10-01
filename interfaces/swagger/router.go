package swagger

import (
	_ "desafio/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router(r *gin.RouterGroup) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
