package httpEngine


import "github.com/gin-gonic/gin"

func Run() {
	err := s.HttpEngine.Run(s.Http.Port)
	if err != nil {
		panic(err)
	}
}

func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.HttpEngine.Group(relativePath, handlers...)
}
