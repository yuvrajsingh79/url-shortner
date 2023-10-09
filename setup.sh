#!/bin/bash

# Step 1: Create a Docker network
docker network create my-network

# Step 2: Start a Redis container on the network
docker run --name my-redis-container -d --network my-network -p 6379:6379 redis

# Step 3: Build the Go application Docker image
docker build -t url-shortener .

# Step 4: Run the Go application Docker container on the network
docker run -p 8080:8080 --network my-network url-shortener
