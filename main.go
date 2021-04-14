package main

import (
   "fmt"
)


func main() {
	resp, _ := searchCocktailByName("vodka")
	fmt.Println(resp)
}