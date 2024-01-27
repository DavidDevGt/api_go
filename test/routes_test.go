package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticket_sys_api/routes"
)

func TestDepartamentosRoute(t *testing.T) {
	// Inicializar las rutas
	routes.InitializeRoutes()

	// Crear un servidor de prueba usando DefaultServeMux
	server := httptest.NewServer(nil) // Aquí pasamos nil para usar el DefaultServeMux
	defer server.Close()

	// Realizar una solicitud GET al servidor de prueba
	res, err := http.Get(server.URL + "/departamentos")
	if err != nil {
		t.Errorf("Error al realizar la solicitud: %v", err)
	}

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()	
	if err != nil {
		t.Errorf("Error al leer la respuesta: %v", err)
	}

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	if res.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado 200, se obtuvo: %d", res.StatusCode)
	}

	if len(body) == 0 {
		t.Error("Se esperaba un cuerpo de respuesta no vacío")
	}
}
