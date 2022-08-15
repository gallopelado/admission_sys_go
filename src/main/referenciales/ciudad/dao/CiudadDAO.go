package dao

import (
	"admission_sys_go/src/main/db"
	"admission_sys_go/src/main/referenciales/ciudad/dto"
	"database/sql"
	"fmt"
	"log"
)

func ListarPorId(id int) *dto.CiudadDTO {
	ciudad := dto.NewCiudadDTO(0, "")
	var ciu_id int
	var ciu_descripcion string
	querySQL := "SELECT ciu_id, ciu_descripcion FROM public.ciudades WHERE ciu_id = $1"
	con := db.OpenConnectionDatabase()
	row := con.QueryRow(querySQL, id)
	defer con.Close()

	switch err := row.Scan(&ciu_id, &ciu_descripcion); err {
	case sql.ErrNoRows:
		fmt.Println("No hay filas")
	case nil:
		ciudad.SetCiuId(ciu_id)
		ciudad.SetDescripcion(ciu_descripcion)
	default:
		log.Fatal(err)
	}
	return ciudad
}

func ListarTodos() []dto.CiudadDTO {
	ciudades := []dto.CiudadDTO{}
	querySQL := "SELECT ciu_id, ciu_descripcion FROM public.ciudades"
	con := db.OpenConnectionDatabase()
	rows, err := con.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer con.Close()
	for rows.Next() {
		var ciu_id int
		var ciu_descripcion string
		err = rows.Scan(&ciu_id, &ciu_descripcion)
		if err != nil {
			log.Fatal(err)
		}
		ciudades = append(ciudades, *dto.NewCiudadDTO(ciu_id, ciu_descripcion))
	}
	return ciudades
}

func Agregar(descripcion string) bool {
	insertSQL := "INSERT INTO public.ciudades(ciu_descripcion) values($1)"
	con := db.OpenConnectionDatabase()
	_, err := con.Exec(insertSQL, descripcion)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	return true
}

func ModificarPorId(id int, descripcion string) bool {
	updateSQL := "UPDATE public.ciudades SET ciu_descripcion=$1 WHERE ciu_id=$2"
	con := db.OpenConnectionDatabase()
	_, err := con.Exec(updateSQL, descripcion, id)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	return true
}

func EliminarPorId(id int) bool {
	deleteSQl := "DELETE FROM public.ciudades WHERE ciu_id=$1"
	con := db.OpenConnectionDatabase()
	_, err := con.Exec(deleteSQl, id)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	return true
}
