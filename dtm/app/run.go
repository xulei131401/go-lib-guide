package app

import (
	"fmt"
	"holygin/app/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/fvbock/endless"
)


func registerRouter() *gin.Engine {
	r := gin.New()
	router.InitRouter(r)
	return r
}

func NormalServer(){
	server := &http.Server{
		Addr:           ":8001",
		Handler:        registerRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("http.Serve error(%v)", err)
		panic(err)
	}
}

func _graceServer() {
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", 8777)

	server := endless.NewServer(endPoint, registerRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		// save it somehow
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server err: %v", err)
	}
}

func signalHandle() {
	var (
		ch = make(chan os.Signal, 1)
	)

	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			log.Printf("get a signal %s, stop the admin process", si.String())

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
