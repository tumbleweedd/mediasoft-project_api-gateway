package restaurant

import (
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantServiceClient struct {
	orderServiceClient   restaurant.OrderServiceClient
	menuServiceClient    restaurant.MenuServiceClient
	productServiceClient restaurant.ProductServiceClient
}

type ServiceClient struct {
	Client RestaurantServiceClient
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
		Client: RestaurantServiceClient{
			orderServiceClient:   orderServiceClient,
			menuServiceClient:    menuServiceClient,
			productServiceClient: productServiceClient,
		},
	}
}

func ConvertToTempServiceClient(serviceClient *ServiceClient) RestaurantServiceClient {
	return RestaurantServiceClient{
		orderServiceClient:   serviceClient.Client.orderServiceClient,
		menuServiceClient:    serviceClient.Client.menuServiceClient,
		productServiceClient: serviceClient.Client.productServiceClient,
	}
}
