package out

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"
	"log"
)

var amount = 100

func TransOutAction (c *gin.Context) {
	num := common.GetAmount(c)
	amount -= num
	log.Println("---- TransOutAction:", amount)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}

func TransOutConfirmAction(c *gin.Context){
	log.Println("---- TransOutConfirmAction:")
	c.JSON(200, gin.H{"dtm_result": dtmcli.ResultSuccess})
}

func TransOutRevokeAction(c *gin.Context){
	num := common.GetAmount(c)
	amount += num
	log.Println("---- TransOutRevokeAction:", amount)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}