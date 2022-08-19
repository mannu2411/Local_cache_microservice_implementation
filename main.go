package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

type Server struct {
	CP CacheProvider
}

func InitService() *Server {
	cacheProvider := NewCacheProvider(time.Minute * 2)
	return &Server{
		CP: cacheProvider,
	}
}

func main() {
	srv := InitService()
	router := chi.NewRouter()
	router.Get("/add", srv.Add)
	router.Get("/data", srv.Data)
	httpPort := "8081"
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
