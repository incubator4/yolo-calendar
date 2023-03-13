package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/pkg/config"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
	"net/http"
)

var client *lighthouse.Client

func init() {
	credential := common.NewCredential(
		config.GlobalConfig.TencentCloudSecretId,
		config.GlobalConfig.TencentCloudSecretKey,
	)
	client, _ = lighthouse.NewClient(credential, regions.Shanghai, profile.NewClientProfile())
}

func registerStatus(g *gin.RouterGroup) {
	g.GET("/traffic", ListTraffic)
}

func ListTraffic(c *gin.Context) {
	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
	response, err := client.DescribeInstancesTrafficPackages(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
