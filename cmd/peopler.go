package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmaterer/peopler/config"
	"github.com/pmaterer/peopler/internal/sqlite"
	"github.com/pmaterer/peopler/user/controller"
	"github.com/pmaterer/peopler/user/repository"
	"github.com/pmaterer/peopler/user/service"
)

func main() {
	cnf := config.Config{
		Server: config.Server{
			ListenAddress: "127.0.0.1",
			ListenPort:    8721,
		},
	}

	db, err := sqlite.NewSQLiteHandler("./peopler.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	userRepo := repository.NewRepository(db)
	userService := service.NewService(userRepo)
	userController := controller.NewController(userService)

	router := mux.NewRouter()
	router.HandleFunc("/user", userController.CreateUser()).Methods("POST")
	router.HandleFunc("/users", userController.GetAllUsers()).Methods("GET")
	router.HandleFunc("/user/{id}", userController.GetUser()).Methods("GET")
	router.HandleFunc("/user/{id}", userController.UpdateUser()).Methods("PUT")
	router.HandleFunc("/user/{id}", userController.DeleteUser()).Methods("DELETE")

	log.Printf("Starting server on %s:%d\n", cnf.Server.ListenAddress, cnf.Server.ListenPort)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf("%s:%d", cnf.Server.ListenAddress, cnf.Server.ListenPort),
		router))
}
