package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServerWithGracefulShutdown(server *http.Server) {
	
	var wait time.Duration

	go func() {
		log.Println("Server is starting ...")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<- c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	log.Println("Server is shutting down ...")
}

