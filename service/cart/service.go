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
	productsMap := make(map[int]types.Product)
	for _, product := range ps {
		productsMap[product.ID] = product
	}
	// check if all products are actually in stack
	if err := checkIfCartIsInStock(items, productsMap); err != nil {
		return 0, 0, err
	}
	//calculate total price
	totalPrice := calculateTotalPrice(items, productsMap)
	//reduce quantity of products in our db
	for _, item := range items {
		product := productsMap[item.ProductID]
		product.Quantity -= item.Quantity

		err := h.productStore.UpdateProduct(product)
		if err != nil {
			return 0, 0, err
		}
	}

	//create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "soem addres",
	})
	if err != nil {
		return 0, 0, err
	}
	//create order items
	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productsMap[item.ProductID].Price,
		})
	}
	return 0, totalPrice, nil
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

func calculateTotalPrice(cartItems []types.CartCheckoutItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
