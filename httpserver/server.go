package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
import "github.com/gorilla/mux"

type Server interface{
	Run()
	SetReading(sensor int, on bool)
}

type server struct {
	router *mux.Router
	sensors map[int]bool
}

var currentReadings *map[int]bool
var currentError *error

func New() Server {
	s := server{
		router: mux.NewRouter(),
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	index := fmt.Sprintf("%s/static/",dir)

	s.router.Handle("/", http.FileServer(http.Dir(index)))

	js := fmt.Sprintf("%s/static/js/build/static/",dir)
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(js))))

	s.router.HandleFunc("/api/readings.json", ReadingsHandler)

	return &s
}

func (s *server) Run() {
	//http.Handle("/", s.router)

	var wait = time.Second * 15

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: s.router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func (s *server) SetReading(sensor int, on bool) {
	s.sensors[sensor] = on
}

func ReadingsHandler(rw http.ResponseWriter, r *http.Request) {
	var d =  struct {
		Readings map[int]bool
		Error error
	} {
		Readings: *currentReadings,
		Error: *currentError,
	}

	responseString, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	rw.Write(responseString)
}