package menuRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

func GetMenu(ctx *gin.Context, c restaurant.MenuServiceClient) {
	menuOnDateQuery := ctx.Query("on_date")
	t, err := time.Parse(time.RFC3339, menuOnDateQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := c.GetMenu(context.Background(), &restaurant.GetMenuRequest{
		OnDate: timestamppb.New(t),
	})

	ctx.JSON(http.StatusOK, &res)
}
