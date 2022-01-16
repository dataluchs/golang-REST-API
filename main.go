package main

import (
	"golang-rest-api/pkg/configs"
	"golang-rest-api/utils"

	"github.com/gorilla/mux"
)

func main() {
	// init router
	router := mux.NewRouter()

	// init server
	server := configs.ServerConfig(router)

	// start api server
	utils.StartServerWithGracefulShutdown(server)
}