package order

import "context"

// Order domain entity

// Order represents an order
type Order struct {
	ID           string      `json:"id, omitempty"`
	CustomerID   string      `json:"customer_id"`
	Status       string      `json:"status"`
	CreatedOn    int64       `json:"created_on,omitempty"`
	RestaurantId string      `json:"restaurant_id"`
	OrderItems   []OrderItem `json:"oreder_items"`
}

// OrderItem represents items ian order
type OrderItem struct {
	ProuctCode string  `json:"product_code"`
	Name       string  `json:"name"`
	UnitPrice  float32 `json:"unit_pricec"`
	Quantity   int32   `json:"quantity"`
}

// Repository describes order repository
type Repository interface {
	CreateOrder(ctx context.Context, order Order) error
	GetOrder(ctx context.Context, id string) (Order, error)
	ChangeOrderStatus(ctx context.Context, id string, status string) error
}
