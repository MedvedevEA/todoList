package main

import (
	"embed"
	"log"
	"todolist/internal/apiserver"
	"todolist/internal/config"
	"todolist/internal/service"
	"todolist/internal/store"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	//config
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	//store
	store, err := store.New(config.DatabaseConnectString, embedMigrations)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()
	//service
	service := service.New(store)
	//apiserver
	log.Printf("API Server 'Todo List' is started in addr:[%s]", config.BindAddress)
	apiServer := apiserver.New(config.BindAddress, service)
	if err := apiServer.Run(); err != nil {
		log.Fatalf("API Server 'Todo List' error: %s", err)
	}
	log.Printf("API Server 'Todo list' is stoped")

}
