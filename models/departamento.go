package models

import (
	"database/sql"
)

type Departamento struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func GetDepartamentos(db *sql.DB) ([]Departamento, error) {
	query := "SELECT id, nombre FROM departamento"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departamentos []Departamento
	for rows.Next() {
		var d Departamento
		if err := rows.Scan(&d.ID, &d.Nombre); err != nil {
			return nil, err
		}
		departamentos = append(departamentos, d)
	}

	return departamentos, nil
}
