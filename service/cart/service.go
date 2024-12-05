package cart

import (
	"fmt"
	"go-backend-api-jwt-mysql/types"
)

func getCartItemsIDs(item []types.CartCheckoutItem) ([]int, error) {
	productsIDs := make([]int, len(item))
	for i, item := range item {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the products %d", item.ProductID)
		}
		productsIDs[i] = item.ProductID
	}
	return productsIDs, nil
}

func (h Handler) createOrder(ps []types.Product, item []types.CartCheckoutItem, userID int) (int, error) {
	// check if all products are actually in stack
	//calculate total price
	//reduce quantity of products in our db
	//create the order
	//create order items
}
