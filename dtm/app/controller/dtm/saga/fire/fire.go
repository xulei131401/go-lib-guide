package fire

import (
	"github.com/gin-gonic/gin"
	"holygin/app/controller/dtm/common"

	"github.com/dtm-labs/dtm/dtmcli"
)

/**
问题：
1.RM接口的返回值，dtm没有暴露给AP，AP只知道失败了，不知道具体的返回信息
2.其中一个正向操作失败，补偿操作都会被调用，可能造成 补偿多余。。。
3.dtm不支持其他分支的结果作为输入
4.支持没有补偿操作，但是必须保证正向操作不能失败，否则没有意义
5.并发开启的话就不能保证顺序了

适用场景：1.各个分支事务比较独立 2.最好不要出现后边的分支依赖前边分支的结果，如果依赖应该单独提供接口，dtm不支持

 */



// DoAction 入口请求，调用分布式
func DoAction(c *gin.Context) {
	req := &gin.H{"amount": 30} // 微服务的载荷
	gid := dtmcli.MustGenGid(common.DefaultHTTPServer)
	saga := dtmcli.NewSaga(common.DefaultHTTPServer, gid).
		Add(common.SagaTransOutAction, common.SagaTransOutRevokeAction, req).
		Add(common.SagaTransInAction, common.SagaTransInRevokeAction, req).
		EnableConcurrent()

	saga.WaitResult = true
	err := saga.Submit()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"gid":gid})
	return
}