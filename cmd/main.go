package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/hinupurthakur/swapi/config"
	_ "github.com/hinupurthakur/swapi/logging"
	_ "github.com/hinupurthakur/swapi/db"
	"github.com/hinupurthakur/swapi/routes"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Println("initizalizing the server ")
}

func main() {
	port := os.Getenv("SERVER_PORT")
	r := routes.CreateRoutes()
	log.Infof("starting server listening on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Errorln("ListenAndServe Errors:", err)
	}
}
