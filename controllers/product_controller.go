package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lcsgborges/goapi/models"
	"github.com/lcsgborges/goapi/usecase"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
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
	var product models.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := models.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := models.Response{
			Message: "Produto não foi encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := models.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProductById(productId)
	if err != nil {
		if err.Error() == "produto com id não encontrado" {
			response := models.Response{
				Message: "Produto não encontrado",
			}
			ctx.JSON(http.StatusNotFound, response)
		} else {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: "Erro ao deletar produto",
			})
		}
		return
	}
	response := models.Response{
		Message: "Produto deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}
