package usersRoute

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"net/http"
	"time"
)

type GetUsersResponseBody struct {
	Result []*User `json:"result"`
}

type User struct {
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	OfficeUuid string `json:"office_uuid"`
	OfficeName string `json:"office_name"`
	CreatedAt  string `json:"created_at"`
}

func GetUsers(ctx *gin.Context, c customer.UserServiceClient) {
	officeUuid := ctx.Query("office_uuid")
	if officeUuid == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadGateway,
			"error": "bad office_uuid",
		})
		return
	}

	var users []*User

	res, err := c.GetUserList(context.Background(), &customer.GetUserListRequest{
		OfficeUuid: officeUuid,
	})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	usersResponseBody := usersResponse(res, users)

	ctx.JSON(http.StatusOK, usersResponseBody)
}

func usersResponse(res *customer.GetUserListResponse, usersResponseBody []*User) *GetUsersResponseBody {
	for _, user := range res.Result {
		timestamp := time.Unix(user.CreatedAt.GetSeconds(), int64(user.CreatedAt.GetNanos())).UTC()
		timestampFormatted := timestamp.Format("2006-01-02T15:04:05.000Z")

		usersResponseBody = append(usersResponseBody, &User{
			Uuid:       user.Uuid,
			Name:       user.Name,
			OfficeUuid: user.Uuid,
			OfficeName: user.Name,
			CreatedAt:  timestampFormatted,
		})
	}

	response := &GetUsersResponseBody{Result: usersResponseBody}
	return response
}
