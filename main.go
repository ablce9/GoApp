package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ablce9/go-assignment/adapters/http"
	"github.com/ablce9/go-assignment/engine"
	"github.com/ablce9/go-assignment/providers/database"
)

const (
	dbAddr     = "db:5432"
	dbUser     = "postgres"
	dbPassword = "bad-password"
	dbDatabase = "go_assignment_dev"
)

func main() {
	provider := database.NewProvider(
		dbAddr,
		dbUser,
		dbPassword,
		dbDatabase,
	)
	e := engine.NewEngine(provider)

	adapter := http.NewAdapter(e)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
	provider.Close()
}
