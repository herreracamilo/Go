package main

import (
	"fmt"
	"sync"
)

type Contact struct {
	Nombre          string
	Apellido        string
	CorreoElectronico string
	Telefono        string
}

type Agenda struct {
	contactos map[string]Contact
	mutex     sync.Mutex
}

func NuevaAgenda() *Agenda {
	return &Agenda{
		contactos: make(map[string]Contact),
	}
}
// uso el mutex para garantiza que solo una goroutine pueda acceder y modificar la estructura de datos contactos a la vez.
func (a *Agenda) AgregarContacto(contacto Contact) { // agregarContacto agrega un nuevo contacto a la agenda
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.contactos[contacto.CorreoElectronico] = contacto
}


func (a *Agenda) EliminarContacto(correo string) { // elimina un contacto de la agenda dado su correo electrónico
	a.mutex.Lock()
	defer a.mutex.Unlock()
	delete(a.contactos, correo)
}


func (a *Agenda) BuscarContacto(correo string) (Contact, bool) { //  busca y devuelve un contacto dado su correo electrónico
	a.mutex.Lock()
	defer a.mutex.Unlock()
	contacto, existe := a.contactos[correo]
	return contacto, existe
}

func main() {
	agenda := NuevaAgenda() // creo una agenda

	var wg sync.WaitGroup

	// creo y agrego contactos a la agenda
	wg.Add(1)
	go func() {
		defer wg.Done()
		agenda.AgregarContacto(Contact{"santaigo", "herrera", "santiago@mail.com", "335234515"})
		agenda.AgregarContacto(Contact{"camilo", "herrera", "camilo@mail.com", "22453287"})
		agenda.AgregarContacto(Contact{"leo", "herrera", "leo@mail.com", "6463445566"})
	}()

	// elimino un contacto
	wg.Add(1)
	go func() {
		defer wg.Done()
		agenda.EliminarContacto("santiago@mail.com")
	}()

	// busco un contacto
	wg.Add(1)
	go func() {
		defer wg.Done()
		contacto, existe := agenda.BuscarContacto("leo@mail.com")
		if existe {
			fmt.Printf("el contacto fue encontrado: %v\n", contacto)
		} else {
			fmt.Println("el contacto no fue encontrado")
		}
	}()

	// espero que todas las goroutines terminen
	wg.Wait()

	// imprimo la agenda final
	fmt.Println("Agenda final:")
	for _, contacto := range agenda.contactos {
		fmt.Println(contacto)
	}
}
