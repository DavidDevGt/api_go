package main

import (
	"log"
	"net/http"
	"ticket_sys_api/routes"
)

func main() {
	routes.InitializeRoutes()
	log.Println("Servidor iniciado en el puerto :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
