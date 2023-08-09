package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nikitamirzani323/wl_super_backend_frontend/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	app := routers.Init()
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3011"
		}

		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}
