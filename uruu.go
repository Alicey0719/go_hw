package main

import(
	"fmt"
	"time"
)

func p1(year int) bool{
	if (year%4==0 && year%100!=0) || year%400==0 {
		return true
	}else{
		return false
	}
}

func p2(year int) bool{
	if year%400==0 || (year%100==0 && year%4==0) {
		return true
	}else{
		return false
	}
}

func p3(year int) bool{
	if (year%100!=0 && year%4==0) || year%400==0 {
		return true
	}else{
		return false
	}
}

func main(){
	
	years := []int{}
	for i:=0; i<100000000; i++{
		years = append(years, i)
	}
	//fmt.Println(years)

	tmp := true
	t := time.Now()
	for _, y := range years{
		tmp = p1(y)
	}	
	fmt.Println(time.Since(t).Seconds())

	t = time.Now()
	for _, y := range years{
		tmp = p2(y)
	}	
	fmt.Println(time.Since(t).Seconds())

	t = time.Now()
	for _, y := range years{
		tmp = p3(y)
	}	
	_ = tmp
	fmt.Println(time.Since(t).Seconds())


}