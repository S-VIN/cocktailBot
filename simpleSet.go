package main

import(
	//"fmt"
)

type Set struct{
	data map[int]string
}

func (set *Set) Add(input string)bool{
	set.data[set.GetSize()] = input
	return true
}

func (set Set) GetByIndex(index int) (string, error){
	res, ok := set.data[index]
	if ok {
		return res, nil
	}else{
		return "", nil
	}
}

func (set Set) GetSize() int{
	return len(set.data)
}

func (set Set) Find(input string) bool{
	for _, val := range(set.data){
		if(val == input){
			return true
		}
	}
	return false
}

func NewSet() *Set{
	var set Set
	set.data = make(map[int]string)
	return &set
}