package community

import (
	"recruiting_system/internal/dep/httpEngine"
)

func init()  {
	community := httpEngine.Group("/community")

	community.Use(checkCommunity())

	register := community.Group("/register")
	{
		register.GET("/list",communityList)
		register.POST("",communityRegister)
	}

	info := community.Group("/info")
	{
		info.GET("",getCommunityInfo) // 获取当前
		info.POST("",postCommunityInfo) // 写入：初次和修改
		info.POST("/publish",publishCommunityInfo) // 发布
		info.POST("/job",postJob) // 修改招募岗位
	}

	manage := community.Group("/manage")
	{
		manage.GET("") // 获取报名名单
		manage.PUT("") // 修改报名者状态
	}

	msg := community.Group("/msg")
	{
		msg.POST("")
	}
}




