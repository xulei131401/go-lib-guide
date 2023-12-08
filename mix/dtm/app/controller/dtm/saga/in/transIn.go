package in

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"
	"log"
	"time"
)

var amount = 0

func TransInAction (c *gin.Context) {
	num := common.GetAmount(c)
	amount += num
	log.Println("---- TransInAction:", amount)
	time.Sleep(time.Second * 7)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultFailure})
}

func TransInRevokeAction(c *gin.Context){
	num := common.GetAmount(c)
	amount -= num
	log.Println("---- TransInRevokeAction:", amount)
	time.Sleep(time.Second * 3)
	c.JSON(200, gin.H{"amount": amount, "dtm_result": dtmcli.ResultSuccess})
}