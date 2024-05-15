package main

import (
	"fmt"
	"reflect"
)

type Fecha struct {
	dia  int
	mes  int
	anio int
}

type Ingresante struct {
	apellido       string
	nombre         string
	ciudadOrigen   string
	fecha          Fecha
	presentoTitulo bool
	codigoCarrera  string
}

// funciones Stringer para la fecha y para el alumno
func (f Fecha) String() string {
	return fmt.Sprintf("%02d/%02d/%04d", f.dia, f.mes, f.anio)
}

func (i Ingresante) String() string {
	return fmt.Sprintf("Apellido: %s, Nombre: %s, Ciudad de Origen: %s, Fecha de Nacimiento: %s, Presentó Título: %t, Código de Carrera: %s",
		i.apellido, i.nombre, i.ciudadOrigen, i.fecha.String(), i.presentoTitulo, i.codigoCarrera)
}

type Node struct {
	elem    any // interface{}
	next    *Node
}

type List *Node

func New() List {
	return (List(nil));
}

func IsEmpty(self List) bool {
	return (self == nil)
  }

// devuelve la longitud de la lista
func Len(self List) int {
  if (IsEmpty(self)) {
      return 0;
    } else {
    return 1 + Len(Next(self));
  }
}

func FrontElement(self List) any {
	return self.elem
}

func Next(self List) List {
	return self.next
  }

  func ToString(l List) string {
	if IsEmpty(l) {
		return "[]"
	} else {
		s := fmt.Sprintf("(%v)\n", l.elem) // convierte l.elem a una cadena con %v, preguntar la diferencia de usar fmt.println
		if Next(l) == nil {
			return s
		}
		return s + " " + ToString(Next(l))
	}
}


func PushFront(self *List, elem any) {
	aux := new(Node)
	aux.elem = elem
	aux.next = *self
	*self = aux
}

func PushBack(self *List, elem any) {
	if (IsEmpty(*self)) {
		aux := new(Node)
		aux.elem = elem
		aux.next = nil
		*self = aux
	} else {
		PushBack( (*List)(&((*self).next)) ,elem )
	}
}

func Remove(self *List) any {
	elem := (*self).elem
	//aux  := (*self)
	*self = (List)(((*self).next))
	return elem;
}

func Iterate(self List, fp func(int) int) {
	aux := self
	for !IsEmpty(aux) {
		if elem, ok := aux.elem.(int); ok { // Hacemos la aserción de tipo aquí
			aux.elem = fp(elem)
		} else {
			// Manejar el caso en que aux.elem no es un int
			actualType := reflect.TypeOf(aux.elem)
            panic(fmt.Sprintf("aux.elem no es un int, es de tipo %v", actualType))
		}
		aux = Next(aux)
	}
}

func sumoYears(lista List,yearMap map[int]int)  {
		if ingresante,ok := lista.elem.(Ingresante); ok{
			year:= ingresante.fecha.anio
			yearMap[year]++
		}
}

func sumoCarrera(lista List,mapCarrera map[string]int )  {
	if ingresante,ok:=lista.elem.(Ingresante);ok{
		switch ingresante.codigoCarrera{
		case "APU":
			mapCarrera["APU"]++
		case "LI":
			mapCarrera["LI"]++
		case "LS":
			mapCarrera["LS"]++
		}
	}
}

func proceso(lista List, yearMap map[int]int,mapCarrera map[string]int)  {
	aux:=lista
	for !IsEmpty(aux){
		if ingresante, ok := aux.elem.(Ingresante); ok {
			if ingresante.ciudadOrigen == "Bariloche" {
				fmt.Println(ingresante.String())
			}
			sumoYears(aux, yearMap)
			sumoCarrera(aux, mapCarrera)
			
		}
		aux = Next(aux)
	}
}

func main() {
	ingresante1 := Ingresante{
		apellido:       "García",
		nombre:         "Juan",
		ciudadOrigen:   "Buenos Aires",
		fecha:          Fecha{dia: 15, mes: 5, anio: 1990},
		presentoTitulo: true,
		codigoCarrera:  "APU",
	}

	ingresante2 := Ingresante{
		apellido:       "López",
		nombre:         "María",
		ciudadOrigen:   "Córdoba",
		fecha:          Fecha{dia: 20, mes: 10, anio: 1992},
		presentoTitulo: false,
		codigoCarrera:  "LI",
	}

	ingresante3 := Ingresante{
		apellido:       "López",
		nombre:         "María",
		ciudadOrigen:   "Córdoba",
		fecha:          Fecha{dia: 20, mes: 10, anio: 1992},
		presentoTitulo: false,
		codigoCarrera:  "LI",
	}

	ingresante4 := Ingresante{
		apellido:       "Martinez",
		nombre:         "Ana",
		ciudadOrigen:   "Bariloche",
		fecha:          Fecha{dia: 10, mes: 8, anio: 1993},
		presentoTitulo: true,
		codigoCarrera:  "LS",
	}

	ingresante5 := Ingresante{
		apellido:       "Perez",
		nombre:         "Carlos",
		ciudadOrigen:   "Bariloche",
		fecha:          Fecha{dia: 5, mes: 4, anio: 1991},
		presentoTitulo: false,
		codigoCarrera:  "APU",
	}
	var lista List
	mapYear:= make(map[int]int)
	mapCarrera:= map[string]int{
		"APU": 0,
		"LI": 0,
		"LS": 0,
	}

	PushBack(&lista, ingresante1)
	PushBack(&lista, ingresante2)
	PushBack(&lista, ingresante3)
	PushBack(&lista, ingresante4)
	PushBack(&lista, ingresante5)
	
	proceso(lista,mapYear,mapCarrera)
	fmt.Println(" ")
	fmt.Println(mapYear)
	fmt.Println(" ")
	fmt.Println(mapCarrera)
	fmt.Println(" ")
	fmt.Println(lista)


}

