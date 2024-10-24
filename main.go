package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

var NS NameStore

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	query := r.URL.Query()

	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}

	log.Printf("Received %s request for %s\n", method, name)

	var response string
	var status int
	var num int
	switch method {
	case http.MethodGet:
		num = NS.GetName(name)
		response = fmt.Sprintf("Have <%d> entries for <%s>\n", num, name)
		status = http.StatusOK
	case http.MethodPut:
		num = NS.PutName(name)
		response = fmt.Sprintf("Have <%d> entries for <%s>\n", num, name)
		status = http.StatusOK
	case http.MethodDelete:
		num = NS.DeleteName(name)
		response = fmt.Sprintf("Have <%d> entries for <%s>\n", num, name)
		status = http.StatusOK
	default:
		log.Printf("Received unsupported %s request for %s\n", method, name)
		response = fmt.Sprintf("Unsupported request: <%s>\n", method)
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal
	<-interruptChan

	// Deadline for waiting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

// To run the app
// Run `go build` in project dir, ../../go-k8
// Run `./go-k8`
//
// To Print only the status code
// curl -s -o /dev/null -w "%{http_code}" -X GET 'localhost:8080?name=Hasti'
//
// To print out only the response
// curl -s -X PUT 'localhost:8080?name=Hasti'
func main() {
	// Init NameStore
	NS = make(map[string]int)

	// Create server and route handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start server
	go func() {
		log.Println("Starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	waitForShutdown(srv)
}
