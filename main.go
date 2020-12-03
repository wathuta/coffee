package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wathuta/rest-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product-api", log.LstdFlags)
	ph := handlers.NewProduct(l)

	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/get", ph.GetProduct)
	getRouter.HandleFunc("/get/{id:[0-9]+}", ph.GetSingleProduct)

	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/create", ph.CreateProduct)
	postRouter.Use(ph.MiddlewareFunc)

	putRouter := mux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/update/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareFunc)

	deleteRouter := mux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/delete/{id:[0-9]+}", ph.DeleteProduct)

	//adding a rich documentation ui to my api using redoc on the get request
	//redoc has a handler of serving my swagger spec directly from my go code

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	//the middleware package enables us to host our own documentation website within our api
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := http.Server{
		Addr:    ":8080",
		Handler: ch(mux),
	}

	go func() {
		l.Println("starting the server on port:8080")
		log.Fatal(s.ListenAndServe())
	}()
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	c := <-sigChan
	fmt.Println("graceful shutdown ", c)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}
