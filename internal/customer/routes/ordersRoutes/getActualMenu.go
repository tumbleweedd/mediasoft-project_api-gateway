package ordersRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"net/http"
)

func GetActualMenu(ctx *gin.Context, c customer.OrderServiceClient) {
	res, err := c.GetActualMenu(context.Background(), &customer.GetActualMenuRequest{})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, res)
}
