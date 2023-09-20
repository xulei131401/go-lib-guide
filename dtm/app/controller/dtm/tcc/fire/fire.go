package fire

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"holygin/app/controller/dtm/common"
	"log"

	"github.com/dtm-labs/dtm/dtmcli"
)

/**
问题：
1.执行顺序：各自分支的TRY,各自分支的Confirm

 */



// DoAction 入口请求，调用分布式
func DoAction(c *gin.Context) {
	req := &gin.H{"amount": 30}

	gid := dtmcli.MustGenGid(common.DefaultHTTPServer)
	err := dtmcli.TccGlobalTransaction(common.DefaultHTTPServer, gid, func(tcc *dtmcli.Tcc) (resp *resty.Response, rerr error) {
		resp, rerr = tcc.CallBranch(req, common.TccTransOutAction, common.TccTransOutConfirmAction, common.TccTransOutRevokeAction)
		if rerr != nil {
			log.Println("Out出错:", rerr)
			return
		}

		resp, rerr = tcc.CallBranch(req, common.TccTransInAction, common.TccTransInConfirmAction, common.TccTransInRevokeAction)
		if rerr != nil {
			log.Println("In出错:", rerr)
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