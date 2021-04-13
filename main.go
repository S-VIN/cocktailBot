package main

import (
   "fmt"
)


func main() {
	resp, _ := lookUpIngredientById("552")
	fmt.Println(resp)
}