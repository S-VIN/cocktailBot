package main

import (
	"github.com/pkg/errors"
)

type Set struct {
	data map[int]string
}

func (set *Set) Add(input string) {
	if !set.Find(input) {
		set.data[set.GetSize()] = input
	}
}

func (set Set) GetByIndex(index int) (res string, err error) {
	res, ok := set.data[index]
	if ok {
		return
	} else {
		err = errors.New("index not created")
		return
	}
}

func (set Set) GetSize() int {
	return len(set.data)
}

func (set Set) Find(input string) bool {
	for _, val := range set.data {
		if val == input {
			return true
		}
	}
	return false
}

func NewSet() *Set {
	var set Set
	set.data = make(map[int]string)
	return &set
}
