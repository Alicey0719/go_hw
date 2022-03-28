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
/*
結果
$ go run uruu.go
0.0275698
0.0314057
0.0288166
$ go run uruu.go
0.0290247
0.029479
0.029212
$ go run uruu.go
0.0284367
0.0288101
0.029845
$ go run uruu.go
0.0297253
0.0285078
0.0363449
$ go run uruu.go
0.0284563
0.0284155
0.0307394
$ go run uruu.go
0.0299512
0.0393929
0.0337116
$ go run uruu.go
0.0320825
0.0329871
0.0296471
$ go run uruu.go
0.0270864
0.0336461
0.034481
$ go run uruu.go
0.0284298
0.032941
0.0300464
$ go run uruu.go
0.0284317
0.0344659
0.0289382
$ go run uruu.go
0.0305863
0.0285225
0.0289977
$ go run uruu.go
0.0298603
0.0349183
0.0308466

*/