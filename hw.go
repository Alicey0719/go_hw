package main

import "fmt"
import "strconv"


type Person struct {
    name string
	voice string
	grade int
}

func (p *Person) SetPerson(name string, voice string, grade int) {
    p.name = name
    p.voice = voice
	p.grade = grade
}

func (p *Person) GetPerson() (string) {
	//string(int)でキャストできん 空になる
    return "Grade: " + strconv.Itoa(p.grade) + ": "+ p.name + "( CV:" + p.voice + ")"
}


func maxreturn(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func return_duble_int ()(int, int){
	return 10, 100
}

func print_strings(s ... string){
	for _, val := range s{
		fmt.Println(val)
	}
}

func main(){
	fmt.Printf("Hello World\n")

	// 宣言が:=
	str_len := []string{} // スライス。個数不定
	str_len = append(str_len, "あやちねね")
	str_len = append(str_len, "いなばめぐる")
	str_len = append(str_len, "しいばつむぎ")
	str_len = append(str_len, "とがくしとうこ")

	fmt.Println(len(str_len))

	fmt.Println(str_len[0], str_len[1])
	for i:=0; i<len(str_len);i++ {
		fmt.Println(str_len[i])
	}

	var x1 int = 10
	x2 := 100
	fmt.Println(maxreturn(x1,x2))

	test, _ :=  return_duble_int() //要らない返り値は_で受け取る
	fmt.Println(test)  //使ってない変数あると怒られる！

	print_strings("はみだしくりえいてぃぶ", "わがままはいすぺっく")
	
	var p_nene Person
	p_nene.SetPerson("綾地寧々","桐谷華",2)
	fmt.Println(p_nene.GetPerson())
	
}