package controllers

import (
	"encoding/json"
	"net/http"
	"ticket_sys_api/models"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := models.ConectarDB()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error al conectar con la base de datos")
		return
	}
	defer db.Close()

	usuarios, err := models.GetUsuarios(db)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error al obtener los usuarios")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}
