package main

import (
	"fmt"
	"reflect"
)

func main() {
	// lista de ejemplo
	list := New()
	PushBack(&list, 1)
	PushBack(&list, 2)
	PushBack(&list, 3)
	PushBack(&list, "string") // elemento no entero para probar el manejo de errores
	PushBack(&list, 4)

	fmt.Println("Lista original:")
	fmt.Println(ToString(list))

	// Aplicar la función Iterate con una función de incremento
	Iterate(list, func(x int) int {
		return x + 1
	})

	fmt.Println("Lista después de aplicar Iterate:")
	fmt.Println(ToString(list))
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
  

  
  