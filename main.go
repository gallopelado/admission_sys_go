package main

import (
	"admission_sys_go/src/main/db"
	"database/sql"
	"fmt"
)

func main() {
	// ciudad := dto.CiudadDTO{}
	//ciudad.SetCiuId(1)
	//ciudad.SetDescripcion("Villa Elisa")
	//fmt.Println(ciudad)
	getCiudades()
}

func getCiudades() {
	//ciudad := dto.CiudadDTO{}
	querySQL := "SELECT * FROM ciudades"
	db := db.OpenConnectionDatabase()
	rows := db.QueryRow(querySQL)

	var ciu_id sql.NullInt16
	var ciu_descripcion sql.NullString

	switch err := rows.Scan(&ciu_id, &ciu_descripcion); err {
	case sql.ErrNoRows:
		fmt.Println("No hay filas")
	case nil:
		//return ciu_id + " - " + ciu_descripcion
		//ciudad.SetCiuId(ciu_id)
		fmt.Println(rows)
	}
}
