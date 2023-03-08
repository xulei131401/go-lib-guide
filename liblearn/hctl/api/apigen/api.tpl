package apigen

import (
	"log"
	"time"
)

type {{.Name}} struct {
	id   int64
	name string
}

func (receiver {{.Name}}) GetId() int64 {
	return receiver.id
}

func (receiver {{.Name}}) GetName() string {
	return receiver.name
}

func TTT() {
	log.Println("当前时间:", time.Now().Format("2006-01-02 15:03:04"))
}
