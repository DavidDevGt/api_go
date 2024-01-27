package models

import (
	"database/sql"
	"time"
)

type Usuario struct {
	ID              int       `json:"id"`
	Nombre          string    `json:"nombre"`
	Correo          string    `json:"correo"`
	Contrasena      string    `json:"contrasena"`
	RolID           int       `json:"rol_id"`
	FechaCreacion   time.Time `json:"fecha_creacion"`
	UsuarioCreacion int       `json:"usuario_creacion_id"`
	FechaEdicion    time.Time `json:"fecha_edicion,omitempty"`
	UsuarioEdicion  int       `json:"usuario_edicion_id,omitempty"`
	Active          bool      `json:"active"`
}

func GetUsuarios(db *sql.DB) ([]Usuario, error) {
	rows, err := db.Query("SELECT id, nombre, correo, contrasena, rol_id, fecha_creacion, usuario_creacion_id, fecha_edicion, usuario_edicion_id, active FROM usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nombre, &u.Correo, &u.Contrasena, &u.RolID, &u.FechaCreacion, &u.UsuarioCreacion, &u.FechaEdicion, &u.UsuarioEdicion, &u.Active); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}
