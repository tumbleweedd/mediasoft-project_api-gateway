package customer

import (
	"github.com/gin-gonic/gin"
	officesRoutes2 "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/customer/routes/officesRoutes"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/customer/routes/ordersRoutes"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/customer/routes/usersRoute"
)

func RegisterRoutes(r *gin.Engine, customerUrl string) {
	serviceClient := InitServiceClient(customerUrl)
	tempSvc := ConvertToTempServiceClient(serviceClient)

	svc := &ServiceClient{
		Client: tempSvc,
	}

	customer := r.Group("/customer")
	{
		offices := customer.Group("/offices")
		{
			offices.POST("", svc.createOffice)
			offices.GET("", svc.getOffices)
		}
		users := customer.Group("/users")
		{
			users.GET("", svc.getUsers)
			users.POST("", svc.createUser)
		}
		orders := customer.Group("/orders")
		{
			orders.POST("", svc.createOrder)
		}
	}
}

// --- Offices

func (s *ServiceClient) createOffice(ctx *gin.Context) {
	officesRoutes2.CreateOffice(ctx, s.Client.officeServiceClient)
}

func (s *ServiceClient) getOffices(ctx *gin.Context) {
	officesRoutes2.GetOffices(ctx, s.Client.officeServiceClient)
}

// --- Users

func (s *ServiceClient) createUser(ctx *gin.Context) {
	usersRoute.CreateUser(ctx, s.Client.userServiceClient)
}

func (s *ServiceClient) getUsers(ctx *gin.Context) {
	usersRoute.GetUsers(ctx, s.Client.userServiceClient)
}

// --- Orders

func (s *ServiceClient) getActualMenu(ctx *gin.Context) {
	ordersRoutes.GetActualMenu(ctx, s.Client.orderServiceClient)
}

func (s *ServiceClient) createOrder(ctx *gin.Context) {
	ordersRoutes.CreateOrder(ctx, s.Client.orderServiceClient)
}
