package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		c.JSON(code, gin.H{
			"code":    code,
			"message": msg,
		})
	} else {
		c.JSON(code, gin.H{
			"code":    code,
			"message": msg,
			"data":    data,
		})
	}

}

func Success(c *gin.Context) {
	Response(c, 200, "success", nil)
}
func SuccessWithData(c *gin.Context, data interface{}) {
	Response(c, 200, "success", data)
}

func Fail(c *gin.Context, msg string) {
	Response(c, 400, msg, nil)
}
