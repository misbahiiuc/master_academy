package main

import "fmt"
import "reflect"

func main(){

	//students := [...]string{"Lutful", "Kader", "Misbah", "student", "CSE"}
	//fmt.Println(students)
	//fmt.Println(len(students))
	//x := students[0:2]
	//fmt.Println(x)
	//x := make([]string, 0)
	//fmt.Println(x)
	var fruits []string
	//fmt.Println(fruits)
	fruits = append(fruits, "Apple", "Banana")
	//fmt.Println(fruits, len(fruits))
	//fmt.Printf("%T\n", fruits)
	//fmt.Printf("%T", students)
	
	b := reflect.TypeOf(fruits).Kind().String()
	fmt.Println(b)
	

}
