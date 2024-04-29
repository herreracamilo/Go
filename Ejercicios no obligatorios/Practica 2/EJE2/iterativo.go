package main

func factorial(n int) int{
	result:=1
	if (n== 0){
		return result
	}
	for i:= 1;i <= n;i++{
		result*=i
	}
	return result
}