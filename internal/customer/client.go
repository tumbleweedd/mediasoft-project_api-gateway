package customer

import (
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CustomerServiceClient struct {
	officeServiceClient customer.OfficeServiceClient
	userServiceClient   customer.UserServiceClient
	orderServiceClient  customer.OrderServiceClient
}

type ServiceClient struct {
	Client CustomerServiceClient
}

func InitServiceClient(customerUrl string) *ServiceClient {
	conn, err := grpc.Dial(customerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to customerService:", err)
	}

	userServiceClient := customer.NewUserServiceClient(conn)
	officeServiceClient := customer.NewOfficeServiceClient(conn)
	orderServiceClient := customer.NewOrderServiceClient(conn)

	return &ServiceClient{
		Client: CustomerServiceClient{
			officeServiceClient: officeServiceClient,
			userServiceClient:   userServiceClient,
			orderServiceClient:  orderServiceClient,
		},
	}
}

func ConvertToTempServiceClient(serviceClient *ServiceClient) CustomerServiceClient {
	return CustomerServiceClient{
		officeServiceClient: serviceClient.Client.officeServiceClient,
		orderServiceClient:  serviceClient.Client.orderServiceClient,
		userServiceClient:   serviceClient.Client.userServiceClient,
	}
}
