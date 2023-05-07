package errors

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func HandleServiceError(ctx *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		// не удалось распарсить ошибку
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "internal server error",
		})
		return
	}

	switch st.Code() {
	case codes.InvalidArgument:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": st.Message(),
		})
	case codes.NotFound:
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": st.Message(),
		})
	case codes.PermissionDenied:
		ctx.JSON(http.StatusForbidden, gin.H{
			"code":  http.StatusForbidden,
			"error": st.Message(),
		})
	case codes.Internal:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": st.Message(),
		})
	case codes.Unauthenticated:
		ctx.JSON(http.StatusForbidden, gin.H{
			"code":  http.StatusForbidden,
			"error": st.Message(),
		})
	// и т.д. для всех возможных кодов ошибок
	default:
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": st.Message(),
		})
	}
}

type ServiceError struct {
	StatusCode int
	Message    string
}

func (se ServiceError) Error() string {
	return se.Message
}

func (se ServiceError) HTTPStatus() int {
	return se.StatusCode
}
