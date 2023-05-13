package ordersRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"net/http"
)

type CreateOrderRequestBody struct {
	UserUUID  string              `json:"user_uuid"`
	Salads    []CustomerOrderItem `json:"salads"`
	Garnishes []CustomerOrderItem `json:"garnishes"`
	Meats     []CustomerOrderItem `json:"meats"`
	Soups     []CustomerOrderItem `json:"soups"`
	Drinks    []CustomerOrderItem `json:"drinks"`
	Desserts  []CustomerOrderItem `json:"desserts"`
}

type CustomerOrderItem struct {
	Count       int    `json:"count"`
	ProductUUID string `json:"product_uuid"`
}

func CreateOrder(ctx *gin.Context, c customer.OrderServiceClient) {
	body := &CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	createOrderRequest := createOrderRequestFromRequestBody(body)

	res, err := c.CreateOrder(context.Background(), createOrderRequest)
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}

func createOrderRequestFromRequestBody(reqBody *CreateOrderRequestBody) *customer.CreateOrderRequest {
	orderRequest := &customer.CreateOrderRequest{
		UserUuid: reqBody.UserUUID,
	}

	addOrderItems := func(dest []*customer.OrderItem, src []CustomerOrderItem) {
		for _, item := range src {
			dest = append(dest, &customer.OrderItem{
				Count:       int32(item.Count),
				ProductUuid: item.ProductUUID,
			})
		}
	}

	addOrderItems(orderRequest.Salads, reqBody.Salads)
	addOrderItems(orderRequest.Garnishes, reqBody.Garnishes)
	addOrderItems(orderRequest.Meats, reqBody.Meats)
	addOrderItems(orderRequest.Soups, reqBody.Soups)
	addOrderItems(orderRequest.Drinks, reqBody.Drinks)
	addOrderItems(orderRequest.Desserts, reqBody.Desserts)

	return orderRequest
}
