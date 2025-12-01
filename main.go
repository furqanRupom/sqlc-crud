package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/furqanrupom/sqlc-crud/config"
	"github.com/furqanrupom/sqlc-crud/handlers"
	"github.com/furqanrupom/sqlc-crud/repo"
	"github.com/furqanrupom/sqlc-crud/server"
	"github.com/furqanrupom/sqlc-crud/services"
)


func main() {
    conf := config.GetConfig()

    // db connection
    ctx := context.Background()
    dbConnect, db, err := config.DBConnect(ctx, conf.DB)
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }
	if dbConnect == nil {
		log.Print(err)
		os.Exit(1)
	}

    log.Println("Database connected successfully")
    log.Println("Server starting...")

    // repos
    newUserRepo := repo.NewUserRepo(db)

    // services
    newUserService := services.NewUserService(newUserRepo)

    // handlers
    newUserHandler := handlers.NewUserHandler(newUserService)

    // server (chi router)
    newServer := server.NewServer(conf, newUserHandler)
	r := newServer.Start()

    // add root route
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Server is running on %d", conf.HttpPort)
    })

    // start server
    connStr := fmt.Sprintf(":%d", conf.HttpPort)
    log.Println("Listening on", connStr)
    http.ListenAndServe(connStr, r)
}
