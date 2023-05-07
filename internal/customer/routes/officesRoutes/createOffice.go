package officesRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"net/http"
)

type CreateUserRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	//CompanyUuid string `json:"company_uuid"`
}

func CreateOffice(ctx *gin.Context, c customer.OfficeServiceClient) {
	body := CreateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.CreateOffice(context.Background(), &customer.CreateOfficeRequest{
		Name:    body.Name,
		Address: body.Address,
	})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}
