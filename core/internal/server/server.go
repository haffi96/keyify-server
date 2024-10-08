package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"apikeyper/internal/database"
	"apikeyper/internal/events"
	"apikeyper/internal/ratelimit"
)

type Server struct {
	Db          database.Service
	Message     events.MessageService
	RateLimiter *ratelimit.RateLimitService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		Db:          database.New(),
		Message:     events.New(),
		RateLimiter: ratelimit.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
