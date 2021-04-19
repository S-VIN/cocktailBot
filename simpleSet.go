package main

type Set struct{
	data map[int]string
	size int
}

func (set *Set) Add(input string)bool{
	if(set.Find(input)){
		return false
	}
	set.data[set.size] = input
	set.size++
	return true
}

func (set Set) Get(index int) (string, error){
	res, ok := set.data[index]
	if ok {
		return res, nil
	}else{
		return "", nil
	}
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