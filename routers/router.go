package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yeqianmen/go-gin-example/middleware/jwt"
	"github.com/yeqianmen/go-gin-example/pkg/setting"
	"github.com/yeqianmen/go-gin-example/routers/api"
	v1 "github.com/yeqianmen/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.PUT("/tags/:id", v1.EditTags)
		apiv1.DELETE("/tags/:id", v1.DeleteTags)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
