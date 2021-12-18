package utilities

import "github.com/gin-gonic/gin"

const (
	failMessage = "failure"
)

type FailureResponse struct {
	Message string `json:"message" example:"Error message for users"`
	Details string `json:"details" example:"Error message for developers"`
	Status  string `json:"status" example:"failure"`
}

func Failure(message string, details string, code int, ctx *gin.Context) {
	failure := FailureResponse{
		Message: message,
		Status:  failMessage,
		Details: details,
	}
	ctx.JSON(code, failure)
}
