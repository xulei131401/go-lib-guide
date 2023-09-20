package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"log"
)

type OrdersService struct {
}

func (this *OrdersService) NewOrder(ctx context.Context, orderReq *OrderRequest) (*OrderResponse, error) {
	orderMain := orderReq.OrderMain
	if err := orderMain.Validate(); err != nil {
		return &OrderResponse{
			Status: "error",
			Message: err.Error(),
		},
		nil
	}

	if orderMain.OrderTime != nil {
		orderTime, err := ptypes.Timestamp(orderMain.OrderTime)
		if err != nil {
			log.Fatal(err)
		}
		timeGo := orderTime.Format("2006-01-02 15:04:05")
		fmt.Println("goTime:", timeGo)
		fmt.Println("orderTime:", orderTime)
		fmt.Printf("server收到订单信息:订单号:%v, 订单金额:%v,订单时间:%v, 子订单:%v\n", orderMain.OrderNo, orderMain.OrderMoney, timeGo, orderMain.OrderDetails)
	} else {
		fmt.Printf("server收到订单信息:订单号:%v, 订单金额:%v, 子订单:%v\n", orderMain.OrderNo, orderMain.OrderMoney, orderMain.OrderDetails)
	}
	return &OrderResponse{Status: "OK", Message: "success",}, nil
}
