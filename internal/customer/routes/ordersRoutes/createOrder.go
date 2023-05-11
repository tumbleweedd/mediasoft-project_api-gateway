package ordersRoutes

type CreateOrderRequestBody struct {
	UserUUID string `json:"user_uuid"`
}

type CustomerOrderItem struct {
	Count       int    `json:"count"`
	ProductUUID string `json:"product_uuid"`
}
