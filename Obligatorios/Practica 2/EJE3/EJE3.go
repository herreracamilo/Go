package main
import (
	"fmt"
)

type elemento struct{
	numero int
	contador int
}

type OptimumSlice struct{
	elementos []elemento
}

func New(s []int) OptimumSlice{
	if(len(s)==0){
		return OptimumSlice{elementos: []elemento{}}
	}
	
	var auxiliar []elemento
	actual:= s[0]
	cont:= 1

	// ahora recorro el slice s para meter los numeros en el auxiliar
	for i:=1; i < len(s);i++{
		if(s[i]==actual){
			cont++
		}else{ // si el numero no era igual al actual tengo que agrandar el slices y meter el actual y el contador que tenia y despues tomar el nuevo y reiniciar el contador
			auxiliar = append(auxiliar, elemento{numero: actual, contador: cont})
			actual = s[i]
			cont = 1
		}
	}
	auxiliar = append(auxiliar, elemento{numero: actual,contador: cont})
	return OptimumSlice{elementos: auxiliar}
}


func IsEmpty(o OptimumSlice) bool{
	if(len(o.elementos)== 0){
		return true
	}
	return false
}

func Len(o OptimumSlice) int{
	return len(o.elementos)
}

func FrontElement(o OptimumSlice) int{
	if IsEmpty(o) { // si esta vacio tiro una advertencia
		panic("OptimumSlice is empty")
	}
	return o.elementos[0].numero
}

func LastElement(o OptimumSlice) int{
	if IsEmpty(o) { // si esta vacio tiro una advertencia
		panic("OptimumSlice is empty")
	}
	return o.elementos[Len(o)-1].numero
}

func Insert(o *OptimumSlice, element int, position int) {
	// si la posicion es menor a 0 o mas grande que el tamaño del OptimumSlice tiro un panic y salgo porque no puedo operar
	if position < 0 || position > Len(*o) {
		panic("Position out of bounds")
	}

	// posicion actual
	actualPos := 0
	// recorro cada elemento para encontrar donde tengo que insertar
	for i, el := range o.elementos {
		// aca verifica si la posición de insercion cae dentro del rango del elemento actual 
		if position <= actualPos+el.contador {
			//si el elemento a insertar es igual al elemento actual solo se incrementa el contador
			if el.numero == element {
				o.elementos[i].contador++
			} else { // sino necesito separarlo en lado izquierdo y derecho porque el que quiero insertar va en el medio
				izqCont := position - actualPos
				derCont := el.contador - izqCont
				//creo un slice auxiliar
				auxElementos := []elemento{
					//al silce meto:
					//la primera parte del elemento que estaba
					{numero: el.numero, contador: izqCont},
					// el nuevo que quiero meter con el contador en 1
					{numero: element, contador: 1},
					//la otra parte del elemento que estaba con lo que quedaba del contador
					{numero: el.numero, contador: derCont},
				}
				// ahora meto los elementos del nuevo slice 
				o.elementos = append(o.elementos[:i], append(auxElementos, o.elementos[i+1:]...)...)
			}
			return
		}
		//actualizo la posicion actual
		actualPos += el.contador
	}
	//si la posicion es al final solo lo agrego al final con el contador en 1
	o.elementos = append(o.elementos, elemento{numero: element, contador: 1})
}

func SliceArray(o OptimumSlice) []int{
	var sliceNormal []int
	if(IsEmpty(o)){
		return sliceNormal
	}
	
	for _, elem:= range o.elementos{
		for i := 0; i < elem.contador; i++ {
			sliceNormal= append(sliceNormal, elem.numero)
		}
	}
	return sliceNormal
}

func main()  {
	slice:= []int {3,3,3,3,3,7,7,7,7,7,7,7,23,23,23,23,23,23,3,3,3,3,3,3,3,3,7,5,5,5}
	fmt.Println("antes")
	fmt.Println(slice)
	fmt.Println("")
	fmt.Println("ahora")
	newS:= New(slice)
	fmt.Println(newS)
	fmt.Println(Len(newS))
	fmt.Println(LastElement(newS))
	Insert(&newS,88,3)
	fmt.Println(newS)
	optimToNormal:= SliceArray(newS)
	fmt.Println(optimToNormal)
}