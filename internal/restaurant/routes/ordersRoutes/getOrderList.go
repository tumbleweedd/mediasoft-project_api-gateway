package ordersRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"net/http"
)

type RestaurantOrder struct {
	ProductUUID string `json:"product_uuid"`
	ProductName string `json:"product_name"`
	Count       string `json:"count"`
}

type OrderByOffice struct {
	OfficeUUID    string            `json:"company_id"`
	OfficeName    string            `json:"office_name"`
	OfficeAddress string            `json:"office_address"`
	Order         []RestaurantOrder `json:"result"`
}

func GetOrderList(ctx *gin.Context, c restaurant.OrderServiceClient) {
	res, err := c.GetUpToDateOrderList(context.Background(), &restaurant.GetUpToDateOrderListRequest{})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}
