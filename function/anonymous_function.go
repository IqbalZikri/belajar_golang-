// Anonymous Function
// Sebelum setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// hal tersebut dinamakan anonymous function, atau function tanpa nama

package main

import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main(){
	blacklist := func(name string)bool{
		return name == "anjing"
	}

	registerUser("eko", blacklist)
	registerUser("anjing", func(name string)bool{
		return name == "anjing"
	})
}