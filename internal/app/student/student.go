package student

import "recruiting_system/internal/dep/httpEngine"

func init()  {
	student := httpEngine.Group("")

	student.GET("")
}
