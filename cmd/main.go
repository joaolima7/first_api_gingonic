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

	//Endpoints
	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.FindProductByID)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product/update/:id", ProductController.UpdateProduct)
	server.DELETE("/product/del/:id", ProductController.DeleteProduct)

	server.Run(":8000")
}
