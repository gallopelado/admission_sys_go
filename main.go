package main

import (
	"admission_sys_go/src/main/referenciales/ciudad/dao"
	"fmt"
	s "strings"
)

/*
Ejemplo de enums:
Esto es muy parecido a los enums en Java pero con implementación diferente.
Se crea un tipo, en este caso OPERACION, a su vez se definen las difentes opciones.
*/
type OPERACION int

const (
	AGREGAR OPERACION = iota
	MODIFICAR
	ELIMINAR
	LISTARTODOS
	LISTARPORID
)

func procesar(op OPERACION) {
	switch op {
	case AGREGAR:
		resultado := dao.Agregar(s.ToUpper("ñemby"))
		if resultado {
			fmt.Println("Agregado, puntos logrados :)")
		}
	case MODIFICAR:
		resultado := dao.ModificarPorId(11, s.ToUpper("san antonio"))
		if resultado {
			fmt.Println("Modificado, puntos logrados :)")
		}
	case ELIMINAR:
		resultado := dao.EliminarPorId(2)
		if resultado {
			fmt.Println("Eliminado, puntos logrados :)")
		}
	case LISTARPORID:
		fmt.Println(*dao.ListarPorId(12))
	case LISTARTODOS:
		fmt.Println(dao.ListarTodos())
	default:
		fmt.Println("No existe esa operación")
	}
}

// Ejecutar con go run *.go
func main() {
	procesar(LISTARTODOS)
}
