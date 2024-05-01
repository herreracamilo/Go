package main

import "fmt"

type PuntoCardinal int


const (
    Norte PuntoCardinal = iota // pone a norte en = 0, este = 1, sur = 2 y asi consecutivamente 
    Este
    Sur
    Oeste
    Noroeste
    Sureste
    Noreste
    Suroeste
)


func DireccionDelViento(viento string) PuntoCardinal {
    switch viento {
    case "N":
        return Norte
    case "S":
        return Sur
    case "E":
        return Este
    case "O":
        return Oeste
    case "NO":
        return Noroeste
    case "SE":
        return Sureste
    case "NE":
        return Noreste
    case "SO":
        return Suroeste
    default:
        return -1 // Valor inválido
    }
}

func main() {
    var viento string
    fmt.Print("Ingrese el punto cardinal del viento (N, S, E, O, NO, SE, NE, SO): ")
    fmt.Scanln(&viento)

    direccion := DireccionDelViento(viento)
    switch direccion {
    case Norte:
        fmt.Println("El viento se dirige hacia el Norte.")
    case Sur:
        fmt.Println("El viento se dirige hacia el Sur.")
    case Este:
        fmt.Println("El viento se dirige hacia el Este.")
    case Oeste:
        fmt.Println("El viento se dirige hacia el Oeste.")
    case Noroeste:
        fmt.Println("El viento se dirige hacia el Noroeste.")
    case Sureste:
        fmt.Println("El viento se dirige hacia el Sureste.")
    case Noreste:
        fmt.Println("El viento se dirige hacia el Noreste.")
    case Suroeste:
        fmt.Println("El viento se dirige hacia el Suroeste.")
    default:
        fmt.Println("El punto cardinal del viento ingresado es inválido.")
    }
}
