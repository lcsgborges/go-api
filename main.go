package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsgborges/goapi/controllers"
	"github.com/lcsgborges/goapi/db"
	"github.com/lcsgborges/goapi/repository"
	"github.com/lcsgborges/goapi/usecase"
)

func main() {
	
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada Usecase:
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	// Camada de Controllers:
	productController := controllers.NewProductController(ProductUseCase)

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductsById)
	server.DELETE("/product/:productId", productController.DeleteProductById)

	server.Run(":8080")
}
