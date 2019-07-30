package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/anilk1sagar/go_base_backend/src/utils"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello Api</h1>")
}

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/", index)

	server := &http.Server{
		Addr: "127.0.0.1:9110",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("Server Started on port: 9110")
		log.Fatal(server.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	// log.Println("Shutting down Server.")
	utils.Logger().Errorln("Shutting down Server.")
	os.Exit(0)
}
