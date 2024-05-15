package main

import (
	"fmt"
	"sort"
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

// ordenamientos

// por edad
func (i Ingresante) Edad() int {
	return i.fecha.anio
}

type PorEdad []Ingresante

// Len implements sort.Interface.
func (a PorEdad) Len() int {
	return len(a)
}

func (a PorEdad) Less(i, j int) bool {
	return a[i].Edad() < a[j].Edad()
}

func (a PorEdad) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// por nombre
func (i Ingresante) NombreCompleto() string {
	return i.apellido + i.nombre
}

type PorNombre []Ingresante

// Len implements sort.Interface.
func (a PorNombre) Len() int {
	return len(a)
}

func (a PorNombre) Less(i, j int) bool {
	return a[i].NombreCompleto() < a[j].NombreCompleto()
}

func (a PorNombre) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	fechaNacimiento := Fecha{
		dia:  7,
		mes:  9,
		anio: 2001,
	}
	fechaNacimiento2 := Fecha{
		dia:  7,
		mes:  9,
		anio: 1902,
	}

	alumno := Ingresante{
		apellido:       "camilo",
		nombre:         "herrera",
		ciudadOrigen:   "Buenos Aires",
		fecha:          fechaNacimiento,
		presentoTitulo: true,
		codigoCarrera:  "LI",
	}
	fmt.Println(" ")
	fmt.Println(alumno)
	fmt.Println(" ")

	ingresantes := []Ingresante{
		{apellido: "Zeta", nombre: "Gama", ciudadOrigen: "Lima", fecha: fechaNacimiento, presentoTitulo: true, codigoCarrera: "LI"},
		{apellido: "Beta", nombre: "Alfa", ciudadOrigen: "Bogotá", fecha: fechaNacimiento2, presentoTitulo: false, codigoCarrera: "LM"},
		{apellido: "Alfa", nombre: "Beta", ciudadOrigen: "Madrid", fecha: fechaNacimiento, presentoTitulo: true, codigoCarrera: "LI"},
	}

	// Ordenar primero por edad
	sort.Sort(PorEdad(ingresantes))
	fmt.Println("Ordenado por Edad:")
	for _, ing := range ingresantes {
		fmt.Println(ing)
	}

	// Ordenar por apellido y nombre
	sort.Sort(PorNombre(ingresantes))
	fmt.Println("\nOrdenado por Apellido y Nombre:")
	for _, ing := range ingresantes {
		fmt.Println(ing)
	}
}
