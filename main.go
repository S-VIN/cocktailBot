package main

import "fmt"



func main() {

	err := telegram.CreateBot()
	if err != nil{
		fmt.Println(err)
	}
	telegram.CheckUpdates()
}