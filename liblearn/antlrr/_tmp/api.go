package apigen

import (
	"log"
	"time"
)

type Teacher struct {
	id   int64
	name string
}

func (receiver Teacher) GetId() int64 {
	return receiver.id
}

func (receiver Teacher) GetName() string {
	return receiver.name
}

func TTT() {
	log.Println("当前时间:", time.Now().Format("2006-01-02 15:03:04"))
}
