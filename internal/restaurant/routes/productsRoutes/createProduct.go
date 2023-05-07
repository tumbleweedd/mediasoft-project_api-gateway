package productsRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"net/http"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Weight      int32   `json:"weight"`
	Price       float64 `json:"price"`
}

func CreateProduct(ctx *gin.Context, c restaurant.ProductServiceClient) {
	body := CreateProductRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productType := restaurant.ProductType_value[body.Type]
	res, err := c.CreateProduct(context.Background(), &restaurant.CreateProductRequest{
		Name:        body.Name,
		Description: body.Description,
		Type:        restaurant.ProductType(productType),
		Weight:      body.Weight,
		Price:       body.Price,
	})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}
