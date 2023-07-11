package router

import "github.com/gin-gonic/gin"

func runController() {
	r.GET("/", renderIndex)
	apiV1()
}

func apiV1() {
	v1 := r.Group("/v1")
	v1.POST("/tv/:id/:language", getTvInfo)
	apiV1Path(v1)
}

func apiV1Path(v1 *gin.RouterGroup) {
	path := v1.Group("/path")

	path.GET("/", getPaths)
	path.POST("/raw/:path", changeRawPath)
	path.POST("/tv/:path", changeTVPath)
	path.POST("/movie/:path", changeMoviePath)
}
