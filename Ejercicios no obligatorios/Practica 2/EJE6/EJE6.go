package main
import(
	"fmt"
)

func main()  {
	slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{6, 7, 8, 9, 10,11,12,13,14}
	fmt.Println(slice1)
	fmt.Println(slice2)
	sumado := Sum(slice1,slice2)
	fmt.Println(sumado)
}

func Sum(a , b []int) []int{
	var min int
	if(len(a)<len(b)){
		min = len(a)
	}else{
		min = len(b)
	}
	sliceSuma:= make([]int,min)
	for i := 0; i < min; i++ {
		sliceSuma[i] = a[i] + b[i]
	}
	return sliceSuma
}

/*
func Avg(a []int) int{

}
*/