package routes

import (
	"net/http"
	"ticket_sys_api/controllers"
)

func InitializeRoutes() {
	http.HandleFunc("/clientes", controllers.GetClientes)
	http.HandleFunc("/usuarios", controllers.GetUsuarios)
	http.HandleFunc("/departamentos", controllers.GetDepartamentos)
}
