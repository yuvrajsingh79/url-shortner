# URL Shortener Service

## Overview

This is a simple URL shortener service built in Golang that provides a RESTful API for shortening URLs. It allows you to convert long, unwieldy URLs into short, easy-to-share links.

## Features

- Shorten a long URL into a concise and memorable short link.
- Redirect users from the short link to the original URL.
- Store URL mappings in memory.
- Utilize Redis for caching.

## Prerequisites

Before running the URL shortener service, make sure you have the following prerequisites installed:

- Go: The Go programming language.
- Docker: If you plan to use the provided Dockerfile for containerization.
- Redis: If you choose to use Redis for caching (optional).

## Getting Started

Follow these steps to set up and run the URL shortener service:

1. **Clone the Repository:**

```
   git clone https://github.com/yuvrajsingh79/url-shortner.git
   cd url-shortener
```

2. **Configuration:**

     If you plan to use Redis for caching, update the Redis server address in the cache.go file:

```
   const redisAddr = "my-redis-container:6379"
```
Ensure that `my-redis-container` matches the name of your Redis container.

3. **Build and Run the Application:**

	To start the service, use the provided `setup.sh` script, which sets up a Redis container (if needed), make `setup.sh` as executable, builds the Go application, and runs it.
```
   chmod +x setup.sh
   cd url-shortener
```

4. **Perform URL Shortening & verify Redirection:**
   
   **Shorten a URL** :
	
   Use the service to shorten a URL by sending a POST request to the /shorten endpoint. Replace <your-original-url> with the URL you want to shorten.
	
```
  curl -X POST -H "Content-Type: application/json" -d '{"url": "https://www.example.com"}' http://localhost:8080/shorten
```
   Response will look like :
    
```
 {"short_url":"https://example.com/78dbdafd-1767-4832-ae45-f730991cb2f0"}
```

   `78dbdafd-1767-4832-ae45-f730991cb2f0` is the short url that we store in memory.
   Note: It does not allow duplicates.
	
   **Test Redirection** :
	
    To verify the redirection functionality, access the shortened URL in your web browser or use curl:
	
```
  curl -i http://localhost:8080/78dbdafd-1767-4832-ae45-f730991cb2f0
```
  Response : 
```
HTTP/1.1 303 See Other
Content-Type: text/html; charset=utf-8
Location: https://www.example.com
Date: Sun, 08 Oct 2023 19:25:00 GMT
Content-Length: 50
<a href="https://www.example.com">See Other</a>
```

	`https://www.example.com` is the original URL that we retrived.

5. **Testing:**

Run the tests to ensure the correctness of the service:
```
go test ./...
```

## Project Structure

The project directory structure is organized as follows:

- `main.go`: The main entry point of the application.
- `pkg/controller`: Contains the core logic for URL shortening.
  - `shortener.go`: Implements the URL shortening service.
  - `shortener_test.go`: Tests for the URL shortening logic.
  - `storage.go`: Defines the storage interface and provides an in-memory storage implementation.
  - `storage_test.go`: Tests for the storage functionality.
  - `cache.go`: Defines the cache interface and provides a Redis-based cache implementation.
  - `cache_test.go`: Tests for the caching functionality.
  -  `mock.go`: Provides mock implementations for testing purposes.
- `api`: Contains the HTTP API handlers and routes.
  - `handlers.go`: Implements HTTP request handlers.
  - `handler_test.go`: Tests for the http handlers.
  - `routes.go`: Configures API routes.
- `Dockerfile`: Defines a Docker image for the application.
- `go.mod`: The Go module file specifying project dependencies.
- `setup.sh`: A script for setting up the project and starting the application
- `README.md`: This file, providing project documentation.

## Contributing

Contributions to this project are welcome. Feel free to open issues or pull requests for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
