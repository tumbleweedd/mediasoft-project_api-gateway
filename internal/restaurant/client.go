package restaurant

import (
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CustomerServiceClient struct {
	orderServiceClient   restaurant.OrderServiceClient
	menuServiceClient    restaurant.MenuServiceClient
	productServiceClient restaurant.ProductServiceClient
}

type ServiceClient struct {
	Client CustomerServiceClient
}

func InitServiceClient(restaurantUrl string) *ServiceClient {
	conn, err := grpc.Dial(restaurantUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to restaurantService:", err)
	}
	orderServiceClient := restaurant.NewOrderServiceClient(conn)
	menuServiceClient := restaurant.NewMenuServiceClient(conn)
	productServiceClient := restaurant.NewProductServiceClient(conn)

	return &ServiceClient{
		Client: CustomerServiceClient{
			orderServiceClient:   orderServiceClient,
			menuServiceClient:    menuServiceClient,
			productServiceClient: productServiceClient,
		},
	}
}

func ConvertToTempServiceClient(serviceClient *ServiceClient) CustomerServiceClient {
	return CustomerServiceClient{
		orderServiceClient:   serviceClient.Client.orderServiceClient,
		menuServiceClient:    serviceClient.Client.menuServiceClient,
		productServiceClient: serviceClient.Client.productServiceClient,
	}
}
