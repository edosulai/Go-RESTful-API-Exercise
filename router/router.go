package router

import (
	"gorest/controllers"
	"gorest/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gorest/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("user")
	{
		//Register
		userRouter.POST("/register", controllers.Register)
		//Login
		userRouter.POST("/login", controllers.Login)
	}

	socialmediaRouter := r.Group("socialmedia")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		//GetOneSocialMedia
		socialmediaRouter.GET("/:socialmediaID", controllers.GetOneSocialMedia)
		//GetAllComment
		socialmediaRouter.GET("/", controllers.GetAllSocialMedia)
		//CreateSocialMedia
		socialmediaRouter.POST("/", controllers.CreateSocialMedia)
		//UpdateSocialMedia
		socialmediaRouter.PUT("/:socialmediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		//DeleteSocialMedia
		socialmediaRouter.DELETE("/:socialmediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("photo")
	{
		photoRouter.Use(middlewares.Authentication())
		//GetOnePhoto
		photoRouter.GET("/:photoID", controllers.GetOnePhoto)
		//GetAllPhoto
		photoRouter.GET("/", controllers.GetAllPhoto)
		//CreatePhoto
		photoRouter.POST("/", controllers.CreatePhoto)
		//UpdatePhoto
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		//DeletePhoto
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("comment")
	{
		commentRouter.Use(middlewares.Authentication())
		//GetOneComment
		commentRouter.GET("/:commentID", controllers.GetOneComment)
		//GetAllComment
		commentRouter.GET("/", controllers.GetAllComment)
		//CreateComment
		commentRouter.POST("/", controllers.CreateComment)
		//UpdateComment
		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		//DeleteComment
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}
