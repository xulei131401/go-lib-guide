package main

import (
	"context"
	"fmt"
	"holy-go-lib/libother/db/redis/go-redis/common"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	rdb := common.Rdb()
	defer rdb.Close()

	if err := rdb.RPush(ctx, "queue", "message").Err(); err != nil {
		panic(err)
	}

	fmt.Println("关闭")
}
