package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApiParams struct {
	Amount int `form:"amount" json:"amount"`
}

func GetReq(c *gin.Context) *gin.H {
	amount := c.DefaultPostForm("amount", "30")
	newAmount, _ := strconv.Atoi(amount)
	req := &gin.H{"amount": newAmount}
	return req
}

func ParseBody(c *gin.Context) ApiParams {
	var p ApiParams
	err := c.BindJSON(&p)
	fmt.Printf("json: %#v\n", p)
	if err != nil {
		return p
	}

	return p
}

func GetAmount(c *gin.Context) int {
	param := ParseBody(c)
	return param.Amount
}
