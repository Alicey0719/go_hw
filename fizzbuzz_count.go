package main

import(
	"fmt"
)

func fizzbuzz_count(b int, a int) (int,int,int){
	fizz := 0
	buzz := 0
	fizzbuzz :=0 

	for i:=b; i<=a; i++{
		if i <= 0{

		}else if i%15 == 0{
			fizzbuzz++
		}else if i%5 == 0{
			buzz++
		}else if i%3 == 0{
			fizz++
		}
	}

	return fizz,buzz,fizzbuzz
}

func main(){
	b := 1
	a := 1000
	fmt.Println(fizzbuzz_count(b,a))
}