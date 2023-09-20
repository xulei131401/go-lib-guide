package in

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"
	"log"
)

var amount = 0

func TransInAction (c *gin.Context) {
	num := common.GetAmount(c)
	amount += num
	log.Println("---- TransInAction:", amount)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}

func TransInConfirmAction(c *gin.Context){
	log.Println("---- TransInConfirmAction:")
	c.JSON(200, gin.H{"dtm_result": dtmcli.ResultSuccess})
}

func TransInRevokeAction(c *gin.Context){
	num := common.GetAmount(c)
	amount -= num
	log.Println("---- TransInRevokeAction:", amount)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}