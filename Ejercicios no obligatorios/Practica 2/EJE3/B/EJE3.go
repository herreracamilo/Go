package main

import (
	
)

type PuntoCardinal string


const (
	Norte PuntoCardinal = "N"
    Sur   PuntoCardinal = "S"
    Este  PuntoCardinal = "E"
    Oeste PuntoCardinal = "O"
    Noreste PuntoCardinal = "NE"
    Sureste PuntoCardinal = "SE"
    Noroeste PuntoCardinal = "NO"
    Suroeste PuntoCardinal = "SO"
)

func main()  {
	
}

func DireccionDelViento(viento string)PuntoCardinal  {
	var mapViento = map[PuntoCardinal]PuntoCardinal{
		"N":  PuntoCardinal("S"),
        "S":  PuntoCardinal("N"),
        "E":  PuntoCardinal("O"),
        "O":  PuntoCardinal("E"),
        "NO": PuntoCardinal("SE"),
        "SE": PuntoCardinal("NO"),
        "NE": PuntoCardinal("SO"),
        "SO": PuntoCardinal("NE"),
    }
    return mapViento[PuntoCardinal(viento)]
}
