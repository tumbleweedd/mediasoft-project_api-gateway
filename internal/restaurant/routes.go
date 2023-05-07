package restaurant

import (
	"github.com/gin-gonic/gin"
	menuRoutes2 "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/restaurant/routes/menuRoutes"
	productsRoutes2 "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/restaurant/routes/productsRoutes"
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
	}
}

// --- Menu

func (s *ServiceClient) createMenu(ctx *gin.Context) {
	menuRoutes2.CreateMenu(ctx, s.Client.menuServiceClient)
}

func (s *ServiceClient) getMenu(ctx *gin.Context) {
	menuRoutes2.GetMenu(ctx, s.Client.menuServiceClient)
}

// --- Products

func (s *ServiceClient) createProduct(ctx *gin.Context) {
	productsRoutes2.CreateProduct(ctx, s.Client.productServiceClient)
}

func (s *ServiceClient) getProducts(ctx *gin.Context) {
	productsRoutes2.GetProducts(ctx, s.Client.productServiceClient)
}
