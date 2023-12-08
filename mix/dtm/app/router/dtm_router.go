package router

import (
	"log"

	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"

	sagaFire "holygin/app/controller/dtm/saga/fire"
	sagaIn "holygin/app/controller/dtm/saga/in"
	sagaOut "holygin/app/controller/dtm/saga/out"

	tccFire "holygin/app/controller/dtm/tcc/fire"
	tccIn "holygin/app/controller/dtm/tcc/in"
	tccOut "holygin/app/controller/dtm/tcc/out"

	msgFire "holygin/app/controller/dtm/msg/fire"
	msgIn "holygin/app/controller/dtm/msg/in"
	msgOut "holygin/app/controller/dtm/msg/out"

	xaFire "holygin/app/controller/dtm/xa/fire"
	xaIn "holygin/app/controller/dtm/xa/in"
	xaOut "holygin/app/controller/dtm/xa/out"

)

var XaClient *dtmcli.XaClient

func AddDTMRouter(router *gin.Engine) {
	saga := router.Group("/trans/saga")
	{
		saga.POST("/out", sagaOut.TransOutAction)
		saga.POST("/out/revoke", sagaOut.TransOutRevokeAction)
		saga.POST("/in", sagaIn.TransInAction)
		saga.POST("/in/revoke", sagaIn.TransInRevokeAction)
		saga.POST("/fire", sagaFire.DoAction)
	}

	tcc := router.Group("/trans/tcc")
	{
		tcc.POST("/out", tccOut.TransOutAction)
		tcc.POST("/out/confirm", tccOut.TransOutConfirmAction)
		tcc.POST("/out/revoke", tccOut.TransOutRevokeAction)
		tcc.POST("/in", tccIn.TransInAction)
		tcc.POST("/in/confirm", tccIn.TransInConfirmAction)
		tcc.POST("/in/revoke", tccIn.TransInRevokeAction)
		tcc.POST("/fire", tccFire.DoAction)
	}

	msg := router.Group("/trans/msg")
	{
		msg.POST("/out", msgOut.TransOutAction)
		msg.POST("/in", msgIn.TransInAction)
		msg.POST("/fire", msgFire.DoAction)
		msg.POST("/QueryPreparedB", msgFire.QueryPreparedBAction)
	}

	//AddDTMXARouter(router)
}

func AddDTMXARouter(router *gin.Engine) {
	xa := router.Group("/trans/xa")
	{
		xa.POST("/out", xaOut.TransOutAction)
		xa.POST("/in", xaIn.TransInAction)
		xa.POST("/fire", xaFire.DoAction)
	}

	var err error
	XaClient, err = dtmcli.NewXaClient(common.DefaultHTTPServer, dtmcli.DBConf{}, "http://localhost:8001/trans/xa/xa", func(path string, xaClient *dtmcli.XaClient) {
		xa.POST(path, func(c *gin.Context) {
			xaClient.HandleCallback(c.Query("gid"), c.Query("branch_id"), c.Query("action"))
		})
	})

	if err != nil {
		log.Println("xa error:", err)
	}
}