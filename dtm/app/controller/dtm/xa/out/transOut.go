package out

import (
	"github.com/gin-gonic/gin"
)

var amount = 100

func TransOutAction (c *gin.Context) {
	//router.XaClient.XaLocalTransaction(c.Request.URL.Query(), func(db *sql.DB, xa *dtmcli.Xa) (interface{}, error) {
	//	amount := common.GetAmount(c)
	//	_, err := dtmcli.ec(db, "update dtm_busi.user_account set balance=balance-? where user_id=?", amount, 1)
	//	fmt.Println("----TransOutXa执行完毕----")
	//	time.Sleep(time.Second * 5)
	//	return dtmcli.MapSuccess, err
	//})
}