package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(200, Response{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// error response
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(200, Response{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
