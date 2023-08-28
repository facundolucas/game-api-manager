package main

import (
	"game-api-manager/controllers"
	"game-api-manager/middlewares"
	"game-api-manager/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {

	models.ConnectDataBase()

	r := gin.Default()
	r.Use(cors.New(CORSConfig()))

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/user", controllers.CurrentUser)

	protected.POST("/moneda/save", controllers.SaveMoneda)
	protected.GET("/moneda/:tipo", controllers.GetMonedasByTipo)
	protected.POST("/moneda/:id/upload", controllers.UploadImagenMoneda)
	protected.POST("/moneda/:id/models/upload", controllers.UploadImagenMonedaModels)

	r.Run(":8080")

}
