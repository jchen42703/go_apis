package main

import (
	"os"

	"github.com/jchen42703/go-api/internal/auth"
	"github.com/jchen42703/go-api/internal/httprouter"
)

func main() {
	baseKratosUrl := os.Getenv("ORY_KRATOS_BASE_URL")
	clientUrl := os.Getenv("CLIENT_URL")
	oryClient := auth.NewOryAPIClient(baseKratosUrl)
	r := httprouter.New(oryClient, []string{clientUrl})
	// connections, err := db.NewConnections()
	// defer connections.DB.Close()

	v1 := r.Group("/api")
	httprouter.RegisterBaseRoutes(v1)

	serverUrl := os.Getenv("SERVER_URL")
	r.Logger.Fatal(r.Start(serverUrl))
}
