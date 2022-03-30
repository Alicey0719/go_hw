package main

import(
	"fmt"
)

func fizzbuzz(num int) string{
	if num <= 0{
		return "Put 1 or more"
	}else if num%15 == 0{
		return "FizzBuzz"
	}else if num%5 == 0{
		return "Buzz"
	}else if num%3 == 0{
		return "Fizz"
	}else{
		return "no fizzbuzz"
	}
}

func main(){
	var num int
	fmt.Print("num> ")
	fmt.Scan(&num)
	fmt.Println(fizzbuzz(num))
}