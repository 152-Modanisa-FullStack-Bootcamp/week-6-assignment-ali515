package main

import (
	"bootcamp/config"
	"bootcamp/controller"
	"bootcamp/handler"
	"bootcamp/repository"
	"bootcamp/service"
	"fmt"
	"net/http"
)

func main() {
	config.Initialize()
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	controller := controller.NewController(handler)
	http.HandleFunc("/", controller.Routing)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
