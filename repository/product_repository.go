package repository

import (
	"database/sql"
	"fmt"

	"github.com/lcsgborges/goapi/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (name, price) values ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*models.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto models.Product
	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &produto, nil
}

func (pr *ProductRepository) DeleteProductById(id_product int) error {
	// Preparando a query DELETE
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}
	defer query.Close()

	result, err := query.Exec(id_product)
	if err != nil {
		fmt.Println("Erro ao executar delete:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Erro ao verificar registros afetados:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto com id %d não encontrado", id_product)
	}

	return nil
}
