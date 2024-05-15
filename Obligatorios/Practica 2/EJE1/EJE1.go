package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		s := fmt.Sprintf("(%v)", l.elem) // convierte l.elem a una cadena con %v, preguntar la diferencia de usar fmt.println
		if Next(l) == nil {
			return s
		}
		return s + "->" + ToString(Next(l))
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

func main() {
	// Abrir el archivo JSON
	file, err := os.Open("ingresantes.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decodificar el archivo JSON en un slice de ingresantes
	var ingresantes []Ingresante
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ingresantes); err != nil {
		panic(err)
	}

	// Convertir el slice de ingresantes en una lista de nodos
	var lista List
	for _, ingresante := range ingresantes {
		PushBack(&lista, ingresante)
	}

	// Imprimir la lista de ingresantes
	fmt.Println(ToString(lista))
}
