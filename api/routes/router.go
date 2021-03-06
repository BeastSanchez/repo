/**
 * Created by VoidArtanis on 10/22/2017
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/VoidArtanis/go-reactjs-boilerplate/api/controllers"
	"github.com/VoidArtanis/go-reactjs-boilerplate/api/middlewares"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)
	authController := new(controllers.AuthController)
	userController := new(controllers.UserController)
	engine.POST("/login",  authController.HandleLogin)
	engine.POST("/signup",  userController.UserSignUp)
	RegisterProtectedRoutes(engine)
	RegisterPublicRoutes(engine)
	RegisterUtilityRoutes(engine)
}

func InitMiddleware(engine *gin.Engine){
	engine.Use(middlewares.CORSMiddleware());
}