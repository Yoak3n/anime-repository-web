package router

import "github.com/gin-gonic/gin"

func response(c *gin.Context, code int, msg string, data interface{}) {
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

func success(c *gin.Context) {
	response(c, 200, "success", nil)
}
func successWithData(c *gin.Context, data interface{}) {
	response(c, 200, "success", data)
}

func fail(c *gin.Context, msg string) {
	response(c, 400, msg, nil)
}
