package router

import (
	"github.com/PaguhEsatrio/task-5-pbi-btpns-PaguhEsatrio/controllers"
	"github.com/PaguhEsatrio/task-5-pbi-btpns-PaguhEsatrio/database"
	middleware "github.com/PaguhEsatrio/task-5-pbi-btpns-PaguhEsatrio/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	database.ConnectDatabase()

	r.POST("/api/users/register", controllers.Register)
	r.POST("/api/users/login", controllers.Login)
	r.POST("/api/users/logout", controllers.Logout)

	authorized := r.Group("/")
	authorized.Use(middleware.CheckAuth())
	{
		authorized.GET("/api/photos", controllers.GetAllPhotos)
		authorized.POST("/api/photos", controllers.CreatePhoto)
		authorized.PUT("/api/photos/:id", controllers.UpdatePhoto)
		authorized.DELETE("/api/photos/:id", controllers.DeletePhoto)

	}

	return r

}
