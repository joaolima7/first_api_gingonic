package controller

import (
	"first-api-gin/model"
	usecase "first-api-gin/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)

}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) FindProductByID(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "ID is required param.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Param is not a number.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.FindProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{
			Message: "ID is required.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Param is not a number.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var productUpdates map[string]interface{}
	err = ctx.ShouldBindJSON(&productUpdates)
	if err != nil {
		response := model.Response{
			Message: "Invalid JSON body.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updated, err := p.productUseCase.UpdateProduct(productUpdates, productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if updated {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product updated successfully",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
	})
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is required.",
		})
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID not is a number.",
		})
		return
	}

	isDeleted, err := p.productUseCase.DeleteProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in delete Product. " + err.Error(),
		})
		return
	}

	if isDeleted {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product deleted successfully",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
	})
}
