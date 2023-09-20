package out

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"
	"log"
	"time"
)

var amount = 100

func TransOutAction (c *gin.Context) {
	num := common.GetAmount(c)
	amount -= num
	log.Println("---- TransOutAction:", amount)
	time.Sleep(time.Second * 5)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}

func TransOutRevokeAction(c *gin.Context){
	num := common.GetAmount(c)
	amount += num
	log.Println("---- TransOutRevokeAction:", amount)
	time.Sleep(time.Second * 3)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}