package usecase

import (
	"first-api-gin/model"
	"first-api-gin/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) FindProductByID(id int) (*model.Product, error) {
	return pu.repository.FindProductByID(id)
}

func (pu *ProductUseCase) UpdateProduct(updates map[string]interface{}, id int) (bool, error) {
	return pu.repository.UpdateProduct(updates, id)
}

func (pu *ProductUseCase) DeleteProduct(id int) (bool, error) {
	return pu.repository.DeleteProduct(id)
}
