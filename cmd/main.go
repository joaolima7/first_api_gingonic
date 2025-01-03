package main

import (
	"first-api-gin/controller"
	"first-api-gin/db"
	"first-api-gin/repository"
	usecase "first-api-gin/usecases"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Repository layer
	ProductRepository := repository.NewProductRepository(dbConnection)

	//UseCase layer
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	//Controller layer
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.FindProductByID)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")
}
