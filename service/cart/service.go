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

func (h *Handler) createOrder(ps []types.Product, items []types.CartCheckoutItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}
	// check if all products are actually in stack
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	//calculate total price
	totalPrice := calculateTotalPrice(items, productMap)
	//reduce quantity of products in our db
	//create the order
	//create order items
	return 0, 0, nil
}

func checkIfCartIsInStock(cartItems []types.CartCheckoutItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", product)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}
	return nil
}
