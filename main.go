package main

import "fmt"

func main(){

	//var x map[string]string
	x := make(map[string]string)
	x["name"] = "Misbah"
	x["height"] = "5.7"
	delete(x, "height")
	fmt.Println(x)
	

}
