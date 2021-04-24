package main

import "fmt"

func main() {

	err := telegram.CreateBot()
	if err != nil {
		fmt.Println(err)
	}
	telegram.CheckUpdates()
}

func main2(){
	var database Database
	database.Init()
	

	database.Like(81, "second")
	database.Like(82, "third")
	database.Like(83, "fourth")
	
	fmt.Println(database.IsLike(81, "second"))
	fmt.Println(database.IsLike(82, "second"))
}