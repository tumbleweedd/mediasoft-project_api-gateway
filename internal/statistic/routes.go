package statistic

import (
	"github.com/gin-gonic/gin"
	statisticsRoutes "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/statistic/routes"
)

func RegisterRoutes(r *gin.Engine, statURL string) {
	serviceClient := InitServiceClient(statURL)

	svc := &ServiceClient{
		Client: serviceClient.Client,
	}

	statistics := r.Group("/statistics")
	{
		statistics.GET("/amount-of-profit", svc.getAmountOfProfit)
	}
}

func (s *ServiceClient) getAmountOfProfit(ctx *gin.Context) {
	statisticsRoutes.GetAmountOfProfit(ctx, s.Client)
}
