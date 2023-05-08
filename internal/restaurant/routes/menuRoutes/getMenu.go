package menuRoutes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type GetMenuResponseBody struct {
	Menu *Menu `json:"menu"`
}

type Menu struct {
	MenuUUID        string     `json:"uuid"`
	OnDate          string     `json:"on_date"`
	OpeningRecordAt string     `json:"opening_record_at"`
	ClosingRecordAt string     `json:"closing_record_at"`
	Salads          []*Product `json:"salads,omitempty" `
	Garnishes       []*Product `json:"garnishes,omitempty"`
	Meats           []*Product `json:"meats,omitempty"`
	Soups           []*Product `json:"soups,omitempty"`
	Drinks          []*Product `json:"drinks,omitempty"`
	Desserts        []*Product `json:"desserts,omitempty"`
	CreatedAt       string     `json:"created_at,omitempty"`
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

func GetMenu(ctx *gin.Context, c restaurant.MenuServiceClient) {
	menuOnDateQuery := ctx.Query("on_date")
	t, err := time.Parse(time.RFC3339, menuOnDateQuery)
	fmt.Println(t)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := c.GetMenu(context.Background(), &restaurant.GetMenuRequest{
		OnDate: timestamppb.New(t),
	})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	menu := &Menu{}
	menuResponseBody := menuResponse(res, menu)

	ctx.JSON(http.StatusOK, menuResponseBody)
}

func menuResponse(res *restaurant.GetMenuResponse, menu *Menu) *GetMenuResponseBody {
	menu.MenuUUID = res.Menu.Uuid
	menu.OnDate = res.Menu.OnDate.AsTime().Format(time.RFC3339)
	menu.OpeningRecordAt = res.Menu.OpeningRecordAt.AsTime().Format(time.RFC3339)
	menu.ClosingRecordAt = res.Menu.ClosingRecordAt.AsTime().Format(time.RFC3339)
	menu.CreatedAt = res.Menu.CreatedAt.AsTime().Format(time.RFC3339)

	for _, products := range [][]*restaurant.Product{
		res.Menu.Salads, res.Menu.Garnishes, res.Menu.Meats, res.Menu.Soups, res.Menu.Drinks, res.Menu.Desserts,
	} {
		for _, product := range products {
			addProductToMenu(menu, product)
		}

	}
	response := &GetMenuResponseBody{Menu: menu}
	return response
}

func addProductToMenu(menu *Menu, product *restaurant.Product) {
	p := &Product{
		ProductUUID: product.Uuid,
		Name:        product.Name,
		Description: product.Description,
		Type:        product.Type.String(),
		Weight:      product.Weight,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.AsTime().Format(time.RFC3339),
	}

	switch product.Type {
	case restaurant.ProductType_PRODUCT_TYPE_SALAD:
		menu.Salads = append(menu.Salads, p)
	case restaurant.ProductType_PRODUCT_TYPE_GARNISH:
		menu.Garnishes = append(menu.Garnishes, p)
	case restaurant.ProductType_PRODUCT_TYPE_MEAT:
		menu.Meats = append(menu.Meats, p)
	case restaurant.ProductType_PRODUCT_TYPE_SOUP:
		menu.Soups = append(menu.Soups, p)
	case restaurant.ProductType_PRODUCT_TYPE_DRINK:
		menu.Drinks = append(menu.Drinks, p)
	case restaurant.ProductType_PRODUCT_TYPE_DESSERT:
		menu.Desserts = append(menu.Desserts, p)
	}
}
