package main

import (
	"context"
	"log"
	"os"

	redisClient "webhook/redis"

	"github.com/go-redis/redis/v8" // Make sure to use the correct version
)

func main() {
	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize the Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"), // Use an environment variable to set the address
		Password: "",                         // No password
		DB:       0,                          // Default DB
	})

	// Subscribe to the "transactions" channel
	err := redisClient.Subscribe(ctx, client)

	if err != nil {
		log.Println("Error:", err)
	}

}
