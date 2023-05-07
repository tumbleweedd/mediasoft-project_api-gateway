package productsRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"net/http"
	"time"
)

type GetProductsResponseBody struct {
	Result []*Product `json:"result"`
}

type Product struct {
	ProductUUID string  `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Weight      int32   `json:"weight"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
}

func GetProducts(ctx *gin.Context, c restaurant.ProductServiceClient) {
	res, err := c.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	products := make([]*Product, 0, len(res.Result))

	responseBody := productResponse(res, products)

	ctx.JSON(http.StatusOK, responseBody)
}

func productResponse(res *restaurant.GetProductListResponse, products []*Product) *GetProductsResponseBody {
	for _, product := range res.Result {
		products = append(products, &Product{
			ProductUUID: product.Uuid,
			Name:        product.Name,
			Description: product.Description,
			Type:        product.Type.String(),
			Weight:      product.Weight,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt.AsTime().Format(time.RFC3339),
		})
	}

	response := &GetProductsResponseBody{Result: products}
	return response
}
