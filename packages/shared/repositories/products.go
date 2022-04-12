package repositories

import "github.com/BawnX/go-example-2/packages/shared/models"

// GetProducts returns all products from the database
func GetProducts() models.Products {
	return models.ProductList
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*models.ProductModel, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, models.ErrProductNotFound
	}

	return models.ProductList[i], nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(p *models.ProductModel) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return models.ErrProductNotFound
	}

	// update the product in the DB
	models.ProductList[i] = p

	return nil
}

// AddProduct adds a new product to the database
func AddProduct(p *models.ProductModel) *models.ProductModel {
	// get the next id in sequence
	maxID := models.ProductList[len(models.ProductList)-1].ID
	p.ID = maxID + 1
	models.ProductList = append(models.ProductList, p)
	return p
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return models.ErrProductNotFound
	}

	models.ProductList = append(models.ProductList[:i], models.ProductList[i+1])

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range models.ProductList {
		if p.ID == id {
			return i
		}
	}

	return -1
}
