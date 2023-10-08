package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/yuvrajsingh79/url-shortner/api"
	"github.com/yuvrajsingh79/url-shortner/pkg/controller"
)

var httpAddr = flag.String("http.addr", ":8080", "HTTP Listen addresss")

func main() {
	//parse the flag
	flag.Parse()
	// Initialize the URL shortener with in-memory storage and Redis cache.
	storage := controller.NewMemoryStorage()
	redisCache, err := controller.NewRedisCache("my-redis-container:6379")
	if err != nil {
		log.Fatalf("Failed to initialize Redis cache: %v", err)
	}
	shortener := controller.NewShortener(storage, redisCache)

	// Create API routes and start the HTTP server.
	router := api.SetupRoutes(shortener)

	http.Handle("/", router)

	log.Printf("Server is running on port %s...", *httpAddr)
	err = http.ListenAndServe(*httpAddr, nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
