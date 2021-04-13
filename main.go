package main

import (
   "fmt"
)


func main() {
	resp, _ := getRandomCocktail()
	fmt.Println(resp)
}