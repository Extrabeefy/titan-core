package main

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/Extrabeefy/titan_core/lib/config"
	"github.com/Extrabeefy/titan_core/server/auth"
	"github.com/Extrabeefy/titan_core/server/world"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.DebugLevel)

	// Load the world config and object manager.
	// TODO(jeshua): make this load from a command-line flag maybe?
	config := config.NewConfigFromJSON("world.json")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		auth.RunAuthServer(5000, config)
	}()

	go func() {
		defer wg.Done()
		world.RunWorldServer("Sydney", 5001, config)
	}()

	wg.Wait()
}
