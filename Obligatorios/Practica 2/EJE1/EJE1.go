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
	elem any 
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

func New() List {
	return List{nil, nil}
}

func IsEmpty(self List) bool {
	return self.head == nil
}

// devuelve la longitud de la lista
func Len(self List) int {
	length := 0
	for current := self.head; current != nil; current = current.next {
		length++
	}
	return length
}

func FrontElement(self List) any {
	if IsEmpty(self) {
		return nil
	}
	return self.head.elem
}

func Next(self List) *Node {
	if IsEmpty(self) {
		return nil
	}
	return self.head.next
}

func ToString(l List) string {
	if IsEmpty(l) {
		return "[]"
	}
	result := ""
	for current := l.head; current != nil; current = current.next {
		result += fmt.Sprintf("(%v)\n", current.elem)
	}
	return result
}

func PushFront(self *List, elem any) {
	aux := &Node{elem: elem, next: self.head, prev: nil}
	if IsEmpty(*self) {
		self.tail = aux
	} else {
		self.head.prev = aux
	}
	self.head = aux
}

func PushBack(self *List, elem any) {
	aux := &Node{elem: elem, next: nil, prev: self.tail}
	if IsEmpty(*self) {
		self.head = aux
	} else {
		self.tail.next = aux
	}
	self.tail = aux
}

func Remove(self *List, actual *Node) {
	if actual == nil || IsEmpty(*self) {
		return
	}
	if actual == self.head {
		self.head = actual.next
		if self.head != nil {
			self.head.prev = nil
		} else {
			self.tail = nil
		}
	} else if actual == self.tail {
		self.tail = actual.prev
		if self.tail != nil {
			self.tail.next = nil
		} else {
			self.head = nil
		}
	} else {
		actual.prev.next = actual.next
		actual.next.prev = actual.prev
	}
}

func Iterate(self List, fp func(int) int) {
	for current := self.head; current != nil; current = current.next {
		if elem, ok := current.elem.(int); ok {
			current.elem = fp(elem)
		} else {
			actualType := reflect.TypeOf(current.elem)
			panic(fmt.Sprintf("current.elem no es un int, es de tipo %v", actualType))
		}
	}
}

func sumoYears(lista *Node, yearMap map[int]int) {
	if ingresante, ok := lista.elem.(Ingresante); ok {
		year := ingresante.fecha.anio
		yearMap[year]++
	}
}

func sumoCarrera(lista *Node, mapCarrera map[string]int) {
	if ingresante, ok := lista.elem.(Ingresante); ok {
		switch ingresante.codigoCarrera {
		case "APU":
			mapCarrera["APU"]++
		case "LI":
			mapCarrera["LI"]++
		case "LS":
			mapCarrera["LS"]++
		}
	}
}

func proceso(lista *List, yearMap map[int]int, mapCarrera map[string]int) {
	for current := lista.head; current != nil; current = current.next {
		if ingresante, ok := current.elem.(Ingresante); ok {
			if ingresante.ciudadOrigen == "Bariloche" {
				fmt.Println(ingresante.String())
			}
			sumoYears(current, yearMap)
			sumoCarrera(current, mapCarrera)
			if !ingresante.presentoTitulo {
				toRemove := current
				current = current.prev //pongo como actual al aterior
				Remove(lista, toRemove)
			}
		}
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
		apellido:       "Vergara",
		nombre:         "Jose Maria",
		ciudadOrigen:   "Santa Fé",
		fecha:          Fecha{dia: 1, mes: 12, anio: 2003},
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
		fecha:          Fecha{dia: 5, mes: 4, anio: 1993},
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
	
	proceso(&lista,mapYear,mapCarrera)
	fmt.Println(" ")
	fmt.Println(mapYear)
	fmt.Println(" ")
	fmt.Println(mapCarrera)
	fmt.Println(" ")
	fmt.Println(ToString(lista))


}

