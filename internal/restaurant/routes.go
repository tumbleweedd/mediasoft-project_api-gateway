package restaurant

import (
	"github.com/gin-gonic/gin"
	restaurantMenuRoutes "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/restaurant/routes/menuRoutes"
	restaurantOrdersRoutes "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/restaurant/routes/ordersRoutes"
	restaurantProductsRoutes "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/restaurant/routes/productsRoutes"
)

func RegisterRoutes(r *gin.Engine, restaurantUrl string) {
	serviceClient := InitServiceClient(restaurantUrl)
	tempSvc := ConvertToTempServiceClient(serviceClient)

	svc := &ServiceClient{
		Client: tempSvc,
	}

	restaurant := r.Group("/restaurant")
	{
		menu := restaurant.Group("/menu")
		{
			menu.POST("", svc.createMenu)
			menu.GET("", svc.getMenu)
		}
		products := restaurant.Group("/products")
		{
			products.POST("", svc.createProduct)
			products.GET("", svc.getProducts)
		}
		orders := restaurant.Group("/orders")
		{
			orders.GET("", svc.getOrderList)
		}
	}
}

// --- Menu

func (s *ServiceClient) createMenu(ctx *gin.Context) {
	restaurantMenuRoutes.CreateMenu(ctx, s.Client.menuServiceClient)
}

func (s *ServiceClient) getMenu(ctx *gin.Context) {
	restaurantMenuRoutes.GetMenu(ctx, s.Client.menuServiceClient)
}

// --- Products

func (s *ServiceClient) createProduct(ctx *gin.Context) {
	restaurantProductsRoutes.CreateProduct(ctx, s.Client.productServiceClient)
}

func (s *ServiceClient) getProducts(ctx *gin.Context) {
	restaurantProductsRoutes.GetProducts(ctx, s.Client.productServiceClient)
}

// --- Orders
func (s *ServiceClient) getOrderList(ctx *gin.Context) {
	restaurantOrdersRoutes.GetOrderList(ctx, s.Client.orderServiceClient)
}
