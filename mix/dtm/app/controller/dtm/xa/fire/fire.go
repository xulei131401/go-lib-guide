package fire

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"holygin/app/controller/dtm/common"
	"log"
)

/**
问题：
1.执行顺序：各自分支的TRY,各自分支的Confirm

 */

var XaClient *dtmcli.XaClient

// DoAction 入口请求，调用分布式
func DoAction(c *gin.Context) {
	req := &gin.H{"amount": 30}

	gid := dtmcli.MustGenGid(common.DefaultHTTPServer)
	err := XaClient.XaGlobalTransaction(gid, func(xa *dtmcli.Xa) (resp *resty.Response, err error) {
		resp, err = xa.CallBranch(req, common.XATransOutAction)
		if err != nil {
			log.Println("Out出错:", err)
			return
		}

		resp, err = xa.CallBranch(req, common.XATransInAction)
		if err != nil {
			log.Println("In出错:", err)
			return
		}

		return
	})

	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"gid": gid})
	return
}