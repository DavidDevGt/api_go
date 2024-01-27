package models

import (
	"database/sql"
	"time"
)

type Cliente struct {
	ID              int       `json:"id"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion,omitempty"`
	Telefono        string    `json:"telefono,omitempty"`
	FechaCreacion   time.Time `json:"fecha_creacion"`
	UsuarioCreacion int       `json:"usuario_creacion_id"`
	FechaEdicion    time.Time `json:"fecha_edicion,omitempty"`
	UsuarioEdicion  int       `json:"usuario_edicion_id,omitempty"`
	Active          bool      `json:"active"`
}

func GetClientes(db *sql.DB) ([]Cliente, error) {
	rows, err := db.Query("SELECT id, nombre, direccion, telefono, fecha_creacion, usuario_creacion_id, fecha_edicion, usuario_edicion_id, active FROM cliente")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []Cliente
	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.ID, &c.Nombre, &c.Direccion, &c.Telefono, &c.FechaCreacion, &c.UsuarioCreacion, &c.FechaEdicion, &c.UsuarioEdicion, &c.Active); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}
