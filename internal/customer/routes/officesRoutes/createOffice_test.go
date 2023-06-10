package officesRoutes

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_customer "github.com/tumbleweedd/mediasoft-intership/api-gateway/internal/customer/routes/officesRoutes/mocks"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_createOffice(t *testing.T) {
	type mockBehavior func(
		mockClient *mock_customer.MockOfficeServiceClient,
		req *customer.CreateOfficeRequest,
		expectedResponse *customer.CreateOfficeResponse,
	)

	testTable := []struct {
		name               string
		requestBody        map[string]string
		expectedRequest    *customer.CreateOfficeRequest
		expectedResponse   *customer.CreateOfficeResponse
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedJSON       string
	}{
		{
			name: "OK",
			requestBody: map[string]string{
				"name":    "Test name",
				"address": "Test address",
			},
			expectedRequest: &customer.CreateOfficeRequest{
				Name:    "Test name",
				Address: "Test address",
			},
			expectedResponse: &customer.CreateOfficeResponse{},
			mockBehavior: func(
				mockClient *mock_customer.MockOfficeServiceClient,
				req *customer.CreateOfficeRequest,
				expectedResponse *customer.CreateOfficeResponse,
			) {
				mockClient.EXPECT().CreateOffice(gomock.Any(), req).Return(expectedResponse, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedJSON:       `{}`,
		},
		{
			name: "Service Failure",
			requestBody: map[string]string{
				"name":    "Test name",
				"address": "Test address",
			},
			expectedRequest: &customer.CreateOfficeRequest{
				Name:    "Test name",
				Address: "Test address",
			},
			expectedResponse: &customer.CreateOfficeResponse{},
			mockBehavior: func(
				mockClient *mock_customer.MockOfficeServiceClient,
				req *customer.CreateOfficeRequest,
				expectedResponse *customer.CreateOfficeResponse,
			) {
				mockClient.EXPECT().CreateOffice(gomock.Any(), req).Return(nil, status.Error(codes.Internal, errors.New("internal server error").Error()))
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedJSON:       `{"code":500,"error":"internal server error"}null`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock_customer.NewMockOfficeServiceClient(ctrl)
			testCase.mockBehavior(mockClient, testCase.expectedRequest, testCase.expectedResponse)

			// Test Server
			router := gin.Default()
			router.POST("customer/offices", func(ctx *gin.Context) {
				CreateOffice(ctx, mockClient)
			})

			requestJSON, _ := json.Marshal(&testCase.requestBody)
			w := httptest.NewRecorder()

			// Test Request
			req := httptest.NewRequest("POST", "/customer/offices",
				bytes.NewBufferString(string(requestJSON)))
			req.Header.Set("Content-Type", "application/json")

			// Perform Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedJSON, w.Body.String())
		})
	}
}
