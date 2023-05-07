package usersRoute

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"net/http"
)

type CreateUserRequestBody struct {
	Name       string `json:"name"`
	OfficeUuid string `json:"office_uuid"`
}

func CreateUser(ctx *gin.Context, c customer.UserServiceClient) {
	body := &CreateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(body.OfficeUuid)

	res, err := c.CreateUser(context.Background(), &customer.CreateUserRequest{
		Name:       body.Name,
		OfficeUuid: body.OfficeUuid,
	})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	ctx.JSON(http.StatusOK, &res)
}
