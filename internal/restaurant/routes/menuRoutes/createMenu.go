package menuRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type CreateMenuRequest struct {
	OnDate          string   `json:"on_date"`
	OpeningRecordAt string   `json:"opening_record_at"`
	ClosingRecordAt string   `json:"closing_record_at"`
	Salads          []string `json:"salads"`
	Garnishes       []string `json:"garnishes"`
	Meats           []string `json:"meats"`
	Soups           []string `json:"soups"`
	Drinks          []string `json:"drinks"`
	Desserts        []string `json:"desserts"`
}

func CreateMenu(ctx *gin.Context, c restaurant.MenuServiceClient) {
	body := CreateMenuRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timesForCreateMenu := createTimeRequest(ctx, body)

	res, err := c.CreateMenu(context.Background(), &restaurant.CreateMenuRequest{
		OnDate:          timesForCreateMenu[0],
		OpeningRecordAt: timesForCreateMenu[1],
		ClosingRecordAt: timesForCreateMenu[2],
		Salads:          body.Salads,
		Garnishes:       body.Garnishes,
		Meats:           body.Meats,
		Soups:           body.Soups,
		Drinks:          body.Drinks,
		Desserts:        body.Desserts,
	})

	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}

func attributeInTimestamp(timeFromRequest string) (*timestamppb.Timestamp, error) {
	t, err := time.Parse(time.RFC3339, timeFromRequest)
	if err != nil {
		return nil, err
	}

	pb := timestamppb.New(t)
	return pb, nil
}

func createTimeRequest(ctx *gin.Context, body CreateMenuRequest) []*timestamppb.Timestamp {
	result := make([]*timestamppb.Timestamp, 0, 3)

	onDate, err := attributeInTimestamp(body.OnDate)
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}
	result = append(result, onDate)

	openingRecordAt, err := attributeInTimestamp(body.OpeningRecordAt)
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}
	result = append(result, openingRecordAt)

	closingRecordAt, err := attributeInTimestamp(body.ClosingRecordAt)
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}
	result = append(result, closingRecordAt)

	return result
}
