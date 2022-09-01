package main

import (
	"fmt"
	"holy-go-lib/lib/resourcemanager/manager"
)

var connManager = manager.NewResourceManager()

// manager并行
func main() {
	fmt.Printf("manager:%+v\n", connManager)
}
