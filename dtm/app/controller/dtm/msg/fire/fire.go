package fire

import (
	"github.com/dtm-labs/dtm/test/busi"
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"
	"holygin/app/db"

	"github.com/dtm-labs/dtm/dtmcli"
)

/**
问题：

 */


// DoAction 入口请求，调用分布式
func DoAction(c *gin.Context) {
	req := &gin.H{"amount": 30} // 微服务的载荷
	gid := dtmcli.MustGenGid(common.DefaultHTTPServer)
	msg := dtmcli.NewMsg(common.DefaultHTTPServer, gid).
		Add(common.MsgTransOutAction, req).
		Add(common.MsgTransInAction, req)

	err := msg.Prepare(busi.Busi + "/query")
	err = msg.Submit()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"gid":gid})
	return
}

func QueryPreparedBAction(c *gin.Context)  {
	ti, err := dtmcli.BarrierFromQuery(c.Request.URL.Query())
	if err != nil {
		return
	}
	ti.QueryPrepared(db.Get())
}