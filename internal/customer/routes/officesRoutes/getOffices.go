package officesRoutes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/mediasoft-intership/api-gateway/pkg/errors"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"net/http"
)

type GetOfficesResponseBody struct {
	Result []*Offices `json:"result"`
}

type Offices struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}

func GetOffices(ctx *gin.Context, c customer.OfficeServiceClient) {
	res, err := c.GetOfficeList(context.Background(), &customer.GetOfficeListRequest{})
	if err != nil {
		errors.HandleServiceError(ctx, err)
	}

	offices := make([]*Offices, 0, len(res.Result))

	officesResponseBody := officeResponse(res, offices)

	ctx.JSON(http.StatusOK, officesResponseBody)
}

func officeResponse(res *customer.GetOfficeListResponse, officesResponseBody []*Offices) *GetOfficesResponseBody {
	for _, office := range res.Result {
		officesResponseBody = append(officesResponseBody, &Offices{
			Uuid:      office.Uuid,
			Name:      office.Name,
			Address:   office.Address,
			CreatedAt: office.CreatedAt.AsTime().Format("2006-01-02T15:04:05.000Z"),
		})
	}

	response := &GetOfficesResponseBody{Result: officesResponseBody}
	return response
}
