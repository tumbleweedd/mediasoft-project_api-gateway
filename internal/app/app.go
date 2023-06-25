package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/customer"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/restaurant"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/statistic"

	"log"
	"os"
)

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	r := gin.Default()

	customer.RegisterRoutes(r, os.Getenv("CUSTOMER_SVC_URL"))
	restaurant.RegisterRoutes(r, os.Getenv("RESTAURANT_SVC_URL"))
	statistic.RegisterRoutes(r, os.Getenv("STATISTICS_SVC_URL"))

	r.Run(os.Getenv("PORT"))
}
