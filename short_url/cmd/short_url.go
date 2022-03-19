package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shor_url/global"
	"shor_url/router"
)

func main() {
	global.InitGlobal()
	echoServer := router.NewServer()

	err := echoServer.Start(":8080")
	if err != nil {
		log.Fatalln("start", err)
	}
	/*server := &http.Server{
		Addr:           ":8080",
		Handler:        echoServer,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()

	gracefulExitWeb(server)*/
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	log.Println("got a signal", sig)
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("shutdown error", err)
	}

	log.Println("exited,exited cost", time.Since(now))
}
