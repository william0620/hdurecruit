package community

import (
	"github.com/gin-gonic/gin"
	"github.com/go-conflict/toolkit/hduhelp/staff"
	"github.com/go-conflict/toolkit/httpHandler"
	"recruiting_system/internal/pkg/community"
)

func communityList(c *gin.Context) {
	communities := community.GetAllCommunities()
	httpHandler.SetJsonSuccess(c,communities)
}

type Register struct {
	CommunityId string
	CommunityName string
}

func communityRegister(c *gin.Context) {
	staffId := staff.GetStaffId(c)


}

type Info struct {
	Info string `json:"info"`
}

func getCommunityInfo(c *gin.Context) {
	tempInfo := new(Info) // todo:需要洗
	err := c.BindJSON(tempInfo)
	if err != nil {
		httpHandler.SetJsonErr(c,httpHandler.BadRequest40000,nil,err)
		return
	}

}

func postJob(c *gin.Context) {

}

func publishCommunityInfo(c *gin.Context) {

}

func postCommunityInfo(c *gin.Context) {

}