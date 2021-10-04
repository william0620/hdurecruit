package community

import (
	"github.com/gin-gonic/gin"
	"github.com/go-conflict/toolkit/hduhelp/staff"
	"github.com/go-conflict/toolkit/httpHandler"
	"recruiting_system/internal/pkg/community"
)

const COMMUNITY_ID = "communityId"

func checkCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		communityId, err := getCommunityId(c)
		if err != nil {
			httpHandler.SetJsonErr()
			return
		}
		c.Set(COMMUNITY_ID,communityId)
	}
}

func getCommunityId(c *gin.Context) (communityId string,err error){
	staffId := staff.GetStaffId(c)

	communityId, err = community.GetCommunityIdByStaffId(staffId)
	if err != nil {
		return "", err
	}

	return communityId,nil
}
