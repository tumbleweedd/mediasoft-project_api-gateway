package statisticsRoutes

import (
	"context"
	goerrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/statistics"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

func GetAmountOfProfit(ctx *gin.Context, c statistics.StatisticsServiceClient) {
	startDateParam := ctx.Query("start_date")
	endDateParam := ctx.Query("end_date")

	startDate, err := time.Parse(time.RFC3339, startDateParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, goerrors.New("invalid start_date param"))
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, goerrors.New("invalid end_date param"))
		return
	}

	res, err := c.GetAmountOfProfit(context.Background(), &statistics.GetAmountOfProfitRequest{
		StartDate: timestamppb.New(startDate),
		EndDate:   timestamppb.New(endDate),
	})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}
