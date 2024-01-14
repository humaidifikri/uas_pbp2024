package main

import (
	"uas-pbp/handlers/produkhandler"
	"uas-pbp/handlers/userhandler"
	"uas-pbp/models"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
	docs "uas-pbp/docs"
)

// @title Dokumentasi UAS Pemrograman Berbasis Platform
// @version 4.0
// @description Membuat API CRUD untuk User dan Produk 
// @host localhost:8080
// @basePath /api
func main() {
    r := gin.Default()
    models.ConnectDatabase()

    docs.SwaggerInfo.BasePath = "/api"

    v1 := r.Group("/api")
    {

        //Login API
        loginGroup := v1.Group("/login")
        {
            loginGroup.POST("/",userhandler.Login)
        }

        // Produk API
        produkGroup := loginGroup.Group("/produk")
        {
            produkGroup.GET("/", produkhandler.Index)
            produkGroup.GET("/:id", produkhandler.Show)
            produkGroup.POST("/", produkhandler.Create)
            produkGroup.PUT("/:id", produkhandler.Update)
            produkGroup.DELETE("/:id", produkhandler.Delete)
        }

        // User API
        userGroup := loginGroup.Group("/user")
        {
            userGroup.GET("/", userhandler.Index)
            userGroup.GET("/:id", userhandler.Show)
            userGroup.POST("/", userhandler.Create)
            userGroup.PUT("/:id", userhandler.Update)
            userGroup.DELETE("/:id", userhandler.Delete)
        }

    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    r.Run(":8080")
}

	// v1 := r.Group("/api")
	// {
	// 	// Login API
	// 	loginGroup := v1.Group("/login")
	// 	{
	// 		loginGroup.POST("/", loginhandler.Login)
	// 	}

	// 	// Protected Produk API
	// 	authProdukGroup := loginGroup.Group("/")
	// 	authProdukGroup.Use(middleware.AuthMiddleware())
	// 	{
	// 		authProdukGroup.GET("/produk", produkhandler.Index)
	// 		authProdukGroup.GET("/produk/:id", produkhandler.Show)
	// 		authProdukGroup.POST("/produk", produkhandler.Create)
	// 		authProdukGroup.PUT("/produk/:id", produkhandler.Update)
	// 		authProdukGroup.DELETE("/produk/:id", produkhandler.Delete)
	// 	}

	// 	// Protected User API
	// 	authUserGroup := loginGroup.Group("/")
	// 	authUserGroup.Use(middleware.AuthMiddleware())
	// 	{
	// 		authUserGroup.GET("/user", userhandler.Index)
	// 		authUserGroup.GET("/user/:id", userhandler.Show)
	// 		authUserGroup.POST("/user", userhandler.Create)
	// 		authUserGroup.PUT("/user/:id", userhandler.Update)
	// 		authUserGroup.DELETE("/user/:id", userhandler.Delete)
	// 	}
	// }
