package main

import (
	"interview/url-shortner/api"
	"interview/url-shortner/controller"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize the URL shortener with in-memory storage and Redis cache.
	storage := controller.NewMemoryStorage()
	redisCache, err := controller.NewRedisCache("redis://localhost:6379")
	if err != nil {
		log.Fatalf("Failed to initialize Redis cache: %v", err)
	}
	shortener := controller.NewShortener(storage, redisCache)

	// Create API routes and start the HTTP server.
	router := api.SetupRoutes(shortener)

	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s...", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
