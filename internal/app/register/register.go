package register

import (
	"recruiting_system/internal/dep/httpEngine"
)

func init()  {
	register := httpEngine.Group("/register")
	{
		register.GET("/community/list",communityList)
		register.GET("/community",communityRegister)
		register.GET("/student")
		register.GET("/admin")
	}

}






