package usecase

import (
	"github.com/lcsgborges/goapi/models"
	"github.com/lcsgborges/goapi/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*models.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(id_product int) error {
	err := pu.repository.DeleteProductById(id_product)
	if err != nil {
		return err
	}
	return nil
}
