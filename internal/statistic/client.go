package statistic

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/statistics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient struct {
	Client statistics.StatisticsServiceClient
}

func InitServiceClient(statisticsURL string) *ServiceClient {
	conn, err := grpc.Dial(statisticsURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect to customerService:%v", err)
	}

	statServiceClient := statistics.NewStatisticsServiceClient(conn)

	return &ServiceClient{
		Client: statServiceClient,
	}
}
