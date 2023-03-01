package main

import (
	"log"

	"github.com/LanternOfDarkness/SoftceryGolang/internal/config"
	queue "github.com/LanternOfDarkness/SoftceryGolang/internal/queue"
	store "github.com/LanternOfDarkness/SoftceryGolang/internal/storage"
	"github.com/LanternOfDarkness/SoftceryGolang/pkg/http"
)


func main() {
	// Get config
  c := config.NewConfig()
	
	// Init storage
	store.Storage = store.NewStorage(c.FileStoragePath)
	if err := store.Storage.LoadIdFromFiles(); err != nil {
		log.Fatal("Failed to load image id from files:", err)
	}
	
	// Run queue reciever
	go queue.RunReciever()

	// Init router
	r := http.SetupRouter()
	if err := r.Run(c.BindAddr); err != nil {
    log.Fatal("Failed to start server:", err)
  }

}

