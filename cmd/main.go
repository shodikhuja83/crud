package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/shodikhuja83/http/cmd/app"
	"github.com/shodikhuja83/http/pkg/banners"
)

// func main() {
// 	host := "0.0.0.0"
// 	port := "3939"

// 	if err := execute(host, port); err != nil {
// 		os.Exit(1)
// 	}
// }

// func execute(host string, port string) (err error) {
// 	srv := &http.Server{
// 		Addr:    net.JoinHostPort(host, port),
// 		Handler: &handler{},
// 	}

// 	return srv.ListenAndServe()
// }

// type handler struct{}

// func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte("Hello world!"))
// 	if err != nil {
// 		log.Print(err)
// 	}
// }

// func main() {
// 	host := "0.0.0.0"
// 	port := "3939"

// 	if err := execute(host, port); err != nil {
// 		os.Exit(1)
// 	}
// }

// func execute(host string, port string) (err error) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("bannners.getAll", func(w http.ResponseWriter, r *http.Request) {
// 		_, err := w.Write([]byte("demo data"))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	})

// 	srv := &http.Server{
// 		Addr:    net.JoinHostPort(host, port),
// 		Handler: mux,
// 	}

// 	return srv.ListenAndServe()
// }

func main() {
	host := "0.0.0.0"
	port := "3939"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)
	server.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}

	fmt.Println("Server is listening...")

	return srv.ListenAndServe()
}
